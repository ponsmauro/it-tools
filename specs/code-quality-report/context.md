# Context - Code Quality Report

## Goal
Generate a code quality report for the `it-tools` project and save it as an HTML file at `/Users/mpons/git-repo/ponsmauro/it-tools-code-quality-gemini.html`.

## Requirements
- The report must display a total scoring for the app.
- It must include separate scorings for Frontend and Backend.
- The UI should have two tabs: one for Frontend and one for Backend, each displaying their respective scoring and details.

## Approach
1. Analyze the project structure to determine what constitutes Frontend (HTML, CSS, JS in `static/`) and Backend (Go code in `cmd/`, `internal/`).
2. Run static analysis tools or scripts to gather metrics (e.g., `golangci-lint` for Go, basic heuristics for JS/CSS/HTML).
3. Generate an HTML file with a modern, clean UI (using the project's existing design language if possible, or a standalone clean design) containing the tabs and scores.
4. Save the file to the specified absolute path.
