# Range Loops - Answers

## Exercise 1: Sum of Slice Elements

**Goal:** Calculate the sum of all numbers in a slice using range.

**Solution Approach:**
- Initialize a `sum` variable to 0
- Use `for _, num := range numbers` to iterate over values (ignore index with `_`)
- Add each number to the sum
- Return the sum

**Key Concepts:**
- Use `_` to ignore the index when you only need the value
- Range handles empty slices gracefully (loop body doesn't execute)
- Initialize accumulator before the loop

See `solution.go` - the `SumSlice()` function demonstrates this pattern.

**Why ignore the index?** Since we're just summing values, we don't need to know the position of each element. Using `_` makes the code cleaner and signals our intent.

---

## Exercise 2: Find Index of Element

**Goal:** Search for a target string and return its position.

**Solution Approach:**
- Use `for index, item := range items` to get both index and value
- Compare each item to the target
- Return the index immediately when found
- Return -1 if the loop completes without finding the target

**Key Concepts:**
- Need both index and value, so use both variables from range
- Early return when condition is met
- Return a sentinel value (-1) to indicate "not found"

See `solution.go` - the `FindIndex()` function shows this search pattern.

**Alternative Approach:** You could use a traditional for loop with manual indexing, but range is more idiomatic and less error-prone.

---

## Exercise 3: Count Vowels in String

**Goal:** Count vowel characters, properly handling Unicode.

**Solution Approach:**
- Define vowels as a string: `"aeiouAEIOU"`
- Use `for _, char := range s` to iterate over runes
- Use `strings.ContainsRune()` to check if char is a vowel
- Increment counter when a vowel is found

**Key Concepts:**
- Ranging over a string gives you runes (Unicode code points), not bytes
- This correctly handles multi-byte UTF-8 characters like "é"
- `strings.ContainsRune()` is a convenient way to check membership
- Works with both uppercase and lowercase vowels

See `solution.go` - the `CountVowels()` function demonstrates string iteration.

**Why runes matter:** If you iterated over bytes, multi-byte characters like "café" would be counted incorrectly. Range automatically handles UTF-8 decoding.

**Note about accented vowels:** The basic implementation counts standard English vowels (a, e, i, o, u). Accented characters like "é" are not counted. To include accented vowels, you'd need a more comprehensive list or Unicode category checking.

**Alternative Approach:** You could also use a switch statement or a map to check for vowels, but `strings.ContainsRune()` is concise and clear.

---

## Exercise 4: Double Each Element

**Goal:** Transform a slice by doubling each value.

**Solution Approach:**
- Create a new result slice with appropriate capacity
- Use `for _, num := range numbers` (ignore index)
- Append doubled value to result
- Return the new slice

**Key Concepts:**
- Use `_` when you only need the value
- Pre-allocating capacity with `make([]int, 0, len(numbers))` is more efficient
- Build a new slice rather than modifying in place
- Works correctly with empty slices

See `solution.go` - the `DoubleElements()` function shows this transformation pattern.

**Why create a new slice?** This follows functional programming principles and avoids side effects. The original slice remains unchanged.

**Performance tip:** Using `make()` with capacity avoids repeated allocations as the slice grows.

---

## Exercise 5: Reverse String Using Runes

**Goal:** Reverse a string while properly handling Unicode characters.

**Solution Approach:**
- Convert string to a slice of runes: `[]rune(s)`
- Use two-pointer technique to swap elements from both ends
- Work inward until pointers meet
- Convert back to string: `string(runes)`

**Key Concepts:**
- Must convert to runes first to handle multi-byte characters correctly
- Traditional two-pointer swap: `runes[i], runes[j] = runes[j], runes[i]`
- The range loop isn't used for reversal, but rune conversion is crucial
- Multiple assignment in Go makes swapping elegant

See `solution.go` - the `ReverseString()` function demonstrates this approach.

**Why this approach?** Converting to a rune slice allows us to reverse by swapping elements. Trying to reverse a string directly would fail with Unicode characters because strings are immutable and indexed by bytes, not characters.

**Alternative Approach:** You could iterate in reverse order and build a new string by prepending each rune, but swapping in place is more efficient:

```go
func ReverseStringAlternative(s string) string {
    var result string
    for _, char := range s {
        result = string(char) + result  // Prepend
    }
    return result
}
```

This works but is less efficient due to repeated string concatenation.

---

## Common Patterns and Tips

### When to Use Each Range Form

1. **Both index and value:** `for i, v := range slice`
   - When you need the position for indexing or logic
   - Example: updating the original slice

2. **Value only:** `for _, v := range slice`
   - Most common pattern
   - When you just need to process each element

3. **Index only:** `for i := range slice`
   - When you need position but not the value
   - Example: counting elements with a condition

4. **Neither (just iterate):** `for range slice`
   - Rare, but useful for executing code N times
   - Example: `for range make([]int, 10)` runs 10 times

### Performance Considerations

- Range is as fast as traditional for loops (compiler optimizes them)
- Pre-allocating slice capacity avoids reallocation
- Range copies values, but for basic types (int, string) this is negligible
- For large structs, consider ranging over indices and accessing elements directly

### Debugging Tips

- Print index and value together to understand iteration: `fmt.Printf("i=%d v=%v\n", i, v)`
- For strings, print both the rune and its numeric value: `fmt.Printf("%c (%d)\n", r, r)`
- Use the debugger to step through range loops and watch variables

---

## Additional Practice Ideas

1. **Filter a slice:** Create a new slice with only elements matching a condition
2. **Find maximum:** Track the largest element while ranging
3. **Check if all match:** Return false early if any element fails a condition
4. **Pair elements:** Iterate with both current and next element
5. **Count occurrences:** Use a map while ranging to count element frequencies

Each of these builds on the range patterns you've learned!
