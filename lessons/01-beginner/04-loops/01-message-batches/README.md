# Message Batches

Difficulty: Beginner

## Prerequisites

- Conditionals
- Basic integer arithmetic

## Learning Objectives

- write counted `for` loops
- build slices with repeated values
- accumulate totals across iterations
- use conditions inside loops

## Concept Summary

Go uses `for` for all loop styles. This lesson focuses on the classic counted form and on building useful values as the loop progresses.

## Test Instructions

```bash
go test
```

## Exercises

### Exercise 1

Implement `totalMessages(days int) int`.

Problem statement:
If a learner sends 1 message on day 1, 2 on day 2, and so on, return the total number of messages sent after `days` days.

### Exercise 2

Implement `buildCountdown(start int) []int`.

Problem statement:
Return a slice containing all numbers from `start` down to `1`. If `start` is less than `1`, return an empty slice.

### Exercise 3

Implement `sumEvenNumbers(limit int) int`.

Problem statement:
Return the sum of all even numbers from `0` up to and including `limit`.
