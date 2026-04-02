# Strings, Sorting, and Conversion

## Difficulty
Intermediate

## Prerequisites
- Basics
- Functions

## Learning Objectives
- Use `strings` helpers for cleanup
- Sort slices with `slices.Sort`
- Convert string data to numbers with `strconv`

## Concept Summary

A large part of idiomatic Go is knowing the standard library well. Small packages like `strings`, `strconv`, and `slices` solve common problems directly and keep custom code short.

## Run Instructions

```bash
go run starter.go
```

## Test Instructions

```bash
go test -v
```

## Exercises

### Exercise 1: Clean Tags
Implement `CleanTags(tags []string) []string` so it trims whitespace, skips blanks, and lowercases each tag.

### Exercise 2: Sort Scores
Implement `SortedScores(raw []string) ([]int, error)` by parsing the numbers and returning them in ascending order.

### Exercise 3: Join Tags
Implement `JoinTags(tags []string) string` so it reuses the cleaned list and joins with commas.
