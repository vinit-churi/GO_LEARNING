# Multiple Return Values

**Difficulty:** Beginner  
**Prerequisites:** Basic functions, variables, error handling basics

## Learning Objectives

By the end of this lesson, you will:

- Understand how to return multiple values from a function
- Learn the idiomatic pattern of returning value and error
- Know when and how to discard return values using the blank identifier
- Apply multiple return values effectively in real-world scenarios

## Concept Summary

Go allows functions to return multiple values, which is one of its distinguishing features. This is commonly used for error handling, where a function returns both a result and an error value.

Basic syntax:

```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}

result, err := divide(10, 2)
if err != nil {
    // handle error
}
```

You can also discard values you don't need using the blank identifier `_`:

```go
result, _ := divide(10, 2) // ignoring error (generally not recommended)
_, err := divide(10, 0)     // only checking for error
```

## How To Run

Work with the starter code:

```bash
go run -tags starter starter.go
```

Run the solution:

```bash
go run solution.go
```

## How To Test

From this directory, run:

```bash
go test -v
```

## Exercises

### Exercise 1: Return Two Values

Write a function `minMax` that takes a slice of integers and returns both the minimum and maximum values.

**Function Signature:**
```go
func minMax(numbers []int) (int, int)
```

**Example:**
```go
min, max := minMax([]int{3, 7, 1, 9, 2})
// min = 1, max = 9
```

### Exercise 2: Return Value and Error

Write a function `safeDivide` that divides two integers and returns the result and an error. The function should return an error if the divisor is zero.

**Function Signature:**
```go
func safeDivide(a, b int) (int, error)
```

**Example:**
```go
result, err := safeDivide(10, 2)
// result = 5, err = nil

result, err := safeDivide(10, 0)
// result = 0, err = "cannot divide by zero"
```

### Exercise 3: Parse User Input

Write a function `parseUserInput` that takes a string in the format "name:age" and returns the name and age separately. Return an error if the format is invalid or age is not a number.

**Function Signature:**
```go
func parseUserInput(input string) (string, int, error)
```

**Example:**
```go
name, age, err := parseUserInput("Alice:30")
// name = "Alice", age = 30, err = nil

name, age, err := parseUserInput("Bob")
// name = "", age = 0, err = "invalid format"
```

### Exercise 4: Swap Values

Write a function `swap` that takes two values and returns them in swapped order. This demonstrates how multiple returns can simplify operations.

**Function Signature:**
```go
func swap(a, b string) (string, string)
```

**Example:**
```go
x, y := swap("hello", "world")
// x = "world", y = "hello"
```

### Exercise 5: Calculate Statistics

Write a function `calculateStats` that takes a slice of integers and returns the sum, average, and count. The average should be a float64. Return an error if the slice is empty.

**Function Signature:**
```go
func calculateStats(numbers []int) (int, float64, int, error)
```

**Example:**
```go
sum, avg, count, err := calculateStats([]int{10, 20, 30})
// sum = 60, avg = 20.0, count = 3, err = nil
```

## Tips

- Multiple return values are commonly used in Go for error handling
- By convention, the error is typically the last return value
- Use the blank identifier `_` to ignore values you don't need
- Named return values can make function signatures more readable
- Always handle errors - don't silently ignore them in production code
