# Constants Reference Guide

## What Are Constants?

Constants are immutable values that are known at compile time and cannot be changed during program execution. In Go, constants are declared using the `const` keyword.

```go
const Pi = 3.14159
const MaxConnections = 100
const AppVersion = "1.0.0"
```

## The const Keyword

The `const` keyword declares a named constant. Constants can be character, string, boolean, or numeric values.

### Basic Syntax

```go
const name = value
const name type = value
```

Examples:
```go
const Port = 8080                    // untyped constant
const Timeout int = 30               // typed constant
const Message = "Hello, World!"      // string constant
const IsEnabled = true               // boolean constant
```

## Typed vs Untyped Constants

### Untyped Constants

Untyped constants don't have a specific type until they're used. They have higher precision (up to 256 bits) and can be used flexibly across different compatible types.

```go
const Number = 42

var i int = Number        // Works - Number becomes int
var f float64 = Number    // Works - Number becomes float64
var c complex128 = Number // Works - Number becomes complex128
```

**Advantages**:
- Greater flexibility in usage
- Higher precision for numeric calculations
- Implicit type conversion based on context

### Typed Constants

Typed constants have an explicit type and follow the same type rules as variables.

```go
const TypedNumber int = 42

var i int = TypedNumber       // Works
var f float64 = TypedNumber   // Error: type mismatch
var f2 float64 = float64(TypedNumber) // Works with explicit conversion
```

**When to use**:
- When you need to enforce type safety
- When implementing interfaces that require specific types
- When the constant should only be used with a specific type

## Const Blocks

Multiple related constants can be grouped using parentheses:

```go
const (
    StatusOK       = 200
    StatusCreated  = 201
    StatusNotFound = 404
)
```

You can mix typed and untyped constants:

```go
const (
    Pi      = 3.14159          // untyped
    E       = 2.71828          // untyped
    Timeout time.Duration = 5 * time.Second // typed
)
```

## iota - The Constant Generator

`iota` is a predeclared identifier that represents successive untyped integer constants. It starts at 0 and increments by 1 for each constant in a const block.

### Basic iota Usage

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

### iota Patterns

**Starting at 1:**
```go
const (
    January = iota + 1  // 1
    February            // 2
    March               // 3
)
```

**Skipping Values:**
```go
const (
    A = iota  // 0
    _         // 1 (skipped)
    C         // 2
    D         // 3
)
```

**Bit Flags:**
```go
const (
    ReadPermission    = 1 << iota  // 1 (1 << 0)
    WritePermission                 // 2 (1 << 1)
    ExecutePermission               // 4 (1 << 2)
)
```

**Complex Expressions:**
```go
const (
    _  = iota             // 0 (unused)
    KB = 1 << (10 * iota) // 1 << 10 = 1024
    MB                    // 1 << 20 = 1048576
    GB                    // 1 << 30 = 1073741824
)
```

### iota Resets

`iota` resets to 0 in each new const block:

```go
const (
    A = iota  // 0
    B         // 1
)

const (
    C = iota  // 0 (reset)
    D         // 1
)
```

## Why Constants Exist

1. **Immutability**: Protect values from accidental modification
2. **Performance**: Compiler can optimize constant usage
3. **Code Clarity**: Named constants are more readable than magic numbers
4. **Type Safety**: Compile-time checking of constant values
5. **Memory Efficiency**: Constants don't consume memory like variables

## When to Use Constants

**Use constants when**:
- Values never change (π, e, max limits)
- Configuration values known at compile time
- Enumerations and status codes
- Mathematical or physical constants
- Magic numbers that need names

**Don't use constants when**:
- Values come from user input or configuration files
- Values depend on runtime conditions
- Values need to be modified based on program state
- Working with complex data structures

## Common Mistakes

### 1. Trying to Modify Constants

```go
const Max = 100
Max = 200  // Error: cannot assign to Max
```

### 2. Using Runtime Values

```go
const Now = time.Now()  // Error: time.Now() is not a constant
```

### 3. Type Mismatches with Typed Constants

```go
const TypedInt int = 42
var f float64 = TypedInt  // Error: type mismatch
```

### 4. Misunderstanding iota Scope

```go
const (
    A = iota  // 0
    B         // 1
)
const C = iota  // 0 (new declaration, not 2)
```

### 5. Using Constants for Everything

Not everything should be a constant. If the value might change or comes from external sources, use a variable or configuration.

## Best Practices

1. **Use Meaningful Names**: `MaxRetries` instead of `N`
2. **Group Related Constants**: Use const blocks for related values
3. **Use iota for Enumerations**: Makes it easy to add/remove values
4. **Prefer Untyped for Flexibility**: Unless type safety is specifically needed
5. **Document Magic Numbers**: Replace magic numbers with named constants
6. **Use Type Definitions**: Create custom types for type-safe enumerations

Example of type-safe enumeration:
```go
type Weekday int

const (
    Sunday Weekday = iota
    Monday
    Tuesday
    Wednesday
    Thursday
    Friday
    Saturday
)

func IsWeekend(day Weekday) bool {
    return day == Sunday || day == Saturday
}
```

## Official Documentation

- [Go Constants Specification](https://go.dev/ref/spec#Constants)
- [Effective Go - Constants](https://go.dev/doc/effective_go#constants)
- [Go Blog - Constants](https://go.dev/blog/constants)

## Summary

Constants are a fundamental feature in Go that provide immutability, compile-time checking, and code clarity. Understanding the difference between typed and untyped constants, mastering const blocks, and using `iota` effectively will help you write cleaner, more maintainable Go code.
