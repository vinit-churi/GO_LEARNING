# Validation Errors

## Difficulty
Intermediate

## Prerequisites
- Errors Basics
- Interfaces in Practice

## Learning Objectives
- Define a custom error type with context
- Wrap lower-level errors with `%w`
- Use `errors.As` and `errors.Is` in tests

## Concept Summary

Errors in Go are values. A custom error type can carry useful context, and wrapping preserves the original cause so callers can inspect it without string matching.

## Run Instructions

```bash
go run starter.go
```

## Test Instructions

```bash
go test -v
```

## Exercises

### Exercise 1: Error Message
Implement `ValidationError.Error() string` with the field name and problem.

### Exercise 2: Validate Username
Implement `ValidateUsername(name string) error` and return `ValidationError` for blank input.

### Exercise 3: Load Username
Implement `LoadUsername(raw string) (string, error)` so it wraps validation failures with `%w`.
