# Generic Dedup and Frequency Reference

## Concept Explanations

A type parameter with `comparable` allows equality checks and map-key usage. This fits deduplication and frequency counting patterns.

## Syntax Patterns

- Define generic helpers as `func Name[T comparable](items []T) ...`.
- Use `map[T]struct{}` for a compact set.
- Use `map[T]int` for counting.

## Common Mistakes

- Using unconstrained `any` when map keys are needed.
- Forgetting to preserve order in dedup helpers.
- Returning ambiguous zero values without a boolean indicator.

## Why This Feature Exists

Generics remove repeated implementations for each concrete type while retaining compile-time safety.

## When To Use It

Use generic helpers for stable algorithms shared by many types.

## When Not To Use It

Avoid unnecessary abstraction when a single concrete use is clearer.
