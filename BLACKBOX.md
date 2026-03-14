# IT Tools - Project Context

## Project Overview

**IT Tools** is a web application that provides a collection of useful tools for developers. The project uses:

- **Backend**: Go (Golang)
- **Frontend**: Go HTML Templates + CSS
- **Architecture**: Server-side rendering with static files
- **i18n**: Multi-language support enabled

## Project Structure (Clean Architecture)

```
it-tools/
├── cmd/
│   └── server/
│       └── main.go              # Entry point
├── internal/
│   ├── domain/                  # Enterprise business rules
│   │   └── models/              # Core entities (Tool, etc.)
│   ├── application/             # Application business rules
│   │   ├── usecases/            # Use cases
│   │   └── services/            # Application services
│   ├── infrastructure/          # Frameworks & drivers
│   │   ├── handlers/            # HTTP handlers
│   │   ├── templates/           # Template rendering
│   │   └── i18n/                # Internationalization
│   └── utils/                   # Shared utilities
├── pkg/                         # Reusable packages
│   └── path/                    # Cross-platform path utilities
├── templates/                   # HTML templates
│   ├── layout.html
│   ├── index.html
│   └── about.html
├── static/
│   └── css/
│       └── style.css
├── go.mod
└── BLACKBOX.md
```

## Mandatory Development Rules

### 1. Language
- **ALL code and documentation MUST be in English**
- Variable names, function names, comments, commit messages
- User-facing pages must support internationalization (i18n)

### 2. Code Quality Principles
- **DRY** (Don't Repeat Yourself): Reuse methods and components
- **KISS** (Keep It Simple, Stupid): Simple, readable solutions
- **YAGNI** (You Aren't Gonna Need It): No over-engineering

### 3. Clean Architecture
- **Domain Layer**: Core entities and business rules (no external dependencies)
- **Application Layer**: Use cases and application services
- **Infrastructure Layer**: HTTP handlers, database, external services
- **Dependency Rule**: Inner layers never depend on outer layers

### 4. Naming Conventions
- Variables and methods must have clear, descriptive names
- Use meaningful names that reflect functionality
- Example: `GetAllTools()` not `GetTools()`, `FormatFilePath()` not `FFP()`

### 5. Cross-Platform Support
- Use `path/filepath` for file paths (OS-specific)
- Never hardcode paths; use utility functions
- Test on multiple operating systems

### 6. Testing Requirements
- **ALL Go tests MUST use table-driven tests**
- Use `mock.Mock` from testing library
- Example pattern:
```go
func TestFunctionName(t *testing.T) {
    tests := []struct {
        name     string
        input    string
        expected string
    }{
        {"case 1", "input1", "expected1"},
        {"case 2", "input2", "expected2"},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            mock := new(mock.Mock)
            // test implementation
        })
    }
}
```

### 7. Internationalization (i18n)
- All user-facing strings must use i18n keys
- Never hardcode display text in templates
- Use translation files for each supported language

### 8. Code Compactness
- Keep the project as compact as possible
- Avoid unnecessary abstractions
- Combine related functionality when appropriate
- **ALWAYS reuse code** to minimize duplication
- Extract common patterns into shared functions

### 9. No Magic Values
- **NEVER leave magic numbers or magic strings in code**
- Always replace them with named constants
- Example: Use `const DefaultPort = 8080` instead of `8080`
- Example: Use `const DefaultLanguage = "en"` instead of hardcoded "en"
- Constants should be grouped and named descriptively

### 10. Lightweight Code Reuse
- Prioritize code reuse over duplication
- Create utility functions for common operations
- Share functionality across packages when appropriate
- Avoid redundant implementations

### 11. Screen Design Optimization
- **ALWAYS maximize the use of available screen dimensions**
- Use full viewport width (100vw) for main containers
- Use full viewport height (100vh) when appropriate
- Avoid unnecessary margins or padding that waste space
- Design responsive layouts that fill the screen on all devices
- Use CSS Grid and Flexbox to utilize available space efficiently
- Content should never be unnecessarily constrained

### 12. Test Coverage
- **ALWAYS try to reach 100% test coverage for each file**
- Write comprehensive tests that cover all functions, edge cases, and error paths
- Use table-driven tests for thorough coverage
- Every public function should have corresponding tests
- Aim for 100% coverage; if not achievable, document why

### 13. Error Logging
- **ONLY log error cases, never successful operations**
- Place exactly ONE log statement at the exact point where the error occurs
- The log must include ALL context needed to reproduce the erroneous use case
- Include relevant input values, state variables, and any debugging information
- Example: `log.Printf("Error processing request: method=%s, path=%s, error=%v", r.Method, r.URL.Path, err)`
- Do NOT log success cases, routine operations, or informational messages
- Logs should only be used for debugging failures

## 14. BLACKBOX Analysis Rule
**NEW FEATURES/FIXES** → Carpetas `specs/{feature-name}/`
```
specs/{feature}/
├── context.md     # Things to consider for planning
└── plan.md        # Detailed step-by-step development plan
```
**Benefits:**
- Context preservation
- Reproducible planning
- Feature isolation
- Easy rollback/review


## Building and Running

### Prerequisites
- Go 1.26+ installed

### Commands

**Pre-commit workflow:**
```bash
make pre-commit   # test + lint + build
make              # = pre-commit
make test         # tests + coverage
make lint         # golangci-lint
make build        # compile server
make run          # dev server
```

**MANDATORY**: Before EVERY commit/push:
```
make pre-commit  # or make all
```

**First time:**
```bash
make lint-install
```

**The server runs on http://localhost:8080**


### Available Routes
- `/` - Home page (Tools list)
- `/about` - About page
- `/static/*` - Static files (CSS, images, etc.)

## Current Tools (10 Examples)

1. JSON Formatter
2. Base64 Encoder/Decoder
3. UUID Generator
4. Hash Generator (MD5, SHA)
5. URL Encoder/Decoder
6. Cron Expression Parser
7. JWT Decoder
8. Color Converter
9. SQL Formatter
10. Regex Tester

### 15. Tool Architecture (Modular Tabs)
- **Click on tool** → POST `/tools/{id}` opens **dynamic tab** (lazy load).
- Tabs bar: Active tool, close button (X), max 8 tabs.
- **Dynamic content**: Server response HTML/JS/CSS fragment injected via AJAX (no full page reload).
- **Separate PR per tool**: `feature/{tool-id}` branch → `develop`.
- Structure:
  - `internal/tools/{id}/` package: handler.go, template.tmpl (embedded), tool.go (models.Tool).
  - Registration: `tool_registry.Register("{id}", ToolHandlerFunc, Base64Tool)`.
  - POST `/tools/{id}`: `{ "action": "open" }` → render tab content.
- Benefits: Isolation, lazy load, multi-tab workflow, easy PR review.
- Frontend: index.html JS manages tabs (localStorage persist open tabs).

### 16. Branch Policy (MANDATORY)
- **NEVER create blackboxai/ branches or repos**
- Use `feature/{tool-id}` branches only

### 17. Go Naming Convention (MANDATORY)
- Go struct/function names **MUST NOT** repeat package name
- Package `handlers` → `KillPortHandler()` not `HandlersKillPortHandler()`
- File `killport.go` not `kill_port_handler.go`

### 18. Standard Delivery Workflow Template (MANDATORY)
1. Sync base branch:
   - `git checkout develop`
   - `git pull origin develop`
2. Create work branch:
   - **Only** `feature/{feature-name}` format
   - Example: `feature/kill-port`
3. Confirm scope and update context:
   - Update `BLACKBOX.md` when introducing new mandatory rules
   - Create/update `specs/{feature}/context.md`
   - Create/update `specs/{feature}/plan.md`
4. Create/update `TODO.md` with phases and checklist.
5. Implement feature following Clean Architecture and i18n rules.
6. Add/update table-driven tests for affected logic and handlers.
7. Run local validation:
   - `go test ./... -v`
   - `make run`
   - curl checks for impacted endpoints (200/400/404/405)
8. Run functional validation on affected pages/components.
9. Review git changes and exclude generated artifacts from commit.
10. Commit with representative conventional message.
11. Push branch:
   - `git push -u origin feature/{feature-name}`
12. Open PR to `develop`.
13. Address review feedback and rerun validations.
14. Merge PR to `develop`.
15. Post-merge hygiene:
   - Clean branch if policy allows
   - Confirm `develop` includes final changes.

### 19. PR Naming & Commit Convention (MANDATORY)
- PR/commit titles MUST follow:
  - `feat(scope): short summary`
  - `fix(scope): short summary`
  - `refactor(scope): short summary`
  - `test(scope): short summary`
  - `docs(scope): short summary`
- Scope should match feature/tool ID when possible.
- Examples:
  - `feat(kill-port): add localized OS guide tabs`
  - `fix(kill-port): validate unsupported HTTP methods`
- PR body minimum sections:
  1. **Summary**
  2. **Validation**
  3. **Notes**

### 20. Testing Protocol (MANDATORY Before Merge)
#### 20.1 Minimum Validation
- `go test ./... -v` must pass.
- `make run` must start server without startup errors.
- Curl tests must cover impacted endpoints:
  - Happy path(s)
  - Invalid JSON/body
  - Invalid action/input
  - Unknown resource/tool

#### 20.2 Thorough Validation
- **Web/UI**:
  - Navigate all affected pages/sections/components
  - Interact with all links/buttons/inputs
  - Verify i18n labels in supported languages
  - Verify result states and error states
- **API/Backend**:
  - Test all impacted endpoints with:
    - happy paths
    - error paths
    - edge cases (empty body, wrong content-type, malformed payload)
  - Validate unsupported methods return correct status code (prefer `405 Method Not Allowed` where applicable)

#### 20.3 Test Result Handling
- If tests fail:
  - Fix implementation first
  - Rerun tests
  - Only continue when passing or when failure is explicitly documented and approved.

### 21. API Validation Rules (MANDATORY)
- Validate request method per endpoint.
- Validate request payload and JSON decoding errors.
- Validate required fields and accepted action values.
- Unknown tool/resource should return `404`.
- Invalid client input should return `400`.
- Unsupported method should return `405` when endpoint contract is method-specific.
- Internal execution/render failures should return `500`.
- Error logging must include reproduction context and follow section 13 (error-only logging).

### 22. Feature Template Blocks (Copy/Paste)

#### 22.1 `specs/{feature}/context.md`
```md
# Context - {Feature Name}

## Current State
- Relevant existing files/routes/components:
  - ...

## Constraints
- Architecture constraints:
- Security constraints:
- UX constraints:

## Tech Stack
- Backend:
- Frontend:
- i18n:
- Testing:
```

#### 22.2 `specs/{feature}/plan.md`
```md
# Plan - {Feature Name}

## Initial Requirements
- Base branch: develop
- Work branch: feature/{feature-name}
- Target PR branch: develop

## Steps
1. ...
2. ...
3. ...

## Validation Plan
- Unit tests:
- Integration/API tests:
- UI smoke/thorough tests:
```

#### 22.3 `TODO.md` phase template
```md
# {Feature Name} TODO

## Phase 1: Setup
- [ ] Sync develop
- [ ] Create branch feature/{feature-name}
- [ ] Update specs/{feature}/context.md
- [ ] Update specs/{feature}/plan.md

## Phase 2: Implementation
- [ ] Backend
- [ ] Frontend
- [ ] i18n

## Phase 3: Testing
- [ ] go test ./... -v
- [ ] make run
- [ ] curl API checks
- [ ] UI checks

## Phase 4: Delivery
- [ ] Commit conventional message
- [ ] Push feature branch
- [ ] Open PR to develop
- [ ] Address review feedback
- [ ] Merge
```

### 23. Definition of Done (DoD)
A feature is considered done only if all are true:
1. Implemented according to approved scope.
2. Architecture rules respected (Clean Architecture + no inline HTML/JS in handlers).
3. i18n completed for all user-facing text.
4. Table-driven tests added/updated.
5. Minimum and required thorough testing completed (or formally waived).
6. PR title/body follow convention.
7. PR merged to `develop`.
8. Documentation (`BLACKBOX.md`, specs, TODO) updated accordingly.


## Notes

- The server runs on port 8080 by default
- Templates are parsed using `template.ParseGlob`
- Static files are served from the `static/` directory
- All paths must use `filepath.Join()` for cross-platform compatibility

