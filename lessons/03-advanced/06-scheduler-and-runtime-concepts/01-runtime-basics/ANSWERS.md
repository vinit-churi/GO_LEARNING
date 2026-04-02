# Runtime Basics Answers

## Hints

- Start from the smallest helper and compose outward.
- Keep the implementation explicit enough that a reader can trace it quickly.
- Reuse previous functions when the later exercise naturally builds on them.

## Exercise Walkthroughs

- `CPUCount` is a direct wrapper around runtime metadata.
- `ChooseWorkers` avoids returning more workers than the machine can plausibly run in parallel.
- `ParallelSummary` keeps the measurement output stable for tests and examples.
