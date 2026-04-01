# Multiple Return Values - Reference

## Multiple Return Values Syntax

Go functions can return multiple values. The return types are specified in parentheses:

```go
func functionName(params) (returnType1, returnType2) {
    return value1, value2
}
```

Example:

```go
func getName() (string, string) {
    return "John", "Doe"
}

firstName, lastName := getName()
```

## Named Return Values

You can name the return values, which creates variables in the function scope:

```go
func divide(a, b float64) (result float64, err error) {
    if b == 0 {
        err = fmt.Errorf("division by zero")
        return // naked return uses named return values
    }
    result = a / b
    return
}
```

**Note:** While named returns can be useful for documentation, naked returns (return without values) can hurt readability in longer functions. Use them judiciously.

## The Error Pattern

The most common use of multiple return values in Go is returning a value alongside an error:

```go
func operation() (Result, error) {
    if somethingWrong {
        return Result{}, errors.New("something went wrong")
    }
    return result, nil
}

// Using the function
result, err := operation()
if err != nil {
    // handle error
    return err
}
// use result
```

**Convention:** The error is always the last return value.

## The Blank Identifier

Use `_` to discard values you don't need:

```go
// Only care about the error
_, err := someFunction()

// Only care about the first value
value, _ := someFunction()

// Discard middle values
first, _, third := threeValues()
```

**Warning:** Ignoring errors with `_` can lead to bugs. Only do this when:
- You're certain the function cannot fail
- The error is genuinely irrelevant to your use case
- You're in example/test code

## Multiple Return Values for Coordination

Multiple returns are useful for operations that naturally produce multiple related values:

```go
// Min and max together
func minMax(numbers []int) (min, max int) {
    min = numbers[0]
    max = numbers[0]
    for _, n := range numbers[1:] {
        if n < min {
            min = n
        }
        if n > max {
            max = n
        }
    }
    return
}

// Quotient and remainder
func divMod(a, b int) (quotient, remainder int) {
    return a / b, a % b
}

// Parsing with success flag
func parse(s string) (value int, ok bool) {
    if val, err := strconv.Atoi(s); err == nil {
        return val, true
    }
    return 0, false
}
```

## Why Multiple Return Values Exist

1. **Error Handling Without Exceptions:** Go doesn't use exceptions. Multiple returns allow functions to return both results and errors.

2. **Simplicity:** No need for output parameters or special error objects.

3. **Explicit:** Errors are values that must be explicitly checked, making error handling visible in the code.

4. **Performance:** Returning multiple values is efficient - no heap allocations needed.

## Common Patterns

### 1. Value and Error

The most common pattern in Go:

```go
func doSomething() (Result, error) {
    // ...
}

result, err := doSomething()
if err != nil {
    return err // or handle it
}
// use result
```

### 2. Value and Boolean (OK Pattern)

Used when failure is not exceptional:

```go
value, ok := map[key]  // map lookup
value, ok := <-channel // channel receive
value, ok := x.(Type)  // type assertion

if !ok {
    // handle missing value
}
```

### 3. Multiple Related Values

When operations naturally produce multiple values:

```go
func dimensions() (width, height int) {
    return 1920, 1080
}

func parseCoordinate(s string) (x, y int, err error) {
    // parse "x,y" format
    return
}
```

## Common Mistakes

1. **Ignoring errors:**
   ```go
   // Bad
   result, _ := riskyOperation()
   
   // Good
   result, err := riskyOperation()
   if err != nil {
       return err
   }
   ```

2. **Using the wrong order:**
   ```go
   // Wrong - error should be last
   func bad() (error, int)
   
   // Right
   func good() (int, error)
   ```

3. **Returning too many values:**
   ```go
   // Hard to read
   func tooMany() (int, int, int, int, string, error)
   
   // Better - use a struct
   type Result struct {
       A, B, C, D int
       Message    string
   }
   func better() (Result, error)
   ```

4. **Inconsistent error handling:**
   ```go
   // Inconsistent
   func inconsistent(x int) (int, error) {
       if x < 0 {
           return -1, errors.New("negative") // -1 is a valid value!
       }
       return x * 2, nil
   }
   
   // Better
   func consistent(x int) (int, error) {
       if x < 0 {
           return 0, errors.New("negative") // zero value on error
       }
       return x * 2, nil
   }
   ```

## When To Use Multiple Returns

**Use multiple returns when:**
- Returning a value and an error
- Returning closely related values (min/max, x/y coordinates)
- A boolean flag indicates success/presence (ok pattern)
- It makes the API simpler and clearer

**Don't use multiple returns when:**
- You have more than 3-4 return values (use a struct instead)
- The values aren't closely related (use separate functions)
- A single struct would be clearer

## Best Practices

1. **Always check errors:**
   ```go
   result, err := operation()
   if err != nil {
       // handle appropriately
   }
   ```

2. **Return zero values on error:**
   ```go
   func process() (Result, error) {
       if err := validate(); err != nil {
           return Result{}, err // zero value for Result
       }
       return result, nil
   }
   ```

3. **Error is always last:**
   ```go
   func standard() (value Type, err error)  // Yes
   func wrong() (err error, value Type)     // No
   ```

4. **Use named returns for documentation:**
   ```go
   func calculate() (sum, product int, err error)
   ```

## Additional Resources

- Go by Example: Multiple Return Values - https://gobyexample.com/multiple-return-values
- Effective Go: Multiple Returns - https://go.dev/doc/effective_go#multiple-returns
- Go Error Handling - https://go.dev/blog/error-handling-and-go
- Go Proverbs: "Errors are values" - https://go-proverbs.github.io/
