# Golden and Helper Patterns

## Difficulty
Advanced

## Prerequisites
- Intro to Testing
- Table-Driven Tests

## Learning Objectives
- Normalize formatted outputs before comparing them
- Use helper functions to keep tests readable
- Choose stable assertions instead of brittle string checks where possible

## Concept Summary

Advanced testing strategy is usually about choosing the right style for the problem: table-driven cases for combinations, helpers for common assertions, and stable golden-style outputs for larger formatted results.

## Run Instructions

```bash
go run starter.go
```

## Test Instructions

```bash
go test -v
```

## Exercises

### Exercise 1: Normalize Report
Implement `NormalizeReport(input string) string` so spacing and trailing newlines are stable.

### Exercise 2: Build Golden
Implement `BuildGolden(name string, count int) string`.

### Exercise 3: Reports Match
Implement `ReportsMatch(left, right string) bool` using normalization.
