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

## Service Handler 實作規範

每個 Handler 必須包含 OpenTelemetry 追蹤與 Todo 註釋。

### 標準範本
使用 `internal.T` (otel.Tracer) 進行追蹤，並維持與 Laravel 方法名一致的 Go 函數名。

```go
package v1

import (
	"net/http"
	"glintfed.org/internal/service/internal"
)

//對應 Laravel 的 UserController@show
func UserShow(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "ApiV1.UserShow")
	defer span.End()

	// TODO: Implement UserShow
}
```

## 轉換清單 (Cheat Sheet)

| Laravel | Go (chi/internal) |
| :--- | :--- |
| `Route::prefix('...')` | `r.Route("...", func(r chi.Router) { ... })` |
| `{id}` | `{id}` (透過 `chi.URLParam(r, "id")` 取得) |
| `middleware('auth')` | `r.Use(middleware.Auth)` |
| `Controller@method` | `pkg.Method` |
| `Log::info()` | `logs.Info()` |