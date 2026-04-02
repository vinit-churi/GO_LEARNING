# Writer-Based Reporting Answers

## Hints

- Start with the smallest helper before composing the later exercises.
- Keep outputs stable so tests read cleanly.
- Reuse the earlier functions instead of duplicating logic.

## Exercise Walkthroughs

- Exercise 1: `WriteLine` delegates to `fmt.Fprintln`, which already appends a newline and returns an error. See `WriteLine` in `solution.go`.
- Exercise 2: `EmitReport` formats a stable string and reuses `WriteLine` to avoid duplicate output logic. See `EmitReport` in `solution.go`.
- Exercise 3: `BufferReport` writes into `bytes.Buffer` and returns the captured string for tests. See `BufferReport` in `solution.go`.
