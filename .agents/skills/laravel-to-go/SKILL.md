---
name: laravel-to-go
description: 將 Laravel 應用程式遷移到 Go
---

# Laravel 遷移指南

本技能旨在將 Laravel 應用程式遷移到 Go 應用程式。

## API 路由

Laravel 的 API 路由位於 `routes/api.php` 中

### HTTP Methods 路由

根據 `Route::get` 或 `Route::post` 等代表的是不同的 HTTP 方法：

在遷移時應該在 go-chi 中定義相應的路由：

```go
// Route::get('/hello-world', 'GreetController@hello');
// Route::post('/login', 'AuthController@login');

mux.Get("/hello-world", greet.Hello())
mux.Post("/login", auth.Login())
```

### Redirect 路由

`Route::redirect` 代表的是使用 HTTP 302 Found 將用戶重定向到指定的路由

```go
// Rotue::redirect("here", "there")

mux.Get("/here", http.RedirectHandler("/there", http.StatusFound))
```

### 多重匹配

`Route::match` 可以指定多個不同的 HTTP Method 對應相同的路由與 Handler。

```go
// Route::match(['PUT', 'PATCH'], 'account', 'AccountController@update')

mux.Put("/account", account.Update())
mux.Patch("/account", account.Update())
```

## Controller

在 Laravel 中會在 `app/Http/Controllers` 中定義控制器，用於處理 HTTP 路由的請求（等價於 Go 中的 Handler）。

```go
// internal/service/account/service.go
type Service interface {
  Update(w http.ResponseWriter, r *http.Request)
}

type service struct {
  clients *data.Clients
}

func NewService(clients *data.Clients) Service {
  return &service{clients: clients}
}

func (s *service) Update(w http.ResponseWriter, r *http.Request) {
  ctx, span := otel.Tracer("service").Start(r.Context(), "account.Update")
  defer span.End()

  if err := s.DB.UpdateUser(ctx, ...); err != nil {
  }
}
```

### 巢狀結構

部份 Controller 可能位於 `app/Http/Controllers` 底下的其它 namespaces，這時需要維持相同的結構。

```php
// app/Http/Controllers/Api/AdminController.php
class adminApiController extends Controller {
  public function supported(Request $request) {
    // ...
  }
}
```

```go
// internal/service/api/admin/service.go
func (s *service) Supported(w http.ResponseWriter, r *http.Request) {
  ctx, span := otel.Tracer("service").Start(r.Context(), "api.admin.Supported")
  defer span.End()

  // ...
}
```