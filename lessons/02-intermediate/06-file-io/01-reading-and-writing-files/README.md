# Reading and Writing Files

## Difficulty
Intermediate

## Prerequisites
- Standard Library Essentials

## Learning Objectives
- Write files with `os.WriteFile`
- Read and clean file contents with `os.ReadFile`
- Keep file paths explicit in tests with temporary directories

## Concept Summary

Go keeps common file operations simple. For small files, `os.ReadFile` and `os.WriteFile` are often enough, especially in tests and small tools.

## Run Instructions

```bash
go run starter.go
```

## Test Instructions

```bash
go test -v
```

## Exercises

### Exercise 1: Save Note
Implement `SaveNote(path, text string) error` so it writes the text to disk.

### Exercise 2: Load Note
Implement `LoadNote(path string) (string, error)` so it reads the file and trims surrounding whitespace.

### Exercise 3: Duplicate Note
Implement `DuplicateNote(src, dst string) error` so it copies the cleaned file contents into a new file.
