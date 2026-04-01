# Maps

## Difficulty
Beginner

## Prerequisites
- Arrays and slices
- Loops

## Learning Objectives
- Create and update maps
- Check whether a key exists
- Count repeated values

## Concept Summary

Maps store key-value pairs. They are useful when you need fast lookup by name, ID, or some other key.

## Run Instructions

```bash
go run starter.go
```

## Test Instructions

```bash
go test -v
```

## Exercises

### Exercise 1: Build Scoreboard
Implement `BuildScoreboard(names []string) map[string]int` so every name starts with score `0`.

### Exercise 2: Lookup Score
Implement `LookupScore(scores map[string]int, name string) (int, bool)`.

### Exercise 3: Count Words
Implement `CountWords(words []string) map[string]int` so repeated words are counted.
