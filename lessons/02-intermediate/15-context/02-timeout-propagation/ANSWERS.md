# Timeout Propagation Answers

## Hints

- Wire the timeout once, then pass the derived context down.
- `select` should race work completion against cancellation.
- Keep composition helpers small and readable.

## Exercise Walkthroughs

- Exercise 1: `WithRequestTimeout` calls `context.WithTimeout` and returns both values unchanged.
- Exercise 2: `WaitForWork` uses `select` between `workDone` and `ctx.Done`, returning `ctx.Err()` when canceled.
- Exercise 3: `RunWithTimeout` composes both helpers, defers `cancel`, then waits for work.
