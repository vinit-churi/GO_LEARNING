# Writer-Based Reporting Reference

## Concept Explanations

In idiomatic Go, interfaces are often consumed rather than declared up front. Accepting a small standard interface like `io.Writer` makes code easy to test and reuse.

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

- Prefer parameters like `io.Writer` or `io.Reader` when behavior matters more than the concrete type.
- Small interfaces are easier to satisfy and easier to mock in tests.
- Use `bytes.Buffer` when you need an in-memory implementation of `io.Writer`.
- Do not create a custom interface if a standard library one already communicates the contract well.
- Common mistake: returning interfaces from constructors when a concrete type would be simpler.
- Official docs: `io`, `bytes`, and Effective Go interface guidance.
