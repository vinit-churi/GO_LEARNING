# Progress Tracker

Use this file to record what is complete, what is in progress, and what the next actor should do.

## Overall Status

- repository planning: `completed`
- curriculum content creation: `in_progress`

## Phase Status

| Phase | Status | Next Action |
|---|---|---|
| Planning docs | completed | Keep docs in sync with implementation |
| Beginner | completed | Keep the staged beginner lessons stable and expand density later if desired |
| Intermediate | pending | Start phase 2 when ready |
| Advanced | pending | Wait for intermediate coverage |

## Decisions Already Made

- Use staged curriculum folders under `lessons/`
- Keep old lesson content for now; do not migrate it yet
- Use `README.md`, `REFERENCE.md`, and `ANSWERS.md` in every new lesson
- Use separate `starter.go` and `solution.go`
- Default validation path is `go test`
- Some lessons may also include `go run` examples
- Multiple exercises per lesson are required
- Teaching style should be FreeCodeCamp-like: concept, examples, practice, challenge

## Immediate Next Step

The staged beginner curriculum now covers all planned phase 1 modules.

When implementation resumes, start with:

1. [04-phase-2-intermediate-plan.md](/Users/vinitchuri/projects/go_learning/docs/04-phase-2-intermediate-plan.md)
2. `lessons/02-intermediate/` for the first intermediate batch

## Completed So Far

The staged beginner curriculum currently has working lessons in:

1. `lessons/01-beginner/01-setup-and-first-program`
2. `lessons/01-beginner/02-basics`
3. `lessons/01-beginner/03-conditionals`
4. `lessons/01-beginner/04-loops`
5. `lessons/01-beginner/05-functions`
6. `lessons/01-beginner/06-scope-and-packages`
7. `lessons/01-beginner/07-arrays-and-slices`
8. `lessons/01-beginner/08-maps`
9. `lessons/01-beginner/09-strings-runes-and-bytes`
10. `lessons/01-beginner/10-structs`
11. `lessons/01-beginner/11-methods`
12. `lessons/01-beginner/12-interfaces-basics`
13. `lessons/01-beginner/13-errors-basics`
14. `lessons/01-beginner/14-intro-to-testing`

Additional notes:

- `lessons/01-beginner/03-conditionals/03-comparison-operators` was completed during this pass
- `lessons/01-beginner/05-functions/01-basic-functions` was finished and now passes tests
- a new staged lesson was added at `lessons/01-beginner/01-setup-and-first-program/01-hello-world-and-formatting`
- a new staged lesson was added at `lessons/01-beginner/02-basics/01-variables-constants-and-zero-values`
- a new staged lesson was added at `lessons/01-beginner/03-conditionals/01-access-control`
- a new staged lesson was added at `lessons/01-beginner/04-loops/01-message-batches`
- a new staged lesson was added at `lessons/01-beginner/05-functions/01-order-summary`
- a staged lesson now exists for every remaining beginner module from `06-scope-and-packages` through `14-intro-to-testing`
- `go test ./...` passes with the completed staged beginner track
- intermediate and advanced phases remain pending by design

## Update Rule

Whenever work is done:

1. change the relevant phase status
2. list what was created
3. point to the next pending batch
