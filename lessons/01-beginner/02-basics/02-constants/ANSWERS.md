# Exercise Answers and Explanations

## Exercise 1: Basic Constants

### Solution Explanation

This exercise introduces the basic syntax for declaring constants in Go. We create three constants:

```go
const Pi = 3.14159
const MaxRetries int = 5
const AppName = "GoLearning"
```

Notice that `Pi` is an untyped constant (no explicit type), while `MaxRetries` is explicitly typed as `int`. This demonstrates both styles of constant declaration.

The `CalculateCircleArea` function uses the constant `Pi` to calculate the area using the formula: area = π × r²

```go
func CalculateCircleArea(radius float64) float64 {
    return Pi * radius * radius
}
```

**Key Points**:
- Constants are declared with the `const` keyword
- Constants can be typed or untyped
- Untyped constants (like `Pi`) can be used in expressions with different numeric types
- Constants improve code readability by replacing magic numbers

**Common Pitfalls**:
- Trying to modify a constant after declaration (will cause compile error)
- Using runtime values in constant declarations (constants must be compile-time values)

## Exercise 2: Typed vs Untyped Constants

### Solution Explanation

This exercise demonstrates the key difference between typed and untyped constants:

```go
const UntypedNumber = 42
const TypedNumber int64 = 42
```

The `DemonstrateConstantTypes` function shows how untyped constants are more flexible:

```go
func DemonstrateConstantTypes() bool {
    var i int = UntypedNumber        // Works - implicit conversion
    var f float64 = UntypedNumber    // Works - implicit conversion
    var i64 int64 = UntypedNumber    // Works - implicit conversion
    
    return i == 42 && f == 42.0 && i64 == 42
}
```

If you tried to use `TypedNumber` (typed as int64) with an `int` variable, you'd get a compile error:
```go
var i int = TypedNumber  // Error: cannot use TypedNumber (type int64) as type int
```

**Key Points**:
- Untyped constants can implicitly convert to compatible types
- Typed constants follow strict type rules like variables
- Untyped constants have up to 256 bits of precision
- Use untyped for flexibility, typed for type safety

**When to Use Each**:
- **Untyped**: Most mathematical constants, values used across different types
- **Typed**: When you need to enforce a specific type or implement an interface

## Exercise 3: Const Blocks

### Solution Explanation

Const blocks group related constants together, improving organization and readability:

```go
const (
    StatusOK                  = 200
    StatusCreated             = 201
    StatusBadRequest          = 400
    StatusUnauthorized        = 401
    StatusNotFound            = 404
    StatusInternalServerError = 500
)
```

The `GetStatusMessage` function uses a switch statement to map status codes to messages:

```go
func GetStatusMessage(code int) string {
    switch code {
    case StatusOK:
        return "OK"
    case StatusCreated:
        return "Created"
    case StatusBadRequest:
        return "Bad Request"
    case StatusUnauthorized:
        return "Unauthorized"
    case StatusNotFound:
        return "Not Found"
    case StatusInternalServerError:
        return "Internal Server Error"
    default:
        return "Unknown Status"
    }
}
```

**Key Points**:
- Const blocks use parentheses to group related constants
- Each constant in the block is separated by newlines (no commas)
- Const blocks improve code organization and readability
- Using named constants (like `StatusOK`) is better than magic numbers (200)

**Benefits**:
- Single place to update related constants
- Self-documenting code
- Easier to refactor
- Better IDE support (autocomplete, go-to-definition)

## Exercise 4: Using iota for Enumerations

### Solution Explanation

The `iota` identifier is a powerful feature for creating enumerated constants. It starts at 0 and increments automatically:

```go
const (
    Sunday = iota  // 0
    Monday         // 1
    Tuesday        // 2
    Wednesday      // 3
    Thursday       // 4
    Friday         // 5
    Saturday       // 6
)
```

For file permissions, we use iota with bit shifting to create bit flags:

```go
const (
    ReadPermission    = 1 << iota  // 1 (binary: 001)
    WritePermission                 // 2 (binary: 010)
    ExecutePermission               // 4 (binary: 100)
)
```

This creates powers of 2, which allows combining permissions using bitwise OR:
```go
permissions := ReadPermission | WritePermission  // 3 (binary: 011)
```

**Function Implementations**:

```go
func IsWeekend(day int) bool {
    return day == Sunday || day == Saturday
}

func IsWeekday(day int) bool {
    return day >= Monday && day <= Friday
}

func HasPermission(userPermissions, requiredPermission int) bool {
    return userPermissions&requiredPermission != 0
}
```

**Key Points**:
- `iota` starts at 0 and increments for each constant in the block
- `iota` resets to 0 in each new const block
- Common pattern: `1 << iota` creates bit flags (powers of 2)
- Bit flags can be combined with bitwise OR (`|`) and checked with bitwise AND (`&`)

**Advanced iota Patterns**:

Starting at 1:
```go
const (
    January = iota + 1  // 1
    February            // 2
)
```

Skipping values:
```go
const (
    A = iota  // 0
    _         // 1 (skipped using blank identifier)
    C         // 2
)
```

Complex expressions:
```go
const (
    KB = 1 << (10 * iota)  // 1024
    MB                      // 1048576
    GB                      // 1073741824
)
```

**Bitwise Operations Explanation**:

For permission checking:
- `ReadPermission = 1` (binary: 001)
- `WritePermission = 2` (binary: 010)
- `userPermissions = 3` (binary: 011)

When checking: `userPermissions & ReadPermission`:
```
  011 (userPermissions)
& 001 (ReadPermission)
-----
  001 (result is non-zero, so permission exists)
```

When checking: `userPermissions & ExecutePermission`:
```
  011 (userPermissions)
& 100 (ExecutePermission)
-----
  000 (result is zero, so permission doesn't exist)
```

**Common Pitfalls**:
- Forgetting that `iota` resets in each const block
- Using `iota` in separate const declarations (each gets 0)
- Not understanding that `iota` represents the position, not necessarily the value
- Confusion with bit operations when using `1 << iota`

## Summary

These exercises cover the essential aspects of constants in Go:

1. **Basic constants**: Understanding const declaration and usage
2. **Typed vs untyped**: Knowing when to use each for flexibility vs type safety
3. **Const blocks**: Organizing related constants together
4. **iota**: Creating enumerations and bit flags efficiently

Constants are a fundamental feature that makes Go code more readable, maintainable, and type-safe. Mastering these concepts will help you write better Go programs.

## Further Practice

Try these additional challenges:
1. Create a color enum using iota with values starting at 1
2. Implement a state machine using const values
3. Create file size constants (KB, MB, GB) using iota with bit shifting
4. Build a priority enum (Low, Medium, High) with custom values
5. Create a day-of-week type with methods like `IsBusinessDay()` and `Next()`
