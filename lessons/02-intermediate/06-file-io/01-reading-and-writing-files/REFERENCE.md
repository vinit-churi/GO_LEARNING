# Reading and Writing Files Reference

## Concept Explanations

Go keeps common file operations simple. For small files, `os.ReadFile` and `os.WriteFile` are often enough, especially in tests and small tools.

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

- Use `os.WriteFile` and `os.ReadFile` for small-file workflows.
- Tests should use `t.TempDir()` so file operations stay isolated and self-cleaning.
- Trim file contents only when your program explicitly wants normalized input.
- For large streams, reach for `os.Open` with buffered readers and writers instead of loading the whole file.
- Common mistake: hardcoding absolute paths in tests or examples.
- Official docs: `os`, `path/filepath`, and `testing` temporary directories.
