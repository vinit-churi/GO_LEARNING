# Measuring String Builders

## Difficulty
Advanced

## Prerequisites
- Benchmarks and Fuzzing
- Standard Library Essentials

## Learning Objectives
- Use `strings.Builder` for repeated concatenation
- Keep optimization claims anchored to benchmarks
- Compare readable implementations before reaching for lower-level tricks

## Concept Summary

Optimization work should start with measurement. In Go, a small benchmark around a hot helper is often enough to compare two straightforward implementations before making any broader claims.

## Run Instructions

```bash
go run starter.go
```

## Test Instructions

```bash
go test -v
```

## Exercises

### Exercise 1: Join With Builder
Implement `JoinWithBuilder(parts []string) string`.

### Exercise 2: Join With Plus
Implement `JoinWithPlus(parts []string) string` for comparison.

### Exercise 3: Report Line
Implement `ReportLine(parts []string) string` using the builder helper.
