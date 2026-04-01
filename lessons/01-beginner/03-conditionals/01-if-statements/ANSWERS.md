# If Statements - Answers

## Exercise 1: Basic If Statement

**Goal:** Use a basic if statement to check if a number is positive.

**Solution Approach:**
- Check if `n > 0`
- Return "positive" if true
- Return "not positive" otherwise (can be done without else by defaulting the return)

See `solution.go` - the `CheckPositive()` function demonstrates this.

**Key Concept:** The simplest if statement just checks a condition. Since we need to return something in both cases, we can use if-else or set a default value and only change it if the condition is true.

**Alternative Approach:**
```go
// Using if-else explicitly
if n > 0 {
    return "positive"
} else {
    return "not positive"
}

// Using default value
result := "not positive"
if n > 0 {
    result = "positive"
}
return result
```

Both approaches are valid. The solution uses early return for clarity.

---

## Exercise 2: If-Else

**Goal:** Determine if a number is even or odd.

**Solution Approach:**
- Use the modulo operator `%` to check if the number is divisible by 2
- If `n % 2 == 0`, the number is even
- Otherwise, it's odd

See `solution.go` - the `IsEven()` function shows this pattern.

**Key Concept:** If-else is perfect when you have exactly two mutually exclusive outcomes. The modulo operator `%` gives the remainder after division, so an even number has remainder 0 when divided by 2.

**Why This Works:**
- `4 % 2 = 0` (even)
- `5 % 2 = 1` (odd)
- `0 % 2 = 0` (even - zero is considered even)

---

## Exercise 3: If-Else If-Else Chain

**Goal:** Compare two numbers and determine their relationship.

**Solution Approach:**
- First check if `a == b` (equal)
- Then check if `a > b` (greater)
- Otherwise, a must be less than b (less)

See `solution.go` - the `CompareNumbers()` function demonstrates this.

**Key Concept:** When you have multiple mutually exclusive conditions, use if-else if-else chains. The conditions are checked in order, and only the first true condition executes.

**Order Matters:**
The order of conditions can affect readability and efficiency:
```go
// This works
if a == b {
    return "equal"
} else if a > b {
    return "greater"
} else {
    return "less"
}

// This also works but checks more conditions
if a > b {
    return "greater"
} else if a < b {
    return "less"
} else {
    return "equal"
}
```

Both are correct, but checking for equality first may be more intuitive.

---

## Exercise 4: If With Short Statement

**Goal:** Use an if statement with a short variable declaration to limit scope.

**Solution Approach:**
- Declare and initialize `result` in the if statement: `if result := a / b`
- Check the result in the same if statement
- Use else if and else for the remaining conditions
- The variable `result` is available in all branches (if, else if, else)

See `solution.go` - the `DivideAndCheck()` function shows this pattern.

**Key Concepts:**
1. **Short Variable Declaration:** `if result := a / b; ...` declares `result` and makes it available only within the if statement scope
2. **Scope Benefits:** The variable doesn't clutter the outer scope and clearly shows it's only used for this decision
3. **Integer Division:** Go performs integer division when both operands are integers (5 / 2 = 2, not 2.5)

**Why Use Short Statement:**
```go
// Without short statement
result := a / b
if result > 10 {
    return "large"
}
// result is still accessible here (may not want this)

// With short statement
if result := a / b; result > 10 {
    return "large"
}
// result is NOT accessible here (cleaner scope)
```

**Common Pattern - Error Handling:**
This pattern is extremely common in Go for error handling:
```go
if err := doSomething(); err != nil {
    return err
}
// err is not accessible here
```

---

## Exercise 5: Complex Condition

**Goal:** Check if a number falls within a range using a compound condition.

**Solution Approach:**
- Use the AND operator `&&` to check both conditions
- `n >= 10` ensures the number is at least 10
- `n <= 20` ensures the number is at most 20
- Both must be true for the number to be in range

See `solution.go` - the `CheckRange()` function demonstrates this.

**Key Concept:** Use logical operators to combine multiple conditions:
- `&&` (AND): Both conditions must be true
- `||` (OR): At least one condition must be true
- `!` (NOT): Negates a condition

**Alternative Approaches:**

```go
// Using nested ifs (more verbose)
if n >= 10 {
    if n <= 20 {
        return "in range"
    }
}
return "out of range"

// Using compound condition (clearest)
if n >= 10 && n <= 20 {
    return "in range"
}
return "out of range"

// Using negative logic (less clear)
if n < 10 || n > 20 {
    return "out of range"
}
return "in range"
```

The compound condition with `&&` is the most readable and idiomatic approach.

---

## General Tips

1. **Keep It Simple:** Don't nest if statements too deeply. Use early returns or break logic into functions.

2. **Be Explicit:** In Go, conditions must be boolean expressions. You can't use `if count` like in some languages - you must write `if count > 0`.

3. **Scope Is Your Friend:** Use short statement declarations (`if x := getValue(); ...`) to keep variables scoped tightly.

4. **Else Placement:** The `else` keyword must be on the same line as the closing brace of the if block due to Go's automatic semicolon insertion.

5. **Test Edge Cases:** Always test boundary conditions (0, negative numbers, equal values, etc.).

## Common Pitfalls

1. **Forgetting Braces:** Go requires braces even for single-line if bodies
   ```go
   // Wrong
   if x > 0
       return x
   
   // Correct
   if x > 0 {
       return x
   }
   ```

2. **Wrong Else Placement:**
   ```go
   // Wrong - won't compile
   if x > 0 {
       return "positive"
   }
   else {
       return "negative"
   }
   
   // Correct
   if x > 0 {
       return "positive"
   } else {
       return "negative"
   }
   ```

3. **Assignment vs. Comparison:**
   ```go
   // Wrong - assigns instead of compares (won't compile in Go)
   if x = 5 {
       // ...
   }
   
   // Correct
   if x == 5 {
       // ...
   }
   ```

4. **Scope Confusion:**
   ```go
   if x := 5; x > 0 {
       fmt.Println(x) // OK
   }
   fmt.Println(x) // Error: undefined
   ```

## Practice Suggestions

1. Try writing different conditions for each exercise
2. Practice combining conditions with `&&` and `||`
3. Experiment with short statement declarations
4. Look at Go standard library code to see idiomatic patterns
5. Refactor nested if statements into clearer structures
