# Intro To Testing

## Difficulty
Beginner

## Prerequisites
- Functions
- Basic errors

## Learning Objectives
- Understand how `go test` runs test files
- Read table-style expectations in tests
- Improve function behavior when tests fail

## Concept Summary

Go looks for files ending in `_test.go` and runs functions that start with `Test`. Tests give you a repeatable way to check logic and catch regressions.

## Run Instructions

```bash
go run starter.go
```

## Test Instructions

```bash
go test -v
```

## Exercises

### Exercise 1: Is Even
Implement `IsEven(n int) bool`.

### Exercise 2: Grade Label
Implement `GradeLabel(score int) string`.

### Exercise 3: Average
Implement `Average(values []int) float64`.
