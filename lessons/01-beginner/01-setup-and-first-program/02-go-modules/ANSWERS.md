# Go Modules - Answers

## Exercise 1: Module Path

**Goal:** Return a valid module path string.

**Solution Approach:**
- Create a function that returns a string
- The string should be a valid module path format
- Common pattern: `domain.com/username/project`

See `solution.go` - the `GetModulePath()` function returns a hardcoded module path.

**Key Concept:** The module path is just a string identifier. It doesn't have to point to a real URL, but it's conventional to use a format that reflects where the code is hosted.

---

## Exercise 2: Package Information

**Goal:** Return formatted information about the package.

**Solution Approach:**
- Use a multi-line string or concatenate strings
- Include package name and purpose
- Format clearly for readability

See `solution.go` - the `GetPackageInfo()` function builds a formatted string.

**Key Concepts:**
- You can build strings using the `+` operator
- `\n` creates a newline
- Keep information clear and concise

**Alternative Approach:** You could use `fmt.Sprintf` for formatted strings, but simple concatenation works fine for beginners.

---

## Exercise 3: Module Version

**Goal:** Return a semantic version string.

**Solution Approach:**
- Return a string in the format `vMAJOR.MINOR.PATCH`
- Start with `v1.0.0` for simplicity
- Follow semantic versioning conventions

See `solution.go` - the `GetModuleVersion()` function returns a semantic version.

**Key Concept:** Semantic versioning uses three numbers: major.minor.patch. The `v` prefix is standard in Go.

---

## Exercise 4: Full Module Info

**Goal:** Combine all information into a comprehensive report.

**Solution Approach:**
- Call the previous functions to get individual pieces
- Format them into a nice report
- Use borders or separators for visual clarity
- Use `fmt.Println` to print the output

See `solution.go` - the `PrintModuleReport()` function demonstrates this.

**Key Concepts:**
- Functions can call other functions to build complex behavior
- Breaking functionality into small functions makes code more maintainable
- Formatting output well improves user experience

**Design Pattern:** This exercise demonstrates the Single Responsibility Principle - each function does one thing well, and they compose together to create more complex behavior.

---

## General Tips

1. **Module paths are identifiers:** They don't need to be real URLs during learning
2. **Consistency matters:** Use consistent naming across your module
3. **Version numbers communicate change:** Follow semver to help users understand updates
4. **Small functions compose well:** Breaking tasks into functions makes code reusable

## Common Pitfalls

- **Invalid module paths:** Avoid spaces and special characters
- **Forgetting the v prefix:** Versions should be `v1.0.0` not `1.0.0`
- **Hardcoding vs. configuration:** In real applications, version info might come from build flags or configuration files
- **String formatting:** Be careful with newlines and spacing for readable output

## Real-World Connection

In production applications:
- Module paths match your repository location
- Version numbers come from git tags
- Package documentation is generated from comments
- Build tools inject version info at compile time

For now, hardcoded values help you understand the concepts before adding complexity.
