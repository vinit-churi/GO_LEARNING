# JSON Struct Tags Reference

## Concept Explanations

The `encoding/json` package maps between JSON documents and Go structs. Tags let you control names and optional output while keeping the Go type expressive.

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

- Use exported struct fields for JSON marshaling and unmarshaling.
- Field tags like ``json:"service"`` map Go field names to JSON keys.
- Use `omitempty` when zero values should be left out of output.
- Do not hand-build JSON strings; let the encoder handle quoting and escaping.
- Common mistake: forgetting that unexported fields are ignored by `encoding/json`.
- Official docs: `encoding/json` package documentation.
