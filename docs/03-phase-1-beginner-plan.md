# Phase 1 Plan: Beginner

Status: `completed`

## Goal

Build the beginner path first so the repo quickly becomes useful for a learner starting from near zero.

## Scope

This phase should cover:

1. setup and first program
2. variables, constants, and types
3. formatting and input/output
4. conditionals
5. loops
6. functions
7. scope and packages
8. arrays and slices
9. maps
10. strings, runes, and bytes
11. structs
12. methods
13. interfaces basics
14. errors basics
15. intro to testing

## Recommended Build Order

Implement in batches:

1. Core syntax and flow
2. Collections and strings
3. Structs, methods, and interfaces
4. Errors and testing

## Batch 1 Detail

Start with these modules:

1. `01-setup-and-first-program`
2. `02-basics`
3. `03-conditionals`
4. `04-loops`
5. `05-functions`

For each module:

- create 3 to 5 lessons
- create at least 3 exercises per lesson
- ensure tests are present before marking the lesson complete

## Batch 1 Status

Completed and verified:

1. `01-setup-and-first-program`
2. `02-basics`
3. `03-conditionals`
4. `04-loops`
5. `05-functions`

Verification:

- full lesson file sets exist for the staged lessons in these modules
- `go test ./...` passes across the repository

## Final Beginner Batch Status

The remaining beginner modules after functions have now been implemented:

1. `06-scope-and-packages`
2. `07-arrays-and-slices`
3. `08-maps`
4. `09-strings-runes-and-bytes`
5. `10-structs`
6. `11-methods`
7. `12-interfaces-basics`
8. `13-errors-basics`
9. `14-intro-to-testing`

Implementation notes:

- each new module now exists under `lessons/01-beginner/`
- each implemented lesson has `README.md`, `REFERENCE.md`, `ANSWERS.md`, `starter.go`, `solution.go`, and `lesson_test.go`
- the scope and packages lesson also includes a small local support package to demonstrate importing exported identifiers
- `go test ./...` passes across the repository with the new staged lessons included

## Output Standard

Every lesson created in this phase must follow [02-content-spec.md](/Users/vinitchuri/projects/go_learning/docs/02-content-spec.md).

## Completion Criteria

Phase 1 is complete when:

- all beginner modules listed in scope exist
- each lesson has the full file set
- the beginner curriculum can be navigated in order
- tests pass for created lessons

Phase 1 is now complete.

## Notes For Future LLMs

- Phase 1 is complete.
- Continue with [04-phase-2-intermediate-plan.md](/Users/vinitchuri/projects/go_learning/docs/04-phase-2-intermediate-plan.md) for the next staged work.
- Keep [06-progress-tracker.md](/Users/vinitchuri/projects/go_learning/docs/06-progress-tracker.md) in sync with future intermediate and advanced implementation.
