# If Statements - Reference

## Basic If Statement

The simplest form checks a condition and executes code if true:

```go
if x > 0 {
    fmt.Println("x is positive")
}
```

**Key Points:**
- No parentheses needed around the condition
- Braces are required even for single statements
- The condition must be a boolean expression

## If-Else Statement

Execute one block if the condition is true, another if false:

```go
if x > 0 {
    fmt.Println("positive")
} else {
    fmt.Println("not positive")
}
```

**Key Points:**
- The `else` must be on the same line as the closing brace of the `if` block
- This is required by Go's automatic semicolon insertion

## If-Else If-Else Chain

Check multiple conditions in sequence:

```go
if x > 0 {
    fmt.Println("positive")
} else if x < 0 {
    fmt.Println("negative")
} else {
    fmt.Println("zero")
}
```

**Key Points:**
- Conditions are checked in order from top to bottom
- Only the first matching condition's block executes
- The final `else` is optional and catches all other cases

## If With Short Statement

Declare and initialize a variable within the if statement:

```go
if x := getValue(); x > 0 {
    fmt.Println(x, "is positive")
}
// x is not accessible here
```

**Key Points:**
- The variable is scoped only to the if block (including else if and else)
- Useful for limiting variable scope and avoiding cluttering the outer scope
- Common pattern when checking results of function calls

### Common Use Case: Error Handling

```go
if err := doSomething(); err != nil {
    return err
}
// err is not accessible here
```

## Boolean Operators

### Comparison Operators
- `==` equal to
- `!=` not equal to
- `<` less than
- `<=` less than or equal to
- `>` greater than
- `>=` greater than or equal to

### Logical Operators
- `&&` logical AND (both conditions must be true)
- `||` logical OR (at least one condition must be true)
- `!` logical NOT (negates a boolean)

### Examples

```go
// AND: both must be true
if x > 0 && x < 10 {
    fmt.Println("x is between 0 and 10")
}

// OR: at least one must be true
if x < 0 || x > 100 {
    fmt.Println("x is out of normal range")
}

// NOT: negates the condition
if !(x > 0) {
    fmt.Println("x is not positive")
}

// Complex combinations
if (x > 0 && x < 10) || x == 100 {
    fmt.Println("x is 0-10 or exactly 100")
}
```

## Parentheses (Optional)

While not required, you can use parentheses for clarity:

```go
if (x > 0) && (x < 10) {
    fmt.Println("x is between 0 and 10")
}
```

## Why If Statements Exist

If statements are fundamental to programming because:
- Programs need to make decisions based on data
- Different inputs require different processing
- They enable dynamic behavior based on runtime conditions
- They're the building blocks for complex logic

## When To Use Each Pattern

### Basic If
Use when you only need to do something in the true case:
```go
if isValid {
    process()
}
```

### If-Else
Use when you have exactly two outcomes:
```go
if isEven {
    fmt.Println("even")
} else {
    fmt.Println("odd")
}
```

### If-Else If-Else
Use when you have multiple mutually exclusive conditions:
```go
if score >= 90 {
    grade = "A"
} else if score >= 80 {
    grade = "B"
} else if score >= 70 {
    grade = "C"
} else {
    grade = "F"
}
```

### If With Short Statement
Use when:
- A variable is only needed for the if condition
- You want to keep scope limited
- Checking function return values

```go
if result := compute(); result > threshold {
    return result
}
```

## When NOT To Use

### Avoid Deep Nesting
Don't create deeply nested if statements:

```go
// Bad: hard to read and maintain
if condition1 {
    if condition2 {
        if condition3 {
            if condition4 {
                // do something
            }
        }
    }
}
```

Instead, use early returns or combine conditions:

```go
// Better
if !condition1 {
    return
}
if !condition2 {
    return
}
if !condition3 {
    return
}
if !condition4 {
    return
}
// do something
```

### Switch Is Better For Many Values
If checking the same variable against many values, use `switch` instead:

```go
// Awkward
if day == "Monday" {
    // ...
} else if day == "Tuesday" {
    // ...
} else if day == "Wednesday" {
    // ...
}

// Better
switch day {
case "Monday":
    // ...
case "Tuesday":
    // ...
case "Wednesday":
    // ...
}
```

## Common Mistakes

1. **Forgetting braces:**
   ```go
   // Wrong: won't compile
   if x > 0
       fmt.Println("positive")
   
   // Correct
   if x > 0 {
       fmt.Println("positive")
   }
   ```

2. **Putting else on wrong line:**
   ```go
   // Wrong: won't compile
   if x > 0 {
       fmt.Println("positive")
   }
   else {
       fmt.Println("not positive")
   }
   
   // Correct
   if x > 0 {
       fmt.Println("positive")
   } else {
       fmt.Println("not positive")
   }
   ```

3. **Using assignment instead of comparison:**
   ```go
   // Wrong: assigns 5 to x (won't compile in Go)
   if x = 5 {
       // ...
   }
   
   // Correct: compares x to 5
   if x == 5 {
       // ...
   }
   ```

4. **Non-boolean condition:**
   ```go
   // Wrong: integers are not booleans in Go
   if count {
       // ...
   }
   
   // Correct: explicit comparison
   if count > 0 {
       // ...
   }
   ```

5. **Confusing scope with short statement:**
   ```go
   if x := 5; x > 0 {
       fmt.Println(x) // OK
   }
   fmt.Println(x) // Error: x is not defined here
   ```

## Additional Resources

- Go Tour on Flow Control: https://go.dev/tour/flowcontrol/1
- Effective Go - Control Structures: https://go.dev/doc/effective_go#control-structures
- Go Spec - If Statements: https://go.dev/ref/spec#If_statements
