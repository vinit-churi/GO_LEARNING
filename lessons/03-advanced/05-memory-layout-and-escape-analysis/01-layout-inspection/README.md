# Layout Inspection

## Difficulty
Advanced

## Prerequisites
- Structs
- Unsafe Package Overview

## Learning Objectives
- Inspect struct size and field offsets
- Recognize when field ordering introduces padding
- Discuss escape analysis without relying on unstable compiler internals in tests

## Concept Summary

Memory layout topics are useful when diagnosing performance or interoperability issues, not for everyday application code. The safest entry point is to inspect layout rather than trying to outsmart the compiler.

## Run Instructions

```bash
go run starter.go
```

## Test Instructions

```bash
go test -v
```

## Exercises

### Exercise 1: Record Size
Implement `RecordSize() uintptr` using `unsafe.Sizeof`.

### Exercise 2: Value Offset
Implement `ValueOffset() uintptr` using `unsafe.Offsetof`.

### Exercise 3: Has Padding
Implement `HasPadding() bool` based on the field offset.
