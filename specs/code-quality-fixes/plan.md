# Plan - Code Quality Fixes

## Initial Requirements
- Base branch: develop
- Work branch: fix/code-quality-report
- Target PR branch: develop

## Phase 1: Setup
- [x] Sync develop
- [x] Create branch fix/code-quality-report
- [x] Update specs/code-quality-fixes/context.md
- [x] Update specs/code-quality-fixes/plan.md

## Phase 2: Implementation Steps
- [x] Step 1: Remove all inline styles (`style="..."`) from HTML templates and move them to `static/css/style.css`.
- [x] Step 2: Extract all embedded `<style>` tags from HTML templates into `static/css/style.css`.
- [x] Step 3: Extract all embedded `<script>` tags from HTML templates into separate `.js` files in `static/js/`.
- [x] Step 4: Write unit tests for backend packages (`handlers`, `usecases`, `repositories`, `models`) to achieve >80% coverage.
- [x] Step 5: Update `AGENTS.md` with new rules to prevent these issues in the future.

## Phase 3: Testing & Validation
- [x] `go test ./... -v`
- [x] `make run`
- [x] curl API checks
- [x] UI checks

## Phase 4: Delivery
- [ ] Commit conventional message
- [ ] Push feature branch
- [ ] Open PR to develop
- [ ] Address review feedback
- [ ] Merge

## Follow-up Fix (Current Branch)
- [x] Validate real backend coverage with `go tool cover -func=coverage.out`.
- [x] Fix report generation to compute total coverage from coverage profile.
- [x] Regenerate `it-tools-code-quality-gpt-codex.html` with corrected metrics.
- [ ] Add preventive rule to `AGENTS.md` for future coverage-report generation.
