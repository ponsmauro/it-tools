# Plan: Standardize Title Sizes

1. **Update `static/css/style.css`**
   - Add global styles for `h1` and `h2` inside the `.main` or `.container` context, or globally.
   - `h1 { font-size: 18px; margin-bottom: 8px; font-weight: 600; color: var(--text-primary); }`
   - `h2 { font-size: 16px; margin-bottom: 6px; font-weight: 600; color: var(--text-primary); }`

2. **Remove inline `h1` and `h2` styles from templates**
   - `static/templates/cron-parser.html`
   - `static/templates/case-converter.html`
   - `static/templates/base64.html`
   - `static/templates/csv-parser.html`
   - `static/templates/env-formatter.html`
   - `static/templates/gzip-tool.html`
   - `static/templates/ip-calculator.html`
   - `static/templates/uuid-generator.html`
   - `static/templates/text-utils.html`
   - `static/templates/diff-checker.html`
   - `static/templates/password-generator.html`
   - `static/templates/html-escape.html`
   - `static/templates/kill-port.html`
   - `static/templates/timestamp-converter.html`
   - `static/templates/lorem-generator.html`
   - `static/templates/url-encoder.html`

3. **Commit and Push**
   - Commit message: `style: standardize h1 and h2 font sizes across all tools`
   - Push to `feature/title-sizes`
   - Create PR to `develop`
