# Hello World and Formatting

Difficulty: Beginner

## Prerequisites

- A working Go installation
- Basic terminal usage

## Learning Objectives

- write and run a small Go program
- return strings from functions
- use `fmt.Sprintf` for formatting
- make simple boolean decisions

## Concept Summary

This lesson introduces the shape of a small Go program. You will create strings, format values into messages, and return results from functions that are easy to test.

## Run Instructions

```bash
go run .
```

## Test Instructions

```bash
go test
```

## Worked Example

`fmt.Sprintf("Hello, %s!", "Gopher")` returns `"Hello, Gopher!"`.

## Exercises

### Exercise 1

Implement `greeting(name string) string`.

Problem statement:
Return a greeting in the exact format `Hello, <name>!`.

### Exercise 2

Implement `learnerProfile(name string, lessonsCompleted int) string`.

Problem statement:
Return a sentence in the exact format `<name> has completed <n> lessons.`.

### Exercise 3

Implement `isReadyForNextLesson(completed int) bool`.

Problem statement:
Return `true` if the learner has completed at least 3 lessons. Otherwise return `false`.
