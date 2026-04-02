# Status Classifier

## Difficulty
Intermediate

## Prerequisites
- Intro to Testing
- Standard Library Essentials

## Learning Objectives
- Group similar test cases into tables
- Write production helpers that are easy to probe with multiple inputs
- Use subtests to keep failures specific

## Concept Summary

Table-driven tests are common in Go because they keep repeated test structure small while making input and output combinations easy to scan.

## Run Instructions

```bash
go run starter.go
```

## Test Instructions

```bash
go test -v
```

## Exercises

### Exercise 1: Classify Score
Implement `ClassifyScore(score int) string` for low, medium, and high statuses.

### Exercise 2: Is Terminal
Implement `IsTerminal(status string) bool` for final statuses.

### Exercise 3: Build Status Line
Implement `BuildStatusLine(name string, score int) string` using the classifier.
