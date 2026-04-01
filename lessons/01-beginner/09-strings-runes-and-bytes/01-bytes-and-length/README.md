# Strings, Runes, And Bytes

## Difficulty
Beginner

## Prerequisites
- Slices
- Basic loops

## Learning Objectives
- Measure a string in bytes
- Read Unicode text as runes
- Convert between strings, bytes, and runes

## Concept Summary

A Go string stores bytes. For plain ASCII text, bytes and characters often line up. For Unicode text, a character may use more than one byte, which is why runes matter.

## Run Instructions

```bash
go run starter.go
```

## Test Instructions

```bash
go test -v
```

## Exercises

### Exercise 1: Byte Length
Implement `ByteLength(text string) int`.

### Exercise 2: First Rune
Implement `FirstRune(text string) rune` so it returns the first Unicode character.

### Exercise 3: Reverse Runes
Implement `ReverseRunes(text string) string` so Unicode text is reversed safely.
