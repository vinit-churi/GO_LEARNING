# Variables, Constants, and Zero Values

Difficulty: Beginner

## Prerequisites

- Hello world
- Returning values from simple functions

## Learning Objectives

- declare and use constants
- understand zero values in structs
- update and return struct values
- reason about integers and booleans

## Concept Summary

Go gives every variable a default zero value. Strings default to `""`, integers to `0`, and booleans to `false`. This makes struct initialization predictable.

## Run Instructions

```bash
go run .
```

## Test Instructions

```bash
go test
```

## Exercises

### Exercise 1

Implement `newPracticeSession(name string) PracticeSession`.

Problem statement:
Return a new session with the given name and the remaining fields left at their zero values.

### Exercise 2

Implement `remainingPracticeMinutes(minutesSpent int) int`.

Problem statement:
Use the constant `maxDailyPracticeMinutes` and return how many minutes are left. If the learner already spent more than the maximum, return `0`.

### Exercise 3

Implement `hasStarted(session PracticeSession) bool`.

Problem statement:
Return `true` if the session has recorded any minutes or is marked complete.
