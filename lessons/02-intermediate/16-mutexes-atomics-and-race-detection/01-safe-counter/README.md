# Safe Counter

## Difficulty
Intermediate

## Prerequisites
- Context-Aware Work
- Goroutines

## Learning Objectives
- Protect shared state with a mutex
- Use atomic values for simple counters
- Understand when `go test -race` is the right validation tool

## Concept Summary

Shared mutable state needs coordination. Mutexes protect composite state, while atomic operations fit narrow cases like independent counters.

## Run Instructions

```bash
go run starter.go
```

## Test Instructions

```bash
go test -v
```

## Exercises

### Exercise 1: Increment Safely
Implement `SafeCounter.Inc()` using a mutex.

### Exercise 2: Read Count
Implement `SafeCounter.Value() int` safely.

### Exercise 3: Increment Atomic
Implement `IncrementAtomic(counter *int64)` using `sync/atomic`.
