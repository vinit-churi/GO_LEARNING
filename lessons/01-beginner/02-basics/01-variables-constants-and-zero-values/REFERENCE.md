# Reference

## Constants

Use `const` for values that should not change during program execution.

```go
const maxDailyPracticeMinutes = 90
```

## Zero Values

Go initializes declared values automatically.

- `string` -> `""`
- `int` -> `0`
- `bool` -> `false`

This is especially useful for structs because missing fields still receive safe defaults.

## Struct Initialization

Named field initialization is explicit and readable:

```go
session := PracticeSession{
	Name: "loops",
}
```

## Common Mistakes

- expecting uninitialized variables to be `nil` when the type is not a reference type
- forgetting to clamp a result to `0`
- mutating input when a fresh value is easier to reason about
