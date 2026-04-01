# Reference

## Function Signatures

Functions declare parameter types and return types:

```go
func subtotal(prices []float64) float64
```

## Composition

Instead of writing one large function, create smaller functions and call them from the final function.

Example flow:

1. compute subtotal
2. compute final total
3. format the summary string

## Working With Slices

Slices are passed by descriptor, so iterating over them in a function is efficient for small exercises like this.

## Common Mistakes

- duplicating subtotal logic inside `orderSummary`
- forgetting to format floats with `%.2f`
- applying the discount in the wrong direction
