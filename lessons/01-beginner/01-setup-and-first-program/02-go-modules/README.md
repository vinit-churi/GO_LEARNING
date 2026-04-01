# Go Modules and Project Structure

**Difficulty:** Beginner  
**Prerequisites:** Hello World lesson

## Learning Objectives

By the end of this lesson, you will:

- Understand what Go modules are and why they're used
- Initialize a Go module with `go mod init`
- Understand the structure of `go.mod` files
- Work with packages in your own module
- Import and use your own packages

## Concept Summary

A Go module is a collection of Go packages with a `go.mod` file at its root. The `go.mod` file defines the module's name and its dependencies. Modules are Go's way of managing code organization and dependencies.

Every Go project should start with:

```bash
go mod init example.com/myproject
```

This creates a `go.mod` file that tracks your module and its dependencies.

## How To Run

From this directory, run:

```bash
go run solution.go
```

## How To Test

From this directory, run:

```bash
go test -v
```

## Exercises

### Exercise 1: Module Path

Create a function that returns the module path string. The module path should be `"github.com/yourname/go-learning"` (or any valid module path you choose).

**Function Signature:**
```go
func GetModulePath() string
```

**Expected Output:**
```
github.com/yourname/go-learning
```

### Exercise 2: Package Information

Create a function that returns information about the current package in a formatted string.

**Function Signature:**
```go
func GetPackageInfo() string
```

**Expected Output:**
```
Package: main
Purpose: Learning Go modules and packages
```

### Exercise 3: Module Version

Create a function that returns a version string for your module following semantic versioning (e.g., "v1.0.0").

**Function Signature:**
```go
func GetModuleVersion() string
```

**Expected Output:**
```
v1.0.0
```

### Exercise 4: Full Module Info

Create a function that combines all the above information into a comprehensive module report.

**Function Signature:**
```go
func PrintModuleReport()
```

**Expected Output:**
```
=== Module Information ===
Module Path: github.com/yourname/go-learning
Package: main
Version: v1.0.0
Purpose: Learning Go modules and packages
========================
```

## Tips

- Module paths typically follow the pattern: `domain.com/username/project`
- For personal projects, you can use paths like `example.com/myproject`
- For GitHub projects, use `github.com/username/repository`
- The `main` package is special - it creates an executable
- Other packages are importable libraries
