# Hello World - Reference

## The Package Declaration

Every Go file begins with a `package` declaration:

```go
package main
```

- The `main` package is special - it defines an executable program
- Other packages create reusable libraries
- The package name must match across all files in the same directory

## The Import Statement

To use code from other packages, you import them:

```go
import "fmt"
```

Multiple imports can be grouped:

```go
import (
    "fmt"
    "strings"
)
```

## The Main Function

```go
func main() {
    // program starts here
}
```

- `func` is the keyword to declare a function
- `main` is the entry point - execution starts here
- The main function takes no arguments and returns nothing
- Only the `main` package can have a `main` function

## The fmt Package

`fmt` stands for "format" and is part of Go's standard library. Common functions:

- `fmt.Println()` - prints arguments with a newline
- `fmt.Print()` - prints arguments without a newline
- `fmt.Printf()` - prints with formatting (covered later)

Example:

```go
fmt.Println("Hello")        // prints: Hello\n
fmt.Print("Hello")          // prints: Hello
fmt.Print("Hello", "World") // prints: HelloWorld
fmt.Println("Hi", "there")  // prints: Hi there\n (space between args)
```

## String Literals

Strings are enclosed in double quotes:

```go
"This is a string"
```

Go also supports raw string literals using backticks:

```go
`This is a raw string
It can span multiple lines
And doesn't interpret escape sequences`
```

## Why This Structure?

- **Explicit package system:** Makes dependencies clear and avoids naming conflicts
- **Standard entry point:** Every Go developer knows where program execution begins
- **No global code:** All code must be inside functions, making programs more predictable
- **Standard library:** Rich built-in packages like `fmt` reduce the need for external dependencies

## Common Mistakes

1. **Forgetting package declaration:** Every `.go` file needs `package` at the top
2. **Wrong package name:** Executables must use `package main`
3. **No main function:** The `main` package must have a `main()` function
4. **Unused imports:** Go won't compile if you import packages you don't use
5. **Single quotes for strings:** Use `"double quotes"` not `'single quotes'` (single quotes are for runes)

## When To Use

- Every executable Go program needs a `main` package and `main` function
- `fmt.Println` is perfect for:
  - Debugging output
  - Simple programs
  - Learning exercises
  - Quick scripts

For production applications, consider using structured logging libraries instead of printing directly to console.

## Additional Resources

- Official Go Documentation: https://go.dev/doc/
- fmt package: https://pkg.go.dev/fmt
- Effective Go: https://go.dev/doc/effective_go
