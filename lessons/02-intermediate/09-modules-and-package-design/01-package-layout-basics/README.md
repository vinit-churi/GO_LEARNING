# Package Layout Basics

## Difficulty
Intermediate

## Prerequisites
- Scope and Packages
- Composition and Embedding

## Learning Objectives
- Split reusable helpers into a small package
- Keep exported APIs focused and named for callers
- Test package-facing behavior from the main package

## Concept Summary

Package design in Go is driven by simplicity at the call site. Small packages with a narrow exported surface are easier to understand than large grab-bag utility packages.

## Run Instructions

```bash
go run starter.go
```

## Test Instructions

```bash
go test -v
```

## Exercises

### Exercise 1: Slug Label
Implement `slugutil.Slug` so it trims, lowercases, and hyphenates a title.

### Exercise 2: Build Module Path
Implement `BuildModulePath(owner, repo string) string` using the helper package.

### Exercise 3: Is Public Package
Implement `IsPublicPackage(name string) bool` so package names starting with `internal-` are treated as non-public.
