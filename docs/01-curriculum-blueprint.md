# Curriculum Blueprint

This blueprint defines the intended learning path and the recommended lesson groupings.

## Stage 1: Beginner

Goal: become comfortable writing small Go programs, understanding syntax, and solving focused problems with tests.

Recommended modules:

1. Setup and first program
2. Variables, constants, and basic types
3. Input, output, and formatting
4. Conditionals
5. Loops
6. Functions
7. Scope and packages
8. Arrays and slices
9. Maps
10. Strings, runes, and bytes
11. Structs
12. Methods
13. Interfaces basics
14. Error basics
15. Intro to testing

## Stage 2: Intermediate

Goal: write idiomatic multi-file programs and understand core Go design patterns.

Recommended modules:

1. Pointers and value vs reference semantics
2. Composition and embedding
3. Interfaces in practice
4. Custom errors and error wrapping
5. Standard library essentials
6. File I/O
7. JSON and encoding
8. Time, randomness, and utilities
9. Modules and package design
10. Table-driven tests
11. Benchmarks and fuzz tests
12. Goroutines
13. Channels
14. Select, timeouts, and cancellation
15. Context
16. Mutexes, atomics, and race detection
17. HTTP clients
18. HTTP servers

## Stage 3: Advanced

Goal: understand how Go behaves at scale and how to use advanced features responsibly.

Recommended modules:

1. Generics fundamentals
2. Generic constraints and design tradeoffs
3. Reflection
4. Advanced interface design
5. Memory layout and escape analysis
6. Go scheduler and runtime concepts
7. Memory model
8. Profiling and optimization
9. Build tags and cross-compilation
10. Workspaces and larger codebase structure
11. Unsafe package
12. cgo overview and boundaries
13. Advanced testing strategy
14. Practical mini-projects

## Cross-Cutting Content Rules

- Every module should contain multiple lessons.
- Every lesson should contain multiple exercises.
- Exercises should progress from direct syntax practice to slightly open-ended problem solving.
- Use real naming and real scenarios where possible, not only toy arithmetic examples.
- Keep examples small enough for focused learning.

## Proposed Implementation Sequence

Build the curriculum in batches:

1. Beginner foundations
2. Beginner data structures and methods
3. Beginner interfaces, errors, and tests
4. Intermediate program design and standard library
5. Intermediate concurrency and networking
6. Advanced language internals and performance
7. Capstone mini-projects

## Minimum Lesson Density

Recommended target for planning purposes:

- 3 to 6 lessons per module
- 3 to 5 exercises per lesson

This target can be adjusted later, but future work should avoid creating thin one-exercise lessons unless the topic is intentionally tiny.
