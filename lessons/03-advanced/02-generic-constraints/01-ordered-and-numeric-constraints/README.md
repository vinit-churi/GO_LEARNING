# Ordered and Numeric Constraints

## Difficulty
Advanced

## Prerequisites
- Generics Fundamentals

## Learning Objectives
- Define reusable constraints for numeric operations
- Use ordered constraints when comparison is required
- Keep constrained generic code readable

## Concept Summary

Constraints describe which operations a generic implementation is allowed to use. They should communicate the narrowest real requirement rather than acting like a catch-all type system shortcut.

## Run Instructions

```bash
go run starter.go
```

## Test Instructions

```bash
go test -v
```

## Exercises

### Exercise 1: Sum Numbers
Implement `SumNumbers[T Number](items []T) T`.

### Exercise 2: Max Ordered
Implement `MaxOrdered[T cmp.Ordered](items []T) T`.

### Exercise 3: Clamp Ordered
Implement `ClampOrdered[T cmp.Ordered](value, low, high T) T`.
