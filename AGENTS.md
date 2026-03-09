# GlintFed Agent 指南

本文件為 AI Agent 提供專案的技術背景、開發規範及操作指引。

## 專案概述

GlintFed 是基於 PixelFed 的專案，將原本由 PHP/Laravel 的實作改為 Go 的實作。

預期 GlintFed 的實作會完全相容於 PixelFed 所提供的 API，為此 Agents 可以參考 `pixelfed` 資料夾下的程式碼，但 `pixelfed` 資料夾下的程式碼是唯讀的，**絕對不要 MUST NOT** 修改它。

### 技術棧
- **語言**: Go (1.26+)
- **路由**: [chi](https://github.com/go-chi/chi/v5)
- **可觀測性**: OpenTelemetry (OTel)
- **配置管理**: `gookit/config/v2`

## 專案結構

- `cmd/`: 包含所有執行檔的入口點。
  - `api/`: 主要的 API 伺服器。
- `internal/`: 內部邏輯。
  - `data/`: 配置讀取、資料庫客戶端與資料模型封裝。
  - `server/`: HTTP 伺服器初始化與路由定義。
  - `service/`: 各個業務邏輯模組 (如 `healthcheck`)，相當於 Laravel 中的 Controller，移植時務必保持結構相同。
  - `lib/`: 內部通用函式庫 (如 `logs`)。
- `pixelfed/`: (子目錄) 原始 PixelFed (PHP) 的相關程式碼，用於移植 API 時的參考，**絕對不要 MUST NOT** 修改它。

## 開發規範

- **程式碼風格**: 遵循 [Google Go Style Guide](https://google.github.io/styleguide/go/)，可以使用 `go-style-guide` 這個 Agent Skill 進行確認
- **錯誤處理**: 使用 `log/slog` 進行日誌記錄，錯誤屬性應使用 `logs.ErrAttr(err)`。
- **格式化**: 提交前必須運行 `make lint` 確保格式符合 `gofmt` 與 `goimports` 標準。
- **可觀測性**：在 `internal/service/` 中的程式碼，需要加入相應的 Tracer 以利 Open Telemetry 追蹤，如下範例所示：

```go
func (s *Service) Login(w http.ResponseWriter, r *http.Request) {
	// ctx, span := otel.Tracer("Service").Start(r.Context(), "{ServiceName}.{FunctionName}")
	ctx, span := internal.T.Start(r.Context(), "{ServiceName}.{FunctionName}")
	defer span.End()
	
	s.clients.DB.GetActiveUsers(ctx) // use ctx from otel.Tracer("Service").Start(), not r.Context() 
}
```

## 自動化工具 (Makefile)

專案根目錄提供了 `Makefile` 用於常用操作：

- `make init`: 安裝開發所需工具 (如 `goimports`)。
- `make gen`: 執行 `go generate ./...` (用於生成 `ent` 代碼)。
- `make lint`: 格式化程式碼。
- `make test`: 執行所有測試。
- `make build`: 編譯所有 `cmd/` 下的執行檔到 `bin/` 目錄。
- `make all`: 執行 `gen` -> `lint` -> `test` -> `build` 的完整流程。
- `make clean`: 清除編譯產物。

## Agent 注意事項

2. **依賴檢查**: 加入新套件前應確認 `go.mod` 及其與現有技術棧的相容性。
3. **環境配置**: 預設配置文件為 `config.yaml`，範例參考 `config.example.yaml`。
