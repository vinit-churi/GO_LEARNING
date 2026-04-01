# Basic Functions

## Difficulty
Beginner

## Prerequisites
- Variables and basic types
- Basic control flow

## Learning Objectives
By the end of this lesson, you will be able to:
- Declare functions with parameters and return values
- Create functions with multiple parameters
- Write functions without return values
- Call functions from other functions
- Understand function scope and execution flow

## Concept Summary

Functions are reusable blocks of code that perform specific tasks. In Go, functions are first-class citizens and are declared using the `func` keyword. Functions can accept zero or more parameters, return zero or more values, and can call other functions.

Basic function syntax:
```go
func functionName(param1 type1, param2 type2) returnType {
    // function body
    return value
}
```

## Run Instructions

To work on the exercises:
```bash
go run starter.go
```

To see the solution:
```bash
go run -tags=solution solution.go
```

## Test Instructions

Run the tests to validate your solutions:
```bash
go test -v
```

To test with the solution:
```bash
go test -v -tags=solution
```

## Exercises

### Exercise 1: Simple Greeting Function
Create a function named `Greet` that takes a `name` (string) as a parameter and returns a greeting message (string) in the format: "Hello, [name]!"

**Expected behavior:**
- Input: "Alice"
- Output: "Hello, Alice!"

### Exercise 2: Add Two Numbers
Create a function named `Add` that takes two integers as parameters and returns their sum.

**Expected behavior:**
- Input: 5, 3
- Output: 8

### Exercise 3: Print Information (No Return Value)
Create a function named `PrintInfo` that takes a `name` (string) and `age` (int) as parameters and prints the information to the console. This function should not return anything.

**Expected behavior:**
- Input: "Bob", 25
- Output (printed): "Name: Bob, Age: 25"

### Exercise 4: Calculate Rectangle Area Using Helper Functions
Create two functions:
1. `Multiply` - takes two integers and returns their product
2. `CalculateRectangleArea` - takes `width` and `height` (both integers) as parameters, calls `Multiply` internally, and returns the area

**Expected behavior:**
- Input to CalculateRectangleArea: 5, 4
- Output: 20

### Exercise 5: Temperature Converter
Create a function named `CelsiusToFahrenheit` that takes a temperature in Celsius (float64) and returns the temperature in Fahrenheit (float64). Use the formula: F = C * 9/5 + 32

**Expected behavior:**
- Input: 0.0
- Output: 32.0
- Input: 100.0
- Output: 212.0
