# 方案一：ory/hydra — 外掛式 OAuth2 Authorization Server

## 概述

將 ory/hydra 作為獨立的 OAuth2/OIDC Authorization Server 部署於 glintfed 旁側。glintfed 同時扮演兩個角色：

1. **Login & Consent Provider**：hydra 在需要驗證使用者身份時，將使用者 redirect 到 glintfed 的登入 UI 與授權同意 UI，由 glintfed 完成驗證後通知 hydra。
2. **Mastodon-compatible OAuth Proxy**：glintfed 暴露 `/oauth/authorize`、`/oauth/token` 等 Mastodon 慣例的 endpoint，在背後與 hydra 互動。

```
Client App
    │
    ├─ POST /oauth/token (password grant)
    │       ↓
    │   glintfed (自行驗帳密 → 呼叫 hydra Admin API 強制核發 token)
    │
    └─ GET  /oauth/authorize (authorization_code grant)
            ↓
        hydra /oauth2/auth
            ↓ redirect to login
        glintfed /oauth/login  (驗帳密)
            ↓ accept login via hydra Admin API
        hydra /oauth2/auth  (繼續流程)
            ↓ redirect to consent
        glintfed /oauth/consent (顯示授權同意)
            ↓ accept consent via hydra Admin API
        hydra → redirect to client with code
            ↓
        Client App POST /oauth/token (authorization_code)
            ↓
        hydra /oauth2/token → access_token
```

---

## 基礎設施變更

### docker-compose.yml 新增

```yaml
services:
  hydra-migrate:
    image: oryd/hydra:v2
    command: migrate sql --yes $DSN
    environment:
      DSN: postgres://hydra:secret@postgres:5432/hydra?sslmode=disable
    depends_on:
      - postgres

  hydra:
    image: oryd/hydra:v2
    command: serve all --dev
    ports:
      - "4444:4444"  # Public API (OAuth2 endpoints)
      - "4445:4445"  # Admin API
    environment:
      DSN: postgres://hydra:secret@postgres:5432/hydra?sslmode=disable
      URLS_SELF_ISSUER: https://example.com
      URLS_CONSENT: https://example.com/oauth/consent
      URLS_LOGIN: https://example.com/oauth/login
      URLS_LOGOUT: https://example.com/oauth/logout
      SECRETS_SYSTEM: youReallyNeedToChangeThis
      OAUTH2_EXPOSE_INTERNAL_ERRORS: "true"
    depends_on:
      - hydra-migrate
      - postgres

  postgres:
    image: postgres:16-alpine
    environment:
      POSTGRES_USER: hydra
      POSTGRES_PASSWORD: secret
      POSTGRES_DB: hydra
```

> **注意**：hydra 不支援 SQLite，必須使用 PostgreSQL 或 MySQL。glintfed 本身的業務資料庫（SQLite）與 hydra 的資料庫是分開的兩套儲存。

---

## 新增 Go 依賴

```bash
go get github.com/ory/hydra-client-go/v2
```

---

## Config 變更

### `internal/data/config.go`

在 `AuthConfig` 中新增 hydra 相關設定：

```go
type AuthConfig struct {
    EnableRegistration bool        `mapstructure:"enable_registration" env:"OPEN_REGISTRATION"`
    EnableOAuth        bool        `mapstructure:"enable_oauth" env:"OAUTH_ENABLED"`
    InAppRegistration  bool        `mapstructure:"in_app_registration" env:"APP_REGISTER"`
    Hydra              HydraConfig `mapstructure:"hydra"`
}

type HydraConfig struct {
    AdminURL  string `mapstructure:"admin_url"  env:"HYDRA_ADMIN_URL"`
    PublicURL string `mapstructure:"public_url" env:"HYDRA_PUBLIC_URL"`
}
```

---

## 實作步驟

### Step 1：Hydra Client Wrapper（`internal/lib/hydra/`）

建立一個薄薄的 wrapper 封裝 hydra Admin API 呼叫，供 usecase 層使用。

**`internal/lib/hydra/client.go`**

```go
package hydra

import (
    hydraclient "github.com/ory/hydra-client-go/v2"
)

type Client struct {
    admin *hydraclient.APIClient
}

func NewClient(adminURL string) *Client {
    cfg := hydraclient.NewConfiguration()
    cfg.Servers = hydraclient.ServerConfigurations{
        {URL: adminURL},
    }
    return &Client{admin: hydraclient.NewAPIClient(cfg)}
}
```

**`internal/lib/hydra/token.go`**

```go
package hydra

import (
    "context"
    "time"

    hydraclient "github.com/ory/hydra-client-go/v2"
)

type TokenResult struct {
    AccessToken  string
    RefreshToken string
    ExpiresIn    int64
    ClientID     string
    ClientSecret string
}

// CreatePersonalAccessToken 直接透過 Admin API 為指定 subject 核發 token，
// 對應 pixelfed AppRegister onboarding 的「直接取得 token」場景。
func (c *Client) CreatePersonalAccessToken(ctx context.Context, subject string, scopes []string, clientID string) (*TokenResult, error) {
    expiry := time.Now().Add(365 * 24 * time.Hour)
    req := c.admin.OAuth2API.CreateOAuth2Token(ctx)
    // hydra v2 Admin API: POST /admin/oauth2/token (implicit grant via admin)
    // 實際上需要透過 AcceptOAuth2ConsentRequest + 特製 client credentials flow
    // 詳見 Step 1 注意事項
    _ = req
    _ = expiry
    panic("see implementation notes below")
}
```

> **Step 1 重要注意事項**：hydra v2 的 Admin API 並沒有直接「為任意 subject 核發 access token + refresh token」的端點（這與 hydra v1 不同）。
> 替代方案有兩種：
> **方案 A**：為每個使用者建立一個 hydra OAuth2 client，然後用 `client_credentials` grant 核發 token，但這樣 token 的 `sub` 會是 client ID 而非 user ID，語意不符。
> **方案 B**：實作一個內部 Authorization Code flow shortcut，由 glintfed 代替使用者完成 login accept + consent accept，取得 authorization code 後再換 token。這是唯一語意正確的方式，但流程複雜。

---

### Step 2：OAuthUsecase 實作（`internal/usecase/oauth/oauth.go`）

```go
package oauth

import (
    "context"
    "fmt"

    "glintfed.org/internal/lib/hydra"
)

type Usecase struct {
    hydra    *hydra.Client
    clientID string
    // 系統預設 OAuth client（對應 pixelfed 的 personal_access_client）
}

func NewUsecase(hydraClient *hydra.Client, clientID string) *Usecase {
    return &Usecase{hydra: hydraClient, clientID: clientID}
}

// CreateTokens 實作 AppRegister onboarding 的直接核發 token 場景。
// 對應 pixelfed 的 $user->createToken('Pixelfed App', [...scopes])。
func (uc *Usecase) CreateTokens(ctx context.Context, userID uint64, scopes []string) (*TokenResult, error) {
    subject := fmt.Sprintf("%d", userID)
    result, err := uc.hydra.CreatePersonalAccessToken(ctx, subject, scopes, uc.clientID)
    if err != nil {
        return nil, err
    }
    return &TokenResult{
        AccessToken:  result.AccessToken,
        RefreshToken: result.RefreshToken,
        ClientID:     result.ClientID,
        ClientSecret: result.ClientSecret,
        ExpiresIn:    result.ExpiresIn,
    }, nil
}
```

---

### Step 3：OAuth Service（`internal/service/oauth/`）

這個 service 負責：
- `GET  /oauth/login`：Login UI（Login Provider）
- `POST /oauth/login`：接受帳密，呼叫 hydra Admin API `AcceptOAuth2LoginRequest`
- `GET  /oauth/consent`：Consent UI
- `POST /oauth/consent`：呼叫 hydra Admin API `AcceptOAuth2ConsentRequest`
- `GET  /oauth/logout`：呼叫 hydra Admin API `AcceptOAuth2LogoutRequest`
- `POST /oauth/token`：Mastodon-compatible token endpoint，處理 `password` grant（hydra 不支援）

**`internal/service/oauth/service.go`**

```go
package oauth

import (
    "context"
    "net/http"

    "glintfed.org/internal/lib/hydra"
)

type Service interface {
    // Login Provider endpoints (hydra redirects here)
    LoginPage(w http.ResponseWriter, r *http.Request)
    LoginSubmit(w http.ResponseWriter, r *http.Request)
    ConsentPage(w http.ResponseWriter, r *http.Request)
    ConsentSubmit(w http.ResponseWriter, r *http.Request)
    LogoutPage(w http.ResponseWriter, r *http.Request)

    // Mastodon-compatible endpoints
    Authorize(w http.ResponseWriter, r *http.Request)  // GET  /oauth/authorize → redirect to hydra
    Token(w http.ResponseWriter, r *http.Request)      // POST /oauth/token
    Revoke(w http.ResponseWriter, r *http.Request)     // POST /oauth/revoke
}

//go:generate go tool moq -rm -out mock_user_authenticator.go . UserAuthenticator
type UserAuthenticator interface {
    // Authenticate 驗證帳密，成功回傳 userID
    Authenticate(ctx context.Context, username, password string) (uint64, error)
}

type svc struct {
    hydra     *hydra.Client
    auth      UserAuthenticator
    publicURL string // hydra public URL，用於 Authorize redirect
}
```

**`internal/service/oauth/token.go`（處理 password grant）**

```go
// Token 處理 POST /oauth/token
// 對 grant_type=password 的請求，glintfed 自行驗帳密後透過 hydra Admin API 核發 token。
// 對 grant_type=authorization_code 的請求，直接 proxy 到 hydra。
func (s *svc) Token(w http.ResponseWriter, r *http.Request) {
    ctx, span := internal.T.Start(r.Context(), "OAuth.Token")
    defer span.End()

    grantType := r.FormValue("grant_type")
    switch grantType {
    case "password":
        s.handlePasswordGrant(ctx, w, r)
    case "authorization_code", "refresh_token", "client_credentials":
        // proxy 到 hydra public token endpoint
        s.proxyToHydra(w, r)
    default:
        http.Error(w, `{"error":"unsupported_grant_type"}`, http.StatusBadRequest)
    }
}
```

---

### Step 4：OOB（Out-of-Band）支援

Mastodon 的 OOB flow 使用 `redirect_uri=urn:ietf:wg:oauth:2.0:oob`，hydra 不認識此 URI。

處理方式：在 `GET /oauth/authorize` handler 中，若偵測到 OOB redirect URI，改用 `http://glintfed.example.com/oauth/oob-callback` 作為實際 redirect URI 傳給 hydra，然後在該 callback 頁面顯示 authorization code 給使用者複製（OOB 的原意）。

```go
func (s *svc) Authorize(w http.ResponseWriter, r *http.Request) {
    redirectURI := r.URL.Query().Get("redirect_uri")
    if redirectURI == "urn:ietf:wg:oauth:2.0:oob" {
        // 替換為 glintfed 自己的 OOB callback URI
        q := r.URL.Query()
        q.Set("redirect_uri", s.cfg.App.Url+"/oauth/oob")
        r.URL.RawQuery = q.Encode()
    }
    // redirect to hydra /oauth2/auth
    http.Redirect(w, r, s.publicURL+"/oauth2/auth?"+r.URL.RawQuery, http.StatusFound)
}

func (s *svc) OOBCallback(w http.ResponseWriter, r *http.Request) {
    code := r.URL.Query().Get("code")
    // render 一個顯示 code 的靜態頁面供使用者複製
    w.Header().Set("Content-Type", "text/html")
    fmt.Fprintf(w, `<pre>%s</pre>`, code)
}
```

---

### Step 5：Token 驗證 Middleware

每個需要 `auth:api` 的請求，需要驗證 Bearer token。

**方案 A（呼叫 hydra introspect，較慢）**：

```go
func HydraTokenAuth(hydraClient *hydra.Client) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            token := extractBearerToken(r)
            introspect, err := hydraClient.IntrospectToken(r.Context(), token)
            if err != nil || !introspect.Active {
                http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
                return
            }
            // 將 subject、scopes 注入 context
            ctx := context.WithValue(r.Context(), ctxKeySubject, introspect.Sub)
            next.ServeHTTP(w, r.WithContext(ctx))
        })
    }
}
```

**方案 B（本地驗證 JWT，較快）**：
設定 hydra 以 JWT 格式核發 access token（`strategies.access_token: jwt`），然後在 glintfed 從 hydra 的 JWKS endpoint 取公鑰做本地驗證，避免每次 introspect 呼叫。

---

### Step 6：Route 註冊（`internal/server/api.go`）

```go
// OAuth Routes（Mastodon-compatible）
mux.Get("/oauth/authorize", svcs.OAuth.Authorize)
mux.Post("/oauth/token", svcs.OAuth.Token)
mux.Post("/oauth/revoke", svcs.OAuth.Revoke)

// Hydra Login/Consent Provider
mux.Get("/oauth/login", svcs.OAuth.LoginPage)
mux.Post("/oauth/login", svcs.OAuth.LoginSubmit)
mux.Get("/oauth/consent", svcs.OAuth.ConsentPage)
mux.Post("/oauth/consent", svcs.OAuth.ConsentSubmit)
mux.Get("/oauth/logout", svcs.OAuth.LogoutPage)
mux.Get("/oauth/oob", svcs.OAuth.OOBCallback)
```

---

### Step 7：DI 註冊（`cmd/api/kessoku.go`）

```go
// 新增 hydra client 到 DI
kessoku.Provide(func(cfg *data.Config) *hydra.Client {
    return hydra.NewClient(cfg.App.Auth.Hydra.AdminURL)
}),

// OAuth service
kessoku.Bind[oauth.Service](kessoku.Provide(oauth.New)),

// OAuth usecase（取代現有的 oauth.NewUsecase）
kessoku.Bind[appregister.OAuthUsecase](kessoku.Provide(oauth.NewUsecase)),
```

---

## 測試策略

- **Usecase 層**：mock `hydra.Client`（在 wrapper 上定義 interface）
- **Service 層**：mock `UserAuthenticator`；hydra Admin API 互動用 `httptest.Server` 模擬
- **整合測試**：需要啟動真實 hydra container（使用 testcontainers-go）

---

## 已知限制與風險

| 項目 | 說明 |
|------|------|
| `grant_type=password` 支援 | hydra 不支援，需在 glintfed 層自行處理，邏輯繞過 hydra 的 consent flow |
| OOB redirect URI | 需 glintfed 攔截並轉換，額外的 callback 頁面 |
| 直接核發 token（AppRegister） | hydra v2 無直接 Admin API，需繞道實作（見 Step 1 注意事項），是本方案最大的技術障礙 |
| 儲存分離 | 業務資料在 SQLite，OAuth 資料在 PostgreSQL，運維複雜度提升 |
| 網路延遲 | 每次 token introspect 增加 RTT；JWT 模式可緩解但需維護 JWKS 快取 |
| Login/Consent UI | 需額外實作 HTML 頁面，pixelfed 原本有完整的 Laravel Blade UI |
| 部署複雜度 | 新增 hydra + PostgreSQL 兩個服務，開發環境 docker-compose 更複雜 |
