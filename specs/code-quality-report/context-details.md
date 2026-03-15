# Context - Code Quality Report Details

## Goal
Enhance the existing code quality report (`generate_report.py`) to include detailed explanations and context for the scores, so the user knows exactly what to improve.

## Requirements
- Add a "Recommendations" or "Details" section for both Frontend and Backend.
- For Frontend: Explain why inline styles and embedded scripts/styles are penalized (e.g., separation of concerns, maintainability, CSP). List the files that contain these issues if possible, or at least provide actionable advice.
- For Backend: Explain the importance of test coverage and `go vet` issues. Provide actionable advice on how to improve the score (e.g., write more tests, fix vet warnings).
- Update the HTML template to display these details clearly.

## Approach
1. Modify `generate_report.py` to gather more specific data (e.g., which files have inline styles).
2. Add logic to generate actionable recommendations based on the gathered metrics.
3. Update the HTML template to include a new section within each tab for "Recommendations" or "Areas for Improvement".
