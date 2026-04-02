# Pointer Updates

## Difficulty
Intermediate

## Prerequisites
- Structs
- Methods

## Learning Objectives
- Update shared state through pointers
- Compare pointer and value behavior
- Use pointer receivers for mutating operations

## Concept Summary

Pointers let a function or method work with the original value instead of a copy. They matter when you need mutation, want to avoid copying larger structs, or need method sets that expose in-place updates.

## Run Instructions

```bash
go run starter.go
```

## Test Instructions

```bash
go test -v
```

## Exercises

### Exercise 1: Apply Discount
Implement `ApplyDiscount(order *Order, percent int)` so it reduces `Cents` in place.

### Exercise 2: Rename Customer
Implement `RenameCustomer(order *Order, name string)` so the stored customer name is trimmed and updated.

### Exercise 3: Snapshot Label
Implement `SnapshotLabel(order Order) string` so it returns a string view of the copied value without mutating the original order.
