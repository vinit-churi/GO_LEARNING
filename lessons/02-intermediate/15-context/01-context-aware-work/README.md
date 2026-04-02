# Context-Aware Work

## Difficulty
Intermediate

## Prerequisites
- Select, Timeouts, and Cancellation

## Learning Objectives
- Accept `context.Context` as the first parameter for request-scoped work
- Stop work early when cancellation happens
- Use context values sparingly for metadata

## Concept Summary

Contexts carry deadlines, cancellation, and request-scoped values across API boundaries. They are not a generic parameter bag; they are for lifetime control and small metadata.

## Run Instructions

```bash
go run starter.go
```

## Test Instructions

```bash
go test -v
```

## Exercises

### Exercise 1: Trace ID
Implement `TraceID(ctx context.Context) string` using a typed context key.

### Exercise 2: Work Or Cancel
Implement `WorkOrCancel(ctx context.Context) error` so it returns `context.Canceled` when the context is done first.

### Exercise 3: Label Request
Implement `LabelRequest(ctx context.Context, name string) string` so it prefixes the name with the trace id.
