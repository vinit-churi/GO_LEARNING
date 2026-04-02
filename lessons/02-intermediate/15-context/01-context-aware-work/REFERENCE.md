# Context-Aware Work Reference

## Concept Explanations

Contexts carry deadlines, cancellation, and request-scoped values across API boundaries. They are not a generic parameter bag; they are for lifetime control and small metadata.

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

- Take `context.Context` as the first parameter when a function performs request-scoped work.
- Use `ctx.Done()` and `ctx.Err()` to observe cancellation and deadlines.
- Use custom unexported key types for context values to avoid collisions.
- Do not store optional function parameters or long-lived application state in a context.
- Common mistake: ignoring cancellation in code that blocks on channels or I/O.
- Official docs: `context` package and Go blog articles on context.
