# Native Boundary Design Answers

## Hints

- Start from the smallest helper and compose outward.
- Keep the implementation explicit enough that a reader can trace it quickly.
- Reuse previous functions when the later exercise naturally builds on them.

## Exercise Walkthroughs

- `NativeCall` depends only on the interface, which makes it testable without native code.
- `SumPair` shows how higher-level helpers can stay unaware of the underlying implementation.
- `ReportPair` formats results at the boundary instead of mixing string handling with the driver.
