# Switch Statements - Reference Guide

## What is a Switch Statement?

A `switch` statement is a control structure that allows you to compare a value against multiple possibilities and execute different code based on which case matches. It's a cleaner and more readable alternative to long chains of if-else statements.

## Basic Syntax

```go
switch expression {
case value1:
    // executed when expression == value1
case value2:
    // executed when expression == value2
case value3, value4:
    // executed when expression == value3 OR value4
default:
    // executed when no case matches (optional)
}
```

## Key Features of Go's Switch

### 1. No Automatic Fallthrough

Unlike C, Java, or JavaScript, Go's switch cases **do not fall through** by default. This is a safety feature that prevents common bugs.

```go
switch day {
case "Monday":
    fmt.Println("Start of week")
    // Automatically breaks here - does NOT continue to Tuesday!
case "Tuesday":
    fmt.Println("Second day")
}
```

### 2. Multiple Values in One Case

You can test multiple values in a single case using commas:

```go
switch day {
case "Saturday", "Sunday":
    fmt.Println("Weekend!")
case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
    fmt.Println("Weekday")
}
```

### 3. Switch Without Expression (Expression-less Switch)

When you omit the expression, each case evaluates a boolean condition. This is like a cleaner if-else chain:

```go
score := 85

switch {
case score >= 90:
    grade = "A"
case score >= 80:
    grade = "B"
case score >= 70:
    grade = "C"
case score >= 60:
    grade = "D"
default:
    grade = "F"
}
```

This is more readable than:
```go
if score >= 90 {
    grade = "A"
} else if score >= 80 {
    grade = "B"
} else if score >= 70 {
    grade = "C"
// ... and so on
```

### 4. The Fallthrough Keyword

If you explicitly want a case to continue to the next case, use `fallthrough`:

```go
switch num {
case 1:
    fmt.Println("One")
    fallthrough
case 2:
    fmt.Println("Two or less")
}

// When num == 1, prints:
// One
// Two or less
```

**Important:** `fallthrough` transfers control to the next case **without** checking its condition. Use it sparingly.

### 5. Switch with Short Statement

You can include a short statement before the switch expression:

```go
switch value := getSomeValue(); value {
case 1:
    fmt.Println("It's one")
case 2:
    fmt.Println("It's two")
}
// value is only in scope within the switch
```

## When to Use Switch vs If-Else

### Use Switch When:
- Comparing one value against many constants
- You have 3+ conditions checking the same variable
- Cases are discrete values or simple ranges
- Code readability would improve with a switch

### Use If-Else When:
- Conditions are complex or unrelated
- You're checking different variables in each condition
- You have only 1-2 simple conditions

## Common Patterns

### Pattern 1: Type Switch (for interfaces)
```go
switch v := value.(type) {
case int:
    fmt.Println("Integer:", v)
case string:
    fmt.Println("String:", v)
default:
    fmt.Println("Unknown type")
}
```

### Pattern 2: Range Checking
```go
switch {
case age < 13:
    category = "child"
case age < 20:
    category = "teenager"
case age < 60:
    category = "adult"
default:
    category = "senior"
}
```

### Pattern 3: String Matching with Multiple Options
```go
switch command {
case "start", "run", "begin":
    startProcess()
case "stop", "end", "terminate":
    stopProcess()
default:
    fmt.Println("Unknown command")
}
```

## Common Mistakes

### Mistake 1: Expecting Fallthrough
```go
// Wrong expectation - this only prints "Weekend"
switch day {
case "Saturday":
    fmt.Println("Weekend")
case "Sunday":
    fmt.Println("Still weekend")
}
```

**Fix:** Use multiple values in one case:
```go
switch day {
case "Saturday", "Sunday":
    fmt.Println("Weekend")
}
```

### Mistake 2: Using Break Unnecessarily
```go
// Go doesn't need 'break' - it's automatic
switch day {
case "Monday":
    fmt.Println("Start")
    break  // This is unnecessary in Go!
}
```

### Mistake 3: Misusing Fallthrough
```go
// Bad - fallthrough doesn't check the next condition!
switch x {
case 1:
    fallthrough
case 2:
    // This executes even if x == 1
    fmt.Println("Case 2")
}
```

### Mistake 4: Unreachable Cases
```go
// Bad - case 5 will never be reached
switch {
case x > 0:
    fmt.Println("Positive")
case x > 5:  // Unreachable! x > 0 already matched
    fmt.Println("Greater than 5")
}
```

**Fix:** Order matters - put more specific cases first:
```go
switch {
case x > 5:
    fmt.Println("Greater than 5")
case x > 0:
    fmt.Println("Positive")
}
```

## Why Switch Exists

1. **Readability:** Easier to read than nested if-else chains
2. **Performance:** Compilers can optimize switch statements better than if-else
3. **Safety:** No automatic fallthrough prevents common bugs
4. **Flexibility:** Expression-less switch provides power without complexity

## When NOT to Use Switch

- **Complex boolean logic:** If cases involve `&&` and `||` operators extensively, if-else might be clearer
- **Few cases:** For just 1-2 conditions, a simple if statement is more readable
- **Unrelated conditions:** When each case checks completely different variables

## Official Documentation

For more details, see:
- [Go by Example: Switch](https://gobyexample.com/switch)
- [Effective Go: Switch](https://go.dev/doc/effective_go#switch)
- [Go Spec: Switch statements](https://go.dev/ref/spec#Switch_statements)

## Additional Notes

### Empty Case Bodies

Empty cases are valid but do nothing:
```go
switch x {
case 1:
    // This case does nothing
case 2:
    fmt.Println("Two")
}
```

### Switch in Loops

You can use switch inside loops, and `break` will break the switch, not the loop:
```go
for i := 0; i < 10; i++ {
    switch i {
    case 5:
        break  // Only breaks the switch, not the for loop
    }
}
```

To break the loop from within a switch, use a label:
```go
Loop:
for i := 0; i < 10; i++ {
    switch i {
    case 5:
        break Loop  // Breaks the for loop
    }
}
```
