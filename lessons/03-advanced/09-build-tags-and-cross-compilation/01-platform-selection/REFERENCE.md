# Platform Selection Reference

## Concept Explanations

Build tags let Go choose files for specific environments at compile time. They are useful for platform differences, but they should remain narrow and easy to reason about.

## Syntax Patterns

- Keep low-level mechanics behind small helpers or types.
- Prefer measurement and explicit contracts over folklore-driven decisions.
- Reach for advanced features only when simpler code is clearly insufficient.

## Common Mistakes

- Generalizing too early without a concrete repeated use case.
- Hiding complexity behind clever APIs that are harder to debug.
- Treating advanced features as defaults instead of specialized tools.

## Why This Feature Exists

It solves specific scaling, abstraction, interoperability, or performance problems that simpler language features cannot address as clearly.

## When To Use It

Use it when the tradeoff is justified by clearer reuse, better measurement, or a real system constraint.

## When Not To Use It

Do not add it just because it is available or because it looks more sophisticated than a direct approach.

## Notes

- Put the smallest possible amount of code behind build tags.
- Shared logic should stay in ordinary files so most of the package remains testable on every platform.
- Cross-compilation is usually driven by `GOOS` and `GOARCH`, but your code should not need many runtime checks for that.
- Official docs: Go build constraints and cross-compilation docs.
