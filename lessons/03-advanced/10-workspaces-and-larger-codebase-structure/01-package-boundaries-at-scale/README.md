# Package Boundaries At Scale

## Difficulty
Advanced

## Prerequisites
- Modules and Package Design

## Learning Objectives
- Keep package responsibilities explicit
- Use helper packages for boundary-focused logic
- Model workspace layout concerns without overcomplicating lesson code

## Concept Summary

Larger Go codebases stay navigable when package boundaries remain clear. Workspaces and multi-module layouts help with local development, but they do not replace thoughtful package ownership.

## Run Instructions

```bash
go run starter.go
```

## Test Instructions

```bash
go test -v
```

## Exercises

### Exercise 1: Root Module Label
Implement `RootModuleLabel(owner, repo string) string` using the helper package.

### Exercise 2: Is Internal Path
Implement `IsInternalPath(path string) bool`.

### Exercise 3: Package Summary
Implement `PackageSummary(path string) string` using the helper package.
