# Runtime Basics Reference

## Concept Explanations

Scheduler and runtime concepts matter when you need to explain concurrency behavior or tune throughput. They should guide observation and measurement, not turn into folklore-driven tweaks.

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

- Scheduler behavior is observable through tools and APIs like `runtime.NumCPU`, traces, and profiles.
- Simple worker heuristics should be conservative and explicit, not magical.
- Do not assume more goroutines always means more throughput.
- Official docs: `runtime` package and Go execution tracer docs.
