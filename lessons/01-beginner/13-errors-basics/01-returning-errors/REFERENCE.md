# Errors Basics Reference

## Returning Errors

```go
func parse(input string) (int, error) {
    return 0, nil
}
```

## Creating Errors

Use `errors.New` for a fixed message and `fmt.Errorf` when the message should include values.

## Handling Errors

Check the error before using the returned data.

```go
value, err := parse("12")
if err != nil {
    return err
}
```

## Common Mistakes

- Ignoring an error return
- Returning vague messages that do not explain the problem
