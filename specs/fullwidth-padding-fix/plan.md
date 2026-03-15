# Plan - Fullwidth and Padding Fixes

## Initial Requirements
- Base branch: develop
- Work branch: feature/fullwidth-padding-fix
- Target PR branch: develop

## Phase 1: Setup
- [x] Sync develop
- [x] Create branch feature/fullwidth-padding-fix
- [x] Update specs/fullwidth-padding-fix/context.md
- [x] Update specs/fullwidth-padding-fix/plan.md

## Phase 2: Implementation Steps
- [ ] Inspect and fix `static/templates/password-generator.html`
  - Remove `max-width` constraints.
  - Ensure `.tool-page` or equivalent full-width classes are used.
  - Remove unnecessary top padding/margin.
- [ ] Inspect and fix `static/templates/kill-port.html`
  - Remove `max-width` constraints.
  - Ensure `.tool-page` or equivalent full-width classes are used.
  - Remove unnecessary top padding/margin.
- [ ] Inspect and fix `static/templates/regex-tester.html`
  - Remove `max-width` constraints.
  - Ensure `.tool-page` or equivalent full-width classes are used.
  - Remove unnecessary top padding/margin.

## Phase 3: Testing & Validation
- [ ] UI checks: Verify the three tools occupy 100% width and have minimal top padding.
- [ ] `make pre-commit`

## Phase 4: Delivery
- [ ] Commit conventional message
- [ ] Push feature branch
- [ ] Open PR to develop
- [ ] Merge via squash
