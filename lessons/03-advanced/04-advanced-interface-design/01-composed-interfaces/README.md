# Composed Interfaces

## Difficulty
Advanced

## Prerequisites
- Interfaces in Practice
- HTTP Clients

## Learning Objectives
- Compose behavior from small interfaces
- Accept interfaces at the call site instead of returning them broadly
- Separate validation from persistence concerns

## Concept Summary

Advanced interface design in Go is usually about making interfaces smaller, more focused, and closer to where they are used. Composition works best when each part expresses a clear responsibility.

## Run Instructions

```bash
go run starter.go
```

## Test Instructions

```bash
go test -v
```

## Exercises

### Exercise 1: Save Reader
Implement `SaveReader(store ReaderStore, id string) error` using the composed interface.

### Exercise 2: Upsert Label
Implement `UpsertLabel(store ReaderStore, id string, value string) string`.

### Exercise 3: Copy If Missing
Implement `CopyIfMissing(store ReaderStore, src, dst string) error`.
