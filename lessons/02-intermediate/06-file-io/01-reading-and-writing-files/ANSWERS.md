# Reading and Writing Files Answers

## Hints

- Start with the smallest helper before composing the later exercises.
- Keep outputs stable so tests read cleanly.
- Reuse the earlier functions instead of duplicating logic.

## Exercise Walkthroughs

- Exercise 1: `SaveNote` delegates to `os.WriteFile` with a standard permission mode. See `SaveNote` in `solution.go`.
- Exercise 2: `LoadNote` reads bytes and trims surrounding whitespace before returning a string. See `LoadNote` in `solution.go`.
- Exercise 3: `DuplicateNote` composes the first two helpers so copy behavior stays small and testable. See `DuplicateNote` in `solution.go`.
