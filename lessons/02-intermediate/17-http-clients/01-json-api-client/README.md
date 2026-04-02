# JSON API Client

## Difficulty
Intermediate

## Prerequisites
- JSON Struct Tags
- Context-Aware Work

## Learning Objectives
- Build an HTTP request with context
- Decode JSON responses with `json.Decoder`
- Test client behavior with `httptest.Server`

## Concept Summary

HTTP client code should be explicit about request creation, status handling, and response decoding. `httptest.Server` makes these flows easy to test without external network calls.

## Run Instructions

```bash
go run starter.go
```

## Test Instructions

```bash
go test -v
```

## Exercises

### Exercise 1: Fetch Status
Implement `FetchStatus(ctx context.Context, client *http.Client, url string) (APIStatus, error)`.

### Exercise 2: Status Label
Implement `StatusLabel(status APIStatus) string`.

### Exercise 3: Fetch Status Label
Implement `FetchStatusLabel(ctx context.Context, client *http.Client, url string) (string, error)` using the other helpers.
