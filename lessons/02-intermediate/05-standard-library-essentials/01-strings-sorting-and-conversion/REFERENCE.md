# Strings, Sorting, and Conversion Reference

## Concept Explanations

A large part of idiomatic Go is knowing the standard library well. Small packages like `strings`, `strconv`, and `slices` solve common problems directly and keep custom code short.

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

- Reach for standard library helpers before writing loops by hand.
- Use `strings.TrimSpace` and `strings.ToLower` for predictable text cleanup.
- Use `strconv.Atoi` or `strconv.ParseInt` to parse numeric input.
- Use `slices.Sort` for in-place ascending sorts of ordered slices.
- Common mistake: ignoring parse errors and silently treating bad input as zero.
- Official docs: `strings`, `strconv`, and `slices`.
