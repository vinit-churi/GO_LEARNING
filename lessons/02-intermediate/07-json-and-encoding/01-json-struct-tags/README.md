# JSON Struct Tags

## Difficulty
Intermediate

## Prerequisites
- Structs
- File I/O

## Learning Objectives
- Marshal structs into JSON with field tags
- Unmarshal JSON into Go values
- Model optional JSON fields with `omitempty`

## Concept Summary

The `encoding/json` package maps between JSON documents and Go structs. Tags let you control names and optional output while keeping the Go type expressive.

## Run Instructions

```bash
go run starter.go
```

## Test Instructions

```bash
go test -v
```

## Exercises

### Exercise 1: Encode Event
Implement `EncodeEvent(event Event) (string, error)` so it returns compact JSON.

### Exercise 2: Decode Event
Implement `DecodeEvent(input string) (Event, error)` using `json.Unmarshal`.

### Exercise 3: Summary Label
Implement `SummaryLabel(event Event) string` so it uses the decoded struct fields rather than parsing JSON directly.
