# Access Control

Difficulty: Beginner

## Prerequisites

- Variables and constants
- Boolean values

## Learning Objectives

- write `if`, `else if`, and `else`
- combine multiple conditions
- return different outputs based on program state
- use conditionals for realistic decisions

## Concept Summary

Conditionals let a program branch. In Go, the condition must evaluate to a boolean. There are no parentheses required around the condition.

## Test Instructions

```bash
go test
```

## Exercises

### Exercise 1

Implement `accessMessage(role string, active bool) string`.

Problem statement:
Return:

- `admin access granted` for an active admin
- `member access granted` for an active member
- `account inactive` for any inactive user
- `access denied` for any other active role

### Exercise 2

Implement `shippingCost(subtotal float64, premium bool) float64`.

Problem statement:
Return `0` for premium members. For everyone else, return `0` if subtotal is at least `50`, otherwise return `4.99`.

### Exercise 3

Implement `passOrRetry(score int) string`.

Problem statement:
Return:

- `excellent` for scores 90 and above
- `pass` for scores from 60 to 89
- `retry` for lower scores
