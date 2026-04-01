# Hello World - Answers

## Exercise 1: Hello, World!

**Goal:** Print the classic "Hello, World!" message.

**Solution Approach:**
- Import the `fmt` package
- Call `fmt.Println()` with the string "Hello, World!"

See `solution.go` - the `exercise1()` function demonstrates this.

**Key Concept:** This is the simplest complete Go program. The `main` function is the entry point, and `fmt.Println` outputs text to the console.

---

## Exercise 2: Greet Yourself

**Goal:** Print a personalized greeting.

**Solution Approach:**
- Use multiple `fmt.Println()` calls for multiple lines
- Each call automatically adds a newline

See `solution.go` - the `exercise2()` function shows two separate print statements.

**Key Concept:** Each `fmt.Println()` call is independent. You can call it as many times as you need.

---

## Exercise 3: Multiple Lines

**Goal:** Print three different messages on separate lines.

**Solution Approach:**
- Use three `fmt.Println()` calls
- Each statement prints one complete thought
- The messages build a coherent narrative

See `solution.go` - the `exercise3()` function demonstrates this pattern.

**Key Concept:** Sequential calls create sequential output. This is how you build up complex output from simple pieces.

---

## Exercise 4: ASCII Art

**Goal:** Create a visual welcome message.

**Solution Approach:**
- Use multiple `fmt.Println()` calls with ASCII art characters
- Empty `fmt.Println()` adds a blank line for spacing
- Backslashes and special characters may need to be escaped

See `solution.go` - the `exercise4()` function shows a simple ASCII art example.

**Key Concepts:**
- Visual output is just text arranged carefully
- Empty `fmt.Println()` creates spacing
- You can use any characters in your strings

**Alternative Approach:** You could also use a single raw string literal (backticks) with embedded newlines for multi-line output, but multiple `Println` calls are clearer for beginners.

---

## General Tips

1. **Start Simple:** The best way to learn is to type the code yourself
2. **Experiment:** Try changing the strings and see what happens
3. **Read Error Messages:** If something doesn't work, Go's error messages are usually helpful
4. **Run Often:** Use `go run solution.go` frequently to see your changes

## Common Pitfalls

- **Forgetting to import fmt:** If you use `fmt.Println` but don't import `fmt`, the code won't compile
- **Using single quotes:** Strings use double quotes `"text"`, not single quotes
- **Capitalization matters:** It's `Println` not `println` - Go is case-sensitive
