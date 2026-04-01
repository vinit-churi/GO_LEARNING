# Scope And Packages

## Difficulty
Beginner

## Prerequisites
- Functions
- Basic strings

## Learning Objectives
- Understand package boundaries and exported names
- Use local scope to keep temporary variables contained
- Import a package from the same lesson directory

## Concept Summary

Go organizes code into packages. Names that start with an uppercase letter are exported and can be used from another package. Local variables belong to the block where they are declared, which helps you avoid leaking temporary state.

## Run Instructions

```bash
go run starter.go
```

## Test Instructions

```bash
go test -v
```

## Exercises

### Exercise 1: Desk Header
Implement `DeskHeader(team string) string` so it returns the exported header produced by the local `labelutil` package.

### Exercise 2: Member Code
Implement `BuildMemberCode(team, name string) string` so it trims both inputs, uppercases them, and returns `TEAM-NAME`.

### Exercise 3: Short Name Check
Implement `HasShortName(name string) bool` so it returns `true` when the trimmed name has length 4 or less.
