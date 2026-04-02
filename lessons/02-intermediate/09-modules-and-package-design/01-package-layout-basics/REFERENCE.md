# Package Layout Basics Reference

## Concept Explanations

Package design in Go is driven by simplicity at the call site. Small packages with a narrow exported surface are easier to understand than large grab-bag utility packages.

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

- A package should have one clear reason to exist from the caller perspective.
- Export only what other packages need; keep helper details unexported where possible.
- Names should read well at the use site, such as `slugutil.Slug(title)` rather than a vague utility package.
- Avoid circular dependencies by keeping packages small and layered.
- Common mistake: creating a `utils` package that grows without a clear boundary.
- Official docs: Go blog posts on package names and modules.
