# Reference

## `if` Statements

Basic form:

```go
if condition {
	// do something
} else if otherCondition {
	// do something else
} else {
	// fallback
}
```

## Ordering Matters

The first matching branch wins. Put more specific conditions before more general ones.

## Combining Conditions

Use:

- `&&` for and
- `||` for or
- `!` for not

Example:

```go
if role == "admin" && active {
	return "admin access granted"
}
```

## Common Mistakes

- checking role before checking whether the account is inactive
- forgetting that premium shipping should short-circuit to `0`
- overlapping score conditions in the wrong order
