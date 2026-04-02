# Package Layout Basics Answers

## Hints

- Start with the smallest helper before composing the later exercises.
- Keep outputs stable so tests read cleanly.
- Reuse the earlier functions instead of duplicating logic.

## Exercise Walkthroughs

- Exercise 1: `slugutil.Slug` performs the shared normalization rules once in a reusable package. See `slugutil/slugutil.go`.
- Exercise 2: `BuildModulePath` composes owner and repo slugs into a GitHub-style module path. See `BuildModulePath` in `solution.go`.
- Exercise 3: `IsPublicPackage` encodes a naming policy in one small helper. See `IsPublicPackage` in `solution.go`.
