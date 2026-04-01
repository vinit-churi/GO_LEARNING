# Interfaces Basics

## Difficulty
Beginner

## Prerequisites
- Structs
- Methods

## Learning Objectives
- Define a small interface
- See that implementation is implicit in Go
- Use interface values in a practical helper function

## Concept Summary

An interface describes behavior. A type implements an interface by defining the required methods. No extra keyword is needed.

## Run Instructions

```bash
go run starter.go
```

## Test Instructions

```bash
go test -v
```

## Exercises

### Exercise 1: Email Preview
Implement `Preview()` for `EmailAlert`.

### Exercise 2: SMS Preview
Implement `Preview()` for `SMSAlert`.

### Exercise 3: Send Preview
Implement `SendPreview(p Previewer) string`.
