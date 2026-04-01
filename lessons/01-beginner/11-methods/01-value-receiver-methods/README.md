# Methods

## Difficulty
Beginner

## Prerequisites
- Structs
- Functions

## Learning Objectives
- Attach behavior to structs with methods
- Use value receivers for read-only behavior
- Use pointer receivers when a method needs to change data

## Concept Summary

Methods are functions attached to a type. They help you keep data and related behavior together.

## Run Instructions

```bash
go run starter.go
```

## Test Instructions

```bash
go test -v
```

## Exercises

### Exercise 1: Summary
Implement `Summary()` for `Counter`.

### Exercise 2: Add
Implement `Add(amount int)` with a pointer receiver so the count changes.

### Exercise 3: Reset
Implement `Reset()` with a pointer receiver.
