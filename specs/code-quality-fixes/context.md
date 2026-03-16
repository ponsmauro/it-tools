# Context - Code Quality Fixes

## Current State
- Previous frontend issues (inline styles, embedded `<style>`, embedded `<script>`) were already fixed.
- The latest generated report still showed backend coverage as 0.0%, but this was a report-generation bug.
- Real backend coverage measured with `go tool cover -func=coverage.out` is 93.1%.
- Root cause: coverage parser in `generate_report.py` was reading the first package coverage line from `go test` output instead of the total coverage line from the coverage profile.

## Constraints
- Architecture constraints: Clean Architecture must be maintained.
- Security constraints: No inline scripts/styles to support future CSP.
- UX constraints: UI must remain identical after refactoring.

## Tech Stack
- Backend: Go (Golang)
- Frontend: HTML Templates, CSS, Vanilla JS
- i18n: Supported
- Testing: Go testing package (Table-Driven Tests required)
