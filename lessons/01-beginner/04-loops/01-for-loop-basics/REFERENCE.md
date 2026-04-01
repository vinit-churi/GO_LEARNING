# For Loop Reference

## Overview

Go's `for` loop is the only looping construct in the language, but it's powerful enough to handle all iteration scenarios. This simplicity is by design—it reduces cognitive load and makes code more consistent.

## Basic Syntax

### Classic Three-Component For Loop

```go
for initialization; condition; post {
    // loop body
}
```

**Example:**
```go
for i := 0; i < 5; i++ {
    fmt.Println(i)  // Prints 0, 1, 2, 3, 4
}
```

### While-Style Loop (Condition Only)

```go
for condition {
    // loop body
}
```

**Example:**
```go
count := 0
for count < 5 {
    fmt.Println(count)
    count++
}
```

### Infinite Loop

```go
for {
    // loop body
    if someCondition {
        break
    }
}
```

**Example:**
```go
for {
    input := getInput()
    if input == "quit" {
        break
    }
    process(input)
}
```

### Range-Based Loop (covered in future lessons)

```go
for index, value := range collection {
    // loop body
}
```

## Control Flow Statements

### Break

Exits the innermost loop immediately.

```go
for i := 0; i < 10; i++ {
    if i == 5 {
        break  // Loop stops when i is 5
    }
    fmt.Println(i)  // Prints 0, 1, 2, 3, 4
}
```

### Continue

Skips the rest of the current iteration and moves to the next one.

```go
for i := 0; i < 5; i++ {
    if i == 2 {
        continue  // Skip when i is 2
    }
    fmt.Println(i)  // Prints 0, 1, 3, 4
}
```

## Nested Loops

Loops can be nested to handle multi-dimensional data structures.

```go
for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
        fmt.Printf("(%d,%d) ", i, j)
    }
    fmt.Println()
}
// Output:
// (0,0) (0,1) (0,2)
// (1,0) (1,1) (1,2)
// (2,0) (2,1) (2,2)
```

**Important:** `break` and `continue` only affect the innermost loop.

## Common Patterns

### Counting Up
```go
for i := 0; i < n; i++ {
    // Process i
}
```

### Counting Down
```go
for i := n - 1; i >= 0; i-- {
    // Process i
}
```

### Stepping by Custom Amount
```go
for i := 0; i < 100; i += 10 {
    fmt.Println(i)  // Prints 0, 10, 20, ..., 90
}
```

### Loop Until Success
```go
for {
    result, err := tryOperation()
    if err == nil {
        break  // Success, exit loop
    }
    time.Sleep(time.Second)  // Wait before retry
}
```

## Common Mistakes

### 1. Off-by-One Errors

**Wrong:**
```go
for i := 0; i <= len(slice); i++ {  // Will panic on last iteration
    fmt.Println(slice[i])
}
```

**Correct:**
```go
for i := 0; i < len(slice); i++ {
    fmt.Println(slice[i])
}
```

### 2. Modifying Loop Variable Incorrectly

**Wrong:**
```go
for i := 0; i < 10; i++ {
    i = 0  // Infinite loop! Resets counter each time
}
```

**Correct:**
```go
for i := 0; i < 10; i++ {
    // Don't modify i inside the loop
}
```

### 3. Scope Issues

The loop variable is scoped to the loop:

```go
for i := 0; i < 5; i++ {
    // i is accessible here
}
// i is NOT accessible here
```

### 4. Forgetting to Advance in While-Style Loops

**Wrong:**
```go
i := 0
for i < 5 {
    fmt.Println(i)
    // Forgot i++, infinite loop!
}
```

**Correct:**
```go
i := 0
for i < 5 {
    fmt.Println(i)
    i++
}
```

### 5. Breaking Out of Wrong Loop

In nested loops, `break` only exits the innermost loop:

```go
for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
        if j == 1 {
            break  // Only exits inner loop
        }
    }
}
```

Use labeled breaks for outer loops:

```go
outer:
for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
        if j == 1 {
            break outer  // Exits both loops
        }
    }
}
```

## Why For Loops Exist

Loops are fundamental to programming because they allow you to:

1. **Process collections**: Iterate through arrays, slices, and other data structures
2. **Repeat operations**: Execute code multiple times without duplication
3. **Search and filter**: Find specific elements or create subsets
4. **Generate sequences**: Create patterns, tables, or computed values
5. **Wait for conditions**: Poll until a state changes

## When to Use Each Pattern

### Classic Three-Component For Loop
- When you need an index counter
- When iterating a specific number of times
- When you need precise control over initialization and increment

### While-Style Loop (Condition Only)
- When the loop count isn't known in advance
- When waiting for an external condition
- When the loop logic doesn't fit the three-component pattern

### Infinite Loop
- For event loops or servers
- When exit conditions are complex or multiple
- When using break makes the logic clearer

### Nested Loops
- Processing 2D data (matrices, grids)
- Generating combinations
- Comparing all pairs of elements

## Performance Considerations

1. **Cache loop bounds**: If calculating the condition is expensive, cache it

```go
// Less efficient
for i := 0; i < expensiveFunction(); i++ {
    // ...
}

// More efficient
limit := expensiveFunction()
for i := 0; i < limit; i++ {
    // ...
}
```

2. **Avoid unnecessary work inside loops**: Move invariant calculations outside

3. **Consider early exits**: Use break when you've found what you need

## Further Reading

- [Go Tour: For loops](https://go.dev/tour/flowcontrol/1)
- [Effective Go: For loops](https://go.dev/doc/effective_go#for)
- [Go Specification: For statements](https://go.dev/ref/spec#For_statements)
