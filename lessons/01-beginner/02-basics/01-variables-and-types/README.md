# Variables and Types

## Difficulty
Beginner

## Prerequisites
- Basic understanding of programming concepts
- Go installed and working
- Completed "Hello World" lesson

## Learning Objectives
By the end of this lesson, you will be able to:
- Declare variables using different syntaxes (`var`, `:=`)
- Understand Go's basic types (int, float64, string, bool)
- Use type inference effectively
- Understand zero values in Go
- Declare multiple variables at once

## Concept Summary

Variables are named storage locations that hold values. Go is a statically-typed language, meaning every variable has a specific type that is determined at compile time.

Go provides several ways to declare variables:
- `var` keyword with explicit type
- `var` keyword with type inference
- Short variable declaration (`:=`) inside functions

Every type in Go has a "zero value" - the default value a variable gets when declared without initialization.

## Run Instructions

This lesson uses build tags to separate starter and solution code.

To work on the exercises:
```bash
go test -v -tags=starter
```

To see the solution:
```bash
go test -v -tags=solution
```

Or to run without tags (uses solution by default):
```bash
go test -v
```

## Exercises

### Exercise 1: Basic Variable Declaration
Declare a variable named `name` of type `string` and assign it your name. Declare another variable `age` of type `int` and assign it your age. Use the `var` keyword with explicit types.

**Function to implement:** `Exercise1() (string, int)`

**Expected behavior:**
- Returns a string and an integer
- String should be a valid name
- Integer should be a positive number

### Exercise 2: Short Variable Declaration
Use the short variable declaration syntax (`:=`) to create three variables:
- `temperature` with value `98.6` (float64)
- `isRaining` with value `false` (bool)
- `city` with value `"San Francisco"` (string)

**Function to implement:** `Exercise2() (float64, bool, string)`

**Expected behavior:**
- Returns three values using type inference
- No explicit type declarations needed
- All values should be correctly typed

### Exercise 3: Zero Values
Declare variables without initialization to observe Go's zero values. Create:
- An `int` variable (zero value is 0)
- A `float64` variable (zero value is 0.0)
- A `string` variable (zero value is "")
- A `bool` variable (zero value is false)

**Function to implement:** `Exercise3() (int, float64, string, bool)`

**Expected behavior:**
- All variables declared with `var` but no initial value
- Returns the zero values for each type

### Exercise 4: Multiple Variable Declarations
Declare multiple variables in a single statement. Create a set of variables representing a product:
- `productName` (string) = "Laptop"
- `price` (float64) = 999.99
- `inStock` (bool) = true
- `quantity` (int) = 5

Use a single `var` block to declare all four variables.

**Function to implement:** `Exercise4() (string, float64, bool, int)`

**Expected behavior:**
- All variables declared in a single `var ()` block
- Returns all four product-related values
- Values should match the specifications above

## Testing

Run the tests to verify your solutions:
```bash
go test -v -tags=starter
```

Each test will check:
- Correct return types
- Appropriate values
- Proper variable declaration syntax (checked by compilation)

## Next Steps

After completing this lesson, you should:
1. Review `REFERENCE.md` for detailed explanations of Go types
2. Check `ANSWERS.md` if you get stuck
3. Experiment with different types and declarations
4. Move on to the next lesson on constants and operators
