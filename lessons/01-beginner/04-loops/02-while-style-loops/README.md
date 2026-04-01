# While-Style Loops

**Difficulty:** Beginner  
**Prerequisites:** Basic Go syntax, variables, conditional statements

## Learning Objectives

By the end of this lesson, you will:

- Understand how Go's `for` loop can act as a while loop
- Implement loops with only condition checks
- Use sentinel values to control loop termination
- Read and process input until a condition is met
- Apply loop control patterns effectively

## Concept Summary

Go doesn't have a separate `while` keyword. Instead, the `for` loop can be used with just a condition, making it equivalent to while loops in other languages.

The basic syntax is:

```go
for condition {
    // code executes while condition is true
}
```

This is cleaner than the full `for` syntax when you don't need initialization or post-iteration statements. It's perfect for situations where you need to repeat until a specific condition is met.

## How To Run

From this directory, run:

```bash
go run solution.go
```

## How To Test

From this directory, run:

```bash
go test -v
```

## Exercises

### Exercise 1: Countdown

Write a function `Countdown(start int) []int` that counts down from a starting number to 1 using a while-style loop. The function should return a slice containing all the numbers in the countdown.

**Example:**
```go
Countdown(5) // returns []int{5, 4, 3, 2, 1}
```

### Exercise 2: Sum Until Threshold

Write a function `SumUntilThreshold(numbers []int, threshold int) int` that adds numbers from a slice until the sum reaches or exceeds the threshold. Return the final sum. If the threshold is reached before processing all numbers, stop early.

**Example:**
```go
SumUntilThreshold([]int{5, 10, 15, 20}, 25) // returns 30 (5 + 10 + 15)
```

### Exercise 3: Read Until Sentinel

Write a function `ReadUntilSentinel(values []string, sentinel string) []string` that reads values from a slice until it encounters the sentinel value. The sentinel should NOT be included in the result.

**Example:**
```go
ReadUntilSentinel([]string{"hello", "world", "STOP", "ignored"}, "STOP")
// returns []string{"hello", "world"}
```

### Exercise 4: Find First Divisor

Write a function `FindFirstDivisor(n int, start int) int` that finds the first number >= start that evenly divides n (has no remainder). Use a while-style loop to check each candidate until you find one.

**Example:**
```go
FindFirstDivisor(20, 3) // returns 4 (first divisor >= 3)
FindFirstDivisor(17, 2) // returns 17 (17 is prime, only divisible by itself)
```

### Exercise 5: Collatz Sequence Length

Write a function `CollatzLength(n int) int` that returns the length of the Collatz sequence starting from n. The Collatz sequence follows these rules:
- If n is even: divide by 2
- If n is odd: multiply by 3 and add 1
- Continue until n becomes 1

Return the number of steps it took.

**Example:**
```go
CollatzLength(10) // returns 7
// Sequence: 10 -> 5 -> 16 -> 8 -> 4 -> 2 -> 1 (7 steps)
```

## Tips

- While-style loops are best when you don't know how many iterations you'll need
- Always ensure your condition will eventually become false to avoid infinite loops
- Consider using `break` to exit early when a condition is met
- Sentinel values are common in data processing and file reading
- Test your loops with edge cases (empty input, threshold already met, etc.)
