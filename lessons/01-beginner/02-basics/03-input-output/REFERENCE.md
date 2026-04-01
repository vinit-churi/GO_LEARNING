# Input/Output Reference

## Overview

Go's `fmt` package provides functionality for formatted I/O, similar to C's `printf` and `scanf`. It's part of the standard library and is one of the most commonly used packages in Go programs.

## Print Functions

### fmt.Print

```go
fmt.Print(a ...interface{}) (n int, err error)
```

- Prints its arguments with their default format
- Spaces are added between operands when neither is a string
- Returns the number of bytes written and any error

**Example:**
```go
fmt.Print("Hello")
fmt.Print(" ")
fmt.Print("World")
// Output: Hello World
```

### fmt.Println

```go
fmt.Println(a ...interface{}) (n int, err error)
```

- Similar to `Print` but adds spaces between arguments and a newline at the end
- Most convenient for simple output

**Example:**
```go
fmt.Println("Hello", "World")
// Output: Hello World
// (with newline)
```

### fmt.Printf

```go
fmt.Printf(format string, a ...interface{}) (n int, err error)
```

- Formats according to a format specifier and writes to standard output
- Does not add a newline unless specified in the format string
- Most flexible printing function

**Example:**
```go
name := "Alice"
age := 30
fmt.Printf("Name: %s, Age: %d\n", name, age)
// Output: Name: Alice, Age: 30
```

## Format Verbs

Format verbs are used with `fmt.Printf`, `fmt.Sprintf`, and similar functions.

### Common Format Verbs

| Verb | Description | Example |
|------|-------------|---------|
| `%v` | Default format | `fmt.Printf("%v", 42)` → `42` |
| `%+v` | Default format with field names (for structs) | `fmt.Printf("%+v", person)` |
| `%#v` | Go-syntax representation | `fmt.Printf("%#v", "hello")` → `"hello"` |
| `%T` | Type of value | `fmt.Printf("%T", 42)` → `int` |
| `%%` | Literal percent sign | `fmt.Printf("%%")` → `%` |

### String and Byte Slice Verbs

| Verb | Description | Example |
|------|-------------|---------|
| `%s` | String or slice of bytes | `fmt.Printf("%s", "hello")` → `hello` |
| `%q` | Quoted string (escaped) | `fmt.Printf("%q", "hello")` → `"hello"` |
| `%x` | Hex encoding | `fmt.Printf("%x", "hello")` → `68656c6c6f` |

### Integer Verbs

| Verb | Description | Example |
|------|-------------|---------|
| `%d` | Decimal | `fmt.Printf("%d", 42)` → `42` |
| `%b` | Binary | `fmt.Printf("%b", 42)` → `101010` |
| `%o` | Octal | `fmt.Printf("%o", 42)` → `52` |
| `%x` | Hexadecimal (lowercase) | `fmt.Printf("%x", 42)` → `2a` |
| `%X` | Hexadecimal (uppercase) | `fmt.Printf("%X", 42)` → `2A` |

### Floating-Point Verbs

| Verb | Description | Example |
|------|-------------|---------|
| `%f` | Decimal point, no exponent | `fmt.Printf("%f", 3.14)` → `3.140000` |
| `%e` | Scientific notation | `fmt.Printf("%e", 3.14)` → `3.140000e+00` |
| `%g` | `%e` or `%f`, whichever is shorter | `fmt.Printf("%g", 3.14)` → `3.14` |

### Width and Precision

You can specify width and precision:

```go
fmt.Printf("%5d", 42)        // "   42" (width 5, right-aligned)
fmt.Printf("%-5d", 42)       // "42   " (width 5, left-aligned)
fmt.Printf("%05d", 42)       // "00042" (width 5, zero-padded)
fmt.Printf("%.2f", 3.14159)  // "3.14" (2 decimal places)
fmt.Printf("%8.2f", 3.14159) // "    3.14" (width 8, 2 decimal places)
```

## Input Functions

### fmt.Scan

```go
fmt.Scan(a ...interface{}) (n int, err error)
```

- Scans text from standard input
- Stores values into arguments (must be pointers)
- Whitespace-delimited values

**Example:**
```go
var name string
var age int
fmt.Scan(&name, &age)
```

### fmt.Scanln

```go
fmt.Scanln(a ...interface{}) (n int, err error)
```

- Similar to `Scan` but stops at newline
- Requires all items to be on one line

### fmt.Scanf

```go
fmt.Scanf(format string, a ...interface{}) (n int, err error)
```

- Scans with a format string
- More control over input parsing

**Example:**
```go
var name string
var age int
fmt.Scanf("%s %d", &name, &age)
```

## Common Mistakes

### 1. Forgetting the Newline

```go
// Wrong - no newline
fmt.Printf("Hello")

// Correct
fmt.Printf("Hello\n")
// or
fmt.Println("Hello")
```

### 2. Wrong Format Verb

```go
age := 30
// Wrong - %s is for strings
fmt.Printf("Age: %s", age) // Will print: Age: %!s(int=30)

// Correct
fmt.Printf("Age: %d", age)
```

### 3. Forgetting Pointers with Scan

```go
var name string
// Wrong - passing value instead of pointer
fmt.Scan(name) // Will not modify name

// Correct
fmt.Scan(&name)
```

### 4. Not Checking Scan Errors

```go
// Better practice
var age int
n, err := fmt.Scan(&age)
if err != nil {
    // Handle error
}
```

## Why These Features Exist

### Type Safety with Format Verbs

Go's format verbs provide type-safe formatted output. The compiler can't check format strings at compile time, but at runtime, Go will detect mismatches and print error strings like `%!d(string=hello)`.

### Variadic Functions

The `Print` family accepts variadic arguments (`...interface{}`), making them flexible and easy to use for any combination of types.

### Default Formatting with %v

The `%v` verb uses each type's default format, which is defined by the type's `String()` method if it exists. This makes debugging easier without remembering specific verbs.

## When to Use What

### Use fmt.Println when:
- You want simple output with automatic spacing and newlines
- You're debugging and want quick output
- Output formatting doesn't matter

### Use fmt.Printf when:
- You need precise control over output format
- You're creating user-facing output
- You need to align output in columns
- You're formatting numbers with specific precision

### Use fmt.Print when:
- You want to print multiple items without automatic newlines
- You're building output incrementally
- Less common than the other two

### Use %v when:
- You want the default format
- You're debugging and don't care about specific formatting
- You want code that works with any type

### Use specific verbs (%d, %s, %f) when:
- You need precise formatting
- You want to catch type errors
- You're formatting user-facing output

## Further Reading

- [Official fmt package documentation](https://pkg.go.dev/fmt)
- [Effective Go - Printing](https://go.dev/doc/effective_go#printing)
