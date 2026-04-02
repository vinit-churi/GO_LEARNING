# Context-Aware Runner Reference

## Concept Explanations

This mini-project combines context cancellation with a small runner abstraction. The emphasis is on making stop conditions explicit and easy to test.

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

- Context-aware code should make cancellation observable in the function result.
- A small wrapper type or helper can keep lifecycle rules centralized.
- Do not start background work you cannot stop or account for.
