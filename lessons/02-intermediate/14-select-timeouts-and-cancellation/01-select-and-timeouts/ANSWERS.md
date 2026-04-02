# Select and Timeouts Answers

## Hints

- Start with the smallest helper before composing the later exercises.
- Keep outputs stable so tests read cleanly.
- Reuse the earlier functions instead of duplicating logic.

## Exercise Walkthroughs

- Exercise 1: `ReceiveWithTimeout` returns the incoming value or the string `timeout` if the timer wins. See `ReceiveWithTimeout` in `solution.go`.
- Exercise 2: `ImmediateStatus` uses `default` to avoid blocking. See `ImmediateStatus` in `solution.go`.
- Exercise 3: `FirstResult` demonstrates the simplest fan-in race between two channels. See `FirstResult` in `solution.go`.
