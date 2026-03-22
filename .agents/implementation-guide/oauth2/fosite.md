# 方案二：ory/fosite — 嵌入式 OAuth2 Server

## 概述

將 ory/fosite 作為 library 直接嵌入 glintfed 程序中，由 glintfed 自行實作 OAuth2 Authorization Server 的完整功能。類似 Laravel Passport 在 pixelfed 的模式，所有 OAuth2 邏輯、token 儲存都在 glintfed 程序內完成，不依賴外部服務。

```
Client App
    │
    ├─ POST /oauth/token (任何 grant type，包含 password)
    │       ↓
    │   glintfed fosite handler
    │       ↓
    │   fosite core（驗證 grant、簽發 token）
    │       ↓
    │   FositeStore（實作 fosite storage interfaces，底層為 ent）
    │
    └─ GET  /oauth/authorize (authorization_code)
            ↓
        glintfed fosite handler
            ↓
        Login UI（glintfed 自行渲染或由 client app 重導）
            ↓
        fosite.NewAuthorizeResponse（完成 authorize）
            ↓
        redirect to client with code
```

---

## 新增 Go 依賴

```bash
go get github.com/ory/fosite
go get github.com/ory/fosite/compose
go get github.com/go-jose/go-jose/v3
```

> fosite 使用 `go-jose` 做 JWT 簽署，需確認 Go 版本相容性。

---

## 核心概念

fosite 要求實作兩個主要 interface：

1. **`fosite.Storage`**：儲存 OAuth2 clients、authorize codes、access tokens、refresh tokens、PKCE sessions。
2. **`fosite.OAuth2Provider`**：由 `compose.Compose(...)` 建立，注入 storage 與各 grant handler。

所有 OAuth2 資料儲存在 glintfed 的主資料庫（SQLite/PostgreSQL），透過 ent 存取。

---

## Ent Schema 新增

需新增以下 ent schema（對應 Laravel Passport 的 5 張 migration table）：

### `ent/schema/oauthauthorizationcode.go`

```go
// Fields: id, client_id, user_id, scopes (JSON), redirect_uri, code_challenge,
//         code_challenge_method, session (JSON/blob), revoked, expires_at, created_at
```

### `ent/schema/oauthaccesstoken.go`

```go
// Fields: id (signature), client_id, user_id, scopes (JSON),
//         session (JSON/blob), revoked, expires_at, created_at
```

### `ent/schema/oauthrefreshtoken.go`

```go
// Fields: id (signature), access_token_id, client_id, user_id, scopes (JSON),
//         session (JSON/blob), revoked, expires_at, created_at
```

### `ent/schema/oauthpkce.go`

```go
// Fields: id (authorize_code_id), code_challenge, code_challenge_method,
//         session (JSON/blob), revoked, expires_at
```

現有的 `ent/schema/oauthclient.go` 需擴充以支援 fosite 所需欄位（`public` flag、`allowed_cors_origins` 等）。

執行 `make gen` 重新產生 ent code。

---

## 實作步驟

### Step 1：FositeStore（`internal/lib/fositestore/`）

實作 fosite 所有 storage interfaces，底層使用 ent。這是工程量最大的部份。

**`internal/lib/fositestore/store.go`**

```go
package fositestore

import (
    "glintfed.org/ent"
    fosite "github.com/ory/fosite"
)

// Store 實作 fosite 所有 storage interface：
//   - fosite.Storage
//   - oauth2.CoreStorage (AuthorizeCodeStorage + AccessTokenStorage + RefreshTokenStorage)
//   - pkce.PKCERequestStorage
//
// 每個 interface 對應一個 ent entity。
type Store struct {
    db *ent.Client
}

func NewStore(db *ent.Client) *Store {
    return &Store{db: db}
}

// 確保 Store 實作了所有必要 interface（編譯期檢查）
var (
    _ fosite.Storage = (*Store)(nil)
    _ interface {
        fosite.ClientManager
        fosite.AuthorizeCodeStorage
        fosite.AccessTokenStorage
        fosite.RefreshTokenStorage
    } = (*Store)(nil)
)
```

**`internal/lib/fositestore/client.go`（ClientManager）**

```go
// GetClient
//
//  SELECT * FROM oauth_clients WHERE id = ? LIMIT 1
func (s *Store) GetClient(ctx context.Context, id string) (fosite.Client, error) {
    client, err := s.db.OauthClient.Get(ctx, parseUint64(id))
    if ent.IsNotFound(err) {
        return nil, fosite.ErrNotFound
    }
    if err != nil {
        return nil, err
    }
    return toFositeClient(client), nil
}
```

**`internal/lib/fositestore/authorize_code.go`**

```go
// CreateAuthorizeCodeSession
//
//  INSERT INTO oauth_authorization_codes ...
func (s *Store) CreateAuthorizeCodeSession(ctx context.Context, code string, req fosite.Requester) error { ... }

// GetAuthorizeCodeSession
//
//  SELECT * FROM oauth_authorization_codes WHERE id = ? LIMIT 1
func (s *Store) GetAuthorizeCodeSession(ctx context.Context, code string, session fosite.Session) (fosite.Requester, error) { ... }

// InvalidateAuthorizeCodeSession
//
//  UPDATE oauth_authorization_codes SET revoked = true WHERE id = ?
func (s *Store) InvalidateAuthorizeCodeSession(ctx context.Context, code string) error { ... }
```

> 依此模式實作 `access_token.go`、`refresh_token.go`、`pkce.go`。每個 method 都加 SQL godoc comment，符合 glintfed 慣例。

---

### Step 2：fosite Provider 建立（`internal/lib/fositestore/provider.go`）

```go
package fositestore

import (
    "github.com/ory/fosite"
    "github.com/ory/fosite/compose"
    "github.com/ory/fosite/token/jwt"
)

// NewOAuth2Provider 建立 fosite OAuth2Provider，啟用所有需要的 grant types。
func NewOAuth2Provider(store *Store, secret []byte) fosite.OAuth2Provider {
    cfg := &fosite.Config{
        AccessTokenLifespan:   365 * 24 * time.Hour,
        RefreshTokenLifespan:  400 * 24 * time.Hour,
        AuthorizeCodeLifespan: 10 * time.Minute,
        GlobalSecret:          secret,
        // 允許 OOB redirect URI
        AllowedPromptValues: []string{"login", "none", "consent"},
    }

    privateKey := loadOrGenerateRSAKey() // 從 config 讀取或自動產生

    return compose.Compose(
        cfg,
        store,
        &compose.CommonStrategy{
            CoreStrategy: compose.NewOAuth2JWTStrategy(jwt.NewRS256JWTStrategy(privateKey), compose.NewOAuth2HMACStrategy(cfg), cfg),
        },
        compose.OAuth2AuthorizeExplicitFactory,       // authorization_code grant
        compose.OAuth2ResourceOwnerPasswordFactory,   // password grant（ROPC）
        compose.OAuth2ClientCredentialsGrantFactory,  // client_credentials grant
        compose.OAuth2RefreshTokenGrantFactory,       // refresh_token grant
        compose.OAuth2TokenRevocationFactory,         // token revocation
        compose.OAuth2TokenIntrospectionFactory,      // token introspection
        compose.OAuth2PKCEFactory,                    // PKCE
    )
}
```

---

### Step 3：OAuthUsecase 實作（`internal/usecase/oauth/oauth.go`）

```go
package oauth

import (
    "context"
    "fmt"
    "net/http"
    "time"

    "github.com/ory/fosite"
    "glintfed.org/internal/lib/fositestore"
)

type Usecase struct {
    provider fosite.OAuth2Provider
    store    *fositestore.Store
    clientID string // 系統預設 client（對應 pixelfed 的 personal_access_client）
}

func NewUsecase(provider fosite.OAuth2Provider, store *fositestore.Store, clientID string) *Usecase {
    return &Usecase{provider: provider, store: store, clientID: clientID}
}

// CreateTokens 直接在 store 層建立 access token + refresh token，
// 繞過 OAuth2 authorize flow，對應 pixelfed 的 $user->createToken(...)。
//
// 這與 fosite 的 AccessTokenStorage.CreateAccessTokenSession 直接互動。
func (uc *Usecase) CreateTokens(ctx context.Context, userID uint64, scopes []string) (*TokenResult, error) {
    subject := fmt.Sprintf("%d", userID)

    // 取得系統預設 client
    client, err := uc.store.GetClient(ctx, uc.clientID)
    if err != nil {
        return nil, err
    }

    session := &fosite.DefaultSession{
        Subject: subject,
        Claims: &jwt.IDTokenClaims{
            Subject:   subject,
            ExpiresAt: time.Now().Add(365 * 24 * time.Hour),
            IssuedAt:  time.Now(),
        },
    }

    // 建立一個 synthetic OAuth2 request
    req := fosite.NewAccessRequest(session)
    req.SetClient(client)
    req.SetRequestedScopes(fosite.Arguments(scopes))
    req.GrantScope(scopes...)

    // 直接寫入 store，取得 access token signature
    accessToken, accessSig, err := uc.provider.(*compose.Fosite).
        Config.GetAccessTokenStrategy(ctx).
        GenerateAccessToken(ctx, req)
    // ...

    // 此處需直接呼叫 store.CreateAccessTokenSession / CreateRefreshTokenSession
    // 詳見 fosite internal token creation pattern
    return &TokenResult{
        AccessToken:  accessToken,
        RefreshToken: refreshToken,
        ClientID:     uc.clientID,
        ClientSecret: client.(*fositestore.FositeClient).Secret,
        ExpiresIn:    int64(365 * 24 * 60 * 60),
    }, nil
}
```

> **實作說明**：fosite 沒有公開「直接核發 token 給 subject」的便利方法，因為這繞過了標準 OAuth2 flow。最乾淨的實作方式是：在 `Usecase` 中直接操作 `fositestore.Store`（`CreateAccessTokenSession` + `CreateRefreshTokenSession`），手動建構 `fosite.Requester`，不透過 provider 的 HTTP handler。這與 Passport 的 `$user->createToken()` 在語意上完全等價。

---

### Step 4：OAuth Service（`internal/service/oauth/`）

fosite 提供 HTTP handler helper，讓 service 層非常薄。

**`internal/service/oauth/service.go`**

```go
package oauth

import (
    "net/http"
    "github.com/ory/fosite"
    "glintfed.org/internal/lib/fositestore"
)

type Service interface {
    Authorize(w http.ResponseWriter, r *http.Request)  // GET /oauth/authorize
    Token(w http.ResponseWriter, r *http.Request)      // POST /oauth/token
    Revoke(w http.ResponseWriter, r *http.Request)     // POST /oauth/revoke
    Introspect(w http.ResponseWriter, r *http.Request) // POST /oauth/introspect
}

//go:generate go tool moq -rm -out mock_user_authenticator.go . UserAuthenticator
type UserAuthenticator interface {
    Authenticate(ctx context.Context, username, password string) (uint64, error)
}

type svc struct {
    provider fosite.OAuth2Provider
    auth     UserAuthenticator
}
```

**`internal/service/oauth/token.go`**

```go
func (s *svc) Token(w http.ResponseWriter, r *http.Request) {
    ctx, span := internal.T.Start(r.Context(), "OAuth.Token")
    defer span.End()

    // fosite 的 password grant 需要提供 ResourceOwnerPasswordCredentialsGrantHandler
    // 該 handler 會呼叫一個 ResourceOwnerPasswordCredentialsGrantHandler.HandleTokenEndpointRequest
    // 需傳入能驗證帳密的 session
    // 在 compose 階段注入自訂的 ResourceOwnerPasswordCredentialsGrantHandler

    ctx = context.WithValue(ctx, fositestore.CtxKeyAuthenticator, s.auth)

    accessReq, err := s.provider.NewAccessRequest(ctx, r, &fosite.DefaultSession{})
    if err != nil {
        s.provider.WriteAccessError(ctx, w, accessReq, err)
        return
    }

    accessResp, err := s.provider.NewAccessResponse(ctx, accessReq)
    if err != nil {
        s.provider.WriteAccessError(ctx, w, accessReq, err)
        return
    }

    s.provider.WriteAccessResponse(ctx, w, accessReq, accessResp)
}
```

**`internal/service/oauth/authorize.go`（處理 OOB）**

```go
func (s *svc) Authorize(w http.ResponseWriter, r *http.Request) {
    ctx, span := internal.T.Start(r.Context(), "OAuth.Authorize")
    defer span.End()

    // 偵測 OOB redirect_uri，fosite 不認識 urn:ietf:wg:oauth:2.0:oob
    // 需在 FositeClient 的 RedirectURIs 中預先加入此 URI，
    // 並在 fosite Config 中加入 RedirectSecureChecker 的豁免
    redirectURI := r.URL.Query().Get("redirect_uri")
    isOOB := redirectURI == "urn:ietf:wg:oauth:2.0:oob"

    authReq, err := s.provider.NewAuthorizeRequest(ctx, r)
    if err != nil {
        s.provider.WriteAuthorizeError(ctx, w, authReq, err)
        return
    }

    // 對於 authorization_code flow，通常需要 login session
    // pixelfed 的 mobile app flow 在此處重導到登入頁面
    // 對於已有 session 的使用者，直接 accept
    session := buildSession(r) // 從 cookie/session store 取得已登入使用者資訊
    response, err := s.provider.NewAuthorizeResponse(ctx, authReq, session)
    if err != nil {
        s.provider.WriteAuthorizeError(ctx, w, authReq, err)
        return
    }

    if isOOB {
        code := response.GetCode()
        w.Header().Set("Content-Type", "text/html")
        fmt.Fprintf(w, `<pre>%s</pre>`, code)
        return
    }

    s.provider.WriteAuthorizeResponse(ctx, w, authReq, response)
}
```

---

### Step 5：Password Grant — 自訂帳密驗證

fosite 的 `OAuth2ResourceOwnerPasswordFactory` 需要一個實作 `fosite.ResourceOwnerPasswordCredentialsGrantHandler` 的 handler，其中需要能驗證帳密。

**方式**：在 `compose.Compose(...)` 之後，將 store 中的 `PasswordHandler` 注入 `UserAuthenticator`：

```go
// internal/lib/fositestore/password_handler.go

type PasswordGrantHandler struct {
    fosite.HandleHelper
    auth UserAuthenticator
}

// ValidateResourceOwnerCredentials 驗證 username + password
func (h *PasswordGrantHandler) ValidateResourceOwnerCredentials(
    ctx context.Context, username, password string, request fosite.AccessRequester,
) error {
    userID, err := h.auth.Authenticate(ctx, username, password)
    if err != nil {
        return fosite.ErrNotFound.WithWrap(err)
    }
    request.GetSession().SetSubject(fmt.Sprintf("%d", userID))
    return nil
}
```

---

### Step 6：Token 驗證 Middleware

因為 fosite 嵌入在同一程序，token 驗證完全在本地完成，不需要任何網路呼叫：

```go
// internal/service/middleware/auth.go

func OAuth2Auth(provider fosite.OAuth2Provider) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            ctx := r.Context()
            tokenReq, err := provider.IntrospectToken(ctx, fosite.AccessTokenFromRequest(r), fosite.AccessToken, &fosite.DefaultSession{})
            if err != nil {
                http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
                return
            }
            ctx = context.WithValue(ctx, ctxKeySubject, tokenReq.GetSession().GetSubject())
            ctx = context.WithValue(ctx, ctxKeyScopes, tokenReq.GetGrantedScopes())
            next.ServeHTTP(w, r.WithContext(ctx))
        })
    }
}
```

---

### Step 7：Config 變更

### `internal/data/config.go`

```go
type AuthConfig struct {
    EnableRegistration bool         `mapstructure:"enable_registration" env:"OPEN_REGISTRATION"`
    EnableOAuth        bool         `mapstructure:"enable_oauth" env:"OAUTH_ENABLED"`
    InAppRegistration  bool         `mapstructure:"in_app_registration" env:"APP_REGISTER"`
    OAuth              OAuthConfig  `mapstructure:"oauth"`
}

type OAuthConfig struct {
    // RSA private key PEM 路徑，用於簽署 JWT access token
    PrivateKeyPath      string `mapstructure:"private_key_path" env:"OAUTH_PRIVATE_KEY_PATH"`
    // HMAC secret，用於 HMAC-based token（fallback）
    HMACSecret          string `mapstructure:"hmac_secret" env:"OAUTH_HMAC_SECRET"`
    // 預設 personal access client ID（AppRegister onboarding 用）
    PersonalClientID    string `mapstructure:"personal_client_id" env:"OAUTH_PERSONAL_CLIENT_ID"`
    AccessTokenLifespan int    `mapstructure:"access_token_lifespan_days" env:"OAUTH_TOKEN_EXPIRATION"`
    RefreshTokenLifespan int   `mapstructure:"refresh_token_lifespan_days" env:"OAUTH_REFRESH_EXPIRATION"`
}
```

---

### Step 8：Route 註冊（`internal/server/api.go`）

```go
// OAuth Routes（Mastodon-compatible）
mux.Get("/oauth/authorize", svcs.OAuth.Authorize)
mux.Post("/oauth/token", svcs.OAuth.Token)
mux.Post("/oauth/revoke", svcs.OAuth.Revoke)
mux.Post("/oauth/introspect", svcs.OAuth.Introspect)  // 可選，供管理工具使用
```

---

### Step 9：DI 註冊（`cmd/api/kessoku.go`）

```go
// fosite store
kessoku.Provide(func(client *data.Client) *fositestore.Store {
    return fositestore.NewStore(client.Ent)
}),

// fosite provider
kessoku.Provide(func(cfg *data.Config, store *fositestore.Store) fosite.OAuth2Provider {
    return fositestore.NewOAuth2Provider(store, []byte(cfg.App.Auth.OAuth.HMACSecret))
}),

// OAuth service
kessoku.Bind[oauthsvc.Service](kessoku.Provide(oauthsvc.New)),

// OAuth usecase（取代現有的 oauth.NewUsecase stub）
kessoku.Bind[appregister.OAuthUsecase](kessoku.Provide(oauthuc.NewUsecase)),
```

---

## 測試策略

- **Store 層**：使用 `data.NewTestClient(t)`（in-memory SQLite）做完整整合測試，不需要 mock
- **Usecase 層**：`FositeStore` 用 in-memory SQLite；無需 mock 外部服務
- **Service 層**：mock `UserAuthenticator`；`fosite.OAuth2Provider` 可以用真實的 in-memory 版本（fosite 提供 `storage/memory`）
- **E2E**：直接對 `/oauth/token` 發送 HTTP 請求，驗證 token 格式與 introspection 結果

```go
// 測試 password grant
func TestTokenPasswordGrant(t *testing.T) {
    // setup in-memory fosite + store
    // POST /oauth/token grant_type=password
    // assert 200 + valid access_token
}
```

---

## 已知限制與風險

| 項目 | 說明 |
|------|------|
| fosite Storage 實作量大 | 需實作約 15–20 個 interface method，每個都要對應 ent query；是此方案的主要工程量 |
| Session 序列化 | fosite session 以 JSON blob 存入 DB，欄位設計需謹慎；migration 後難以修改 |
| `grant_type=password` | 完整支援，但 ROPC 在 OAuth 2.1 中已廢棄，若未來客戶端升級可能需要遷移 |
| OOB redirect URI | fosite 的 redirect URI 驗證器需客製化，允許 `urn:ietf:wg:oauth:2.0:oob` 通過白名單檢查 |
| Login/Consent UI | authorization_code flow 仍需要 UI，但 pixelfed mobile app 主要走 password grant，短期內影響較小 |
| Key rotation | RSA key 輪替需要對應策略（雙 key 期間的 token 仍需可驗證） |
| fosite 文件品質 | 官方文件較少，主要依賴原始碼與 hydra 的實作作為參考 |
