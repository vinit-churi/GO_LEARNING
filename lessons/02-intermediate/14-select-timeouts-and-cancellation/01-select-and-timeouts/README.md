# Select and Timeouts

## Difficulty
Intermediate

## Prerequisites
- Channels
- Time Formatting and Deadlines

## Learning Objectives
- Wait on multiple channel operations with `select`
- Use `time.After` for simple timeouts
- Model timeout behavior explicitly in function results

## Concept Summary

The `select` statement lets concurrent code react to whichever communication happens first. It is especially useful for timeouts and early exits.

## Run Instructions

```bash
go run starter.go
```

## Test Instructions

```bash
go test -v
```

## Exercises

### Exercise 1: Receive With Timeout
Implement `ReceiveWithTimeout(ch <-chan string, timeout time.Duration) string`.

### Exercise 2: Immediate Status
Implement `ImmediateStatus(ch <-chan string) string` so it returns `idle` when no value is ready.

### Exercise 3: First Result
Implement `FirstResult(left, right <-chan string) string` so it returns whichever value arrives first.
