# Multiple Return Values - Answers

## Exercise 1: Return Two Values

**Goal:** Return both minimum and maximum values from a slice.

**Solution Approach:**
- Initialize `min` and `max` to the first element
- Iterate through the slice, updating min and max as needed
- Return both values

See `solution.go` - the `minMax()` function.

**Key Concept:** This demonstrates the simplest case of multiple returns - two related values that are calculated together. Computing them in a single pass is more efficient than making two separate function calls.

**Edge Case:** The solution assumes a non-empty slice. In production code, you'd want to handle empty slices, perhaps by returning an error as a third value.

---

## Exercise 2: Return Value and Error

**Goal:** Implement safe division with error handling.

**Solution Approach:**
- Check if divisor is zero
- If zero, return 0 and an error
- Otherwise, perform division and return result with nil error

See `solution.go` - the `safeDivide()` function.

**Key Concept:** This is the most idiomatic use of multiple returns in Go. The pattern is:
1. Return value and error
2. Error is always the last return value
3. On error, return zero value and error
4. On success, return result and nil

**Why This Matters:** Go doesn't have exceptions. This pattern makes error handling explicit and forces you to consider error cases.

---

## Exercise 3: Parse User Input

**Goal:** Parse a formatted string and return multiple values with error handling.

**Solution Approach:**
- Use `strings.Split()` to separate the input
- Validate the format (should have exactly 2 parts)
- Use `strconv.Atoi()` to convert age string to int
- Handle conversion errors
- Return all three values: name, age, and error

See `solution.go` - the `parseUserInput()` function.

**Key Concepts:**
- Functions can return more than 2 values
- Each parsing step can fail, requiring error checks
- Build up the result incrementally
- Return zero values ("", 0) when there's an error

**Alternative Approach:** You could return a struct instead:
```go
type User struct {
    Name string
    Age  int
}
func parseUser(input string) (User, error)
```

This is a matter of preference. Use multiple returns for 2-3 simple values; use structs for more complex data.

---

## Exercise 4: Swap Values

**Goal:** Swap two values in a single function call.

**Solution Approach:**
- Simply return the parameters in reversed order

See `solution.go` - the `swap()` function.

**Key Concept:** This showcases how multiple returns can simplify operations that would require temporary variables or complex logic in other languages.

In other languages:
```java
// Java requires a temporary variable
String temp = a;
a = b;
b = temp;
```

In Go with multiple returns:
```go
a, b = swap(a, b)  // or just: a, b = b, a
```

**Note:** While this function works, in practice you'd usually just use: `a, b = b, a` directly. This exercise demonstrates the concept.

---

## Exercise 5: Calculate Statistics

**Goal:** Return multiple statistical values in one pass.

**Solution Approach:**
- Check for empty slice and return error if needed
- Initialize sum to 0 and count to slice length
- Iterate once to calculate sum
- Calculate average as float64(sum) / float64(count)
- Return all four values

See `solution.go` - the `calculateStats()` function.

**Key Concepts:**
- Single iteration to gather multiple statistics is efficient
- When returning 3+ values, consider if a struct would be clearer
- Empty input is an error condition
- Converting to float64 for average calculation

**When Multiple Returns Get Unwieldy:**

This function returns 4 values. That's approaching the limit of readability. Consider:

```go
type Stats struct {
    Sum   int
    Avg   float64
    Count int
}

func calculateStats(numbers []int) (Stats, error)
```

**Benefits of struct approach:**
- Named fields are self-documenting
- Easier to extend with more statistics
- Can add methods to the Stats type
- Cleaner call site: `stats, err := calculateStats(nums)`

**Trade-offs:**
- Multiple returns: Simpler for small number of related values
- Structs: Better for complex return data or when you might extend it later

---

## General Patterns and Best Practices

### 1. Error Handling Pattern

Always check errors before using the result:

```go
result, err := someFunction()
if err != nil {
    // handle error first
    return err
}
// now safe to use result
```

### 2. Ignoring Values Selectively

Use `_` for values you don't need:

```go
// Only want the max
_, max := minMax(numbers)

// Only checking if operation succeeded
_, err := someOperation()
if err != nil {
    return err
}
```

### 3. Unpacking All Values

Sometimes you need all values:

```go
min, max := minMax(numbers)
fmt.Printf("Range: %d to %d\n", min, max)
```

### 4. Named Returns for Documentation

Named returns make the function signature clearer:

```go
func calculateStats(numbers []int) (sum int, avg float64, count int, err error)
```

Anyone reading this immediately knows what each return value represents.

---

## Common Pitfalls Explained

### 1. Not Checking Errors

```go
// Dangerous!
result, _ := safeDivide(x, y)
// If y was 0, result is meaningless but we continue anyway
```

### 2. Using Returned Value When There's an Error

```go
result, err := safeDivide(10, 0)
fmt.Println(result) // result is 0 (zero value), but we shouldn't use it!
if err != nil {
    return err
}
```

**Rule:** When a function returns an error, only use the other return values if err is nil.

### 3. Wrong Order of Return Values

```go
// Wrong - error should be last by convention
func badStyle() (error, int) {
    return nil, 42
}

// Correct
func goodStyle() (int, error) {
    return 42, nil
}
```

---

## Testing Tips

When testing functions with multiple returns:

```go
func TestMinMax(t *testing.T) {
    min, max := minMax([]int{1, 5, 3})
    
    if min != 1 {
        t.Errorf("expected min=1, got %d", min)
    }
    
    if max != 5 {
        t.Errorf("expected max=5, got %d", max)
    }
}
```

For error testing:

```go
func TestSafeDivideError(t *testing.T) {
    _, err := safeDivide(10, 0)
    
    if err == nil {
        t.Error("expected error for division by zero")
    }
}
```

---

## Next Steps

After mastering multiple return values, explore:

1. **Defer statements** - useful with cleanup after error checking
2. **Custom error types** - creating more informative errors
3. **Error wrapping** - `fmt.Errorf` with `%w` for error chains
4. **Variadic functions** - functions that accept variable number of arguments
5. **Method receivers** - functions attached to types

Multiple return values are fundamental to idiomatic Go. They appear everywhere in Go code, especially in the standard library. Practice using them until the pattern feels natural!
