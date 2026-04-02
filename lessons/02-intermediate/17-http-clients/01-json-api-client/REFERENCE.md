# JSON API Client Reference

## Concept Explanations

HTTP client code should be explicit about request creation, status handling, and response decoding. `httptest.Server` makes these flows easy to test without external network calls.

## Syntax Patterns

- Keep helpers small and name them for the caller.
- Prefer standard library packages when they already model the behavior you need.
- Use tests to lock down edge cases and make the intended contract explicit.

## Common Mistakes

- Skipping error handling when the lesson API returns an `error`.
- Hiding side effects that should be visible in the function signature.
- Repeating logic instead of composing the helper functions from the lesson.

## Why This Feature Exists

It gives Go programs a small, explicit tool for this part of application design without needing framework magic.

## When To Use It

Use it when the code is clearer, safer, or more testable with this feature.

## When Not To Use It

Do not force the pattern when a simpler sequential or concrete approach would be easier to understand.

## Notes

- Use `http.NewRequestWithContext` so request cancellation follows the caller context.
- Check `resp.StatusCode` before decoding when non-200 responses need special handling.
- Prefer `json.NewDecoder(resp.Body)` for streaming decode from HTTP responses.
- Use `httptest.Server` in tests instead of real network dependencies.
- Common mistake: forgetting to close `resp.Body`.
- Official docs: `net/http`, `httptest`, and `encoding/json`.
