# In-Memory Status Service

## Difficulty
Advanced

## Prerequisites
- HTTP Servers
- Generic Safe Cache

## Learning Objectives
- Combine shared state with an HTTP handler
- Return JSON from a tiny service surface
- Use `httptest` to verify end-to-end behavior

## Concept Summary

This mini-project combines HTTP handling, JSON responses, and shared in-memory state. It stays intentionally small so the design tradeoffs remain visible.

## Run Instructions

```bash
go run starter.go
```

## Test Instructions

```bash
go test -v
```

## Exercises

### Exercise 1: Set Status
Implement `(*StatusService).SetStatus(service, status string)`.

### Exercise 2: Get Status
Implement `(*StatusService).GetStatus(service string) (string, bool)`.

### Exercise 3: Handler
Implement `(*StatusService).Handler() http.HandlerFunc`.
