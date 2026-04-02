# Status Classifier Answers

## Hints

- Start with the smallest helper before composing the later exercises.
- Keep outputs stable so tests read cleanly.
- Reuse the earlier functions instead of duplicating logic.

## Exercise Walkthroughs

- Exercise 1: `ClassifyScore` maps score ranges to stable labels, which works well with table-driven coverage. See `ClassifyScore` in `solution.go`.
- Exercise 2: `IsTerminal` keeps final-state policy in one helper instead of scattering string checks. See `IsTerminal` in `solution.go`.
- Exercise 3: `BuildStatusLine` combines the helpers into one formatted string. See `BuildStatusLine` in `solution.go`.
