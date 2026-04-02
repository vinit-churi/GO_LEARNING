# JSON Struct Tags Answers

## Hints

- Start with the smallest helper before composing the later exercises.
- Keep outputs stable so tests read cleanly.
- Reuse the earlier functions instead of duplicating logic.

## Exercise Walkthroughs

- Exercise 1: `EncodeEvent` marshals the struct and converts bytes to string. See `EncodeEvent` in `solution.go`.
- Exercise 2: `DecodeEvent` unmarshals into a typed `Event`, keeping field validation separate. See `DecodeEvent` in `solution.go`.
- Exercise 3: `SummaryLabel` formats a stable label from struct data, which is easier to test than ad hoc JSON inspection. See `SummaryLabel` in `solution.go`.
