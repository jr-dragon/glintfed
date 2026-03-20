---
name: laravel-to-go
description: 將 Laravel 應用程式遷移到 Go
---

# Laravel 遷移指南

本技能旨在將 Laravel 應用程式遷移到 Go 應用程式。

## 目錄與檔案規範

- **目錄結構**：Go 的 package 結構應對應 Laravel 的 Namespace。
    - 例如：`App\Http\Controllers\Api\V1` 應遷移至 `internal/service/api/v1/`。
- **檔案命名**：主要的 Handler 實作檔案應統一命名為 `service.go`。
- **Package 命名**：使用簡短且具代表性的 package 名稱（例如 `v1`, `federation`, `healthcheck`）。

## API 路由映射

使用 `go-chi/chi` 作為路由框架。

### 1. 基礎 HTTP 方法
- `Route::get` -> `mux.Get`
- `Route::post` -> `mux.Post`
- `Route::put` -> `mux.Put`
- `Route::delete` -> `mux.Delete`

### 2. 路由群組與前綴 (Groups & Prefixes)
Laravel 的 `Route::group` 或 `Route::prefix` 應映射至 chi 的 `r.Route` 或 `r.Group`。

```php
// Laravel
Route::prefix('api/v1')->group(function () {
    Route::get('status', 'StatusController@show');
});
```

```go
// Go (chi)
mux.Route("/api/v1", func(r chi.Router) {
    r.Get("/status", v1.StatusShow)
})
```

### 3. 路徑參數 (Route Parameters)
Laravel 使用 `{param}`，chi 同樣使用 `{param}`。

```php
// Route::get('users/{id}', 'UserController@show')
mux.Get("/users/{id}", user.Show)
```

### 4. 重定向 (Redirects)
```php
// Route::redirect(".well-known/change-password", "/settings/password");
mux.Get("/.well-known/change-password", func(w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r, "/settings/password", http.StatusFound)
})
```

## Service Layer 實作規範

每個 Service 應定義為 Interface，並透過建構子注入依賴。

### 標準範本 (Kessoku DI)
使用 `internal.T` (otel.Tracer) 進行追蹤，並維持與 Laravel 方法名一致的 Go 函數名。

```go
package v1

import (
	"net/http"
	"glintfed.org/internal/data"
	"glintfed.org/internal/service/internal"
)

type Service interface {
	UserShow(w http.ResponseWriter, r *http.Request)
}

func New(cfg *data.Config) Service {
	return &svc{cfg: cfg}
}

type svc struct {
	cfg *data.Config
}

// 對應 Laravel 的 UserController@show
func (s *svc) UserShow(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "ApiV1.UserShow")
	defer span.End()

	// TODO: Implement UserShow
}
```

**注意**：修改 `New` 建構子參數後，必須在專案根目錄執行 `make gen` 以更新 Kessoku DI 生成代碼。

## Model 網域與 URL 邏輯遷移

Laravel Model 中與 URL 生成或 ActivityPub JSON 相關的邏輯（例如 `getActor()`, `permalink()`）應遷移至 `ent/` 套件下的獨立檔案。

### 規範與命名
- **檔案命名**：`<model>_url.go` (例如 `instanceactor_url.go`)。
- **實作方式**：定義為 `ent` model pointer 的 method。
- **配置傳遞**：由於 `ent` 套件無法直接存取全域配置，這類 method 應明確接收需要的參數 (如 `appURL`, `appDomain`)。

```go
// ent/instanceactor_url.go
func (ia *InstanceActor) GetActor(appURL, appDomain string) map[string]any {
    return map[string]any{
        "id": ia.Permalink(appURL),
        // ...
    }
}
```

## 測試規範 (Testing Strategy)

每個遷移的組件都必須包含對應的測試。

| 組件 | 測試檔案 | 測試重點 | 工具/方法 |
| :--- | :--- | :--- | :--- |
| **Service Handler** | `service_test.go` | HTTP 回應碼、Content-Type、JSON 欄位 | `httptest`, `moq` |
| **Model DB 邏輯** | `model_test.go` | CRUD 操作、Query 邏輯 | `data.NewTestClient` (SQLite In-memory) |
| **Model URL 邏輯** | `<model>_url_test.go` | URL 拼接、ActivityPub 結構 | 格式化斷言 |

## 轉換清單 (Cheat Sheet)

| Laravel | Go (chi/internal) |
| :--- | :--- |
| `Route::prefix('...')` | `r.Route("...", func(r chi.Router) { ... })` |
| `{id}` | `{id}` (透過 `chi.URLParam(r, "id")` 取得) |
| `middleware('auth')` | `r.Use(middleware.Auth)` |
| `Controller@method` | `pkg.Method` |
| `Log::info()` | `logs.Info()` |
| `config('app.url')` | 透過建構子注入 `*data.Config` |

## Eloquent Model 遷移至 ent

將 Laravel 的 Eloquent Model 遷移至 Go 的 `ent` schema 時，應遵循以下模式與規範。

### 0. 快速生成 Schema

可以使用專案提供的腳本，根據 Laravel 的 Model 自動批量生成空白的 `ent` schema 檔案。

```bash
# 在專案根目錄執行
./.agents/skills/laravel-to-go/scripts/gen-ent-schema.sh <LARAVEL_APP_ROOT>
```

### 0.1 透過資料庫取得屬性映射資訊

可以使用以下腳本來取得資料表的欄位定義，協助進行屬性映射。

```bash
./.agents/skills/laravel-to-go/scripts/get-columns-from-mysql-cli.sh <LARAVEL_APP_ROOT> [specific_table]
```

### 1. 型別映射 (Field Types)

| Laravel (Migration) | ent (Go Type) | 範例 |
| :--- | :--- | :--- |
| `id()`, `bigIncrements('id')` | `field.Uint64("id")` | `field.Uint64("id").Unique()` |
| `unsignedBigInteger('xxx_id')` | `field.Uint64("xxx_id")` | `field.Uint64("profile_id").Optional()` |
| `string('...')` | `field.String("...")` | `field.String("username").Unique()` |
| `text('...')` | `field.Text("...")` | `field.Text("bio").Optional()` |
| `boolean('...')` | `field.Bool("...")` | `field.Bool("is_private").Default(false)` |
| `timestamp()`, `datetime()` | `field.Time("...")` | `field.Time("last_active_at").Optional()` |
| `json('...')` | `field.JSON("...", T{})` | `field.JSON("media_ids", []uint64{})` |
| `enum('...', [...])` | `field.Enum("...").Values(...)` | `field.Enum("visibility").Values("public", "private")` |

### 2. Eloquent 屬性映射

- **`$fillable`**: 所有出現在 `$fillable` 的欄位都應定義在 `Fields()` 中。
- **`$hidden`**: 敏感資訊加上 `.Sensitive()`。
- **軟刪除 (`SoftDeletes`)**: 包含 `field.Time("deleted_at").Optional()`。

### 3. 命名規範與特殊處理

- **CamelCase**: Go 的欄位名稱使用 CamelCase，`ent` 會自動轉換。
- **數字開頭欄位**: 使用 `StorageKey`。
  ```go
  field.Bool("two_fa_enabled").StorageKey("2fa_enabled")
  ```
- **資料表名稱**: 使用 `Annotations` 指定。

### 4. 標準元數據欄位 (Standard Metadata)

```go
func (X) Fields() []ent.Field {
    return []ent.Field{
        field.Uint64("id").Unique(),
        field.Time("created_at").Default(time.Now).Immutable(),
        field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
        field.Time("deleted_at").Optional(),
    }
}
```
