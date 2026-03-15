# Context: Standardize Title Sizes

## Current State
The `h1` and `h2` titles across different tools have inconsistent font sizes. Some tools like "Text Utils" have `h1` set to `18px`, while others like "Regex Tester" do not have explicit sizes and inherit the browser's default large sizes.

## Goal
Standardize the font sizes of `h1` and `h2` across all tools to follow the compact UI standard.
- `h1` should be `18px` with a small bottom margin.
- `h2` should be `16px` with a small bottom margin.

## Approach
1. Add global styles for `h1` and `h2` in `static/css/style.css`.
2. Remove inline `h1` and `h2` style overrides from all tool templates in `static/templates/*.html`.
