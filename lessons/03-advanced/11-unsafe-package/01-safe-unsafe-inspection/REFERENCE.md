# Safe Unsafe Inspection Reference

## Concept Explanations

The `unsafe` package is for low-level interop and layout inspection, not routine application logic. The safest way to teach it is to focus on read-only facts before any pointer conversions.

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

- Inspecting layout is much safer than mutating memory through converted pointers.
- Most Go code should never need `unsafe` at all.
- If a package uses `unsafe`, keep it small, documented, and heavily tested.
- Official docs: `unsafe` package and package documentation warnings.
