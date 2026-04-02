# Generic Safe Cache

## Difficulty
Advanced

## Prerequisites
- Generics Fundamentals
- Mutexes and Atomics

## Learning Objectives
- Combine generics with mutex-protected shared state
- Design a small reusable type with clear methods
- Keep concurrency concerns contained inside the type

## Concept Summary

This mini-project ties together generics and synchronization in a small reusable cache. The goal is not to build a production cache, but to apply the language features in a constrained design.

## Run Instructions

```bash
go run starter.go
```

## Test Instructions

```bash
go test -v
```

## Exercises

### Exercise 1: Set
Implement `(*Cache[K, V]).Set(key K, value V)`.

### Exercise 2: Get
Implement `(*Cache[K, V]).Get(key K) (V, bool)`.

### Exercise 3: Size
Implement `(*Cache[K, V]).Size() int`.
