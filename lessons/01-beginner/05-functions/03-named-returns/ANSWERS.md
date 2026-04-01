# Named Return Values - Answers

## Exercise 1: Basic Named Returns

**Goal:** Implement a function that calculates rectangle dimensions using named returns and a naked return.

**Solution Approach:**
1. Declare named returns `length` and `width`
2. Loop through possible values to find dimensions that match both area and perimeter
3. Use the constraint that length ≥ width
4. Assign to the named return variables
5. Use a naked return

See `solution.go` - the `rectangleDimensions()` function demonstrates this.

**Key Concepts:**
- Named return values are automatically initialized to zero (0 for int)
- You can assign to them like regular variables throughout the function
- A naked `return` automatically returns these named variables
- This pattern works well for simple functions with clear return values

**Mathematical Note:**
Given area A and perimeter P, for a rectangle:
- A = length × width
- P = 2(length + width)

The solution iterates through possible dimensions to find a match.

---

## Exercise 2: Named Returns with Error Handling

**Goal:** Demonstrate the common Go pattern of named returns for result and error.

**Solution Approach:**
1. Declare named returns `result float64` and `err error`
2. Check for division by zero condition
3. Set `err` and use naked return for error case
4. Set `result` and use naked return for success case
5. Both variables are available throughout the function

See `solution.go` - the `safeDivide()` function shows this pattern.

**Key Concepts:**
- The `(result Type, err error)` pattern is very common in Go
- Named returns are automatically initialized: `result=0.0`, `err=nil`
- You can return early with just `return` when error occurs
- This pattern is clean for simple functions but can be problematic in complex ones

**Common Pitfall to Avoid:**
```go
// BAD - shadowing the named return
result, err := someOtherFunction() // := creates NEW variables
return // returns the zero values, NOT the values from someOtherFunction!

// GOOD - assigning to named returns
result, err = someOtherFunction() // = assigns to named returns
return
```

---

## Exercise 3: When to Use Named Returns

**Goal:** Compare named returns vs explicit returns to understand when each is better.

**Solution Approach:**

**Implementation A (Named):**
- Use named returns with descriptive names
- Assign to the named variables
- Use naked returns
- Good for documentation

**Implementation B (Explicit):**
- No named returns in signature
- Explicitly return values at each return point
- Clearer what's being returned at each point
- Better for functions with multiple return paths

See `solution.go` - both `parseNameNamed()` and `parseNameExplicit()` are implemented.

**Key Concepts:**
- Named returns document what the function returns
- Explicit returns make it crystal clear what values are being returned
- For simple functions, either approach works
- For functions with multiple return points, explicit returns are often clearer

**Comparison:**

Named returns pros:
- Self-documenting in function signature
- Less repetition
- Works well with defer

Named returns cons:
- Naked returns can obscure what's being returned
- Risk of accidentally returning wrong values
- Can be confusing in longer functions

Explicit returns pros:
- Crystal clear at every return point
- No ambiguity about what's being returned
- Harder to make mistakes

Explicit returns cons:
- More verbose
- Less documentation in signature
- More repetition

**Recommendation:** For functions like this with multiple return points and error handling, explicit returns are often clearer. Named returns work best in short, simple functions.

---

## Exercise 4: Complex Named Returns

**Goal:** Handle multiple named returns and understand when to use naked vs explicit returns.

**Solution Approach:**
1. Use descriptive named returns: `sum`, `count`, `avg`, `err`
2. Check for empty slice first - use explicit return for error case (clearer)
3. Calculate sum using a loop
4. Set count from slice length
5. Calculate average from sum and count
6. Use naked return for success case (all values already set)

See `solution.go` - the `calculateStats()` function demonstrates this.

**Key Concepts:**
- You can mix naked returns and explicit returns in the same function
- Use explicit returns when it makes the code clearer (like error cases)
- Use naked returns when all values are already set appropriately
- Multiple named returns can serve as documentation

**Trade-off Discussion:**

This function shows both approaches:
```go
if len(numbers) == 0 {
    return 0, 0, 0.0, errors.New("empty slice") // explicit - very clear
}
// ... calculate values ...
return // naked - values already set, clean
```

**Alternative Approach:**
You could use all explicit returns:
```go
return sum, count, avg, nil
```

This is more verbose but some teams prefer it for consistency and clarity.

**Best Practice:** In real code, prefer explicit returns for:
- Functions with multiple return paths
- When return values come from different logic branches
- When clarity is paramount

Use naked returns for:
- Very short functions
- When all return values are clearly set just before return
- When the pattern is established and clear to your team

---

## Exercise 5: Named Returns Best Practices

**Goal:** Demonstrate good practices: using named returns for documentation but explicit returns for clarity.

**Solution Approach:**
1. Use named returns in signature for documentation: `(output string, valid bool)`
2. Trim the input string
3. Check length validation
4. Use explicit returns at each return point (NOT naked returns)
5. This shows the best of both worlds

See `solution.go` - the `validateAndFormat()` function shows this pattern.

**Key Concepts:**
- Named returns provide documentation value in the function signature
- Explicit returns provide clarity at each return point
- This hybrid approach is often the best balance
- The function signature tells you what you're getting
- Each return statement tells you exactly what values are being returned

**Why This Approach Works:**
```go
func validateAndFormat(input string) (output string, valid bool) {
    // Names in signature document what we return
    trimmed := strings.TrimSpace(input)
    
    if len(trimmed) < 3 {
        return "", false // Clear: invalid case returns empty and false
    }
    
    return strings.ToUpper(trimmed), true // Clear: valid case returns formatted and true
}
```

**Comparison to Alternatives:**

1. **No named returns + explicit returns:**
```go
func validateAndFormat(input string) (string, bool)
```
Less documentation, but explicit returns are clear.

2. **Named returns + naked returns:**
```go
output = ""
valid = false
return
```
More confusing - requires reading the whole function to see what's returned.

3. **Named returns + explicit returns (our approach):**
Best of both worlds - documentation + clarity!

---

## General Best Practices

### When to Use Named Returns

✅ **DO use named returns when:**
- The function is short (fits on one screen)
- The names add documentation value
- You're using `defer` to modify returns
- Following a common pattern like `(result Type, err error)`

❌ **DON'T use named returns when:**
- The function is long or complex
- Multiple return points with different logic paths
- The meaning is obvious without names
- You're just trying to save a few characters

### When to Use Naked Returns

✅ **DO use naked returns when:**
- The function is very short (< 10 lines)
- Return values are set clearly just before return
- The pattern is simple and obvious

❌ **DON'T use naked returns when:**
- The function is longer than one screen
- There are many lines between setting values and returning
- Multiple return paths make it unclear
- You have to scroll to see what's being returned

### The Golden Rule

**Clarity beats brevity.** When in doubt:
1. Use named returns for documentation in the signature
2. Use explicit returns for clarity at each return point
3. This hybrid approach gives you the best of both worlds

---

## Common Mistakes to Avoid

### 1. Shadowing Named Returns

```go
// BUG
func example() (result int, err error) {
    result, err := doSomething() // := shadows!
    return // returns zero values!
}

// FIX
func example() (result int, err error) {
    result, err = doSomething() // = assigns correctly
    return
}
```

### 2. Naked Returns in Long Functions

```go
// BAD - unclear what's being returned
func process() (a, b, c int) {
    // ... 50 lines ...
    return
}

// GOOD - explicit and clear
func process() (a, b, c int) {
    // ... 50 lines ...
    return a, b, c
}
```

### 3. Unnecessary Named Returns

```go
// BAD - no benefit
func add(x, y int) (sum int) {
    return x + y
}

// GOOD - simpler
func add(x, y int) int {
    return x + y
}
```

---

## Further Reading

- Review REFERENCE.md for detailed explanations of named returns
- Practice writing both styles and see which feels more natural
- Read real Go code to see how experienced developers use named returns
- Check Go's standard library for examples (like `io.Reader.Read`)
- Remember: the Go community values clarity and simplicity
