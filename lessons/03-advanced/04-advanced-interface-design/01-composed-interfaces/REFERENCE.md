# Composed Interfaces Reference

## Concept Explanations

Advanced interface design in Go is usually about making interfaces smaller, more focused, and closer to where they are used. Composition works best when each part expresses a clear responsibility.

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

- Declare interfaces where they are consumed, not in distant packages without a caller in mind.
- A composed interface should still be easy to understand from its name and methods.
- Do not create giant "manager" interfaces just to avoid passing two collaborators.
- Official docs: Effective Go interface advice.
