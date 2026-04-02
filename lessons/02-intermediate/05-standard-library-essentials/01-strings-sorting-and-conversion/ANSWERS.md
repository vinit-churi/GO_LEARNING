# Strings, Sorting, and Conversion Answers

## Hints

- Start with the smallest helper before composing the later exercises.
- Keep outputs stable so tests read cleanly.
- Reuse the earlier functions instead of duplicating logic.

## Exercise Walkthroughs

- Exercise 1: `CleanTags` normalizes each tag and skips empty results, keeping the API focused on clean output. See `CleanTags` in `solution.go`.
- Exercise 2: `SortedScores` parses each score and sorts the resulting integers in place. See `SortedScores` in `solution.go`.
- Exercise 3: `JoinTags` reuses `CleanTags` so cleanup rules stay centralized. See `JoinTags` in `solution.go`.
