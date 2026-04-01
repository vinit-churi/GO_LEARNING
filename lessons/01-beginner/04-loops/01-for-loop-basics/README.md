# For Loop Basics

## Difficulty
Beginner

## Prerequisites
- Variables and data types
- Basic operators
- Understanding of program flow

## Learning Objectives
By the end of this lesson, you will:
- Understand the classic three-component for loop syntax
- Write nested loops to solve multi-dimensional problems
- Use `break` to exit loops early
- Use `continue` to skip iterations
- Create infinite loops with proper exit conditions
- Choose the right loop pattern for different scenarios

## Concept Summary

Go has only one looping construct: the `for` loop. However, it's incredibly versatile and can handle all looping scenarios you'd need `while` or `do-while` loops for in other languages.

The classic for loop has three components:
```go
for initialization; condition; post {
    // loop body
}
```

- **Initialization**: Runs once before the first iteration (e.g., `i := 0`)
- **Condition**: Evaluated before each iteration; loop continues while true
- **Post**: Runs after each iteration (e.g., `i++`)

All three components are optional, giving you flexibility to create different loop patterns.

## Exercises

### Exercise 1: Classic For Loop (CountToN)
Write a function `CountToN(n int) []int` that returns a slice containing numbers from 1 to n (inclusive).

**Requirements:**
- Use a classic for loop with init, condition, and post components
- Return an empty slice if n is less than 1
- Examples:
  - `CountToN(5)` → `[1, 2, 3, 4, 5]`
  - `CountToN(0)` → `[]`

### Exercise 2: Nested Loops (MultiplicationTable)
Write a function `MultiplicationTable(size int) [][]int` that generates a multiplication table.

**Requirements:**
- Use nested for loops
- Return a size×size 2D slice
- Each element at position [i][j] should contain (i+1) * (j+1)
- Example for size=3:
  ```
  [[1, 2, 3],
   [2, 4, 6],
   [3, 6, 9]]
  ```

### Exercise 3: Break Statement (FindFirstDivisible)
Write a function `FindFirstDivisible(numbers []int, divisor int) int` that finds the first number divisible by the divisor.

**Requirements:**
- Use a for loop with break
- Return the first number that is divisible by divisor
- Return -1 if no number is found or divisor is 0
- Examples:
  - `FindFirstDivisible([3, 5, 10, 15], 5)` → `5`
  - `FindFirstDivisible([1, 2, 3], 7)` → `-1`

### Exercise 4: Continue Statement (SumOddNumbers)
Write a function `SumOddNumbers(numbers []int) int` that sums only the odd numbers in a slice.

**Requirements:**
- Use a for loop with continue to skip even numbers
- Return the sum of all odd numbers
- Examples:
  - `SumOddNumbers([1, 2, 3, 4, 5])` → `9` (1+3+5)
  - `SumOddNumbers([2, 4, 6])` → `0`

### Exercise 5: Infinite Loop (ReadUntilNegative)
Write a function `ReadUntilNegative(numbers []int) []int` that reads numbers from a slice until it encounters a negative number.

**Requirements:**
- Use an infinite loop pattern with a break condition
- Return all numbers before the first negative number
- If no negative number is found, return all numbers
- Examples:
  - `ReadUntilNegative([1, 2, -1, 3])` → `[1, 2]`
  - `ReadUntilNegative([5, 10, 15])` → `[5, 10, 15]`

## Run Instructions

To test your solution:
```bash
go test -v
```

To run the starter code:
```bash
go run -tags starter .
```

To run the solution:
```bash
go run .
```

## Test Instructions

The tests validate that your functions produce correct results for various inputs, including edge cases. Each test is named after the function it validates, making it easy to identify which exercise needs work.

Run with verbose output to see detailed test results:
```bash
go test -v
```

Run a specific test:
```bash
go test -v -run TestCountToN
```
