# Comparison and Logical Operators

**Difficulty:** Beginner  
**Prerequisites:** Variables, basic types, and if statements

## Learning Objectives

By the end of this lesson, you will:

- Understand and use all comparison operators (==, !=, <, >, <=, >=)
- Master logical operators (&&, ||, !)
- Apply operator precedence rules correctly
- Combine multiple conditions effectively
- Write clear and maintainable boolean expressions

## Concept Summary

Comparison and logical operators are the foundation of decision-making in programs. They let you compare values and combine conditions to create complex logic.

**Comparison Operators:**
```go
==  // equal to
!=  // not equal to
<   // less than
>   // greater than
<=  // less than or equal to
>=  // greater than or equal to
```

**Logical Operators:**
```go
&&  // AND - both conditions must be true
||  // OR - at least one condition must be true
!   // NOT - negates a boolean value
```

**Operator Precedence (high to low):**
1. `!` (NOT)
2. `==`, `!=`, `<`, `>`, `<=`, `>=` (Comparison)
3. `&&` (AND)
4. `||` (OR)

## How To Run

From this directory, run:

```bash
go run solution.go
```

Or to run with the starter file:

```bash
go run -tags=starter starter.go
```

## How To Test

From this directory, run:

```bash
go test -v
```

## Exercises

### Exercise 1: Basic Comparisons

Write a function `CompareIntegers(a, b int) map[string]bool` that returns a map containing the results of all six comparison operators between a and b.

The map should have keys: "equal", "notEqual", "lessThan", "greaterThan", "lessOrEqual", "greaterOrEqual"

**Example:**
```go
CompareIntegers(5, 3)
// returns map[string]bool{
//     "equal": false,
//     "notEqual": true,
//     "lessThan": false,
//     "greaterThan": true,
//     "lessOrEqual": false,
//     "greaterOrEqual": true,
// }
```

### Exercise 2: Logical AND Operator

Write a function `IsValidAge(age int) bool` that returns true if the age is between 0 and 150 (inclusive). Use the `&&` operator.

**Example:**
```go
IsValidAge(25)   // returns true
IsValidAge(-5)   // returns false
IsValidAge(200)  // returns false
IsValidAge(0)    // returns true
IsValidAge(150)  // returns true
```

### Exercise 3: Logical OR Operator

Write a function `IsWeekend(day string) bool` that returns true if the day is "Saturday" or "Sunday". Use the `||` operator.

**Example:**
```go
IsWeekend("Saturday")  // returns true
IsWeekend("Sunday")    // returns true
IsWeekend("Monday")    // returns false
IsWeekend("Friday")    // returns false
```

### Exercise 4: Logical NOT Operator

Write a function `IsNotEmpty(s string) bool` that returns true if the string is not empty. Use the `!` operator to negate the empty check.

**Example:**
```go
IsNotEmpty("hello")  // returns true
IsNotEmpty("")       // returns false
IsNotEmpty(" ")      // returns true (space is not empty)
```

### Exercise 5: Complex Conditions with AND

Write a function `CanVote(age int, isCitizen bool) bool` that returns true if a person can vote. A person can vote if they are 18 or older AND they are a citizen.

**Example:**
```go
CanVote(25, true)   // returns true
CanVote(16, true)   // returns false
CanVote(25, false)  // returns false
CanVote(16, false)  // returns false
```

### Exercise 6: Complex Conditions with OR

Write a function `NeedsUmbrella(isRaining bool, isSnowing bool) bool` that returns true if you need an umbrella. You need an umbrella if it's raining OR snowing.

**Example:**
```go
NeedsUmbrella(true, false)   // returns true
NeedsUmbrella(false, true)   // returns true
NeedsUmbrella(true, true)    // returns true
NeedsUmbrella(false, false)  // returns false
```

### Exercise 7: Combining Multiple Logical Operators

Write a function `IsValidPassword(password string, confirmPassword string, length int) bool` that returns true if:
- password equals confirmPassword AND
- length is at least 8 AND
- length is not more than 128

**Example:**
```go
IsValidPassword("secret123", "secret123", 9)   // returns true
IsValidPassword("secret", "secret123", 9)      // returns false (not equal)
IsValidPassword("secret123", "secret123", 5)   // returns false (too short)
IsValidPassword("secret123", "secret123", 200) // returns false (too long)
```

### Exercise 8: Operator Precedence

Write a function `CheckAccessLevel(age int, isMember bool, isPremium bool) string` that returns:
- "full access" if the person is premium OR (is a member AND is 18 or older)
- "limited access" if the person is a member but not premium and under 18
- "no access" otherwise

This exercise requires understanding operator precedence where:
- `age >= 18` is evaluated first
- `isMember && age >= 18` is evaluated next
- `isPremium || (isMember && age >= 18)` is evaluated last

**Example:**
```go
CheckAccessLevel(25, true, true)    // returns "full access" (premium)
CheckAccessLevel(25, true, false)   // returns "full access" (member AND 18+)
CheckAccessLevel(16, true, false)   // returns "limited access" (member but under 18)
CheckAccessLevel(16, false, false)  // returns "no access"
CheckAccessLevel(25, false, false)  // returns "no access" (not a member)
```

### Exercise 9: Negation with Complex Conditions

Write a function `IsOutsideRange(value int, min int, max int) bool` that returns true if the value is outside the range [min, max]. Use the `!` operator to negate an "in range" check.

**Example:**
```go
IsOutsideRange(5, 1, 10)   // returns false (5 is inside range)
IsOutsideRange(15, 1, 10)  // returns true (15 is outside range)
IsOutsideRange(1, 1, 10)   // returns false (1 is at boundary)
IsOutsideRange(0, 1, 10)   // returns true (0 is outside range)
```

## Tips

- Comparison operators work with numbers, strings, and booleans
- `&&` short-circuits: if the first condition is false, the second is not evaluated
- `||` short-circuits: if the first condition is true, the second is not evaluated
- Use parentheses to make complex conditions more readable
- When in doubt about precedence, use parentheses
- Test boundary conditions (equal values, edge cases)
