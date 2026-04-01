# Interfaces Basics Reference

## Small Interfaces

Prefer focused interfaces with one clear job.

```go
type Previewer interface {
    Preview() string
}
```

## Implicit Implementation

If a type has the required methods, it satisfies the interface automatically.

## Common Mistakes

- Making interfaces too large too early
- Thinking a type needs an `implements` keyword in Go
