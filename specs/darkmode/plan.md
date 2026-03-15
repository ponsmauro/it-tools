# Plan - Dark Mode Tests

## Initial Requirements
- Base branch: develop
- Work branch: feature/darkmode
- Target PR branch: develop

## Steps
1. Create `internal/infrastructure/templates/darkmode_test.go`.
2. Implement a table-driven test that reads all `.html` files in `static/templates/`.
3. For each file, verify that it does NOT contain light mode hex codes (e.g., `#ffffff`, `#f8fafc`, `#e2e8f0`, `#1e293b` as text color).
4. For each file, verify that it DOES contain dark mode hex codes (e.g., `#0f172a`, `#1e293b` as background, `#334155`, `#f1f5f9`).
5. Run `go test ./... -v` to ensure the tests pass.

## Validation Plan
- Unit tests: `darkmode_test.go` will validate the templates.
- Integration/API tests: N/A
- UI smoke/thorough tests: N/A (already done manually)
