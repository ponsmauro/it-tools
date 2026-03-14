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
- **Separate PR per tool**: `blackboxai/{tool-id}` branch → `debelop`.
- Structure:
  - `internal/tools/{id}/` package: handler.go, template.tmpl (embedded), tool.go (models.Tool).
  - Registration: `tool_registry.Register("{id}", ToolHandlerFunc, Base64Tool)`.
  - POST `/tools/{id}`: `{ "action": "open" }` → render tab content.
- Benefits: Isolation, lazy load, multi-tab workflow, easy PR review.
- Frontend: index.html JS manages tabs (localStorage persist open tabs).

## Notes

- The server runs on port 8080 by default
- Templates are parsed using `template.ParseGlob`
- Static files are served from the `static/` directory
- All paths must use `filepath.Join()` for cross-platform compatibility

