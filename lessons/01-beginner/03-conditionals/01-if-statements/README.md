# If Statements

**Difficulty:** Beginner  
**Prerequisites:** Variables and basic types

## Learning Objectives

By the end of this lesson, you will:

- Understand how to write basic if statements in Go
- Use if-else to handle two conditions
- Chain if-else if-else for multiple conditions
- Use if statements with short variable declarations
- Know when to use each pattern

## Concept Summary

If statements control the flow of your program based on conditions. Go's if statements don't require parentheses around the condition but do require braces around the body.

Basic syntax:

```go
if condition {
    // code runs if condition is true
}
```

If statements can also include a short statement before the condition:

```go
if x := getValue(); x > 0 {
    // x is only available in this if block
}
```

## How To Run

From this directory, run:

```bash
go run solution.go
```

Or to run with the starter file:

```bash
go run -tags=starter starter.go
```

## How To Test

From this directory, run:

```bash
go test -v
```

## Exercises

### Exercise 1: Basic If Statement

Write a function `CheckPositive(n int) string` that returns "positive" if the number is greater than zero, otherwise returns "not positive".

**Example:**
```go
CheckPositive(5)   // returns "positive"
CheckPositive(-3)  // returns "not positive"
CheckPositive(0)   // returns "not positive"
```

### Exercise 2: If-Else

Write a function `IsEven(n int) string` that returns "even" if the number is even, otherwise returns "odd".

**Example:**
```go
IsEven(4)  // returns "even"
IsEven(7)  // returns "odd"
IsEven(0)  // returns "even"
```

### Exercise 3: If-Else If-Else Chain

Write a function `CompareNumbers(a, b int) string` that:
- Returns "equal" if a equals b
- Returns "greater" if a is greater than b
- Returns "less" if a is less than b

**Example:**
```go
CompareNumbers(5, 3)  // returns "greater"
CompareNumbers(2, 7)  // returns "less"
CompareNumbers(4, 4)  // returns "equal"
```

### Exercise 4: If With Short Statement

Write a function `DivideAndCheck(a, b int) string` that:
- Divides a by b and stores the result
- Returns "large" if the result is greater than 10
- Returns "medium" if the result is between 1 and 10 (inclusive)
- Returns "small" if the result is less than 1
- Use an if statement with a short variable declaration

**Example:**
```go
DivideAndCheck(100, 5)  // returns "large" (100/5 = 20)
DivideAndCheck(20, 4)   // returns "medium" (20/4 = 5)
DivideAndCheck(5, 10)   // returns "small" (5/10 = 0)
```

### Exercise 5: Complex Condition

Write a function `CheckRange(n int) string` that:
- Returns "in range" if n is between 10 and 20 (inclusive)
- Returns "out of range" otherwise

**Example:**
```go
CheckRange(15)  // returns "in range"
CheckRange(5)   // returns "out of range"
CheckRange(10)  // returns "in range"
CheckRange(20)  // returns "in range"
CheckRange(25)  // returns "out of range"
```

## Tips

- Go's if statements don't use parentheses around conditions (but you can add them)
- Braces `{}` are always required, even for single-line bodies
- Variables declared in the if statement's initialization are scoped to the if block
- Use `&&` for AND, `||` for OR, and `!` for NOT in conditions
