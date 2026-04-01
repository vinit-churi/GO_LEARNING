# Arrays And Slices Reference

## Arrays

Arrays store a fixed number of values.

```go
scores := [3]int{10, 20, 30}
```

## Slices

Slices are flexible and support `append`.

```go
tasks := []string{"plan", "code"}
tasks = append(tasks, "test")
```

## Slicing

Use `a[start:end]` to create a slice window.

```go
weekend := days[5:7]
```

## Common Mistakes

- Mixing up array types with different lengths
- Assuming `append` changes the original variable without assignment
- Using the wrong slice indexes
