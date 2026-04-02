# Platform Selection

## Difficulty
Advanced

## Prerequisites
- Packages
- HTTP Servers

## Learning Objectives
- Use build tags to provide platform-specific implementations
- Keep shared code outside tagged files
- Discuss cross-compilation at the package boundary instead of scattering conditionals everywhere

## Concept Summary

Build tags let Go choose files for specific environments at compile time. They are useful for platform differences, but they should remain narrow and easy to reason about.

## Run Instructions

```bash
go run starter.go
```

## Test Instructions

```bash
go test -v
```

## Exercises

### Exercise 1: Tagged Platform
Provide a platform-specific implementation of `taggedPlatform()` using build-tagged files.

### Exercise 2: Build Artifact Name
Implement `BuildArtifactName(base string) string` with platform-specific extensions.

### Exercise 3: Target Triple
Implement `TargetTriple(goos, goarch string) string`.
