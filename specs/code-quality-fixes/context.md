# Context - Code Quality Fixes

## Current State
- The project has a low code quality score (19.0/100) according to the generated report.
- Frontend issues: 18 inline styles, 25 embedded `<style>` tags, 23 embedded `<script>` tags.
- Backend issues: 0.0% test coverage.

## Constraints
- Architecture constraints: Clean Architecture must be maintained.
- Security constraints: No inline scripts/styles to support future CSP.
- UX constraints: UI must remain identical after refactoring.

## Tech Stack
- Backend: Go (Golang)
- Frontend: HTML Templates, CSS, Vanilla JS
- i18n: Supported
- Testing: Go testing package (Table-Driven Tests required)
