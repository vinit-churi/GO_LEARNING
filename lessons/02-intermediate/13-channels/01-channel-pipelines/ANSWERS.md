# Channel Pipelines Answers

## Hints

- Start with the smallest helper before composing the later exercises.
- Keep outputs stable so tests read cleanly.
- Reuse the earlier functions instead of duplicating logic.

## Exercise Walkthroughs

- Exercise 1: `Feed` launches a producer goroutine that sends all values then closes the output channel. See `Feed` in `solution.go`.
- Exercise 2: `SquareStream` reads each input value, transforms it, and closes its own output when upstream ends. See `SquareStream` in `solution.go`.
- Exercise 3: `Collect` gives tests a plain slice result so assertions stay simple. See `Collect` in `solution.go`.
