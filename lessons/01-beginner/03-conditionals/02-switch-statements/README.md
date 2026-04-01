# Switch Statements

**Difficulty:** Beginner  
**Prerequisites:** Variables, Basic Types, If Statements

## Learning Objectives

By the end of this lesson, you will:

- Understand how to use basic switch statements in Go
- Use switch with multiple cases
- Write switch statements without an expression (expression-less switch)
- Understand and use the `fallthrough` keyword
- Know when to prefer switch over if-else chains

## Concept Summary

The `switch` statement provides a clean way to compare a value against multiple possibilities. Unlike many other languages, Go's switch cases **do not fall through by default** - execution stops after the first matching case. This makes the code safer and more predictable.

Basic switch syntax:

```go
switch value {
case option1:
    // code when value == option1
case option2:
    // code when value == option2
default:
    // code when no case matches
}
```

Switch statements in Go are more powerful than in many other languages:
- No automatic fallthrough (safer by default)
- Can switch without an expression (like a cleaner if-else chain)
- Cases can be expressions, not just constants
- Multiple values can be tested in a single case

## How To Run

From this directory, run:

```bash
# Run with starter code
go run -tags starter starter.go

# Run with solution
go run solution.go
```

## How To Test

From this directory, run:

```bash
go test -v
```

## Exercises

### Exercise 1: Day of the Week

Write a function `GetDayType(day string) string` that takes a day of the week and returns:
- "Weekday" for Monday through Friday
- "Weekend" for Saturday and Sunday
- "Invalid day" for any other input

**Example:**
```go
GetDayType("Monday")    // returns "Weekday"
GetDayType("Saturday")  // returns "Weekend"
GetDayType("Funday")    // returns "Invalid day"
```

### Exercise 2: Grade Evaluator

Write a function `GetGradeMessage(grade string) string` that returns messages for grades:
- "A" or "B" → "Excellent work!"
- "C" → "Good job!"
- "D" → "You passed, but consider studying more."
- "F" → "You need to retake this course."
- Any other input → "Invalid grade"

Use a single case for both "A" and "B".

**Example:**
```go
GetGradeMessage("A")  // returns "Excellent work!"
GetGradeMessage("B")  // returns "Excellent work!"
GetGradeMessage("C")  // returns "Good job!"
```

### Exercise 3: Number Classifier

Write a function `ClassifyNumber(n int) string` that uses a switch **without an expression** to classify a number:
- If n < 0 → "Negative"
- If n == 0 → "Zero"
- If n > 0 and n <= 10 → "Small positive"
- If n > 10 and n <= 100 → "Medium positive"
- If n > 100 → "Large positive"

**Example:**
```go
ClassifyNumber(-5)   // returns "Negative"
ClassifyNumber(0)    // returns "Zero"
ClassifyNumber(5)    // returns "Small positive"
ClassifyNumber(50)   // returns "Medium positive"
ClassifyNumber(200)  // returns "Large positive"
```

### Exercise 4: Season Describer

Write a function `DescribeSeason(season string) string` that returns descriptions:
- "Spring" → "Flowers bloom"
- "Summer" → "Hot and sunny" (then fallthrough)
- "Fall" or "Autumn" → "Leaves change color"
- "Winter" → "Cold and snowy"
- Any other input → "Unknown season"

For "Summer", use the `fallthrough` keyword so it returns both "Hot and sunny" and "Leaves change color" concatenated with a newline.

**Example:**
```go
DescribeSeason("Spring")   // returns "Flowers bloom"
DescribeSeason("Summer")   // returns "Hot and sunny\nLeaves change color"
DescribeSeason("Fall")     // returns "Leaves change color"
DescribeSeason("Winter")   // returns "Cold and snowy"
```

### Exercise 5: HTTP Status Classifier

Write a function `ClassifyHTTPStatus(code int) string` that classifies HTTP status codes:
- 200-299 → "Success"
- 300-399 → "Redirection"
- 400-499 → "Client Error"
- 500-599 → "Server Error"
- Any other value → "Unknown Status"

Use an expression-less switch with range checks.

**Example:**
```go
ClassifyHTTPStatus(200)  // returns "Success"
ClassifyHTTPStatus(404)  // returns "Client Error"
ClassifyHTTPStatus(500)  // returns "Server Error"
```

## Tips

- Switch statements in Go do NOT fall through automatically (unlike C/Java)
- Use `fallthrough` only when you explicitly need the next case to execute
- Expression-less switch (just `switch {`) is great for replacing long if-else chains
- Multiple values in one case: `case "A", "B", "C":`
- Cases are evaluated top to bottom; the first match wins
- The `default` case is optional but recommended for handling unexpected values
