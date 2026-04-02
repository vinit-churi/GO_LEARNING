# Concurrent Workers

## Difficulty
Intermediate

## Prerequisites
- Functions
- Table-Driven Tests

## Learning Objectives
- Launch work with goroutines
- Wait for completion with `sync.WaitGroup`
- Keep concurrent code deterministic in tests by collecting results explicitly

## Concept Summary

A goroutine is a lightweight concurrent function. The hard part is coordinating results and knowing when all launched work is done.

## Run Instructions

```bash
go run starter.go
```

## Test Instructions

```bash
go test -v
```

## Exercises

### Exercise 1: Double Async
Implement `DoubleAsync(values []int) []int` so it doubles values concurrently and preserves input order.

### Exercise 2: Sum Result
Implement `Sum(values []int) int` for later aggregation checks.

### Exercise 3: Double And Sum
Implement `DoubleAndSum(values []int) int` using the concurrent helper.
