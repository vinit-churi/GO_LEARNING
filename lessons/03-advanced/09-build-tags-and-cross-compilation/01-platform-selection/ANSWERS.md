# Platform Selection Answers

## Hints

- Start from the smallest helper and compose outward.
- Keep the implementation explicit enough that a reader can trace it quickly.
- Reuse previous functions when the later exercise naturally builds on them.

## Exercise Walkthroughs

- `taggedPlatform` is implemented in separate files selected by the compiler.
- `BuildArtifactName` shows how shared logic can layer on top of the tagged helper.
- `TargetTriple` keeps the representation stable and easy to test.
