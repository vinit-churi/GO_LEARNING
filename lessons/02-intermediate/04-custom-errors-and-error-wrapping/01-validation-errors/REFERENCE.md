# Validation Errors Reference

## Concept Explanations

Errors in Go are values. A custom error type can carry useful context, and wrapping preserves the original cause so callers can inspect it without string matching.

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

- Implement the `error` interface by defining an `Error() string` method.
- Wrap errors with `fmt.Errorf("context: %w", err)` so callers can still inspect the original cause.
- Use `errors.As` for custom types and `errors.Is` for sentinel values.
- Avoid comparing error strings in tests when type or wrapped identity is what you actually care about.
- Custom errors are useful when callers need structured context such as a failing field or operation.
- Official docs: `errors` package and Go 1.13 error wrapping notes.
