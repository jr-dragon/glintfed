# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

GlintFed is a PixelFed-compatible photo sharing service written in Go — a ground-up reimplementation of the PHP/Laravel PixelFed project, maintaining full API compatibility. When porting APIs, refer to the `pixelfed/` directory (PHP source) as a reference. **Never modify files under `pixelfed/`.**

## Commands

```bash
make init    # Install dev tools (goimports, ent, moq, kessoku)
make gen     # Run go generate ./... (regenerates ent ORM + kessoku DI code)
make lint    # Format with gofmt + goimports
make test    # Run all tests
make build   # Build binaries from cmd/ to bin/
make all     # gen → lint → test → build
```

Run a single test: `go test ./internal/service/story/... -run TestFoo`

After modifying any service or model, run `make gen` to regenerate DI and mock code.

## Architecture

The project uses a **Service → Model → Ent** layered architecture with compile-time dependency injection via [kessoku](https://github.com/mazrean/kessoku).

```
HTTP Request
    ↓
internal/server/api.go       # chi router, all route registrations
    ↓
internal/service/{module}/   # HTTP handlers, OTel tracing, interface definitions
    ↓
internal/usecase/{module}/   # Complex business logic (e.g. OAuth token issuance)
    ↓
internal/model/{module}/     # DB operations implementing service interfaces
    ↓
ent/                         # Generated ORM code (do not edit manually)
    ↓
SQLite / Redis
```

Use `internal/usecase/` when business logic is too complex for the service layer and does not belong directly to DB operations (e.g. OAuth token creation, external API calls). Use `internal/lib/errs.Todo` to mark stub implementations that are not yet complete.

**Key files:**
- `cmd/api/kessoku.go` — DI wiring; register new services/models here
- `internal/data/client.go` — `*data.Client` holds `Ent`, `DB`, and `RDB` (Redis)
- `internal/data/config.go` — config struct; loaded from `config.yaml`
- `internal/server/api.go` — all route registrations (v1, v1.1, v1.2, v2, admin, federation)

## API Development Flow

### 1. Service Layer (`internal/service/{module}/service.go`)
- Define a `Service` interface and dependency interfaces (e.g., `Getter`, `Storer`)
- Add `//go:generate go tool moq -rm -out mock_{name}.go . {InterfaceName}` for each dependency interface
- Implement handlers on `svc` struct; add OTel tracing in every handler:

```go
func (s *svc) MyHandler(w http.ResponseWriter, r *http.Request) {
    ctx, span := internal.T.Start(r.Context(), "ServiceName.MyHandler")
    defer span.End()
    // ...
}
```

### 2. Model Layer (`internal/model/{module}/model.go`)
- Embed the specific `*ent.{Entity}Client` (not `*data.Client`) in the `Model` struct; `NewModel` accepts `*data.Client` and extracts the client:

```go
type Model struct {
    *ent.AppRegisterClient
}

func NewModel(client *data.Client) *Model {
    return &Model{AppRegisterClient: client.Ent.AppRegister}
}
```

- Add SQL Godoc comments on every method:

```go
// GetByID
//
//	SELECT * FROM stories WHERE id = ? LIMIT 1
func (m *Model) GetByID(ctx context.Context, id uint64) (*ent.Story, error) { ... }
```

### 3. DI Registration (`cmd/api/kessoku.go`)
- Register new constructors with `kessoku.Provide`
- Bind interface implementations with `kessoku.Bind[Interface](kessoku.Provide(NewModel))`
- Run `make gen` after changes

### 4. Route Registration (`internal/server/api.go`)
- Add routes under the appropriate version group (v1, v1.1, v2, etc.)

## Testing Strategy

- **Service tests**: Use `moq`-generated mocks to isolate model dependencies
- **Model tests**: Use `data.NewTestClient(t)` for in-memory SQLite

## Key Conventions

- **Logging**: Use `log/slog`; wrap errors with `logs.ErrAttr(err)`
- **Error attributes**: `slog.Error("msg", logs.ErrAttr(err))`
- **Config**: Default `config.yaml`; see `config.example.yaml` for structure
- **Style**: Follow [Google Go Style Guide](https://google.github.io/styleguide/go/)
- **Federation reference**: `pixelfed/app/Http/Controllers/Api/` for PHP originals
- **Request validation**: Use `github.com/go-playground/validator/v10` struct tags for static constraints; dynamic constraints sourced from `cfg` are checked manually after `validate.Struct()`. Register custom validators (e.g. `username`) once in `New()` via `v.RegisterValidation(...)`.
- **Stub implementations**: Return `errs.Todo` (`internal/lib/errs`) for functions that are not yet implemented, and add a `// TODO:` comment explaining what needs to be done.
