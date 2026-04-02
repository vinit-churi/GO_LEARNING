# Mapping With Type Parameters

## Difficulty
Advanced

## Prerequisites
- Functions
- Slices
- Intermediate Standard Library Work

## Learning Objectives
- Define generic functions with type parameters
- Reuse one algorithm across multiple concrete types
- Keep generic APIs small and readable

## Concept Summary

Generics let you write reusable functions and types without giving up static typing. They are most useful when the logic is truly the same across many element types.

## Run Instructions

```bash
go run starter.go
```

## Test Instructions

```bash
go test -v
```

## Exercises

### Exercise 1: Map Slice
Implement `MapSlice[T, U any](items []T, fn func(T) U) []U`.

### Exercise 2: First Or Zero
Implement `FirstOrZero[T any](items []T) T` so it returns the first value or the zero value.

### Exercise 3: Pair Labels
Implement `PairLabels[T any](items []T, label func(T) string) []string` using `MapSlice`.
