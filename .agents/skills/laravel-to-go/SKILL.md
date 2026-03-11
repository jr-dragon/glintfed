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

## Eloquent Model 遷移至 ent

將 Laravel 的 Eloquent Model 遷移至 Go 的 `ent` schema 時，應遵循以下模式與規範。

### 0. 快速生成 Schema

可以使用專案提供的腳本，根據 Laravel 的 Model 自動批量生成空白的 `ent` schema 檔案。

```bash
# 在專案根目錄執行
./.agents/skills/laravel-to-go/scripts/gen-ent-schema.sh <LARAVEL_APP_ROOT>
```

該腳本會掃描 Laravel 中的 `app/*.php` 與 `app/Models/*.php`，找出所有繼承自 `Model` 或 `Authenticatable` 的類別，並執行 `ent new`。

### 0.1 透過資料庫取得屬性映射資訊

若已經有既有的資料庫（如本地端的 MySQL 或開發環境中的 Docker 容器），可以使用以下腳本來取得資料表的欄位定義，協助進行屬性映射。

**使用本地 MySQL CLI：**
```bash
./.agents/skills/laravel-to-go/scripts/get-columns-from-mysql-cli.sh <LARAVEL_APP_ROOT> [specific_table]
```

**使用 Docker 容器中的 MySQL：**
```bash
./.agents/skills/laravel-to-go/scripts/get-columns-from-mysql-container.sh <LARAVEL_APP_ROOT> [specific_table]
```

這些腳本會列出每個資料表的 `DESCRIBE` 結果，幫助你快速確認每個欄位的型別、長度及是否可為空（NULL）。

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
- **`$hidden`**: 敏感資訊（如 `password`, `email`, `2fa_secret`）應加上 `.Sensitive()`。
- **`$casts`**: 
    - `datetime`, `timestamp` -> `field.Time`
    - `array`, `json` -> `field.JSON`
- **軟刪除 (`SoftDeletes`)**: 必須包含 `field.Time("deleted_at").Optional()`。

### 3. 命名規範與特殊處理

- **CamelCase**: Go 的欄位名稱使用 CamelCase，`ent` 會自動將其轉換為資料庫的 snake_case。
- **數字開頭欄位**: Go 的欄位名稱不能以數字開頭。若 Laravel 欄位為 `2fa_enabled`，應定義為 `TwoFaEnabled` 並使用 `StorageKey`。
  ```go
  field.Bool("two_fa_enabled").StorageKey("2fa_enabled").Default(false)
  ```
- **資料表名稱**: 使用 `Annotations` 明確指定與 Laravel 一致的複數表名。
  ```go
  func (User) Annotations() []schema.Annotation {
      return []schema.Annotation{
          entsql.Annotation{Table: "users"},
      }
  }
  ```

### 4. 關聯處理 (Relationships)

目前專案傾向使用 **基礎欄位關聯** (Field-based Relations) 而非 `ent.Edge`，以簡化初期遷移過程。

- **`belongsTo`**: 在 schema 中新增一個 `Uint64` 欄位並加上 `_id` 後綴。
  - Laravel: `Status -> belongsTo(Profile)`
  - Go: `field.Uint64("profile_id").Optional()`
- **`hasOne` / `hasMany`**: 通常不直接在 schema 中定義，而是透過查詢對方 table 的 `xxx_id` 來實現。

### 5. 標準元數據欄位 (Standard Metadata)

每個 Schema 通常都應包含以下標準欄位：

```go
func (X) Fields() []ent.Field {
    return []ent.Field{
        field.Uint64("id").Unique(),
        // ... 其他欄位
        field.Time("created_at").Default(time.Now).Immutable(),
        field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
        field.Time("deleted_at").Optional(),
    }
}
```