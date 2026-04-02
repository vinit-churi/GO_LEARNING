# Ordered and Numeric Constraints Reference

## Concept Explanations

Constraints describe which operations a generic implementation is allowed to use. They should communicate the narrowest real requirement rather than acting like a catch-all type system shortcut.

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

- A constraint should reflect the operations the body actually needs: arithmetic, ordering, or method calls.
- If the implementation only compares values, `cmp.Ordered` is often enough.
- Overly broad custom constraints can make code harder to read than concrete helpers.
- Official docs: `cmp` package and generic constraints guidance.
