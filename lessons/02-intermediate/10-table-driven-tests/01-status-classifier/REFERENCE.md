# Status Classifier Reference

## Concept Explanations

Table-driven tests are common in Go because they keep repeated test structure small while making input and output combinations easy to scan.

## Syntax Patterns

- Keep helpers small and name them for the caller.
- Prefer standard library packages when they already model the behavior you need.
- Use tests to lock down edge cases and make the intended contract explicit.

## Common Mistakes

- Skipping error handling when the lesson API returns an `error`.
- Hiding side effects that should be visible in the function signature.
- Repeating logic instead of composing the helper functions from the lesson.

## Why This Feature Exists

It gives Go programs a small, explicit tool for this part of application design without needing framework magic.

## When To Use It

Use it when the code is clearer, safer, or more testable with this feature.

## When Not To Use It

Do not force the pattern when a simpler sequential or concrete approach would be easier to understand.

## Notes

- Use table-driven tests when one function should be checked across many input combinations.
- Subtests with `t.Run` make failures easy to localize without duplicating setup.
- Keep test tables close to the assertion logic so learners can read both together.
- Do not force table-driven style for every test; direct tests are still fine for one-off cases.
- Common mistake: hiding test intent behind too many helper layers.
- Official docs: `testing` package examples and Go wiki guidance on table-driven tests.
