# Time Formatting and Deadlines Reference

## Concept Explanations

Time handling is easier when you normalize to explicit locations and use the standard `time` APIs. Durations are typed values, which makes arithmetic readable and safe.

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

- Use `time.Time` for timestamps and `time.Duration` for spans of time.
- Go uses example-based layouts like `time.RFC3339` instead of `%Y` formatting verbs.
- Tests should pick explicit UTC times to avoid timezone-dependent failures.
- Use `time.Until` only when you truly want wall-clock dependency; direct subtraction is better for deterministic code.
- Common mistake: assuming durations are in seconds when the value is actually nanoseconds.
- Official docs: `time` package.
