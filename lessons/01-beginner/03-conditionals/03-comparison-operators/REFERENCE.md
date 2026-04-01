# Comparison and Logical Operators - Reference

## Comparison Operators

Comparison operators compare two values and return a boolean result (`true` or `false`).

### Equal To (`==`)

Tests if two values are equal:

```go
5 == 5        // true
5 == 3        // false
"hello" == "hello"  // true
"hello" == "world"  // false
true == true  // true
```

**Important:** Use `==` for comparison, not `=` (which is assignment).

### Not Equal To (`!=`)

Tests if two values are different:

```go
5 != 3        // true
5 != 5        // false
"hello" != "world"  // true
```

### Less Than (`<`)

Tests if the left value is strictly less than the right:

```go
3 < 5         // true
5 < 5         // false (not strictly less)
5 < 3         // false
```

### Greater Than (`>`)

Tests if the left value is strictly greater than the right:

```go
5 > 3         // true
5 > 5         // false (not strictly greater)
3 > 5         // false
```

### Less Than or Equal To (`<=`)

Tests if the left value is less than or equal to the right:

```go
3 <= 5        // true
5 <= 5        // true
7 <= 5        // false
```

### Greater Than or Equal To (`>=`)

Tests if the left value is greater than or equal to the right:

```go
5 >= 3        // true
5 >= 5        // true
3 >= 5        // false
```

## Logical Operators

Logical operators combine or modify boolean values.

### Logical AND (`&&`)

Returns `true` only if BOTH operands are true:

```go
true && true    // true
true && false   // false
false && true   // false
false && false  // false

// With comparisons
age := 25
age >= 18 && age <= 65  // true (25 is between 18 and 65)

x := 5
x > 0 && x < 10  // true (5 is positive and less than 10)
```

**Short-Circuit Evaluation:**
If the first operand is `false`, the second operand is not evaluated:

```go
false && someExpensiveFunction()  // someExpensiveFunction() is never called
```

This is useful for safe checking:

```go
if user != nil && user.IsActive {
    // user is checked first, so user.IsActive is safe
}
```

### Logical OR (`||`)

Returns `true` if AT LEAST ONE operand is true:

```go
true || true    // true
true || false   // true
false || true   // true
false || false  // false

// With comparisons
day := "Saturday"
day == "Saturday" || day == "Sunday"  // true

x := -5
x < 0 || x > 100  // true (x is negative)
```

**Short-Circuit Evaluation:**
If the first operand is `true`, the second operand is not evaluated:

```go
true || someExpensiveFunction()  // someExpensiveFunction() is never called
```

This is useful for providing defaults:

```go
if useCache || loadFromDatabase() {
    // if cache is available, database is not queried
}
```

### Logical NOT (`!`)

Negates a boolean value (flips true to false, false to true):

```go
!true         // false
!false        // true
!(5 > 3)      // false (5 > 3 is true, negated to false)
!(5 < 3)      // true (5 < 3 is false, negated to true)

// Useful for readability
isEmpty := len(s) == 0
isNotEmpty := !isEmpty
```

**Double Negation:**
```go
!!true   // true (negated twice returns to original)
!!false  // false
```

## Operator Precedence

When multiple operators are used together, Go evaluates them in a specific order:

**Precedence (highest to lowest):**
1. Parentheses `()`
2. Unary operators: `!` (NOT)
3. Comparison operators: `==`, `!=`, `<`, `>`, `<=`, `>=`
4. Logical AND: `&&`
5. Logical OR: `||`

### Examples

```go
// Without parentheses
!false && true || false
// Evaluated as: ((!false) && true) || false
// Step 1: !false = true
// Step 2: true && true = true
// Step 3: true || false = true
// Result: true

// With comparisons
x := 5
x > 0 && x < 10 || x == 100
// Evaluated as: ((x > 0) && (x < 10)) || (x == 100)
// Step 1: 5 > 0 = true
// Step 2: 5 < 10 = true
// Step 3: true && true = true
// Step 4: 5 == 100 = false
// Step 5: true || false = true
// Result: true
```

### Using Parentheses for Clarity

Even when not needed, parentheses improve readability:

```go
// Hard to read
age >= 18 && isMember || isPremium

// Clearer with parentheses
(age >= 18 && isMember) || isPremium
```

**Critical Case - Changing Meaning:**
```go
// These are DIFFERENT
age >= 18 && isMember || isPremium
age >= 18 && (isMember || isPremium)

// First: (age >= 18 AND isMember) OR isPremium
// Second: age >= 18 AND (isMember OR isPremium)
```

## Short-Circuit Evaluation

Go uses short-circuit evaluation for `&&` and `||`:

### AND Short-Circuit

```go
false && expensiveFunction()
// expensiveFunction() is NOT called because first operand is false
```

**Practical Example - Nil Checking:**
```go
if user != nil && user.IsActive() {
    // Safe: user.IsActive() only called if user is not nil
}

// This would crash if user is nil
if user.IsActive() && user != nil {
    // BAD: IsActive() called before nil check
}
```

### OR Short-Circuit

```go
true || expensiveFunction()
// expensiveFunction() is NOT called because first operand is true
```

**Practical Example - Default Values:**
```go
if cachedValue != "" || loadFromDatabase(&cachedValue) {
    // Database only queried if cache is empty
}
```

## Comparing Different Types

### Numbers

All numeric types can be compared, but they must be the same type:

```go
var a int = 5
var b int = 3
a > b  // OK

var c int = 5
var d int64 = 3
c > d  // ERROR: cannot compare int and int64

// Must convert first
c > int(d)  // OK
```

### Strings

Strings are compared lexicographically (alphabetically):

```go
"apple" < "banana"  // true
"apple" < "apricot" // true
"abc" == "abc"      // true
"Abc" == "abc"      // false (case-sensitive)
```

### Booleans

Booleans can only be compared for equality:

```go
true == true    // OK: true
true != false   // OK: true
true < false    // ERROR: cannot use < with booleans
```

## Common Patterns

### Range Checking

```go
// Check if value is in range
if value >= min && value <= max {
    fmt.Println("in range")
}

// Check if value is outside range
if value < min || value > max {
    fmt.Println("out of range")
}
```

### Multiple Equality Checks

```go
// Check if value matches any of several options
if day == "Saturday" || day == "Sunday" {
    fmt.Println("weekend")
}

// For many values, consider using a switch statement instead
```

### Combining Different Conditions

```go
// Access control logic
if (age >= 18 && hasLicense) || isSupervised {
    fmt.Println("can drive")
}

// Form validation
if username != "" && len(password) >= 8 && password == confirmPassword {
    fmt.Println("valid registration")
}
```

### Negating Complex Conditions

```go
// Check if NOT in range (clear with negation)
if !(value >= min && value <= max) {
    fmt.Println("out of range")
}

// Equivalent but using De Morgan's law
if value < min || value > max {
    fmt.Println("out of range")
}
```

## Why These Operators Exist

1. **Decision Making:** Programs need to compare values and make decisions
2. **Validation:** Check if data meets requirements
3. **Control Flow:** Determine which code path to execute
4. **Filtering:** Select items that match criteria
5. **Logic Representation:** Express real-world rules in code

## When To Use Each

### Use `&&` When:
- ALL conditions must be true
- You need to combine requirements
- Performing safety checks (nil checks before access)

### Use `||` When:
- ANY condition being true is sufficient
- Providing alternatives
- Default or fallback logic

### Use `!` When:
- You need the opposite of a condition
- Checking for absence or falseness
- Making code more readable ("isNotEmpty" vs "!(isEmpty)")

## When NOT To Use

### Avoid Complex Nested Logic

```go
// Hard to understand
if (a && b || c) && (d || e && !f) {
    // what does this even check?
}

// Better: break into named booleans
hasBasicAccess := a && b || c
hasAdvancedAccess := d || e && !f
if hasBasicAccess && hasAdvancedAccess {
    // clearer intent
}
```

### Don't Compare Booleans to true/false

```go
// Redundant
if isActive == true {
    // ...
}

// Better
if isActive {
    // ...
}

// Redundant
if isActive == false {
    // ...
}

// Better
if !isActive {
    // ...
}
```

### Avoid Complicated Precedence Dependencies

```go
// Relies on precedence knowledge
if !a && b || c

// Clearer with parentheses
if (!a && b) || c
```

## Common Mistakes

### 1. Assignment Instead of Comparison

```go
// Wrong: assigns 5 to x (won't compile in if statement)
if x = 5 {
    // ...
}

// Correct: compares x to 5
if x == 5 {
    // ...
}
```

### 2. Comparing Different Types

```go
var age int = 25
var limit int64 = 21

// Wrong: type mismatch
if age > limit {  // ERROR
    // ...
}

// Correct: convert types
if age > int(limit) {
    // ...
}
```

### 3. String Case Sensitivity

```go
// These are NOT equal
"Hello" == "hello"  // false

// Use strings package for case-insensitive comparison
strings.EqualFold("Hello", "hello")  // true
```

### 4. Incorrect Precedence Assumptions

```go
// This is: (age > 18 && isMember) || isPremium
if age > 18 && isMember || isPremium {
    // ...
}

// If you meant: age > 18 && (isMember || isPremium)
if age > 18 && (isMember || isPremium) {
    // ...
}
```

### 5. Not Using Short-Circuit Safety

```go
// Dangerous: may panic if pointer is nil
if pointer.IsValid() && pointer != nil {
    // WRONG ORDER: IsValid() called first
}

// Safe: nil check first
if pointer != nil && pointer.IsValid() {
    // Correct: IsValid() only called if pointer is not nil
}
```

### 6. Redundant Boolean Comparisons

```go
// Redundant
if (x > 5) == true {
    // ...
}

// Simpler
if x > 5 {
    // ...
}
```

## Additional Resources

- Go Tour - More Types: https://go.dev/tour/moretypes/1
- Go Spec - Comparison Operators: https://go.dev/ref/spec#Comparison_operators
- Go Spec - Logical Operators: https://go.dev/ref/spec#Logical_operators
- Go Spec - Operator Precedence: https://go.dev/ref/spec#Operator_precedence
- Effective Go - Control Structures: https://go.dev/doc/effective_go#control-structures
