# Safe Unsafe Inspection

## Difficulty
Advanced

## Prerequisites
- Memory Layout Inspection

## Learning Objectives
- Use `unsafe.Sizeof`, `Alignof`, and `Offsetof` safely
- Recognize why `unsafe` should remain rare and isolated
- Connect low-level facts back to higher-level design choices

## Concept Summary

The `unsafe` package is for low-level interop and layout inspection, not routine application logic. The safest way to teach it is to focus on read-only facts before any pointer conversions.

## Run Instructions

```bash
go run starter.go
```

## Test Instructions

```bash
go test -v
```

## Exercises

### Exercise 1: Header Size
Implement `HeaderSize() uintptr` for the sample header struct.

### Exercise 2: Header Alignment
Implement `HeaderAlignment() uintptr`.

### Exercise 3: Payload Offset
Implement `PayloadOffset() uintptr`.
