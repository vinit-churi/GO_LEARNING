# Embedded Services Reference

## Concept Explanations

Go favors composition over inheritance. Embedding lets one struct reuse fields and methods from another while keeping the relationship explicit in the type definition.

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

- Composition means building larger types from smaller focused types.
- Embedding a field without a name promotes its fields and methods, so `service.Owner` can refer to `service.AuditInfo.Owner`.
- Embedding is useful when the outer type should naturally expose the inner behavior.
- Use a named field instead of embedding when you want a clearer boundary or when the inner API should stay explicit.
- Common mistake: embedding just to save typing when the relationship is not conceptually strong.
- Official docs: see Effective Go on embedding.
