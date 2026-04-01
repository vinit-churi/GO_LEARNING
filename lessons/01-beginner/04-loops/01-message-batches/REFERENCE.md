# Reference

## Counted Loops

```go
for i := 0; i < 5; i++ {
	// repeated work
}
```

This loop has:

- initialization: `i := 0`
- condition: `i < 5`
- post statement: `i++`

## Accumulators

When you want a final total, create a variable before the loop and update it each iteration.

```go
total := 0
for i := 1; i <= days; i++ {
	total += i
}
```

## Building Slices

Use `append` to grow a slice:

```go
values := []int{}
values = append(values, 3)
```

## Common Mistakes

- off-by-one errors in loop boundaries
- forgetting to initialize the result slice
- summing odd numbers when the problem asks for even ones only
