# Named Return Values - Reference

## What Are Named Return Values?

In Go, you can give names to the return values in a function signature. These names act as variables that are:
1. Automatically declared at the beginning of the function
2. Initialized to their zero values
3. Available for use throughout the function body
4. Returned when you use a naked `return` statement

### Basic Syntax

```go
// Without named returns
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

// With named returns
func divide(a, b int) (result int, err error) {
    if b == 0 {
        err = errors.New("division by zero")
        return // naked return
    }
    result = a / b
    return // naked return
}
```

## Naked Returns

A "naked return" or "bare return" is a `return` statement without explicit values. When you use a naked return, Go automatically returns the current values of the named return variables.

```go
func getCoordinates() (x, y int) {
    x = 10
    y = 20
    return // returns x=10, y=20
}
```

**Important:** You can still use explicit returns with named return values:

```go
func calculate(n int) (result int, err error) {
    if n < 0 {
        return 0, errors.New("negative input") // explicit return
    }
    result = n * 2
    return // naked return
}
```

## Why Named Returns Exist

### 1. Documentation

Named returns can serve as documentation for what the function returns:

```go
// Clear what each return value means
func openFile(path string) (file *File, err error)

// Less clear
func openFile(path string) (*File, error)
```

### 2. Automatic Initialization

Named returns are initialized to zero values, which can simplify code:

```go
func process() (count int, success bool) {
    // count is 0, success is false by default
    for _, item := range items {
        if doSomething(item) {
            count++
            success = true
        }
    }
    return
}
```

### 3. Defer Statements

Named returns work well with `defer` when you need to modify return values:

```go
func readFile(path string) (content []byte, err error) {
    f, err := os.Open(path)
    if err != nil {
        return
    }
    defer f.Close()
    
    content, err = io.ReadAll(f)
    return
}
```

With `defer`, you can even modify return values after they're set:

```go
func example() (err error) {
    defer func() {
        if err != nil {
            err = fmt.Errorf("wrapped: %w", err)
        }
    }()
    
    // function body
    return someOperation()
}
```

## When to Use Named Returns

### Good Use Cases

1. **Short Functions** - When the function fits on one screen and logic is simple
2. **Documentation Value** - When names clarify what's being returned
3. **Error Handling Pattern** - Common in Go to have `(result Type, err error)`
4. **Defer Modifications** - When you need to modify returns in defer statements

```go
// Good: Short and clear
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return
}

// Good: Clear documentation
func getUser(id int) (user *User, found bool, err error) {
    // implementation
}
```

### When NOT to Use Named Returns

1. **Long Functions** - Naked returns make it hard to track what's being returned
2. **Complex Logic** - Multiple return points with different values get confusing
3. **Obvious Returns** - When the meaning is clear without names

```go
// Bad: Function too long for naked returns
func complexOperation() (result int, err error) {
    // ... 50 lines of code ...
    result = something
    // ... 50 more lines ...
    return // What values are being returned? Unclear!
}

// Good: Explicit returns in long functions
func complexOperation() (result int, err error) {
    // ... 50 lines of code ...
    result = something
    // ... 50 more lines ...
    return result, err // Clear!
}

// Bad: Unnecessary names
func add(a, b int) (sum int) {
    sum = a + b
    return
}

// Better: Obvious without naming
func add(a, b int) int {
    return a + b
}
```

## Common Mistakes

### 1. Shadowing Named Returns

```go
// Bug: Shadowing the named return
func example() (result int, err error) {
    result = 5
    
    if needsProcessing() {
        // BUG: This declares a NEW local variable err, 
        // doesn't set the named return err
        result, err := process() // := creates new variables
        return // returns result=5, err=nil (NOT the err from process!)
    }
    
    return
}

// Fix: Use = instead of :=
func example() (result int, err error) {
    result = 5
    
    if needsProcessing() {
        result, err = process() // = assigns to named returns
        return
    }
    
    return
}
```

### 2. Forgetting Naked Return Returns Something

```go
// Bug: Naked return with unset values
func calculate(n int) (result int, err error) {
    if n < 0 {
        err = errors.New("negative")
        return // Returns result=0, err=error (result is unintended!)
    }
    // ... never set result ...
    return // Returns result=0, err=nil
}
```

### 3. Overusing Naked Returns

```go
// Bad: Naked return in long function is unclear
func process() (a, b, c int, err error) {
    // ... 100 lines of code ...
    return // What are we returning?
}

// Better: Explicit return is clearer
func process() (a, b, c int, err error) {
    // ... 100 lines of code ...
    return a, b, c, err // Crystal clear
}
```

## Go Style Guidelines

The official Go code review comments suggest:

- **Use named returns sparingly**
- Naked returns are acceptable **only in short functions**
- Named returns can improve documentation and work well with `defer`
- Explicit returns are generally clearer in longer functions
- **Clarity beats brevity**

From the Go wiki:
> "Don't name result parameters just to avoid declaring a var inside the function; that trades off a minor implementation brevity at the cost of unnecessary API verbosity."

## Comparison: Named vs Explicit

```go
// Named returns with naked return
func divide1(a, b float64) (result float64, err error) {
    if b == 0 {
        err = errors.New("division by zero")
        return
    }
    result = a / b
    return
}

// Named returns with explicit return (hybrid approach)
func divide2(a, b float64) (result float64, err error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    result = a / b
    return result, nil
}

// No named returns, explicit returns
func divide3(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}
```

All three are valid. Choose based on:
- Function length and complexity
- Team conventions
- Documentation needs
- Personal/team preference for clarity

## Additional Resources

- Go Tour - Named Results: https://go.dev/tour/basics/7
- Effective Go - Named Results: https://go.dev/doc/effective_go#named-results
- Go Code Review Comments: https://github.com/golang/go/wiki/CodeReviewComments#named-result-parameters
