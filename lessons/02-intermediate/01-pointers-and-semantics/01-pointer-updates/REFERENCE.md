# Pointer Updates Reference

## Concept Explanations

Pointers let a function or method work with the original value instead of a copy. They matter when you need mutation, want to avoid copying larger structs, or need method sets that expose in-place updates.

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

- Use `*T` when a function must mutate a caller-owned value.
- Passing a struct by value copies the fields, so later writes only affect the local copy.
- Pointer receivers are common when methods update state or when the struct is large enough that copying is undesirable.
- Common mistake: mixing value and pointer receivers without a reason, which can make method sets harder to reason about.
- Use pointers for shared mutable state, but do not use them just to imitate class-based designs.
- Official docs: see the Go tour sections on pointers and methods.
