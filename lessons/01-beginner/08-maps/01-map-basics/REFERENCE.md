# Maps Reference

## Creating Maps

```go
scores := map[string]int{
    "Ava": 10,
}
```

## Writing And Reading

```go
scores["Ava"] = 12
value, ok := scores["Ava"]
```

## Counting Pattern

Increment the current count on every visit.

```go
counts[word]++
```

## Common Mistakes

- Forgetting that reading a missing key returns the zero value
- Assuming map iteration order is stable
- Writing to a nil map
