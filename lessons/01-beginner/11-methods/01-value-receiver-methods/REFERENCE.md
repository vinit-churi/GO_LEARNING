# Methods Reference

## Value Receivers

Use value receivers for behavior that reads data but does not need to change the original value.

```go
func (c Counter) Summary() string { return "" }
```

## Pointer Receivers

Use pointer receivers when the method should update the struct.

```go
func (c *Counter) Add(n int) { c.Value += n }
```

## Common Mistakes

- Expecting a value receiver to mutate the original struct
- Mixing receiver styles without a reason
