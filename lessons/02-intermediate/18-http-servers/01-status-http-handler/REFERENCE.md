# Status HTTP Handler Reference

## Concept Explanations

Go HTTP servers are mostly ordinary functions over `http.ResponseWriter` and `*http.Request`. Keeping handlers small makes them easy to test and evolve.

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

- Handlers can be plain functions with the signature `func(http.ResponseWriter, *http.Request)`.
- Set response headers before writing the body.
- Use `httptest.NewRecorder` and synthetic requests to unit test handlers directly.
- Keep routing and response formatting separate when possible so handler code stays small.
- Common mistake: writing headers after the body has already been written.
- Official docs: `net/http`, `httptest`, and `encoding/json`.
