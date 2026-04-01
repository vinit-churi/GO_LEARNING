# Order Summary

Difficulty: Beginner

## Prerequisites

- Loops
- Formatting with `fmt.Sprintf`

## Learning Objectives

- break logic into small reusable functions
- pass slices into functions
- return formatted output
- compose multiple helper functions together

## Concept Summary

Functions are the main unit of reuse in Go. A small function should do one thing clearly, and larger behavior can be built by composing several smaller ones.

## Test Instructions

```bash
go test
```

## Exercises

### Exercise 1

Implement `subtotal(prices []float64) float64`.

Problem statement:
Return the sum of all prices in the slice.

### Exercise 2

Implement `applyDiscount(amount float64, isMember bool) float64`.

Problem statement:
If `isMember` is `true`, apply a 10% discount. Otherwise return the amount unchanged.

### Exercise 3

Implement `orderSummary(name string, prices []float64, isMember bool) string`.

Problem statement:
Build the final sentence in the exact format `<name>: subtotal=<subtotal>, total=<total>` using 2 decimal places.
