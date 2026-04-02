# Inspecting Structs

## Difficulty
Advanced

## Prerequisites
- Structs
- Interfaces in Practice

## Learning Objectives
- Inspect struct metadata with `reflect.Type`
- Read field names and kinds safely
- Use reflection for narrow tooling-style helpers

## Concept Summary

Reflection lets code inspect types and values at runtime. It is powerful, but it trades away compile-time clarity, so it should usually sit behind a narrow helper instead of spreading through a codebase.

## Run Instructions

```bash
go run starter.go
```

## Test Instructions

```bash
go test -v
```

## Exercises

### Exercise 1: Describe Struct
Implement `DescribeStruct(input any) string` to report the struct type name and field count.

### Exercise 2: Field Names
Implement `FieldNames(input any) []string` for exported and unexported fields.

### Exercise 3: Is Zero Value
Implement `IsZeroValue(input any) bool` using reflection.
