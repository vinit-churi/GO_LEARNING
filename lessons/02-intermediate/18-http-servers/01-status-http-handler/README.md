# Status HTTP Handler

## Difficulty
Intermediate

## Prerequisites
- HTTP Clients
- JSON Struct Tags

## Learning Objectives
- Implement an `http.HandlerFunc`
- Write JSON responses with explicit headers
- Test handlers with `httptest.NewRecorder`

## Concept Summary

Go HTTP servers are mostly ordinary functions over `http.ResponseWriter` and `*http.Request`. Keeping handlers small makes them easy to test and evolve.

## Run Instructions

```bash
go run starter.go
```

## Test Instructions

```bash
go test -v
```

## Exercises

### Exercise 1: Write JSON
Implement `WriteJSON(w http.ResponseWriter, status APIStatus)` so it sets the content type and writes JSON.

### Exercise 2: Status Handler
Implement `StatusHandler(service string) http.HandlerFunc` to return a handler that reports an `ok` status.

### Exercise 3: Route Label
Implement `RouteLabel(r *http.Request) string` using the request method and path.
