# Named Return Values

**Difficulty:** Beginner  
**Prerequisites:** Basic functions, Multiple return values

## Learning Objectives

By the end of this lesson, you will:

- Understand how named return values work in Go
- Use naked returns appropriately
- Recognize when named returns improve code clarity
- Understand the trade-offs between named returns and explicit returns
- Avoid common pitfalls with named returns

## Concept Summary

Named return values allow you to specify names for the values that a function returns. These names act as variables that are automatically initialized to their zero values and can be used throughout the function. When you use a bare `return` statement (called a "naked return"), Go returns the current values of these named return variables.

Example:

```go
func divide(a, b float64) (result float64, err error) {
    if b == 0 {
        err = errors.New("division by zero")
        return // naked return
    }
    result = a / b
    return // naked return
}
```

Named returns can make code more readable in some cases, but they can also reduce clarity if overused. The key is knowing when to use them.

## How To Run

From this directory, run:

```bash
go run solution.go
```

Or with the starter code:

```bash
go run -tags=starter starter.go
```

## How To Test

From this directory, run:

```bash
go test -v
```

## Exercises

### Exercise 1: Basic Named Returns

Write a function `rectangleDimensions(area, perimeter int) (length, width int)` that calculates the length and width of a rectangle given its area and perimeter. Assume the rectangle has integer dimensions and length ≥ width.

For simplicity, if there's no valid solution or multiple solutions, return length=0, width=0.

**Hints:**
- Use named returns: `length` and `width`
- Set these variables in your function body
- Use a naked return at the end

**Example:**
```go
length, width := rectangleDimensions(24, 20) // length=6, width=4 (6*4=24, 2*(6+4)=20)
```

### Exercise 2: Named Returns with Error Handling

Write a function `safeDivide(a, b float64) (result float64, err error)` that divides two numbers and returns an error if the divisor is zero.

**Requirements:**
- Use named return values `result` and `err`
- Return an error with message "division by zero" when b is 0
- Use naked returns (bare `return` statements)
- When successful, set result and err appropriately

**Example:**
```go
result, err := safeDivide(10, 2)  // result=5.0, err=nil
result, err := safeDivide(10, 0)  // result=0.0, err="division by zero"
```

### Exercise 3: When to Use Named Returns

Write TWO implementations of a function that parses a user's full name into first and last name:

**Implementation A:** `parseNameNamed(fullName string) (firstName, lastName string, err error)`
- Use named returns with naked return

**Implementation B:** `parseNameExplicit(fullName string) (string, string, error)`
- Use explicit returns (no named returns)

The function should split on the first space character. If there's no space, return an error "invalid name format".

**Example:**
```go
first, last, err := parseNameNamed("John Doe")     // first="John", last="Doe", err=nil
first, last, err := parseNameExplicit("Madonna")   // first="", last="", err="invalid name format"
```

**Reflection:** After implementing both, consider which version is clearer and why.

### Exercise 4: Complex Named Returns (Clarity vs Brevity)

Write a function `calculateStats(numbers []int) (sum, count int, avg float64, err error)` that calculates statistics for a slice of integers.

**Requirements:**
- Use named returns: `sum`, `count`, `avg`, `err`
- Return an error if the slice is empty (message: "empty slice")
- Calculate the sum, count, and average of all numbers
- Use at least one explicit return and one naked return appropriately

**Example:**
```go
sum, count, avg, err := calculateStats([]int{2, 4, 6, 8})
// sum=20, count=4, avg=5.0, err=nil

sum, count, avg, err := calculateStats([]int{})
// sum=0, count=0, avg=0.0, err="empty slice"
```

### Exercise 5: Named Returns Best Practices

Write a function `validateAndFormat(input string) (output string, valid bool)` that:
- Trims whitespace from the input
- Checks if the trimmed string has at least 3 characters
- Converts valid strings to uppercase
- Returns the formatted string and whether it was valid

**Requirements:**
- Use named returns
- Use explicit returns (NOT naked returns) - demonstrate when explicit is better
- Show good practices: clear logic flow and return points

**Example:**
```go
output, valid := validateAndFormat("  hello  ")  // output="HELLO", valid=true
output, valid := validateAndFormat("  hi  ")     // output="", valid=false
```

## Tips

- Named returns are initialized to their zero values automatically
- Naked returns should be used sparingly and only in short functions
- Named returns can serve as documentation for what a function returns
- Explicit returns are often clearer, especially in longer functions
- When in doubt, favor clarity over brevity
