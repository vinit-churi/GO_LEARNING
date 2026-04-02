# Safe Counter Reference

## Concept Explanations

Shared mutable state needs coordination. Mutexes protect composite state, while atomic operations fit narrow cases like independent counters.

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

- Use `sync.Mutex` when multiple goroutines access shared composite state.
- Use `sync/atomic` for very small pieces of independent numeric state.
- Race detection is a runtime tool: `go test -race ./...` helps catch unsynchronized access.
- A mutex usually communicates intent more clearly than atomics when several fields change together.
- Common mistake: mixing atomic and non-atomic access to the same variable.
- Official docs: `sync`, `sync/atomic`, and Go race detector docs.
