# Arrays And Slices

## Difficulty
Beginner

## Prerequisites
- Loops
- Functions

## Learning Objectives
- Work with fixed-size arrays
- Append to slices
- Use slicing and `range`

## Concept Summary

Arrays have a fixed length that is part of the type. Slices are lightweight views over arrays and are used much more often in everyday Go code.

## Run Instructions

```bash
go run starter.go
```

## Test Instructions

```bash
go test -v
```

## Exercises

### Exercise 1: Sum First Three
Implement `SumFirstThree(numbers [3]int) int`.

### Exercise 2: Add Task
Implement `AddTask(tasks []string, task string) []string` using `append`.

### Exercise 3: Weekend Window
Implement `WeekendWindow(days []string) []string` so it returns the Saturday and Sunday portion of the slice.
