# Constants in Go

## Difficulty
Beginner

## Prerequisites
- Basic Go syntax
- Understanding of variables
- Familiarity with data types

## Learning Objectives
By the end of this lesson, you will be able to:
- Declare and use constants in Go
- Understand the difference between typed and untyped constants
- Use const blocks to group related constants
- Use `iota` to create enumerations and sequential values

## Concept Summary

Constants in Go are immutable values known at compile time. Unlike variables, constants cannot be changed once declared. Go provides the `const` keyword to declare constants, and they can be either typed or untyped. Untyped constants have greater flexibility and precision than typed constants.

The `iota` identifier is a special constant generator that simplifies the creation of enumerated constants. It starts at 0 and increments by 1 for each constant in a const block.

## Run Instructions

To work with the starter code:
```bash
go test -tags=starter -v
```

To run the solution:
```bash
go test -v
```

## Test Instructions

Run all tests with:
```bash
go test -v
```

For a specific test:
```bash
go test -v -run TestExercise1
```

## Exercises

### Exercise 1: Basic Constants
**Objective**: Declare and use basic constants.

Create the following constants:
- `Pi` as an untyped constant with value 3.14159
- `MaxRetries` as a typed int constant with value 5
- `AppName` as a string constant with value "GoLearning"

Implement the function `CalculateCircleArea(radius float64) float64` that uses the `Pi` constant to calculate the area of a circle.

**Expected behavior**:
- `CalculateCircleArea(2.0)` should return approximately `12.56636`
- `CalculateCircleArea(5.0)` should return approximately `78.53975`

### Exercise 2: Typed vs Untyped Constants
**Objective**: Understand the difference between typed and untyped constants.

Create two constants:
- `UntypedNumber` as an untyped constant with value 42
- `TypedNumber` as an int64 constant with value 42

Implement the function `DemonstrateConstantTypes()` that:
- Assigns `UntypedNumber` to variables of types `int`, `float64`, and `int64`
- Demonstrates that untyped constants can be used flexibly across different numeric types
- Returns true if all assignments work correctly

**Expected behavior**:
- The function should return `true` showing that untyped constants are flexible
- The untyped constant should work with multiple numeric types without explicit conversion

### Exercise 3: Const Blocks
**Objective**: Use const blocks to group related constants.

Create a const block for HTTP status codes with the following constants:
- `StatusOK` = 200
- `StatusCreated` = 201
- `StatusBadRequest` = 400
- `StatusUnauthorized` = 401
- `StatusNotFound` = 404
- `StatusInternalServerError` = 500

Implement the function `GetStatusMessage(code int) string` that returns a human-readable message for each status code:
- 200: "OK"
- 201: "Created"
- 400: "Bad Request"
- 401: "Unauthorized"
- 404: "Not Found"
- 500: "Internal Server Error"
- Any other code: "Unknown Status"

**Expected behavior**:
- `GetStatusMessage(StatusOK)` returns `"OK"`
- `GetStatusMessage(StatusNotFound)` returns `"Not Found"`
- `GetStatusMessage(999)` returns `"Unknown Status"`

### Exercise 4: Using iota for Enumerations
**Objective**: Use `iota` to create enumerated constants.

Create a const block using `iota` for days of the week:
- `Sunday` = 0 (using iota)
- `Monday` = 1
- `Tuesday` = 2
- `Wednesday` = 3
- `Thursday` = 4
- `Friday` = 5
- `Saturday` = 6

Also create a const block for file permissions using iota with bit shifting:
- `ReadPermission` = 1 << iota (equals 1)
- `WritePermission` (equals 2)
- `ExecutePermission` (equals 4)

Implement the following functions:
- `IsWeekend(day int) bool` - returns true if the day is Saturday or Sunday
- `IsWeekday(day int) bool` - returns true if the day is Monday through Friday
- `HasPermission(userPermissions, requiredPermission int) bool` - checks if user has the required permission using bitwise AND

**Expected behavior**:
- `IsWeekend(Sunday)` returns `true`
- `IsWeekend(Wednesday)` returns `false`
- `IsWeekday(Monday)` returns `true`
- `HasPermission(3, ReadPermission)` returns `true` (3 has read bit set)
- `HasPermission(3, ExecutePermission)` returns `false` (3 doesn't have execute bit set)

## Additional Notes

- Constants must be compile-time values (cannot use runtime expressions)
- Untyped constants have precision up to 256 bits
- Constants can be used in places where their type can be inferred
- `iota` resets to 0 in each const block
- You can skip values in iota using the blank identifier `_`
