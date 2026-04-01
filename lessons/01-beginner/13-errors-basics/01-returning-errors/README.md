# Errors Basics

## Difficulty
Beginner

## Prerequisites
- Functions
- Conditionals

## Learning Objectives
- Return errors from functions
- Create clear error values
- Handle errors before using the successful result

## Concept Summary

Go uses explicit error returns instead of exceptions for routine failure cases. A function can return a value and an `error`, and the caller decides how to handle it.

## Run Instructions

```bash
go run starter.go
```

## Test Instructions

```bash
go test -v
```

## Exercises

### Exercise 1: Parse Quantity
Implement `ParseQuantity(input string) (int, error)`.

### Exercise 2: Safe Divide
Implement `SafeDivide(a, b int) (int, error)`.

### Exercise 3: Build Order Message
Implement `BuildOrderMessage(name string, qty int) (string, error)`.
