# Context - Base64 Encoder/Decoder

## Current State
- Relevant existing files/routes/components:
  - Repository already includes base64 in tool list (ID: "base64")
  - No existing implementation in static/templates/
  - Similar tools: html-escape.html, url-encoder.html (for UI patterns)

## Constraints
- Architecture constraints:
  - Pure client-side implementation (no backend processing needed)
  - Must follow template pattern like other tools
  - Should use native browser APIs (btoa/atob)
- Security constraints:
  - Large files should be handled client-side only
  - No server-side processing of user data
- UX constraints:
  - Should support both text and file input
  - Must provide clear feedback on encoding/decoding
  - Copy buttons for easy result transfer

## Tech Stack
- Backend: None needed (pure client-side)
- Frontend: HTML, CSS, JavaScript (btoa/atob)
- i18n: Template variables for all user-facing text
- Testing: Manual UI testing
