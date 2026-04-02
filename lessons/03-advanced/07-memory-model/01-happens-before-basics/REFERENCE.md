# Happens-Before Basics Reference

## Concept Explanations

The Go memory model explains which synchronization operations make writes visible across goroutines. The practical goal is not memorizing theory, but knowing which tools establish safe ordering.

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

- Channel send/receive and mutex operations establish ordering that plain reads and writes do not.
- A `sync.Once` call guarantees one initialization and safe publication of the resulting state.
- Sleeping is not synchronization.
- Official docs: Go memory model and `sync` package.
