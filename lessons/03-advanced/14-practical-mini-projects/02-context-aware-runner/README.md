# Context-Aware Runner

## Difficulty
Advanced

## Prerequisites
- Context
- Select and Timeouts

## Learning Objectives
- Wrap context-aware work in a small type
- Handle cancellation without leaking goroutines
- Expose a deterministic API for tests

## Concept Summary

This mini-project combines context cancellation with a small runner abstraction. The emphasis is on making stop conditions explicit and easy to test.

## Run Instructions

```bash
go run starter.go
```

## Test Instructions

```bash
go test -v
```

## Exercises

### Exercise 1: Run Once
Implement `RunOnce(ctx context.Context, work func() string) (string, error)`.

### Exercise 2: Label Result
Implement `LabelResult(value string) string`.

### Exercise 3: Run Labeled
Implement `RunLabeled(ctx context.Context, work func() string) (string, error)`.
