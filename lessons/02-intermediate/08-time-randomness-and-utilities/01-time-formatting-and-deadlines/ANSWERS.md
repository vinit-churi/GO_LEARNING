# Time Formatting and Deadlines Answers

## Hints

- Start with the smallest helper before composing the later exercises.
- Keep outputs stable so tests read cleanly.
- Reuse the earlier functions instead of duplicating logic.

## Exercise Walkthroughs

- Exercise 1: `RFC3339Label` converts to UTC before formatting so test output is stable. See `RFC3339Label` in `solution.go`.
- Exercise 2: `ExtendDeadline` adds a duration built from whole minutes. See `ExtendDeadline` in `solution.go`.
- Exercise 3: `MinutesRemaining` subtracts times and divides by `time.Minute` to get a whole-number result. See `MinutesRemaining` in `solution.go`.
