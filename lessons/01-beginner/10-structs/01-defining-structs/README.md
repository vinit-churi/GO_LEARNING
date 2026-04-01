# Structs

## Difficulty
Beginner

## Prerequisites
- Functions
- Slices

## Learning Objectives
- Define structs with named fields
- Build nested structs for related data
- Use struct values to model real records

## Concept Summary

Structs let you group related fields into one type. They make code easier to read than passing many separate values around.

## Run Instructions

```bash
go run starter.go
```

## Test Instructions

```bash
go test -v
```

## Exercises

### Exercise 1: New Book
Implement `NewBook(title string, pages int) Book`.

### Exercise 2: Total Pages
Implement `TotalPages(books []Book) int`.

### Exercise 3: Pickup Label
Implement `PickupLabel(order PickupOrder) string` using fields from a nested struct.
