# IT Tools - Project Context

> 🚨 **CRITICAL RULE 0 FOR ALL AI AGENTS** 🚨
> **YOU MUST READ THIS ENTIRE FILE BEFORE EXECUTING ANY COMMAND, WRITING ANY CODE, OR STARTING ANY TASK.**
> Failure to follow the rules in this file (especially the Git Flow, PR process, and `specs/` folder creation) is a severe violation of the project's constraints.


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

### 0. Modifying this file
- **NEVER modify `agents.md` without explicit authorization from the user.**
- If you believe a rule needs to be updated or added, you must ask the user for permission first.

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
- **ALWAYS try to reach 100% test coverage for each file** (Minimum acceptable is 80%).
- Write comprehensive tests that cover all functions, edge cases, and error paths.
- Use table-driven tests for thorough coverage.
- Every public function should have corresponding tests.
- Aim for 100% coverage; if not achievable, document why.

### 13. Error Logging
- **ONLY log error cases, never successful operations**
- Place exactly ONE log statement at the exact point where the error occurs
- The log must include ALL context needed to reproduce the erroneous use case
- Include relevant input values, state variables, and any debugging information
- Example: `log.Printf("Error processing request: method=%s, path=%s, error=%v", r.Method, r.URL.Path, err)`
- Do NOT log success cases, routine operations, or informational messages
- Logs should only be used for debugging failures

### 14. Security Guidelines
- **Input Sanitization**: Always sanitize user inputs on the backend to prevent SQL Injection and command injection.
- **XSS Prevention**: Always escape HTML content when rendering user input in the frontend. Never use `innerHTML` with unsanitized data.
- **Sensitive Data**: Never log sensitive data (passwords, tokens, keys).

### 15. Accessibility (a11y)
- **Semantic HTML**: Use proper HTML5 tags (`<button>`, `<nav>`, `<main>`, etc.) instead of generic `<div>`s.
- **ARIA Labels**: Use `aria-label` for elements that lack visible text (e.g., icon-only buttons).
- **Keyboard Navigation**: Ensure all interactive elements are reachable and usable via the `Tab` and `Enter` keys.

### 16. CSS & Styling Conventions
- **No Inline Styles**: NEVER use `style="..."` attributes in HTML. All styling must be done via CSS classes in `static/css/style.css`.
- **No Embedded Style Tags**: NEVER use `<style>` tags inside HTML files. All CSS must be in `static/css/style.css`.
- **CSS Variables**: Use CSS variables (`var(--primary-color)`) for theming, especially to support Dark Mode.
- **Methodology**: Keep CSS modular and scoped to avoid global style conflicts.

### 16.1 JavaScript Conventions
- **No Embedded Script Tags**: NEVER use `<script>` tags with inline code inside HTML files. All JavaScript logic must be in separate `.js` files under `static/js/` and linked via `<script src="...">`. This is crucial for maintainability and Content Security Policy (CSP).

### 17. Frontend Error Handling & State
- **Visible Feedback**: Never fail silently. Always show a visible error message (toast, inline text) to the user if an API call fails.
- **Loading States**: Show loading indicators (spinners, disabled buttons) during asynchronous operations.
- **State Management**: Prefer `localStorage` for persistent user preferences (like dark mode) and `sessionStorage` for temporary session data.

### 18. Performance Optimization
- **Non-Blocking**: Do not block the main thread with heavy synchronous JavaScript operations. Use Web Workers if necessary.
- **Lazy Loading**: Load resources only when needed to keep the initial page load fast.

## 19. BLACKBOX Analysis Rule
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

### 20. Tool Architecture (Modular Tabs)
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

### 21. Branch Policy (MANDATORY)
- **NEVER create blackboxai/ branches or repos**
- **Branches from `main`**: Can ONLY be `hotfix/*`, `RC/*`, or `backports/*`.
- **Branches from `develop`**: `feature/*`, `fix/*`, `chore/*`, etc. MUST be created from `develop` unless the user explicitly says otherwise.
- **Merging to `develop`**: All merges to `develop` MUST be done via **squash** with a short, representative message of the PR's functionality.

### 22. Go Naming Convention (MANDATORY)
- Go struct/function names **MUST NOT** repeat package name
- Package `handlers` → `KillPortHandler()` not `HandlersKillPortHandler()`
- File `killport.go` not `kill_port_handler.go`

### 23. Standard Delivery Workflow Template (MANDATORY)
1. Sync base branch:
   - `git checkout develop`
   - `git pull origin develop`
2. Create work branch:
   - **Only** `feature/{feature-name}` format
   - Example: `feature/kill-port`
3. Confirm scope and update context:
   - Update `BLACKBOX.md` when introducing new mandatory rules
   - Create/update `specs/{feature}/context.md`
   - Create/update `specs/{feature}/plan.md` (This file acts as your TODO list. **Do NOT create separate TODO.md files in the root directory**).
4. Implement feature following Clean Architecture and i18n rules, checking off items in `plan.md`.
5. Add/update table-driven tests for affected logic and handlers.
6. Run local validation:
   - `go test ./... -v`
   - `make run`
   - curl checks for impacted endpoints (200/400/404/405)
7. Run functional validation on affected pages/components.
8. Review git changes and exclude generated artifacts from commit.
9. Commit with representative conventional message.
10. Push branch:
   - `git push -u origin feature/{feature-name}`
11. Open PR to `develop` (preferably using GitHub CLI):
   - `gh pr create --base develop --head feature/{feature-name} --title "feat({feature-name}): short summary" --body "Detailed description of changes..."`
12. Address review feedback and rerun validations.
13. Merge PR to `develop`.
14. Post-merge hygiene:
   - Clean branch if policy allows
   - Confirm `develop` includes final changes.

### 24. PR Naming & Commit Convention (MANDATORY)
- **Each PR should be focused on a single tool or feature** to make reviews easier.
- PR/commit titles MUST follow this exact format:
  - `[entidad o funcionalidad] - descripcion`
- Examples:
  - `[kill-port] - add localized OS guide tabs`
  - `[darkmode] - implement consistent dark mode and compact UI`
- PR body minimum sections:
  1. **Summary**
  2. **Validation**
  3. **Notes**

### 25. Testing Protocol (MANDATORY Before Merge)
#### 25.1 Minimum Validation
- `go test ./... -v` must pass.
- `make run` must start server without startup errors.
- Curl tests must cover impacted endpoints:
  - Happy path(s)
  - Invalid JSON/body
  - Invalid action/input
  - Unknown resource/tool

#### 25.2 Thorough Validation
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

#### 25.3 Test Result Handling
- If tests fail:
  - Fix implementation first
  - Rerun tests
  - Only continue when passing or when failure is explicitly documented and approved.

### 26. API Validation Rules (MANDATORY)
- Validate request method per endpoint.
- Validate request payload and JSON decoding errors.
- Validate required fields and accepted action values.
- Unknown tool/resource should return `404`.
- Invalid client input should return `400`.
- Unsupported method should return `405` when endpoint contract is method-specific.
- Internal execution/render failures should return `500`.
- Error logging must include reproduction context and follow section 13 (error-only logging).

### 27. Feature Template Blocks (Copy/Paste)

#### 27.1 `specs/{feature}/context.md`
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

#### 27.2 `specs/{feature}/plan.md`
```md
# Plan - {Feature Name}

## Initial Requirements
- Base branch: develop
- Work branch: feature/{feature-name}
- Target PR branch: develop

## Phase 1: Setup
- [ ] Sync develop
- [ ] Create branch feature/{feature-name}
- [ ] Update specs/{feature}/context.md
- [ ] Update specs/{feature}/plan.md

## Phase 2: Implementation Steps
- [ ] Step 1...
- [ ] Step 2...
- [ ] Step 3...

## Phase 3: Testing & Validation
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

### 28. Definition of Done (DoD)
A feature is considered done only if all are true:
1. Implemented according to approved scope.
2. Architecture rules respected (Clean Architecture + no inline HTML/JS in handlers).
3. i18n completed for all user-facing text.
4. Table-driven tests added/updated.
5. Minimum and required thorough testing completed (or formally waived).
6. PR title/body follow convention.
7. PR merged to `develop`.
8. Documentation (`BLACKBOX.md`, specs) updated accordingly.

### 29. Terminal Management (MANDATORY)
- **ALWAYS close terminals when they are no longer needed.**
- Do not leave multiple terminals open in the background unless a long-running process (like a server) is explicitly required.
- Use the `exit` command or equivalent to close terminals after executing one-off scripts or commands.

### 30. No HTML/JS/CSS Embedded in Go Files (MANDATORY)
- **NEVER embed HTML, CSS, or JavaScript as string literals or constants inside `.go` files.**
- All markup belongs in `.html` template files under `static/templates/`.
- Go handlers must only call `templates.ExecuteTemplate(...)` — never write raw HTML strings to `http.ResponseWriter`.
- Example violation: `const content = "<div style=...><script>...</script></div>"` inside a `.go` file.

### 31. Single Initialization in main.go (MANDATORY)
- Every initialization or registration function (e.g. `tool.Init()`) must be called **exactly once** in `main.go`.
- Group all init calls under a clearly commented section: `// Register tools`.
- Duplicate init calls cause silent bugs if the registry or resource is not idempotent.

### 32. Explicit Errors on Not-Found Lookups (MANDATORY)
- Functions that search by ID or key must **never return `(nil, nil)`**.
- When a resource is not found, return a named sentinel error: `var ErrNotFound = errors.New("not found")`.
- Callers must be able to distinguish "not found" from "internal error" without nil-checking a pointer.
- Example: `return nil, ErrNotFound` instead of `return nil, nil`.

### 33. No time.Now() in Static Data (MANDATORY)
- Hardcoded/static data (e.g. seed lists, in-memory repositories) must **not** use `time.Now()`.
- Use `time.Time{}` (zero value) for timestamps in static records.
- `time.Now()` in static data makes every call non-deterministic and breaks timestamp-based tests.

### 34. No Ignored Parameters in Public Function Signatures (MANDATORY)
- Public functions must **not** have ignored parameters (`_`) in their signature.
- If a parameter is unused, remove it from the signature entirely and update all callers.
- Ignored parameters in public APIs confuse callers and indicate dead interface design.
- Example violation: `func Register(id string, handler HandlerFunc, _ string)`.

### 35. Stub/Placeholder Tools Must Show "Coming Soon" (MANDATORY)
- Any tool that is not fully implemented must display a visible **"Coming Soon"** banner.
- **Never** expose mock data, hardcoded fake results, or `alert()` calls as if they were real functionality.
- Stub tools must clearly communicate their status to the user.
- Mock results (e.g. random PIDs, fake process lists, incorrect cron schedules) are forbidden without an explicit disclaimer.

### 36. No Deprecated JavaScript APIs (MANDATORY)
- **Never use deprecated browser APIs** in frontend templates:
  - Use `TextDecoder` instead of `escape()` / `unescape()`.
  - Use `navigator.clipboard.writeText()` instead of `document.execCommand('copy')`.
  - Use `crypto.randomUUID()` instead of manual UUID generation.
- If a fallback for older browsers is needed, wrap it in a feature-detection block and document it with a comment.

### 37. No External CDNs in Production (MANDATORY)
- **All third-party JavaScript libraries must be served locally** from `/static/js/vendor/`.
- Never load libraries from external CDNs (cdnjs, jsdelivr, unpkg, etc.) in production templates.
- External CDN dependencies cause failures when the network is unavailable and introduce supply-chain risks.
- To add a new vendor library: download the minified file, place it in `static/js/vendor/{lib}-{version}.min.js`, and reference it via `/static/js/vendor/...`.

### 38. No innerHTML with Unescaped Data (MANDATORY)
- **NEVER use `innerHTML` with user-controlled data, error messages, or any dynamic string.**
- Always use `textContent` to set text content of DOM elements.
- To build dynamic HTML structures, use `document.createElement()` + `textContent` + `appendChild()`.
- The only safe use of `innerHTML` is with fully static, developer-controlled strings (no variables).
- This applies to: user inputs, API responses, error messages from `catch(err)`, tool IDs from URLs, generated values (passwords, UUIDs, etc.).
- Example violation: `element.innerHTML = 'Error: ' + err.message`
- Example fix: `element.textContent = 'Error: ' + err.message`

### 39. Validate All HTTP Input at the Boundary (MANDATORY)
- **Every value extracted from an HTTP request must be validated before use.**
- Query parameters (e.g. `?lang=`) must be whitelisted against known valid values. Unknown values must fall back to the default.
- Path segments used as identifiers (e.g. `toolID` from `/tools/{id}`) must be validated against the known set of resources (e.g. via `GetByID`). If not found, return `404 Not Found` — never `500`.
- Never concatenate unvalidated URL path segments into template names, page titles, or any output.
- Example: `lang` param → whitelist against `config.SupportedLangEN`, `config.SupportedLangES`. Anything else → `config.DefaultLanguage`.

### 40. HTTP Security Headers Middleware (MANDATORY)
- **Every HTTP server must include a security headers middleware** applied globally to all routes.
- Minimum required headers:
  - `X-Content-Type-Options: nosniff`
  - `X-Frame-Options: DENY`
  - `Referrer-Policy: strict-origin-when-cross-origin`
- Implement as a single middleware function wrapping the default mux in `main.go`.
- Do NOT set these headers individually in each handler.

### 41. Limit Request Body Size on All POST Handlers (MANDATORY)
- **Every POST handler must wrap `r.Body` with `http.MaxBytesReader` before decoding.**
- Use a size appropriate to the expected payload. For JSON action payloads: `1024` bytes (1KB).
- This prevents trivial memory exhaustion DoS attacks.
- Example: `r.Body = http.MaxBytesReader(w, r.Body, 1024)`
- Place this as the first line in every handler that reads `r.Body`.

### 42. Pre-compute Static Data — No Alloc per Request (MANDATORY)
- **Static data (seed lists, in-memory repositories) must be computed once at construction time**, not on every method call.
- Repositories with static data must initialize their slice and lookup map in `New...()`, not in `GetAll()` or `GetByID()`.
- `GetAll()` must return a copy of the pre-computed slice. `GetByID()` must use a pre-built `map[string]T` for O(1) lookup.
- Never call `sort.Slice` or allocate a new slice inside a method that is called on every HTTP request.

### 43. Depend on Interfaces, Not Concrete Types (MANDATORY)
- **Infrastructure layer handlers must depend on interfaces, not concrete application layer types.**
- Define a port interface (e.g. `ToolUseCasePort`) in the `handlers` package that declares only the methods the handler needs.
- The concrete use case struct implements the interface implicitly (Go duck typing).
- This enables handler unit tests to use mocks without instantiating the real use case or repository.
- Example violation: `type Handler struct { toolUC *usecases.ToolUseCase }`
- Example fix: `type Handler struct { toolUC ToolUseCasePort }` where `ToolUseCasePort` is a local interface.

### 44. Cap All User-Driven Processing Loops (MANDATORY)
- **Any loop that processes user input must have an explicit upper bound.**
- Regex match loops, CSV row loops, diff computation loops, and any other iterative processing of user-provided data must stop at a reasonable maximum (e.g. 500 matches, 10,000 rows).
- When the cap is reached, show a visible notice to the user: "Showing first N results".
- This prevents browser hangs and server-side resource exhaustion from crafted inputs.

### 45. CSS Variables for All Dynamic Styling in JavaScript (MANDATORY)
- **Never hardcode hex color values in JavaScript** for dynamic styling (e.g. hover states, copy feedback, highlights).
- Use CSS classes with CSS variable-based styles instead: `element.classList.add('state-copied')`.
- Define the state styles in `style.css` using CSS variables: `.state-copied { background: var(--success-bg); color: var(--success-text); }`.
- This ensures all dynamic styles respect the theme and are maintainable from a single source.

### 46. Focus Styles Required for All Interactive Elements (MANDATORY)
- **Every interactive element must have a visible `:focus-visible` style** defined in `style.css`.
- This includes: `<button>`, `<a>`, `<input>`, `<textarea>`, `<select>`, tabs, and any `div` with `onclick`.
- Use `outline: 2px solid var(--primary-color); outline-offset: 2px;` as the minimum focus style.
- Never use `outline: none` without providing an alternative focus indicator.
- This is required for WCAG 2.1 SC 2.4.7 (Focus Visible) compliance.

### 47. No Implicit Global Event Object in JavaScript (MANDATORY)
- **Never reference `window.event` or the implicit `event` global** inside named functions.
- `window.event` is deprecated and does not exist in Firefox.
- Always pass the event or the target element explicitly as a function parameter.
- Example violation: `function copyItem(index) { const el = event.currentTarget; }`
- Example fix: `function copyItem(index, element) { ... }` called as `onclick="copyItem(0, this)"`


## Notes

- The server runs on port 8080 by default
- Templates are parsed using `template.ParseGlob`
- Static files are served from the `static/` directory
- All paths must use `filepath.Join()` for cross-platform compatibility

