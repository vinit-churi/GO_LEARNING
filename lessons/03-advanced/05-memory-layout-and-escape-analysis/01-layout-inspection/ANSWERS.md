# Layout Inspection Answers

## Hints

- Start from the smallest helper and compose outward.
- Keep the implementation explicit enough that a reader can trace it quickly.
- Reuse previous functions when the later exercise naturally builds on them.

## Exercise Walkthroughs

- `RecordSize` and `ValueOffset` expose concrete layout facts for a specific struct.
- `HasPadding` checks whether the aligned field starts after the first byte-sized field.
- The lesson keeps layout inspection measurable while leaving compiler-specific escape output to the reference notes.
