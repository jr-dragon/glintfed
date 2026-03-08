---
name: go-style-guide
description: 遵循 Google Go 風格指南編寫地道 (idiomatic) 且具可讀性的 Go 程式碼。適用於 Go 專案重構、代碼審查及新功能開發。
---

# Go 語言風格指南 (Google Style)

本技能旨在引導開發者遵循 Google 的 Go 風格規範，編寫清晰、簡潔且具可維護性的程式碼。

## 核心指導原則

在編寫 Go 程式碼時，應優先考慮以下屬性（按重要性排序）：

1. **清晰 (Clarity)**：程式碼的意圖和原理對讀者來說是明確的。
2. **簡單 (Simplicity)**：以最簡單的方式實現目標。
3. **簡練 (Concision)**：具備高信噪比，避免冗餘。
4. **可維護性 (Maintainability)**：易於長期維護。
5. **一致性 (Consistency)**：與整體程式碼庫風格保持一致。

## 快速參考指南

### 命名規範 (Naming)
- **禁用底線**：標識符中原則上不使用底線（測試/基準測試函數及 cgo 除外）。
- **包名**：僅使用小寫字母和數字，不使用底線或混合大小寫。例如：`tabwriter`, `oauth2`。
- **變數名**：Go 傾向於短變數名，尤其是範圍較小時。
- **MixedCaps**：導出與非導出標識符均使用 PascalCase 或 camelCase。

### 註釋與文檔
- 註釋應解釋「為什麼」這麼做，而非僅僅是「做了什麼」。
- 導出標識符必須有文檔註釋。

## 詳細參考資源

本技能包含以下詳細參考文件，當需要針對特定場景進行深入分析時，請閱讀對應文件：

- **[基礎指南 (Guide)](references/google/guide.md)**：定義了 Go 風格的核心基礎與原則。
- **[風格決定 (Decisions)](references/google/decisions.md)**：收錄了具體的風格決定、理由與示例。
- **[最佳實踐 (Best Practices)](references/google/best-practices.md)**：針對常見問題的演進模式與推薦做法。
- **[總覽與定義 (Index)](references/google/index.md)**：各文檔的導航與術語定義。

## 使用建議

1. **程式碼審查**：使用本技能檢查 PR 是否符合 Go 地道用法。
2. **重構**：將不符合規範的命名或結構按指南進行調整。
3. **新功能開發**：從開始就遵循清晰與簡單的原則。
