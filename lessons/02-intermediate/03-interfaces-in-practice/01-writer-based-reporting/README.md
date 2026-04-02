# Writer-Based Reporting

## Difficulty
Intermediate

## Prerequisites
- Interfaces Basics
- Composition and Embedding

## Learning Objectives
- Use standard library interfaces instead of concrete dependencies
- Inject an `io.Writer` for flexible output
- Keep interfaces small and usage-driven

## Concept Summary

In idiomatic Go, interfaces are often consumed rather than declared up front. Accepting a small standard interface like `io.Writer` makes code easy to test and reuse.

## Run Instructions

```bash
go run starter.go
```

## Test Instructions

```bash
go test -v
```

## Exercises

### Exercise 1: Write Line
Implement `WriteLine(w io.Writer, line string) error` so it writes the line with a trailing newline.

### Exercise 2: Emit Report
Implement `EmitReport(w io.Writer, service string, count int) error` using `WriteLine`.

### Exercise 3: Buffer Report
Implement `BufferReport(service string, count int) string` so it captures the report in memory.
