# Plan - Code Quality Report

## Phase 1: Data Gathering
- [ ] Run a script to analyze Backend (Go) files:
  - Count files, lines of code.
  - Run `go test -cover` to get test coverage.
  - Run `go vet` or similar to find issues.
  - Calculate a score based on coverage and issues.
- [ ] Run a script to analyze Frontend (HTML/CSS/JS) files:
  - Count files, lines of code.
  - Check for inline styles/scripts vs external.
  - Calculate a score based on best practices (e.g., separation of concerns).

## Phase 2: HTML Generation
- [ ] Create a Python script `generate_report.py` that:
  - Executes the analysis commands.
  - Parses the output.
  - Generates an HTML string with:
    - A header showing the Total Score (average of Front and Back).
    - A tabbed interface (Frontend | Backend).
    - Frontend tab: Score, metrics (files, LOC, inline styles count).
    - Backend tab: Score, metrics (files, LOC, test coverage, vet issues).
    - CSS for styling the report (dark mode, clean UI).
  - Writes the HTML string to `/Users/mpons/git-repo/ponsmauro/it-tools-code-quality-gemini.html`.

## Phase 3: Execution & Delivery
- [ ] Run `generate_report.py`.
- [ ] Verify the file is created at the correct location.
- [ ] Commit the script and specs to the repository.
- [ ] Open PR and merge.
