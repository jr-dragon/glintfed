# 方案二：ory/fosite — 嵌入式 OAuth2 Server

## 概述

將 ory/fosite 作為 library 直接嵌入 glintfed 程序中，由 glintfed 自行實作 OAuth2 Authorization Server 的完整功能。類似 Laravel Passport 在 pixelfed 的模式，所有 OAuth2 邏輯、token 儲存都在 glintfed 程序內完成，不依賴外部服務。

```
Client App
    │
    ├─ POST /oauth/token (password / refresh_token / client_credentials)
    │       ↓
    │   svc.Token
    │       ├─ grant_type=password → handlePasswordGrant（service 層手動處理，不走 fosite ROPC）
    │       └─ 其他 grant type → provider.NewAccessRequest / NewAccessResponse
    │
    └─ GET  /oauth/authorize (authorization_code)
            ↓
        svc.Authorize（目前回傳 ErrAccessDenied；Login UI 尚未實作）
```

### 關鍵設計決策

- **HMAC-only token**：使用 `oauth2.HMACSHAStrategy`，token 為不透明的 HMAC 字串，不使用 JWT/RSA。
- **密碼 grant 手動處理**：`handlePasswordGrant` 在 service 層直接驗證帳密，繞過 fosite 的 ROPC handler（`OAuth2ResourceOwnerPasswordFactory` 未納入 `compose.Compose`）。
- **Pixelfed schema 相容**：ent schema 與 pixelfed 的 Laravel Passport MySQL schema 完全相同，方便使用者從 pixelfed 遷移至 glintfed。不儲存 session blob，session 改由 DB 欄位重建。
- **直接核發 token**：`CreatePersonalAccessTokens` 方法讓 Usecase 可在不走 HTTP authorize flow 的情況下直接簽發 access token + refresh token（對應 pixelfed 的 `$user->createToken(...)`）。

---

## Go 依賴

```bash
go get github.com/ory/fosite
```

> fosite 不需要 `go-jose`，因為實作中只使用 HMAC 策略。

---

## 核心概念

fosite 要求實作 storage interfaces：

1. **`fosite.Storage`**：儲存 OAuth2 clients、authorize codes、access tokens、refresh tokens、PKCE sessions。
2. **`fosite.OAuth2Provider`**：由 `compose.Compose(...)` 建立，注入 storage 與各 grant handler。

Session 不以 JSON blob 存入 DB，而是在查詢時由 `user_id`、`scopes`、`expires_at` 欄位重建，以相容 pixelfed 的 schema。

---

## Ent Schema

所有 schema 與 pixelfed 的 Laravel Passport MySQL schema 完全相容：

### `ent/schema/oauthaccesstoken.go` → `oauth_access_tokens`

```go
field.String("id").MaxLen(100).Unique()
field.Uint64("user_id").Optional()           // client_credentials 時可為空
field.Uint64("client_id")
field.String("name").MaxLen(191).Optional()
field.Text("scopes").Optional()              // JSON text: ["read","write"]
field.Bool("revoked")
field.Time("created_at").Default(time.Now).Immutable()
field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now)
field.Time("expires_at").Optional()
// table: oauth_access_tokens; index: user_id
```

### `ent/schema/oauthauthorizationcode.go` → `oauth_auth_codes`

```go
field.String("id").MaxLen(100).Unique()
field.Uint64("user_id")                      // NOT NULL（authorize code 必有使用者）
field.Uint64("client_id")
field.Text("scopes").Optional()
field.Bool("revoked")
field.Time("expires_at").Optional()
// index: user_id
```

### `ent/schema/oauthrefreshtoken.go` → `oauth_refresh_tokens`

```go
field.String("id").MaxLen(100).Unique()
field.String("access_token_id").MaxLen(100)  // 關聯 oauth_access_tokens.id
field.Bool("revoked")
field.Time("expires_at").Optional()
// index: access_token_id
```

### `ent/schema/oauthpersonalaccessclient.go` → `oauth_personal_access_clients`

```go
field.Uint64("id").Unique()
field.Uint64("client_id")
field.Time("created_at").Default(time.Now).Immutable()
field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now)
```

### `ent/schema/oauthpkce.go`（glintfed-only，pixelfed 無此表）

```go
// 儲存 PKCE session blob，僅 glintfed 使用
```

### 現有 `ent/schema/oauthclient.go`

沿用 pixelfed schema，無需新增額外欄位。`FositeClient` 會從現有欄位（`password_client`、`personal_access_client`、`redirect`、`revoked`）推導出 fosite 所需的 metadata。

執行 `make gen` 重新產生 ent code。

---

## 實作

### Step 1：FositeStore（`internal/lib/fositestore/`）

**`store.go`**

```go
// Store 實作 fosite storage interfaces，底層使用 ent ORM。
type Store struct {
    db       *ent.Client
    strategy *oauth2.HMACSHAStrategy
}

// New 建立 Store，並從 config 的 HMACSecret 初始化 HMAC 策略。
func New(client *data.Client, cfg *data.Config) *Store {
    globalSecret := []byte(cfg.App.Auth.OAuth.HMACSecret)
    fositeCfg := &fosite.Config{GlobalSecret: globalSecret}
    strategy := compose.NewOAuth2HMACStrategy(fositeCfg)
    return &Store{db: client.Ent, strategy: strategy}
}

// Strategy 回傳底層 HMAC 策略，供 NewOAuth2Provider 使用。
func (s *Store) Strategy() *oauth2.HMACSHAStrategy { return s.strategy }

// CreatePersonalAccessTokens 直接簽發 access token + refresh token，
// 繞過標準 OAuth2 HTTP flow（對應 pixelfed 的 $user->createToken(...)）。
// 關鍵：呼叫 CreateRefreshTokenSession 前先 req.SetID(atSig)，
// 讓 req.GetID() 作為 refresh token 的 access_token_id。
func (s *Store) CreatePersonalAccessTokens(ctx context.Context, req fosite.Requester) (accessToken, refreshToken string, err error) {
    at, atSig, err := s.strategy.GenerateAccessToken(ctx, req)
    rt, rtSig, err := s.strategy.GenerateRefreshToken(ctx, req)
    s.CreateAccessTokenSession(ctx, atSig, req)
    req.SetID(atSig)  // 設定 req ID = access token 的 signature
    s.CreateRefreshTokenSession(ctx, rtSig, req)
    return at, rt, nil
}

// RevokeAccessToken / RevokeRefreshToken / RotateRefreshToken
// 均透過 UpdateOneID(...).SetRevoked(true) 實作。
```

**`session.go`**

```go
// marshalScopes 將 scopes 序列化為 pixelfed 的 JSON text 格式：["read","write"]
func marshalScopes(scopes []string) string

// unmarshalScopes 解析 pixelfed 的 JSON text scopes 欄位
func unmarshalScopes(text string) []string
```

**`client.go`**

```go
// FositeClient 包裝 *ent.OauthClient，實作 fosite.Client interface。
// 不儲存額外欄位，所有 fosite metadata 從現有 pixelfed 欄位推導：
//   - GetGrantTypes()：從 PasswordClient bool 決定是否包含 "password"
//   - GetScopes()：固定回傳 ["read","write","follow","push"]
//   - IsPublic()：等於 PersonalAccessClient bool
//   - GetRedirectURIs()：Split(Redirect, "\n")
//   - GetID()：strconv.FormatUint(c.ID, 10)
//
// GetClient 解析 string ID → uint64，並在 c.Revoked 時回傳 ErrInvalidClient。
```

**`access_token.go`**

```go
// CreateAccessTokenSession
//   - client_id 必須能解析為 uint64，否則回傳錯誤
//   - subject 為空（client_credentials）時跳過 user_id 設定
//   - subject 非空時必須能解析為 uint64，否則回傳錯誤
//
// GetAccessTokenSession
//   - 從 user_id、scopes、expires_at 重建 fosite.DefaultSession（無 blob）
//
// DeleteAccessTokens(clientID string)
//   - 先將 clientID string 解析為 uint64，否則回傳錯誤
```

**`authorize_code.go`**

```go
// CreateAuthorizeCodeSession
//   - client_id 與 user_id（NOT NULL）均須能解析為 uint64，否則回傳錯誤
//
// GetAuthorizeCodeSession
//   - 從 user_id、scopes、expires_at 重建 session
//   - revoked=true 時回傳 fosite.ErrInvalidatedAuthorizeCode
//
// InvalidateAuthorizeCodeSession：UpdateOneID(...).SetRevoked(true)
```

**`refresh_token.go`**

```go
// CreateRefreshTokenSession
//   - 使用 req.GetID() 作為 access_token_id（由 CreatePersonalAccessTokens 設定）
//
// GetRefreshTokenSession
//   - JOIN oauth_access_tokens 取得 user_id、client_id、scopes
//   - 從這些欄位重建 session（無 blob）
//
// DeleteRefreshTokens(clientID string)
//   - 先將 clientID 解析為 uint64，否則回傳錯誤
//   - 查詢所有 access token ID，再批次 revoke refresh tokens
```

---

### Step 2：fosite Provider（`internal/lib/fositestore/provider.go`）

```go
// NewOAuth2Provider 建立 fosite.OAuth2Provider。
// TTL 從 config 讀取（預設 access=365天、refresh=400天）。
// 使用 HMAC-only 策略（store.Strategy()），不使用 JWT/RSA。
// 注意：不包含 OAuth2ResourceOwnerPasswordFactory，
//       password grant 改由 service 層手動處理。
func NewOAuth2Provider(store *Store, cfg *data.Config) fosite.OAuth2Provider {
    fositeCfg := &fosite.Config{
        AccessTokenLifespan:        time.Duration(tokenDays) * 24 * time.Hour,
        RefreshTokenLifespan:       time.Duration(refreshDays) * 24 * time.Hour,
        AuthorizeCodeLifespan:      10 * time.Minute,
        SendDebugMessagesToClients: false,
    }

    strategy := &compose.CommonStrategy{CoreStrategy: store.Strategy()}

    return compose.Compose(
        fositeCfg,
        store,
        strategy,
        compose.OAuth2AuthorizeExplicitFactory,      // authorization_code
        compose.OAuth2ClientCredentialsGrantFactory, // client_credentials
        compose.OAuth2RefreshTokenGrantFactory,      // refresh_token
        compose.OAuth2TokenRevocationFactory,        // POST /oauth/revoke
        compose.OAuth2TokenIntrospectionFactory,     // token introspection（middleware 使用）
        compose.OAuth2PKCEFactory,                   // PKCE
        // OAuth2ResourceOwnerPasswordFactory 刻意不納入，
        // password grant 在 service 層以 handlePasswordGrant 手動處理。
    )
}
```

---

### Step 3：OAuth Usecase（`internal/usecase/oauth/oauth.go`）

負責在不走 HTTP flow 的情況下直接簽發 token（例如 AppRegister onboarding）：

```go
type Usecase struct {
    store    *fositestore.Store
    clientID string        // PersonalClientID（config.App.Auth.OAuth.PersonalClientID）
    tokenTTL time.Duration // AccessTokenLifespanDays，預設 365 天
}

func NewUsecase(store *fositestore.Store, cfg *data.Config) *Usecase

// CreateTokens 對應 pixelfed 的 $user->createToken('Pixelfed App', scopes)
func (uc *Usecase) CreateTokens(ctx context.Context, userID uint64, scopes []string) (*TokenResult, error) {
    subject := strconv.FormatUint(userID, 10)  // 使用 strconv，不用 fmt.Sprintf

    client, err := uc.store.GetClient(ctx, uc.clientID)

    session := &fosite.DefaultSession{
        Subject:  subject,
        Username: subject,
        ExpiresAt: map[fosite.TokenType]time.Time{
            fosite.AccessToken:  now.Add(uc.tokenTTL),
            fosite.RefreshToken: now.Add(uc.tokenTTL + 35*24*time.Hour),
        },
    }

    req := fosite.NewRequest()
    // 設定 client、scopes、session 後呼叫：
    accessToken, refreshToken, err := uc.store.CreatePersonalAccessTokens(ctx, req)

    // 回傳 ClientSecret = string(fositeClient.GetHashedSecret())
    return &TokenResult{
        AccessToken:  accessToken,
        RefreshToken: refreshToken,
        ClientID:     uc.clientID,
        ClientSecret: string(fositeClient.GetHashedSecret()),
        ExpiresIn:    int64(uc.tokenTTL.Seconds()),
    }, nil
}
```

---

### Step 4：OAuth Service（`internal/service/oauth/`）

**`service.go`**

```go
type Service interface {
    Authorize(w http.ResponseWriter, r *http.Request) // GET /oauth/authorize
    Token(w http.ResponseWriter, r *http.Request)     // POST /oauth/token
    Revoke(w http.ResponseWriter, r *http.Request)    // POST /oauth/revoke
    // 注意：無 Introspect（僅供 middleware 內部使用）
}

//go:generate go tool moq -rm -out mock_user_authenticator.go . UserAuthenticator
type UserAuthenticator interface {
    Authenticate(ctx context.Context, username, password string) (uint64, error)
}

type svc struct {
    provider        fosite.OAuth2Provider
    store           *fositestore.Store
    auth            UserAuthenticator
    appURL          string
    accessTokenTTL  time.Duration // 從 config 讀取，預設 365 天
    refreshTokenTTL time.Duration // 從 config 讀取，預設 400 天
}

func New(provider fosite.OAuth2Provider, store *fositestore.Store, auth UserAuthenticator, cfg *data.Config) Service
```

**`token.go`**

```go
func (s *svc) Token(w http.ResponseWriter, r *http.Request) {
    // password grant 由 handlePasswordGrant 手動處理（不走 fosite）
    if r.FormValue("grant_type") == "password" {
        s.handlePasswordGrant(w, r.WithContext(ctx))
        return
    }
    // 其他 grant type 交給 fosite
    accessReq, err := s.provider.NewAccessRequest(ctx, r, &fosite.DefaultSession{})
    accessResp, err := s.provider.NewAccessResponse(ctx, accessReq)
    s.provider.WriteAccessResponse(ctx, w, accessReq, accessResp)
}

func (s *svc) handlePasswordGrant(w http.ResponseWriter, r *http.Request) {
    // 1. 驗證 username/password → auth.Authenticate
    // 2. 解析 scope、client_id
    // 3. 建立 fosite.DefaultSession（ExpiresAt 使用 s.accessTokenTTL / s.refreshTokenTTL）
    // 4. 呼叫 store.CreatePersonalAccessTokens
    // 5. 回傳 JSON：access_token, refresh_token, token_type, expires_in, scope
    //    expires_in = int64(s.accessTokenTTL.Seconds())
    subject := strconv.FormatUint(userID, 10)  // 使用 strconv
}
```

**`authorize.go`**

```go
// 目前 Authorize 回傳 ErrAccessDenied（Login UI 尚未實作）。
// TODO: 加入使用者驗證 / session 檢查，完成 authorization_code flow。
func (s *svc) Authorize(w http.ResponseWriter, r *http.Request) {
    authReq, err := s.provider.NewAuthorizeRequest(ctx, r)
    s.provider.WriteAuthorizeError(ctx, w, authReq,
        fosite.ErrAccessDenied.WithDescription("user authentication not implemented"))
}
```

---

### Step 5：Token 驗證 Middleware（`internal/server/middleware/oauth.go`）

```go
// OAuth2Auth 驗證 Bearer token，成功後將 subject 與 scopes 注入 context。
func OAuth2Auth(provider fosite.OAuth2Provider) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            token := extractBearerToken(r)  // strings.CutPrefix("Bearer ", ...)
            if token == "" {
                http.Error(w, `{"error":"missing_token"}`, http.StatusUnauthorized)
                return
            }
            _, ar, err := provider.IntrospectToken(
                r.Context(), token, fosite.AccessToken, &fosite.DefaultSession{},
            )
            if err != nil {
                http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
                return
            }
            ctx := context.WithValue(r.Context(), CtxKeySubject, ar.GetSession().GetSubject())
            ctx = context.WithValue(ctx, CtxKeyScopes, ar.GetGrantedScopes())
            next.ServeHTTP(w, r.WithContext(ctx))
        })
    }
}

// SubjectFromContext / ScopesFromContext 供 handler 取用認證資訊。
```

---

### Step 6：Config（`internal/data/config.go`）

```go
type OAuthConfig struct {
    HMACSecret               string `mapstructure:"hmac_secret"               env:"OAUTH_HMAC_SECRET"`
    PersonalClientID         string `mapstructure:"personal_client_id"         env:"OAUTH_PERSONAL_CLIENT_ID"`
    AccessTokenLifespanDays  int    `mapstructure:"access_token_lifespan_days" env:"OAUTH_TOKEN_EXPIRATION"`
    RefreshTokenLifespanDays int    `mapstructure:"refresh_token_lifespan_days" env:"OAUTH_REFRESH_EXPIRATION"`
}
// 注意：無 PrivateKeyPath（HMAC-only，不使用 JWT/RSA）
```

---

### Step 7：Route 註冊（`internal/server/api.go`）

```go
mux.Get("/oauth/authorize", svcs.OAuth.Authorize)
mux.Post("/oauth/token",    svcs.OAuth.Token)
mux.Post("/oauth/revoke",   svcs.OAuth.Revoke)
// 無 /oauth/introspect route（introspection 僅在 middleware 內部使用）
```

---

### Step 8：DI 註冊（`cmd/api/kessoku.go`）

```go
// fositestore
kessoku.Provide(fositestore.New),
kessoku.Provide(fositestore.NewOAuth2Provider),

// OAuth service
kessoku.Bind[oauthsvc.Service](kessoku.Provide(oauthsvc.New)),

// UserAuthenticator 綁定到 *user.Model（實作 Authenticate 方法）
kessoku.Bind[oauthsvc.UserAuthenticator](kessoku.Provide(func(m *usermodel.Model) oauthsvc.UserAuthenticator {
    return m
})),

// OAuth usecase（供 AppRegister onboarding 使用）
kessoku.Bind[appregistersvc.OAuthUsecase](kessoku.Provide(oauthuc.NewUsecase)),
```

---

## 資料流：password grant

```
POST /oauth/token
  grant_type=password&username=foo&password=bar&client_id=1&scope=read+write
    │
    ▼ svc.Token 偵測到 grant_type=password → handlePasswordGrant
    │
    ├─ auth.Authenticate(ctx, "foo", "bar") → userID uint64
    ├─ store.GetClient(ctx, "1") → FositeClient
    ├─ 建立 fosite.DefaultSession{Subject: strconv.FormatUint(userID,10), ExpiresAt: ...}
    ├─ 建立 fosite.Request{ID, Client, Session, Scopes}
    ├─ store.CreatePersonalAccessTokens(ctx, req)
    │     ├─ strategy.GenerateAccessToken  → at + atSig
    │     ├─ strategy.GenerateRefreshToken → rt + rtSig
    │     ├─ CreateAccessTokenSession(ctx, atSig, req)
    │     ├─ req.SetID(atSig)
    │     └─ CreateRefreshTokenSession(ctx, rtSig, req)  // req.GetID() = atSig
    └─ JSON response: {access_token, refresh_token, token_type, expires_in, scope}
```

## 資料流：AppRegister onboarding

```
POST /api/v1/apps/register/onboarding
    │
    ▼ svc.Onboarding
    ├─ 驗證 verify_code
    ├─ um.Create(ctx, ...)  → user
    ├─ ouc.CreateTokens(ctx, user.ID, scopes)  [OAuthUsecase]
    │     └─ 同上 CreatePersonalAccessTokens 流程
    └─ JSON response: {status, token_type, access_token, refresh_token,
                       client_id, client_secret, expires_in, scope, user, ...}
```

---

## 測試策略

- **Store 層**：使用 `data.NewTestClient(t)`（in-memory SQLite）做整合測試，不需要 mock
- **Usecase 層**：`fositestore.Store` 搭配 in-memory SQLite；無需 mock 外部服務
- **Service 層**：mock `UserAuthenticator`；`fosite.OAuth2Provider` 使用真實 in-memory 版本（`fositestore.New` + SQLite）
- **Middleware**：直接用 `httptest` 搭配真實 provider 測試 token 驗證

---

## 已知限制與風險

| 項目 | 說明 |
|------|------|
| `grant_type=password` (ROPC) | 手動處理，完整支援。ROPC 在 OAuth 2.1 已廢棄，未來客戶端升級時需遷移 |
| Login/Consent UI | `Authorize` 目前回傳 `ErrAccessDenied`。authorization_code flow 需實作 Login UI 後才能完整運作 |
| HMAC secret 輪替 | 輪替 `HMACSecret` 會使所有現有 token 失效；需搭配 key rotation 策略（雙 secret 期間的 token 仍需可驗證） |
| OOB redirect URI | `urn:ietf:wg:oauth:2.0:oob` 的支援（pixelfed mobile app 使用）需客製化 fosite redirect URI 驗證邏輯 |
| fosite 文件品質 | 官方文件較少，主要依賴原始碼與 ory/hydra 的實作作為參考 |
| Session 重建限制 | 因無 session blob，refresh token 換新 access token 時，新 access token 的 session 資訊來自 DB JOIN，無法保留原始請求的額外 claim |
