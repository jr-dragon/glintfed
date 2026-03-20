# GlintFed Agent 指南

本文件為 AI Agent 提供專案的技術背景、開發規範及操作指引。

## 專案概述

GlintFed 是基於 PixelFed 的專案，將原本由 PHP/Laravel 的實作改為 Go 的實作。

預期 GlintFed 的實作會完全相容於 PixelFed 所提供的 API，為此 Agents 可以參考 `pixelfed` 資料夾下的程式碼，但 `pixelfed` 資料夾下的程式碼是唯讀的，**絕對不要 MUST NOT** 修改它。

### 技術棧
- **語言**: Go (1.26+)
- **路由**: [chi](https://github.com/go-chi/chi/v5)
- **依賴注入 (DI)**: [kessoku](https://github.com/mazrean/kessoku)
- **ORM**: [ent](https://entgo.io/)
- **可觀測性**: OpenTelemetry (OTel)
- **配置管理**: `gookit/config/v2`

## 專案結構

- `cmd/`: 包含所有執行檔的入口點。
  - `api/`: 主要的 API 伺服器，使用 `kessoku` 進行依賴注入 (見 `kessoku.go`)。
- `internal/`: 內部邏輯。
  - `data/`: 配置讀取、資料庫客戶端與資料模型封裝。提供 `NewClient` 與測試用的 `NewTestClient`。
  - `server/`: HTTP 伺服器初始化與路由定義 (見 `api.go`)。
  - `service/`: 服務層，負責處理 HTTP 請求、路由邏輯與 Tracing。
  - `model/`: 資料模型層，負責具體的資料庫操作與業務邏輯實作。
  - `usecase/`: 業務邏輯層，目前主要用於複雜的背景任務 (如 `worker`)。
  - `lib/`: 內部通用函式庫 (如 `logs`)。
- `ent/`: 由 `ent` 自動生成的 ORM 程式碼。
  - `schema/`: 資料庫 Schema 定義。
- `pixelfed/`: (子目錄) 原始 PixelFed (PHP) 的相關程式碼，用於移植 API 時的參考，**絕對不要 MUST NOT** 修改它。

## 開發規範

- **程式碼風格**: 遵循 [Google Go Style Guide](https://google.github.io/styleguide/go/)，可以使用 `go-style-guide` 這個 Agent Skill 進行確認。
- **依賴注入**: 專案使用 `kessoku` 進行編譯時 DI。新增 Service 或 Model 後，務必在 `cmd/api/kessoku.go` 中註冊，並運行 `make gen` 生成代碼，可以使用 `kessoku-di` 這個 Agent Skill 獲得幫助。
- **錯誤處理**: 使用 `log/slog` 進行日誌記錄，錯誤屬性應使用 `logs.ErrAttr(err)`。
- **格式化**: 提交前必須運行 `make lint` 確保格式符合 `gofmt` 與 `goimports` 標準。
- **可觀測性**：在 `internal/service/` 中的程式碼，需要加入相應的 Tracer 以利 Open Telemetry 追蹤，如下範例所示：

```go
func (s *Service) Login(w http.ResponseWriter, r *http.Request) {
	ctx, span := internal.T.Start(r.Context(), "{ServiceName}.{FunctionName}")
	defer span.End()
	
	s.model.SomeOperation(ctx) 
}
```

## 自動化工具 (Makefile)

專案根目錄提供了 `Makefile` 用於常用操作：

- `make init`: 安裝開發所需工具。
- `make gen`: 執行 `go generate ./...` (用於生成 `ent` 與 `kessoku` 代碼)。
- `make lint`: 格式化程式碼。
- `make test`: 執行所有測試。
- `make build`: 編譯所有 `cmd/` 下的執行檔到 `bin/` 目錄。
- `make all`: 執行 `gen` -> `lint` -> `test` -> `build` 的完整流程。

## API 開發流程

專案採用 **Service-Model-Ent** 的分層架構，開發 API 時請遵循以下流程：

### 1. 服務層 (Service Layer) - `internal/service/{module}`
服務層負責 HTTP 介面定義與 Handler 實作。
- **定義介面**: 在 `service.go` 中定義 `Service` 介面及其需要的依賴介面 (如 `Getter`, `Storer`)。
- **實作邏輯**: 建立 `svc` 結構體並實作 `Service` 介面，透過依賴介面調用資料邏輯。
- **Mocking**: 在 `service.go` 中加入 `//go:generate moq` 指令，生成 Mock 物件用於測試。

### 2. 資料模型層 (Model Layer) - `internal/model/{module}`
資料模型層負責實作 Service 層定義的依賴介面。
- **初始化**: 在 `model.go` 中定義 `Model` 結構體與 `NewModel` 構造函式，通常會接收 `*data.Client`。
- **實作功能**: 直接使用 `ent` 客戶端進行資料庫操作。

### 3. 依賴注入 (DI) - `cmd/api/kessoku.go`
- 在 `kessoku.go` 中將 Service 與 Model 綁定。
- 如果 Service 依賴特定介面，使用 `kessoku.Bind[Interface](kessoku.Provide(NewModel))` 進行綁定。

### 4. 測試策略
- **Service 測試**: 使用 `moq` 生成的 Mock 物件來隔離 Model 依賴，測試 Handler 邏輯。
- **Model 測試**: 使用 `data.NewTestClient(t)` 建立測試用的資料庫實例 (SQLite in-memory)，測試資料庫操作。

## Agent 注意事項

1. **優先參考原始實作**: 移植 API 時，請先閱讀 `pixelfed/app/Http/Controllers/Api/` 下對應的 PHP 實作。
2. **依賴檢查**: 加入新套件前應確認 `go.mod` 及其與現有技術棧的相容性。
3. **環境配置**: 預設配置文件為 `config.yaml`，範例參考 `config.example.yaml`。
