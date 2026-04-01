# Hello World

**Difficulty:** Beginner  
**Prerequisites:** None (This is your first lesson!)

## Learning Objectives

By the end of this lesson, you will:

- Understand the basic structure of a Go program
- Write and run your first Go program
- Use the `fmt` package to print output
- Understand the role of the `main` function and `main` package

## Concept Summary

Every Go program starts execution in the `main` function of the `main` package. The `fmt` package from the standard library provides formatted I/O functions like `Println` for printing to the console.

A minimal Go program looks like this:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

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

### Exercise 1: Hello, World!

Write a program that prints "Hello, World!" to the console.

**Expected Output:**
```
Hello, World!
```

### Exercise 2: Greet Yourself

Modify your program to print a personalized greeting with your name.

**Example Output:**
```
Hello, Alice!
Welcome to Go programming.
```

### Exercise 3: Multiple Lines

Write a program that prints three separate lines:
1. A greeting
2. A message about what you're learning
3. An encouraging statement

**Example Output:**
```
Hello, Go learner!
You are learning the Go programming language.
Keep practicing and you'll master it!
```

### Exercise 4: ASCII Art

Create a simple ASCII art welcome message using multiple `fmt.Println` statements.

**Example Output:**
```
 ____  ____  
|  __||  _ \ 
| |_  | |_) |
|  _| |  _ < 
|_|   |_| \_\

Welcome to Go!
```

## Tips

- Each call to `fmt.Println()` adds a newline automatically
- You can print empty lines by calling `fmt.Println()` with no arguments
- String literals are enclosed in double quotes: `"like this"`
