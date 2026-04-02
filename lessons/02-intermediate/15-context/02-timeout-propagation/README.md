# Timeout Propagation

## Difficulty
Intermediate

## Prerequisites
- Context basics
- Goroutines
- Select statements

## Learning Objectives
- Create child contexts with deadlines
- Stop work when cancellation is signaled
- Compose timeout setup with worker orchestration

## Concept Summary

A context should be passed from the caller to all work that might block. When a timeout expires, all dependent operations should observe cancellation quickly and return a useful error.

## Run Instructions

```bash
go run starter.go
```

## Test Instructions

```bash
go test -v
```

## Exercises

### Exercise 1: Create Request Timeout
Implement `WithRequestTimeout(parent context.Context, d time.Duration)` to derive a timed child context.

### Exercise 2: Respect Cancellation
Implement `WaitForWork(ctx context.Context, workDone <-chan struct{}) error` so it returns `ctx.Err()` if canceled first.

### Exercise 3: Compose Timeout + Work
Implement `RunWithTimeout(parent context.Context, d time.Duration, workDone <-chan struct{}) error` to wire both helpers.
