# Validation Errors Answers

## Hints

- Start with the smallest helper before composing the later exercises.
- Keep outputs stable so tests read cleanly.
- Reuse the earlier functions instead of duplicating logic.

## Exercise Walkthroughs

- Exercise 1: `ValidationError` formats field and message into one stable string. See `ValidationError.Error` in `solution.go`.
- Exercise 2: `ValidateUsername` trims whitespace and returns a custom error when the input is empty. See `ValidateUsername` in `solution.go`.
- Exercise 3: `LoadUsername` reuses `ValidateUsername` and wraps failures so callers keep both the higher-level operation and the root cause. See `LoadUsername` in `solution.go`.
