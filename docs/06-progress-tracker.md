# Progress Tracker

Use this file to record what is complete, what is in progress, and what the next actor should do.

## Overall Status

- repository planning: `completed`
- curriculum content creation: `completed`

## Phase Status

| Phase | Status | Next Action |
|---|---|---|
| Planning docs | completed | Keep docs in sync with implementation |
| Beginner | completed | Keep the staged beginner lessons stable and expand density later if desired |
| Intermediate | completed | Expand lesson density later if desired, but keep the track stable |
| Advanced | completed | Expand density, examples, or projects later if desired |

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

There is no pending phase plan now.

If implementation resumes, the next useful work is:

1. increase lesson density inside completed modules
2. add more optional capstones or project polish where desired

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
- a staged lesson now exists for every planned intermediate module from `01-pointers-and-semantics` through `18-http-servers`
- a staged lesson now exists for every planned advanced module from `01-generics-fundamentals` through `14-practical-mini-projects`
- the advanced track includes three mini-project lessons tying together prior topics
- `go test ./...` passes with the completed staged beginner, intermediate, and advanced tracks

## Update Rule

Whenever work is done:

1. change the relevant phase status
2. list what was created
3. point to the next pending batch
