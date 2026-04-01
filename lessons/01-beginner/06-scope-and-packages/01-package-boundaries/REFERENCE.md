# Scope And Packages Reference

## Packages

Packages group related code. A package can expose functions, types, and variables to other packages by starting the name with an uppercase letter.

```go
package helper

func Header(name string) string { return name }
```

## Imports

Import a package when you want to use exported identifiers from it.

```go
import "strings"
```

## Scope

Variables declared inside a function or block only exist there.

```go
if trimmed := strings.TrimSpace(name); trimmed != "" {
    return len(trimmed) < 5
}
```

## Common Mistakes

- Trying to use an unexported name from another package
- Reusing a variable outside the block where it was declared
- Forgetting to trim input before comparing lengths
