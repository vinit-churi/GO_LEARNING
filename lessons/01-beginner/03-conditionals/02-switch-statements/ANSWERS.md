# Answers and Explanations

## Exercise 1: Day of the Week

### Approach

This exercise demonstrates the most basic form of a switch statement - comparing a single value against multiple constants. The key insight is using multiple values in a single case for weekdays.

### Implementation Strategy

```go
func GetDayType(day string) string {
    switch day {
    case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
        return "Weekday"
    case "Saturday", "Sunday":
        return "Weekend"
    default:
        return "Invalid day"
    }
}
```

### Key Points

- **Multiple values in one case:** The weekday case demonstrates how to group related values using commas
- **Default case:** Handles invalid input gracefully
- **Return in cases:** We can return directly from within a case - no need for extra variables or breaks

### Common Pitfall

Don't write separate cases for each weekday - that's verbose and harder to maintain:
```go
// Don't do this:
case "Monday":
    return "Weekday"
case "Tuesday":
    return "Weekday"
// ... repetitive!
```

---

## Exercise 2: Grade Evaluator

### Approach

This exercise reinforces using multiple values in a single case. It's similar to Exercise 1 but focuses on a different domain (grades).

### Implementation Strategy

```go
func GetGradeMessage(grade string) string {
    switch grade {
    case "A", "B":
        return "Excellent work!"
    case "C":
        return "Good job!"
    case "D":
        return "You passed, but consider studying more."
    case "F":
        return "You need to retake this course."
    default:
        return "Invalid grade"
    }
}
```

### Key Points

- Both "A" and "B" share the same message, demonstrating practical use of multiple values
- Each grade has a distinct message, making the switch statement the natural choice
- Default case provides input validation

### Alternative Approach

You could use a map, but switch is clearer for this scenario:
```go
// Possible but less readable for simple mappings:
messages := map[string]string{
    "A": "Excellent work!",
    "B": "Excellent work!",
    // ...
}
```

---

## Exercise 3: Number Classifier

### Approach

This exercise introduces the **expression-less switch** - one of Go's most powerful features. Instead of comparing a value, each case evaluates a boolean condition.

### Implementation Strategy

```go
func ClassifyNumber(n int) string {
    switch {
    case n < 0:
        return "Negative"
    case n == 0:
        return "Zero"
    case n > 0 && n <= 10:
        return "Small positive"
    case n > 10 && n <= 100:
        return "Medium positive"
    case n > 100:
        return "Large positive"
    }
    return "" // Unreachable, but Go requires it
}
```

### Key Points

- **No expression after switch:** Just `switch {` - this means each case is evaluated as a boolean
- **Order matters:** Cases are checked top-to-bottom; the first true case wins
- **Range checking:** Expression-less switch is perfect for classifying values into ranges
- **Cleaner than if-else:** Compare with the if-else equivalent - the switch is more scannable

### Why This is Better Than If-Else

```go
// Compare with if-else (less readable):
if n < 0 {
    return "Negative"
} else if n == 0 {
    return "Zero"
} else if n > 0 && n <= 10 {
    return "Small positive"
} // ... you get the idea
```

The switch version:
- Has less visual noise
- Shows the parallel structure more clearly
- Is easier to add new cases to

### Design Note

The cases must be mutually exclusive and cover all possibilities. If you have overlapping conditions, the first match wins:

```go
switch {
case n > 0:      // This matches first for n=5
    return "Positive"
case n > 0 && n < 10:  // This never executes!
    return "Small"
}
```

---

## Exercise 4: Season Describer

### Approach

This exercise demonstrates the `fallthrough` keyword - a feature that allows explicit case-to-case flow. Use it sparingly!

### Implementation Strategy

```go
func DescribeSeason(season string) string {
    result := ""
    switch season {
    case "Spring":
        return "Flowers bloom"
    case "Summer":
        result = "Hot and sunny\n"
        fallthrough
    case "Fall", "Autumn":
        result += "Leaves change color"
        return result
    case "Winter":
        return "Cold and snowy"
    default:
        return "Unknown season"
    }
}
```

### Key Points

- **Fallthrough is explicit:** Unlike C/Java, you must use the keyword to fall through
- **Fallthrough ignores conditions:** When "Summer" falls through, it executes the "Fall" case **without checking** if season == "Fall"
- **Building strings:** We use a variable to accumulate the result across cases
- **Multiple names for same season:** "Fall" and "Autumn" are handled together

### Important Warning

`fallthrough` is rarely needed. Most of the time, you want multiple values in one case instead:

```go
// Usually better:
case "Summer", "Fall":
    return "Warm seasons"

// Instead of:
case "Summer":
    fallthrough
case "Fall":
    return "Warm seasons"
```

### When to Use Fallthrough

Legitimate uses are rare, but include:
1. Accumulating results across multiple cases (like this exercise)
2. Sequential processing where each case adds to previous work
3. Implementing state machines with sequential states

---

## Exercise 5: HTTP Status Classifier

### Approach

This exercise combines range checking with expression-less switch. It's a practical example you'll encounter in real web applications.

### Implementation Strategy

```go
func ClassifyHTTPStatus(code int) string {
    switch {
    case code >= 200 && code <= 299:
        return "Success"
    case code >= 300 && code <= 399:
        return "Redirection"
    case code >= 400 && code <= 499:
        return "Client Error"
    case code >= 500 && code <= 599:
        return "Server Error"
    default:
        return "Unknown Status"
    }
}
```

### Key Points

- **Real-world pattern:** HTTP status codes are grouped by ranges - switch is perfect for this
- **Inclusive ranges:** Using `>=` and `<=` makes the ranges clear
- **Default for invalid codes:** Status codes outside 200-599 are caught by default

### Alternative Approaches

**Map approach (doesn't work well for ranges):**
```go
// This only works for specific codes, not ranges
statusMap := map[int]string{
    200: "Success",
    201: "Success",
    // ... way too many entries!
}
```

**Integer division (clever but less readable):**
```go
// Technically works but harder to understand
switch code / 100 {
case 2:
    return "Success"
case 3:
    return "Redirection"
// ...
}
```

The range-checking switch is the clearest approach.

### Practical Usage

In real code, you might combine this with specific checks:
```go
switch {
case code == 404:
    return "Not Found"
case code >= 400 && code <= 499:
    return "Client Error"
// ...
}
```

---

## General Takeaways

### Switch vs If-Else Decision Guide

**Use switch when:**
- Comparing one value against multiple constants (Exercises 1, 2)
- Checking value ranges (Exercises 3, 5)
- You have 3+ related conditions
- You want code that's easy to scan visually

**Use if-else when:**
- Conditions are unrelated
- Complex boolean logic is involved
- You only have 1-2 simple conditions

### Remember

1. **No automatic fallthrough** - This is a feature, not a bug!
2. **Expression-less switch** - One of Go's best features for replacing if-else chains
3. **Multiple values per case** - Use commas to group related values
4. **Default is optional** - But recommended for input validation
5. **Fallthrough is rare** - Think twice before using it

### Testing Your Understanding

Can you answer these?
- What happens if no case matches and there's no default? (Answer: Nothing - execution continues)
- Can a case have multiple statements? (Answer: Yes - any number)
- Does `break` work in switch? (Answer: Yes, but it's unnecessary since there's no automatic fallthrough)
- Can you switch on types? (Answer: Yes - see type switch in REFERENCE.md)

For the complete implementation, see `solution.go`.
