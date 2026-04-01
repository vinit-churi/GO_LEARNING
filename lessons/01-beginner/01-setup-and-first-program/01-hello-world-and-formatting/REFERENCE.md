# Reference

## Key Concepts

Go programs organize logic into packages and functions. For small single-folder exercises, `package main` is the simplest choice and allows you to run the folder with `go run .`.

## `fmt.Sprintf`

`fmt.Sprintf` creates a formatted string without printing it.

Examples:

```go
message := fmt.Sprintf("Hello, %s!", "Gopher")
count := fmt.Sprintf("%d lessons complete", 4)
```

Common verbs:

- `%s` for strings
- `%d` for integers
- `%.2f` for a float with 2 decimal places

## Booleans

Boolean values are `true` and `false`. A comparison such as `completed >= 3` directly returns a boolean.

## Common Mistakes

- using `fmt.Println` when the function needs to return a string
- forgetting punctuation required by the test
- confusing printing output with returning values
