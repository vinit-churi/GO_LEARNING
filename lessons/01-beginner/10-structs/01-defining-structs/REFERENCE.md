# Structs Reference

## Defining A Struct

```go
type Book struct {
    Title string
    Pages int
}
```

## Nested Structs

Struct fields can hold other structs when values belong together.

```go
type Address struct {
    City string
}
```

## Common Mistakes

- Forgetting to initialize all required fields
- Passing too many separate parameters instead of creating a struct
- Using anonymous data when a named struct would be clearer
