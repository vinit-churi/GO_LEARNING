# Generic Dedup and Frequency

## Difficulty
Advanced

## Prerequisites
- Generics fundamentals
- Maps
- Slices

## Learning Objectives
- Use type parameters with `comparable` constraints
- Build reusable collection utilities without reflection
- Preserve insertion order while tracking uniqueness

## Concept Summary

Generics let you write collection helpers once and reuse them for many concrete types. A `comparable` constraint is the right choice when values are map keys.

## Run Instructions

```bash
go run starter.go
```

## Test Instructions

```bash
go test -v
```

## Exercises

### Exercise 1: Unique
Implement `Unique[T comparable](items []T) []T` to return values in first-seen order.

### Exercise 2: Count By Value
Implement `CountBy[T comparable](items []T) map[T]int` to count occurrences.

### Exercise 3: Most Frequent
Implement `MostFrequent[T comparable](items []T) (T, bool)` returning the most frequent value and false for empty input.
