# Exercise Solutions

## Exercise 1: Simple Greeting Function

### Explanation
This exercise introduces the basic structure of a function with a single parameter and a return value. The function takes a string input and returns a formatted string output.

### Hints
- Use the `+` operator or `fmt.Sprintf()` to concatenate strings
- Remember to include the `return` keyword
- The return type must match what you declared in the function signature

### Solution
```go
func Greet(name string) string {
    return "Hello, " + name + "!"
}
```

### Discussion
String concatenation in Go can be done in multiple ways:
1. Using `+` operator (simple and readable for a few strings)
2. Using `fmt.Sprintf()` (more flexible with formatting)
3. Using `strings.Builder` (most efficient for many concatenations)

For this simple case, the `+` operator is perfectly fine.

See `solution.go` for the complete implementation.

---

## Exercise 2: Add Two Numbers

### Explanation
This exercise demonstrates a function with multiple parameters of the same type. When parameters share the same type, Go allows you to declare the type once at the end.

### Hints
- Since both parameters are integers, you can write `Add(a, b int)` instead of `Add(a int, b int)`
- The return value is also an integer
- Simply return the sum using the `+` operator

### Solution
```go
func Add(a, b int) int {
    return a + b
}
```

### Discussion
This is the simplest form of a function with multiple parameters. In Go, when consecutive parameters share the same type, you can omit the type from all but the last parameter. Both forms are correct:
- `Add(a int, b int) int` - explicit
- `Add(a, b int) int` - concise (preferred)

See `solution.go` for the complete implementation.

---

## Exercise 3: Print Information (No Return Value)

### Explanation
This exercise introduces functions that perform actions without returning values. Such functions are useful for operations with side effects, like printing to the console or writing to files.

### Hints
- When a function doesn't return anything, omit the return type
- Use `fmt.Printf()` or `fmt.Println()` to print formatted output
- You don't need a `return` statement at the end (though you can use `return` alone to exit early)

### Solution
```go
func PrintInfo(name string, age int) {
    fmt.Printf("Name: %s, Age: %d\n", name, age)
}
```

### Discussion
Functions without return values are used for side effects. The function signature has no return type specified. Alternative implementations:
- `fmt.Println("Name:", name, ", Age:", age)`
- `fmt.Printf("Name: %s, Age: %d\n", name, age)` (used in solution for precise formatting)

The second option gives you more control over formatting and matches the expected output format exactly.

See `solution.go` for the complete implementation.

---

## Exercise 4: Calculate Rectangle Area Using Helper Functions

### Explanation
This exercise demonstrates function composition - one function calling another. This is a fundamental concept in building larger programs from smaller, reusable pieces.

### Hints
- Create `Multiply` first - it's a simple function that returns the product of two numbers
- In `CalculateRectangleArea`, call `Multiply` with the width and height parameters
- The return value of `Multiply` becomes the return value of `CalculateRectangleArea`

### Solution
```go
func Multiply(a, b int) int {
    return a * b
}

func CalculateRectangleArea(width, height int) int {
    return Multiply(width, height)
}
```

### Discussion
This exercise demonstrates several important concepts:
1. **Function Reusability** - `Multiply` can be used anywhere multiplication is needed
2. **Function Composition** - Building complex operations from simple ones
3. **Abstraction** - `CalculateRectangleArea` expresses intent clearly

While you could calculate the area directly (`width * height`), using a helper function illustrates how functions can work together. In real-world code, helper functions often contain more complex logic that you want to reuse.

Alternative approach (direct calculation):
```go
func CalculateRectangleArea(width, height int) int {
    return width * height
}
```

Both are valid, but the exercise demonstrates function composition.

See `solution.go` for the complete implementation.

---

## Exercise 5: Temperature Converter

### Explanation
This exercise works with floating-point numbers and demonstrates a practical use case for functions - converting between units of measurement.

### Hints
- Use `float64` for both parameter and return type to handle decimal values
- Formula: F = C × 9/5 + 32
- Be careful with integer division - use `9.0` and `5.0` to ensure floating-point division
- You can write the formula directly in the return statement

### Solution
```go
func CelsiusToFahrenheit(celsius float64) float64 {
    return celsius * 9.0 / 5.0 + 32.0
}
```

### Discussion
This exercise introduces working with floating-point numbers in functions. Key points:

1. **Type Consistency** - Both input and output are `float64`
2. **Floating-Point Division** - Using `9.0 / 5.0` instead of `9 / 5` ensures floating-point division
   - `9 / 5` would perform integer division, resulting in `1`
   - `9.0 / 5.0` performs floating-point division, resulting in `1.8`
3. **Order of Operations** - Multiplication and division happen before addition, so no parentheses needed

Common test cases:
- 0°C = 32°F (freezing point of water)
- 100°C = 212°F (boiling point of water)
- 37°C = 98.6°F (normal body temperature)

See `solution.go` for the complete implementation.

---

## General Tips for Writing Functions

1. **Naming** - Use descriptive names that indicate what the function does
2. **Single Responsibility** - Each function should do one thing well
3. **Keep It Simple** - Start with simple functions and build complexity gradually
4. **Test Your Functions** - Write tests to verify correctness
5. **Documentation** - Add comments for complex logic

## Next Steps

After mastering basic functions, you'll learn about:
- Multiple return values
- Variadic functions
- Anonymous functions and closures
- Function types and higher-order functions
- Defer statements
