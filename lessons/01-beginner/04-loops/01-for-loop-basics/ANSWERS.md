# Exercise Answers

## Exercise 1: CountToN

### Explanation
This exercise introduces the classic three-component for loop. The key is understanding how the initialization, condition, and post statement work together.

### Approach
1. Create an empty slice to store the results
2. Use `i := 1` as initialization (we want to start from 1, not 0)
3. Use `i <= n` as the condition (inclusive of n)
4. Use `i++` to increment after each iteration
5. Append each value to the result slice

### Common Pitfalls
- Starting from 0 instead of 1
- Using `i < n` instead of `i <= n` (would miss the last number)
- Forgetting to handle the edge case where n < 1

### Solution Reference
See `CountToN` function in `solution.go`

---

## Exercise 2: MultiplicationTable

### Explanation
This exercise demonstrates nested loops. The outer loop creates rows, and the inner loop creates columns within each row.

### Approach
1. Create a 2D slice with `size` rows
2. Outer loop: iterate through rows (i from 0 to size-1)
3. Inner loop: iterate through columns (j from 0 to size-1)
4. Calculate multiplication: `(i+1) * (j+1)` 
   - We add 1 because we want the table to start from 1×1, not 0×0
5. Each inner loop iteration builds one row

### Key Insights
- The outer loop runs `size` times
- For each outer iteration, the inner loop runs `size` times
- Total iterations: size × size
- Array indexing starts at 0, but multiplication table starts at 1

### Visualization
For size=3:
```
i=0, j=0: (0+1)*(0+1) = 1×1 = 1
i=0, j=1: (0+1)*(1+1) = 1×2 = 2
i=0, j=2: (0+1)*(2+1) = 1×3 = 3
i=1, j=0: (1+1)*(0+1) = 2×1 = 2
...
```

### Solution Reference
See `MultiplicationTable` function in `solution.go`

---

## Exercise 3: FindFirstDivisible

### Explanation
This exercise teaches you when and how to use `break`. We want to stop searching as soon as we find the first match—no need to continue through the rest of the array.

### Approach
1. Check for edge case: divisor is 0 (would cause panic)
2. Loop through the numbers slice
3. Check if current number is divisible by divisor (using modulo operator `%`)
4. If divisible, return immediately (or use break and return after loop)
5. If loop completes without finding anything, return -1

### Two Valid Patterns

**Pattern 1: Return directly**
```go
for _, num := range numbers {
    if num % divisor == 0 {
        return num  // Exit immediately
    }
}
return -1
```

**Pattern 2: Break and return**
```go
result := -1
for _, num := range numbers {
    if num % divisor == 0 {
        result = num
        break
    }
}
return result
```

Both are valid. Direct return is more concise for simple cases.

### Key Concepts
- `break` exits the loop early
- Modulo operator `%` gives the remainder of division
- If `a % b == 0`, then a is divisible by b

### Solution Reference
See `FindFirstDivisible` function in `solution.go`

---

## Exercise 4: SumOddNumbers

### Explanation
This exercise demonstrates `continue`, which skips the rest of the current iteration. When we encounter an even number, we skip it and move to the next iteration.

### Approach
1. Initialize sum to 0
2. Loop through all numbers
3. Check if number is even (num % 2 == 0)
4. If even, use `continue` to skip to next iteration
5. If odd, add to sum
6. Return the final sum

### Alternative Approach (without continue)
```go
sum := 0
for _, num := range numbers {
    if num % 2 != 0 {  // If odd
        sum += num
    }
}
return sum
```

Both work! The `continue` version makes the "skip even numbers" logic more explicit.

### When to Use Continue
- When the skipping logic is more natural to express
- When you have multiple conditions that would skip
- When it makes the code more readable

### Key Concepts
- `continue` skips to the next iteration
- Even numbers: num % 2 == 0
- Odd numbers: num % 2 != 0 (or num % 2 == 1)

### Solution Reference
See `SumOddNumbers` function in `solution.go`

---

## Exercise 5: ReadUntilNegative

### Explanation
This exercise shows how to use an infinite loop with a break condition. It's useful when you don't know in advance how many iterations you need, and the exit condition is checked inside the loop.

### Approach
1. Create an empty result slice
2. Start an infinite loop with `for i := 0; ; i++` or use an index variable
3. Check if we've reached the end of the slice
4. Check if the current number is negative
5. If negative, break out of the loop
6. If not negative, append to result
7. Return the result

### Loop Pattern
```go
for i := 0; ; i++ {  // Infinite: no condition
    if i >= len(numbers) {
        break  // Exit condition 1: end of slice
    }
    if numbers[i] < 0 {
        break  // Exit condition 2: found negative
    }
    // Process the number
}
```

### Alternative Pattern
```go
i := 0
for {  // Truly infinite
    if i >= len(numbers) {
        break
    }
    if numbers[i] < 0 {
        break
    }
    result = append(result, numbers[i])
    i++
}
```

### When to Use Infinite Loops
- When you have multiple exit conditions
- When the exit check is in the middle of the loop logic
- When it makes the control flow clearer

### Tradeoffs
- **Pro**: Very flexible, clear exit conditions
- **Con**: Must ensure break is reachable (avoid true infinite loops)
- **Pro**: Good for event loops, servers, reading until EOF
- **Con**: Can be harder to reason about than bounded loops

### Solution Reference
See `ReadUntilNegative` function in `solution.go`

---

## General Insights

### Choosing the Right Pattern
- Use **classic for loop** when you need an index and know the iteration count
- Use **break** when you want to exit early after finding something
- Use **continue** when you want to skip certain items
- Use **infinite loop** when exit conditions are complex or multiple

### Testing Strategy
For each function:
1. Test normal cases (typical inputs)
2. Test edge cases (empty slices, zero values, boundary conditions)
3. Test error conditions (invalid inputs)

### Further Practice
Try these variations:
1. Modify `CountToN` to count down from n to 1
2. Create a function that finds ALL divisible numbers (not just first)
3. Write a function that sums even numbers using continue for odd ones
4. Create a function that reads until a value greater than a threshold
