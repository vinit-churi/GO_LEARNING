# Concurrent Workers Reference

## Concept Explanations

A goroutine is a lightweight concurrent function. The hard part is coordinating results and knowing when all launched work is done.

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

- Use `go` to launch a concurrent function call.
- Use `sync.WaitGroup` to wait until a fixed set of goroutines has finished.
- Preserving input order usually requires storing results by index rather than appending from multiple goroutines.
- Do not capture loop variables carelessly; pass values into the goroutine or shadow them first.
- Common mistake: writing to shared slices without stable indexing or synchronization.
- Official docs: `sync` package and the Go blog on concurrency.
