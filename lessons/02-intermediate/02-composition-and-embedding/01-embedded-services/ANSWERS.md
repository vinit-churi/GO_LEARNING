# Embedded Services Answers

## Hints

- Start with the smallest helper before composing the later exercises.
- Keep outputs stable so tests read cleanly.
- Reuse the earlier functions instead of duplicating logic.

## Exercise Walkthroughs

- Exercise 1: `FullName` trims and joins the owner and team. See `AuditInfo.FullName` in `solution.go`.
- Exercise 2: `Summary` reads promoted fields from the embedded struct and formats a deployment summary. See `DeployService.Summary` in `solution.go`.
- Exercise 3: `Touch` writes through the embedded field so the outer struct stays updated. See `DeployService.Touch` in `solution.go`.
