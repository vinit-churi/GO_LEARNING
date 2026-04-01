# Variables and Types Reference

## Go's Basic Types

### Integer Types
Go provides several integer types of different sizes:

- **Signed integers:** `int8`, `int16`, `int32`, `int64`, `int`
- **Unsigned integers:** `uint8`, `uint16`, `uint32`, `uint64`, `uint`

The `int` and `uint` types are platform-dependent (32 or 64 bits).

**When to use:**
- Use `int` for most integer operations (it's the most efficient on your platform)
- Use specific sizes (`int32`, `int64`) when you need precise control or when working with binary protocols
- Use unsigned types (`uint`) when values can never be negative (e.g., counters, sizes)

**Common mistake:** Mixing different integer types requires explicit conversion.

```go
var a int32 = 10
var b int64 = 20
// result := a + b  // Error: type mismatch
result := int64(a) + b  // Correct: explicit conversion
```

### Floating-Point Types

- `float32`: 32-bit floating-point number
- `float64`: 64-bit floating-point number (preferred)

**When to use:**
- Use `float64` by default (better precision, and it's Go's default for float literals)
- Use `float32` only when memory is constrained or when interfacing with systems that require it

**Common mistake:** Floating-point arithmetic can have precision issues.

```go
var x float64 = 0.1 + 0.2
fmt.Println(x == 0.3)  // false due to floating-point precision
```

### String Type

`string` represents a sequence of bytes (typically UTF-8 encoded text).

**Key characteristics:**
- Strings are immutable (cannot be changed after creation)
- String literals use double quotes: `"hello"`
- Raw string literals use backticks: `` `hello\nworld` `` (backslash is literal)

**When to use:**
- Use strings for text data
- For byte manipulation, consider `[]byte` instead

**Common mistake:** Strings are not just character arrays.

```go
s := "hello"
// s[0] = 'H'  // Error: cannot assign to s[0]
s = "Hello"    // Correct: assign new string
```

### Boolean Type

`bool` represents a boolean value: `true` or `false`.

**When to use:**
- Use for logical conditions and flags
- Prefer descriptive names: `isValid`, `hasPermission`, `canProceed`

**Common mistake:** Go doesn't automatically convert other types to bool.

```go
var count int = 5
// if count { ... }  // Error: non-bool used as condition
if count > 0 { ... } // Correct: explicit comparison
```

## Variable Declaration Forms

### 1. `var` with Explicit Type

```go
var name string = "Alice"
var age int = 30
var height float64 = 5.8
var isStudent bool = true
```

**When to use:**
- When you want to be explicit about the type
- At package level (outside functions)
- When declaring without initialization

**Why it exists:**
Provides clarity and allows declaring variables at package scope.

### 2. `var` with Type Inference

```go
var name = "Alice"        // inferred as string
var age = 30              // inferred as int
var height = 5.8          // inferred as float64
var isStudent = true      // inferred as bool
```

**When to use:**
- When the type is obvious from the value
- When you prefer `var` style but want conciseness

**Common mistake:** Integer literals are inferred as `int`, not `int32` or `int64`.

```go
var x = 42      // type is int, not int32
var y int32 = x // Error: cannot use int as int32
```

### 3. Short Variable Declaration (`:=`)

```go
name := "Alice"
age := 30
height := 5.8
isStudent := true
```

**When to use:**
- Inside functions only (not at package level)
- When type is clear from context
- For concise, idiomatic Go code

**When NOT to use:**
- At package level (causes compile error)
- When you need an explicit type different from the default

**Common mistake:** `:=` declares a new variable. Using it on an existing variable creates a shadowing issue.

```go
name := "Alice"
name := "Bob"   // Error: no new variables on left side of :=
name = "Bob"    // Correct: assignment to existing variable
```

**Special case:** `:=` can redeclare if at least one variable is new:

```go
name, age := "Alice", 30
name, city := "Bob", "NYC"  // OK: city is new, name is reassigned
```

## Zero Values

In Go, variables declared without an explicit initial value are given their "zero value":

| Type      | Zero Value |
|-----------|------------|
| `int`     | `0`        |
| `float64` | `0.0`      |
| `string`  | `""`       |
| `bool`    | `false`    |
| pointers  | `nil`      |
| slices    | `nil`      |
| maps      | `nil`      |

**Why it exists:**
- Eliminates undefined behavior
- Makes code more predictable
- Reduces initialization bugs

**Example:**

```go
var count int       // count is 0
var message string  // message is ""
var valid bool      // valid is false

// This is safe:
count = count + 1   // count is now 1
```

**Common mistake:** Forgetting that numeric types start at 0, not some undefined value.

```go
var total int
for i := 0; i < 10; i++ {
    total += i  // Works correctly, total starts at 0
}
```

## Multiple Variable Declarations

### Block Style with `var`

```go
var (
    name   string  = "Alice"
    age    int     = 30
    height float64 = 5.8
)
```

**When to use:**
- Package-level variables
- Grouping related variables
- Improving readability

### Single Line Multiple Declaration

```go
var x, y, z int = 1, 2, 3
var name, age = "Alice", 30  // types inferred
```

**When to use:**
- Related variables of the same type
- Parallel assignment patterns

**With `:=`:**

```go
name, age, city := "Alice", 30, "NYC"
```

**When to use:**
- Inside functions
- When you want concise, related declarations

## Type Conversion

Go requires explicit type conversion:

```go
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
```

**Why it exists:**
Prevents accidental precision loss and overflow bugs.

**Common mistake:** Forgetting to convert when mixing types.

```go
var x int = 10
var y float64 = 3.14
// result := x + y           // Error
result := float64(x) + y     // Correct
```

## Best Practices

1. **Use `:=` inside functions** for concise code
2. **Use `var` at package level** (`:=` doesn't work there)
3. **Be explicit with types** when it aids clarity or when you need a specific size
4. **Use type inference** when the type is obvious
5. **Group related `var` declarations** in a block
6. **Choose descriptive names**: `userCount` not `uc`, `isValid` not `iv`
7. **Use zero values intentionally** - they're a feature, not a bug

## Links to Official Documentation

- [A Tour of Go - Variables](https://go.dev/tour/basics/8)
- [Effective Go - Variables](https://go.dev/doc/effective_go#variables)
- [Go Specification - Types](https://go.dev/ref/spec#Types)
- [Go Specification - Variables](https://go.dev/ref/spec#Variables)
