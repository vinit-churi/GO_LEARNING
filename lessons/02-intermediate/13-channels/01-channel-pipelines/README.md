# Channel Pipelines

## Difficulty
Intermediate

## Prerequisites
- Goroutines

## Learning Objectives
- Send values between goroutines with channels
- Close producer channels when work is complete
- Build small pipeline stages that compose cleanly

## Concept Summary

Channels let goroutines communicate through explicit message passing. Small pipeline stages can be easier to reason about than shared mutable state.

## Run Instructions

```bash
go run starter.go
```

## Test Instructions

```bash
go test -v
```

## Exercises

### Exercise 1: Feed Values
Implement `Feed(values []int) <-chan int` so it publishes each value then closes the channel.

### Exercise 2: Square Stream
Implement `SquareStream(in <-chan int) <-chan int` as a pipeline stage.

### Exercise 3: Collect
Implement `Collect(in <-chan int) []int` so tests can drain the pipeline result.
