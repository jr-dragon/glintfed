# OAuth2 實作方案綜合評估

## 快速總覽

| 評估維度 | 方案一：ory/hydra | 方案二：ory/fosite |
|---------|-----------------|-----------------|
| 架構模式 | 外掛式獨立服務 | 嵌入式 library |
| `grant_type=password` | ❌ 需繞道實作 | ✅ 原生支援 |
| OOB redirect URI | ⚠️ 需攔截轉換 | ⚠️ 需客製化 URI 白名單 |
| AppRegister 直接核發 token | ❌ hydra v2 無直接 Admin API，繞道困難 | ✅ 直接操作 store |
| 部署複雜度 | 高（+hydra +PostgreSQL） | 低（無新服務） |
| 運維複雜度 | 高（兩套 DB、多個服務） | 低（單一程序） |
| Token 驗證速度 | 中（JWT 本地驗 or introspect RTT） | 快（同程序） |
| 初始工程量 | 中（Login/Consent UI + proxy） | 高（實作 fosite Storage interfaces） |
| Mastodon API 相容性 | 低（password grant 需自建） | 高（全部 grant 皆支援） |
| 測試難度 | 高（需模擬 hydra Admin API） | 低（in-memory store 即可） |
| 標準合規性 | 高（OAuth 2.1，無 ROPC） | 中（支援舊標準 ROPC） |
| 未來擴展（OIDC 等） | 易（hydra 原生支援） | 需自行在 fosite 上加 OIDC layer |

---

## 詳細分析

### 1. Mastodon API 相容性（關鍵指標）

Mastodon-compatible API 的 OAuth2 使用場景主要有三種：

**場景 A：Web/Desktop App（Authorization Code Grant）**
- 兩方案均可支援，流程標準

**場景 B：Mobile App（Password Grant，`grant_type=password`）**
- pixelfed 官方 app、大量第三方 Mastodon client 仍使用此 flow
- **hydra**：不支援，需在 glintfed 層完全自行實作；token 由 hydra Admin API 核發，但語意上繞過了 hydra 的 consent flow，等於有一半的流程在 hydra 外
- **fosite**：`compose.OAuth2ResourceOwnerPasswordFactory` 原生支援，注入自訂帳密驗證器即可

**場景 C：AppRegister Onboarding（直接核發）**
- 對應 `POST /api/auth/onboarding`，使用者完成 email 驗證後直接取得 token
- **hydra**：v2 Admin API 無法直接替 subject 核發帶 refresh token 的 access token（這是 hydra 刻意的設計決策），是本方案**最根本的技術障礙**
- **fosite**：直接操作 `fositestore.Store.CreateAccessTokenSession` + `CreateRefreshTokenSession` 即可，與 Passport 的 `$user->createToken()` 完全等價

### 2. 工程量評估

**hydra 方案**的主要工作：
- 實作 Login Provider（login page、submit handler）
- 實作 Consent Provider（consent page、submit handler）
- 實作 `/oauth/token` password grant 繞道邏輯
- 解決 AppRegister 直接核發 token 的技術問題（目前無標準解法）
- 實作 OOB redirect URI 攔截

**fosite 方案**的主要工作：
- 新增 4 個 ent schema（authorization_codes、access_tokens、refresh_tokens、pkce）+ `make gen`
- 實作 `fositestore.Store`（約 15–20 個 method，每個約 10–30 行）
- 自訂 `PasswordGrantHandler` 注入帳密驗證邏輯
- 處理 OOB redirect URI 白名單

fosite 的工程量較大但**都是機械性的 CRUD 實作**，且有明確的 interface 規格。hydra 的工程量看似較小，但 AppRegister token 核發問題**沒有乾淨的解法**，可能需要大量 hack。

### 3. 部署與運維

**hydra 方案**：
- docker-compose 新增 hydra + PostgreSQL 兩個服務
- 需維護 hydra 版本、PostgreSQL 備份、hydra config
- 開發環境 setup 更複雜，新進開發者的入門成本提高
- hydra 與 glintfed 之間的網路故障會導致 OAuth 功能完全失效

**fosite 方案**：
- 無新增外部服務
- OAuth 資料與業務資料共用同一 SQLite/PostgreSQL，統一備份
- 單一程序，故障域更小

### 4. 標準合規性

hydra 方案不支援 ROPC，從 OAuth 2.1 標準的角度更正確。然而 glintfed 定位是 Mastodon/pixelfed 相容服務，目標客戶（第三方 app 開發者）的使用習慣仍有大量 ROPC 依賴，放棄相容性的代價超過標準合規的好處。

### 5. 未來演進

若未來需要 OpenID Connect（OIDC）：
- hydra 原生支援 OIDC，可直接啟用
- fosite 的 `compose` package 也有 `OpenIDConnectExplicitFactory`，可追加，但需要額外實作 `openid.OpenIDConnectRequestStorage`

---

## 建議

**建議選擇方案二（fosite）**，原因如下：

1. **AppRegister 直接核發 token 的問題在 hydra 方案中沒有標準解法**，這是 `feat/oauth-support` branch 目前最核心的需求（`OAuthUsecase.CreateTokens`），也是整個 OAuth 功能的起點。

2. **password grant 的 Mastodon 相容性**是 glintfed 作為 pixelfed 替代品的基本要求，hydra 方案需要完整的額外實作，實際上並沒有真正「使用」到 hydra 的核心功能。

3. **部署簡單性**對於開源社群部署（self-hosted）非常重要，pixelfed 本身已因部署複雜度被詬病，glintfed 應優先降低門檻。

4. fosite 的工程量雖大，但**每一步都有明確的規格和對應的測試方式**，風險可控。

---

## 實作優先順序（fosite 方案）

```
Phase 1：基礎 token 核發（unblock AppRegister onboarding）
  1. 新增 ent schema（access_tokens、refresh_tokens）
  2. 實作 fositestore.Store 的 ClientManager + AccessTokenStorage + RefreshTokenStorage
  3. 實作 OAuthUsecase.CreateTokens（直接操作 store）
  4. 驗證 POST /api/auth/onboarding 端對端可運作

Phase 2：完整 OAuth2 flows
  5. 新增 ent schema（authorization_codes、pkce）
  6. 實作剩餘 fositestore interfaces
  7. 實作 OAuth Service（/oauth/authorize、/oauth/token、/oauth/revoke）
  8. 自訂 PasswordGrantHandler（password grant）
  9. OOB redirect URI 支援

Phase 3：Token 驗證 middleware
  10. 實作 OAuth2Auth middleware，保護需要認證的 API endpoints
  11. 將 auth:api middleware 注入 chi router
```
