# Context - Dark Mode

## Current State
- The application has been migrated to dark mode.
- CSS variables in `static/css/style.css` have been updated.
- Inline styles in all HTML templates (`static/templates/*.html`) have been updated to use dark mode colors.
- The user requested unit tests to verify the dark mode implementation.

## Constraints
- Architecture constraints: Go unit tests cannot render HTML/CSS visually.
- Testing constraints: We need to parse the HTML files and verify the presence of dark mode hex codes and the absence of light mode hex codes.

## Tech Stack
- Backend: Go
- Frontend: HTML/CSS
- Testing: Go `testing` package
