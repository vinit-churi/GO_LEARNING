# Measuring String Builders Answers

## Hints

- Start from the smallest helper and compose outward.
- Keep the implementation explicit enough that a reader can trace it quickly.
- Reuse previous functions when the later exercise naturally builds on them.

## Exercise Walkthroughs

- `JoinWithBuilder` preallocates through the builder as values are appended.
- `JoinWithPlus` stays deliberately simple for side-by-side comparison.
- `ReportLine` keeps a realistic call site while reusing the builder implementation.
