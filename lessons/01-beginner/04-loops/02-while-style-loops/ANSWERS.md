# While-Style Loops - Answers

## Exercise 1: Countdown

**Goal:** Count down from a number to 1 using a while-style loop.

**Solution Approach:**
- Start with the given number
- Use a while-style loop with condition `n >= 1`
- Append each number to a result slice
- Decrement n after each iteration

See `solution.go` - the `Countdown()` function demonstrates this.

**Key Concepts:**
- The loop continues while the condition is true
- We decrement the counter inside the loop body
- This is cleaner than a traditional for loop for counting down

**Alternative Approach:** You could use a traditional for loop: `for i := start; i >= 1; i--`, but the while-style demonstrates the pattern more clearly and is more flexible if you need complex decrement logic.

---

## Exercise 2: Sum Until Threshold

**Goal:** Add numbers from a slice until reaching a threshold.

**Solution Approach:**
- Keep track of current sum and index
- Loop while sum is below threshold AND there are more numbers
- Add the next number to the sum
- Move to the next index
- Return when either condition fails

See `solution.go` - the `SumUntilThreshold()` function shows this pattern.

**Key Concepts:**
- Multiple conditions in the loop: check both the threshold and array bounds
- Early termination is the whole point - we don't process all elements
- The condition uses AND (&&) because both must be true to continue

**Edge Cases Handled:**
- Empty slice: returns 0 immediately
- Threshold already met: returns 0
- All numbers processed: returns the full sum

---

## Exercise 3: Read Until Sentinel

**Goal:** Process values until encountering a special "stop" value.

**Solution Approach:**
- Use an index to track position in the slice
- Loop while we haven't found the sentinel AND haven't reached the end
- Check if current value is the sentinel
- If not, add it to results and move to next
- If yes, break out of the loop

See `solution.go` - the `ReadUntilSentinel()` function demonstrates this.

**Key Concepts:**
- Sentinel values are common in data processing
- We check the sentinel BEFORE adding to results (so it's excluded)
- The loop condition could also use `break` inside for cleaner logic

**Real-World Usage:** This pattern is very common when:
- Reading user input until they type "quit"
- Processing file lines until a delimiter
- Parsing data with end markers

**Alternative Approach:** You could check for the sentinel in the loop condition, but checking inside allows for better control and clarity about what happens when found.

---

## Exercise 4: Find First Divisor

**Goal:** Find the first number >= start that divides n evenly.

**Solution Approach:**
- Start checking from the `start` value
- Loop while we haven't found a divisor
- Check if current candidate divides n evenly (remainder is 0)
- If yes, return it
- If no, try the next number
- To avoid infinite loops, stop at n itself (which always divides itself)

See `solution.go` - the `FindFirstDivisor()` function shows this search pattern.

**Key Concepts:**
- The modulo operator `%` gives the remainder
- If `n % candidate == 0`, candidate divides n evenly
- We're searching for the FIRST match, so return immediately when found
- Search loops are a classic use case for while-style iteration

**Edge Cases:**
- Prime numbers: only divisible by 1 and themselves
- Starting at 1: 1 divides everything
- n = 1: returns 1 (only divisor)

---

## Exercise 5: Collatz Sequence Length

**Goal:** Count steps in the Collatz sequence until reaching 1.

**Solution Approach:**
- Start with the given number
- Loop while n is not 1
- If even: divide by 2
- If odd: multiply by 3 and add 1
- Count each step
- Return the count when n reaches 1

See `solution.go` - the `CollatzLength()` function demonstrates this famous sequence.

**Key Concepts:**
- The loop condition is simple: `n != 1`
- The iteration logic is more complex (conditional inside the loop)
- This is perfect for while-style because we don't know how many steps it will take
- The Collatz conjecture states this always reaches 1, though it's unproven!

**Interesting Facts:**
- Also called the 3n+1 problem
- Some numbers take hundreds of steps to reach 1
- No one has proven that all numbers eventually reach 1 (famous unsolved problem)
- The sequence length varies wildly and unpredictably

**Example Trace:**
```
10 -> 5 (even, divide by 2)
5 -> 16 (odd, 3*5+1)
16 -> 8 (even)
8 -> 4 (even)
4 -> 2 (even)
2 -> 1 (even)
Total: 6 steps
```

---

## General Patterns

### Loop Termination Checklist

When writing while-style loops, ensure:
1. ✓ The condition will eventually become false
2. ✓ You're updating the variables that affect the condition
3. ✓ You handle edge cases (empty input, immediate match, etc.)
4. ✓ You consider adding a maximum iteration count for safety

### Debugging While Loops

If your loop seems stuck:
1. Print the condition variables at each iteration
2. Add a counter and break after N iterations
3. Verify your condition logic (especially AND vs OR)
4. Check that you're actually modifying the condition variables

### Common Patterns Summary

- **Countdown/Countup:** Numeric condition with increment/decrement
- **Accumulator:** Sum/collect while below threshold
- **Sentinel:** Read until special marker
- **Search:** Find first match then exit
- **Complex condition:** Check multiple things each iteration

---

## Tips for Success

1. **Start with the condition:** Think about what should keep the loop running
2. **Consider edge cases:** What if the condition is immediately false?
3. **Avoid infinite loops:** Always ensure progress toward termination
4. **Use break for clarity:** Sometimes checking inside and breaking is clearer than complex conditions
5. **Test thoroughly:** Try empty inputs, immediate matches, and normal cases
