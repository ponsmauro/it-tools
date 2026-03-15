# Plan - Code Quality Report Details

## Phase 1: Enhance Data Gathering
- [ ] Update `analyze_backend()` in `generate_report.py`:
  - Capture the actual output of `go vet` to list the specific issues.
  - Generate a list of recommendations based on coverage and vet issues.
- [ ] Update `analyze_frontend()` in `generate_report.py`:
  - Use `grep -l` to get the list of files containing inline styles, `<style>` tags, and `<script>` tags.
  - Generate a list of recommendations explaining why these are bad practices and how to fix them.

## Phase 2: Update HTML Template
- [ ] Modify `generate_html()` in `generate_report.py`:
  - Add a "Recommendations" section below the metrics in both the Frontend and Backend tabs.
  - Style the recommendations section to be clear and readable (e.g., using lists or cards).
  - Include the specific files that need attention for frontend issues.
  - Include the specific `go vet` output for backend issues.

## Phase 3: Execution & Delivery
- [ ] Run `python3 generate_report.py`.
- [ ] Verify the generated HTML file contains the new details and recommendations.
- [ ] Commit the changes and push to the feature branch.
- [ ] Open PR and merge.
