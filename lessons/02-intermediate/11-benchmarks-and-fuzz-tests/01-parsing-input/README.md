# Parsing Input

## Difficulty
Intermediate

## Prerequisites
- Table-Driven Tests
- Standard Library Essentials

## Learning Objectives
- Write benchmark functions for hot paths
- Add a fuzz target for parser robustness
- Keep parsing logic deterministic and side-effect free

## Concept Summary

Benchmarks measure performance trends, and fuzz tests push random inputs through code to discover edge cases. Both work best when the target function is small and deterministic.

## Run Instructions

```bash
go run starter.go
```

## Test Instructions

```bash
go test -v
```

## Exercises

### Exercise 1: Parse Pair
Implement `ParsePair(input string) (string, int, error)` for `name=count` input.

### Exercise 2: Must Parse Count
Implement `MustParseCount(input string) int` so it returns zero on bad input.

### Exercise 3: Format Pair
Implement `FormatPair(name string, count int) string` so it rebuilds the canonical form.
