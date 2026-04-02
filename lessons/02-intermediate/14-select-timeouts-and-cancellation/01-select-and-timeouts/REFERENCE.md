# Select and Timeouts Reference

## Concept Explanations

The `select` statement lets concurrent code react to whichever communication happens first. It is especially useful for timeouts and early exits.

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

- Use `select` when a goroutine needs to wait on more than one channel operation.
- A `default` branch makes a `select` non-blocking.
- Use `time.After` for simple one-off timeouts; for repeated use in hot paths, a reusable timer is often better.
- Timeout handling should be part of the API contract, not an unspoken side effect.
- Common mistake: forgetting that a `default` case can cause busy looping if used in a `for` without delay.
- Official docs: `select` and `time` package timers.
