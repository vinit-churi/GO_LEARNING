# Functions Reference

## What Are Functions?

Functions are self-contained blocks of code that perform a specific task. They are the building blocks of Go programs and help organize code into reusable, maintainable pieces.

## Why Functions Exist

Functions provide several benefits:
1. **Code Reusability** - Write once, use many times
2. **Modularity** - Break complex problems into smaller, manageable pieces
3. **Readability** - Give meaningful names to operations
4. **Testing** - Easier to test small, focused functions
5. **Maintenance** - Changes in one place affect all uses

## Basic Function Syntax

### Function Declaration
```go
func functionName(parameter1 type1, parameter2 type2) returnType {
    // function body
    return value
}
```

### Function with No Parameters
```go
func sayHello() string {
    return "Hello!"
}
```

### Function with No Return Value
```go
func printMessage(message string) {
    fmt.Println(message)
}
```

### Function with Multiple Parameters of the Same Type
When consecutive parameters share the same type, you can omit the type from all but the last:
```go
func add(x, y int) int {
    return x + y
}
```

### Calling Functions
```go
result := add(5, 3)
fmt.Println(result) // Output: 8
```

## Detailed Patterns

### 1. Simple Function with Return Value
```go
func square(n int) int {
    return n * n
}

// Usage
result := square(5) // result = 25
```

### 2. Multiple Parameters
```go
func fullName(firstName, lastName string) string {
    return firstName + " " + lastName
}

// Usage
name := fullName("John", "Doe") // name = "John Doe"
```

### 3. Function Without Return (Side Effects)
```go
func logMessage(message string) {
    fmt.Println("[LOG]:", message)
}

// Usage
logMessage("Application started") // Prints to console
```

### 4. Functions Calling Other Functions
```go
func double(n int) int {
    return n * 2
}

func quadruple(n int) int {
    return double(double(n))
}

// Usage
result := quadruple(3) // result = 12
```

### 5. Named Return Values
Go allows you to name return values, which can make code more readable:
```go
func divide(dividend, divisor int) (result int, remainder int) {
    result = dividend / divisor
    remainder = dividend % divisor
    return // naked return uses named values
}
```

## Common Mistakes

### 1. Forgetting to Return a Value
```go
// ❌ Wrong
func add(a, b int) int {
    sum := a + b
    // Missing return statement
}

// ✅ Correct
func add(a, b int) int {
    sum := a + b
    return sum
}
```

### 2. Type Mismatch
```go
// ❌ Wrong
func add(a, b int) string {
    return a + b // Cannot return int when string is expected
}

// ✅ Correct
func add(a, b int) int {
    return a + b
}
```

### 3. Unused Return Values
```go
func getValue() int {
    return 42
}

// Sometimes you might not need the return value
getValue() // This is fine, but the return value is ignored
```

### 4. Confusing Function Declaration and Call
```go
// ❌ Wrong - Missing parentheses
result := add // This assigns the function itself, not the result

// ✅ Correct
result := add(5, 3) // This calls the function
```

## When to Use Functions

### Use Functions When:
- You have code that will be used in multiple places
- You want to break down a complex operation into simpler steps
- You want to give a meaningful name to a series of operations
- You need to test a specific piece of logic
- You want to hide implementation details

### Don't Overcomplicate:
- Don't create a function for a single line of code used only once
- Don't create deeply nested function calls that hurt readability
- Avoid functions that do too many unrelated things

## Function Scope

Variables declared inside a function are local to that function:
```go
func example() {
    x := 10 // x only exists within this function
    fmt.Println(x)
}

// x is not accessible here
```

## Parameters Are Passed by Value

In Go, function parameters are passed by value, meaning the function receives a copy:
```go
func increment(n int) {
    n = n + 1
    fmt.Println(n) // Prints the incremented value
}

x := 5
increment(x) // Prints: 6
fmt.Println(x) // Still prints: 5 (original unchanged)
```

## Official Documentation

For more information, see:
- [A Tour of Go - Functions](https://go.dev/tour/basics/4)
- [Effective Go - Functions](https://go.dev/doc/effective_go#functions)
- [Go Language Specification - Function declarations](https://go.dev/ref/spec#Function_declarations)
