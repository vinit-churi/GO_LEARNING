# Golden and Helper Patterns Answers

## Hints

- Start from the smallest helper and compose outward.
- Keep the implementation explicit enough that a reader can trace it quickly.
- Reuse previous functions when the later exercise naturally builds on them.

## Exercise Walkthroughs

- `NormalizeReport` turns formatting variance into a stable representation for tests.
- `BuildGolden` keeps expected output generation explicit and deterministic.
- `ReportsMatch` centralizes comparison rules so tests can express intent directly.
