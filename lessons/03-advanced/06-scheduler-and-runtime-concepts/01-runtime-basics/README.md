# Runtime Basics

## Difficulty
Advanced

## Prerequisites
- Goroutines
- Channels

## Learning Objectives
- Read basic runtime information from the standard library
- Make conservative worker-count decisions from CPU data
- Separate measurement-oriented helpers from business logic

## Concept Summary

Scheduler and runtime concepts matter when you need to explain concurrency behavior or tune throughput. They should guide observation and measurement, not turn into folklore-driven tweaks.

## Run Instructions

```bash
go run starter.go
```

## Test Instructions

```bash
go test -v
```

## Exercises

### Exercise 1: CPU Count
Implement `CPUCount() int` with `runtime.NumCPU`.

### Exercise 2: Choose Workers
Implement `ChooseWorkers(limit int) int` so it caps worker count by CPU count.

### Exercise 3: Parallel Summary
Implement `ParallelSummary(limit int) string` to report CPUs and chosen workers.
