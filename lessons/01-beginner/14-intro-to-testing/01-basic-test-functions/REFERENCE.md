# Intro To Testing Reference

## Test Function Shape

```go
func TestSomething(t *testing.T) {
    if got != want {
        t.Fatalf("got %v want %v", got, want)
    }
}
```

## Running Tests

Use `go test` in the lesson directory or `go test ./...` at the repo root.

## What To Improve After A Failure

- Check the input in the failing test
- Compare actual and expected output
- Fix the implementation, then run the tests again
