# Measuring String Builders Reference

## Concept Explanations

Optimization work should start with measurement. In Go, a small benchmark around a hot helper is often enough to compare two straightforward implementations before making any broader claims.

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

- A benchmark should isolate the code under test and avoid unrelated work in the measurement loop.
- Readable code still matters; a tiny speedup is not worth a permanently confusing helper.
- Profiles and benchmarks are evidence, not decorations.
- Official docs: `testing` benchmarks and `pprof` overview.
