# Inspecting Structs Reference

## Concept Explanations

Reflection lets code inspect types and values at runtime. It is powerful, but it trades away compile-time clarity, so it should usually sit behind a narrow helper instead of spreading through a codebase.

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

- Reflection is often best for serializers, validators, test helpers, and framework boundaries.
- A reflected value may represent a pointer, so unwrap it deliberately when needed.
- Prefer direct field access when the types are known at compile time.
- Official docs: `reflect` package.
