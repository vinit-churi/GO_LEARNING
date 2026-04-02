# Native Boundary Design

## Difficulty
Advanced

## Prerequisites
- Interfaces in Practice
- Unsafe Inspection

## Learning Objectives
- Model a narrow boundary for native-backed work
- Keep native concerns isolated behind an interface
- Test Go-side behavior without requiring a cgo toolchain

## Concept Summary

Actual cgo usage depends on local toolchains and platform details, so this lesson focuses on the boundary design you should have before enabling it. The key idea is to isolate native calls behind a small Go-facing seam.

## Run Instructions

```bash
go run starter.go
```

## Test Instructions

```bash
go test -v
```

## Exercises

### Exercise 1: Native Call
Implement `NativeCall(a, b int, driver NativeAdder) int`.

### Exercise 2: Sum Pair
Implement `SumPair(pair Pair, driver NativeAdder) int`.

### Exercise 3: Report Pair
Implement `ReportPair(pair Pair, driver NativeAdder) string`.
