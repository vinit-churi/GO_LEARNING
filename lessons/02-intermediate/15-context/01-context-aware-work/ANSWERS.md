# Context-Aware Work Answers

## Hints

- Start with the smallest helper before composing the later exercises.
- Keep outputs stable so tests read cleanly.
- Reuse the earlier functions instead of duplicating logic.

## Exercise Walkthroughs

- Exercise 1: `TraceID` reads a typed key and safely falls back to an empty string. See `TraceID` in `solution.go`.
- Exercise 2: `WorkOrCancel` uses `select` to prefer a completed context over simulated work. See `WorkOrCancel` in `solution.go`.
- Exercise 3: `LabelRequest` keeps formatting separate from cancellation logic while still using context metadata. See `LabelRequest` in `solution.go`.
