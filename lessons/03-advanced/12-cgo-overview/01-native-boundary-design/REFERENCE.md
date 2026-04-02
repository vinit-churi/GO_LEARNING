# Native Boundary Design Reference

## Concept Explanations

Actual cgo usage depends on local toolchains and platform details, so this lesson focuses on the boundary design you should have before enabling it. The key idea is to isolate native calls behind a small Go-facing seam.

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

- Before using cgo, decide what the smallest possible Go-facing API should be.
- Keep cgo code in a narrow package so callers do not inherit native complexity or build requirements.
- Tests should prefer a pure-Go fake when the goal is behavior, not ABI verification.
- Official docs: cgo overview and Go toolchain docs.
