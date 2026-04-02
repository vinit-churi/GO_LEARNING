# Parsing Input Answers

## Hints

- Start with the smallest helper before composing the later exercises.
- Keep outputs stable so tests read cleanly.
- Reuse the earlier functions instead of duplicating logic.

## Exercise Walkthroughs

- Exercise 1: `ParsePair` splits once, trims the key, and parses the numeric value. See `ParsePair` in `solution.go`.
- Exercise 2: `MustParseCount` intentionally collapses parse failures to zero for a lenient caller. See `MustParseCount` in `solution.go`.
- Exercise 3: `FormatPair` provides a canonical serializer that pairs well with parser tests. See `FormatPair` in `solution.go`.
