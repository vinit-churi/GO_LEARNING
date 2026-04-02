# Main Orchestrator

This is the authoritative planning document for building the repository into a full FreeCodeCamp-style Go study guide.

## Objective

Create a complete Go learning path from beginner to advanced, with multiple exercises per lesson, strong concept explanation, runnable code, and tests. The result should help a learner study progressively while also remaining structured enough for future LLMs to continue implementation safely.

## Final Target Shape

The curriculum should eventually live under `lessons/` and follow a staged structure:

```text
lessons/
  01-beginner/
    01-setup-and-first-program/
      01-hello-world/
      02-project-structure/
    02-basics/
      01-variables-and-constants/
      02-types-and-zero-values/
      ...
  02-intermediate/
    ...
  03-advanced/
    ...
```

Each lesson folder should eventually contain:

- `README.md`: learning goals, problem statements, exercise instructions, run instructions
- `REFERENCE.md`: concept explanation, syntax notes, tradeoffs, idioms, common mistakes
- `ANSWERS.md`: reasoning, hints, solution walkthroughs, and references to solution files
- `starter.go`: incomplete or scaffolded learner-facing code
- `solution.go`: reference implementation
- `lesson_test.go`: tests covering the exercises
- optional extra files for multi-file package lessons

## Why This Structure

- `README.md` is for active learning and tasks.
- `REFERENCE.md` is for studying concepts without giving away the exercise flow immediately.
- `ANSWERS.md` gives explanations plus where to look in the solution code.
- Separate `starter.go` and `solution.go` avoids mixing learner work with the answer.
- `go test` remains the main validation path, while selected lessons may also include `go run` examples where that improves learning.

## Curriculum Policy

The curriculum should cover Go from basic to advanced, including:

- language fundamentals
- control flow
- functions
- packages and modules
- arrays, slices, maps
- strings, runes, bytes
- structs and methods
- interfaces and composition
- error handling
- testing
- concurrency and channels
- context and cancellation
- file I/O and standard library usage
- HTTP clients and servers
- JSON and encoding
- generics
- benchmarking and profiling
- reflection
- memory model and runtime behavior
- unsafe and lower-level topics
- build tooling, workspaces, and practical project structure

Coverage should remain practical. Deep topics like `unsafe`, reflection internals, or memory model details belong later and should emphasize when not to use them.

## Lesson Design Standard

Every lesson should follow this flow:

1. Short concept setup
2. Learning objectives
3. One or more worked examples
4. Multiple exercises with increasing difficulty
5. Clear run/test instructions
6. Solution walkthrough in `ANSWERS.md`
7. Deeper explanation in `REFERENCE.md`

## Implementation Order

Always build in this order:

1. beginner
2. intermediate
3. advanced

Do not jump ahead unless a prerequisite plan explicitly says so.

## Source of Truth Files

Read these in order before doing work:

1. [06-progress-tracker.md](/Users/vinitchuri/projects/go_learning/docs/06-progress-tracker.md)
2. [01-curriculum-blueprint.md](/Users/vinitchuri/projects/go_learning/docs/01-curriculum-blueprint.md)
3. [02-content-spec.md](/Users/vinitchuri/projects/go_learning/docs/02-content-spec.md)
4. The phase plan currently marked `in_progress`

## How Future LLMs Should Continue

When starting a new turn:

1. Read this file.
2. Read [06-progress-tracker.md](/Users/vinitchuri/projects/go_learning/docs/06-progress-tracker.md).
3. Find the earliest incomplete phase.
4. Read that phase plan.
5. Implement one coherent batch only.
6. Update tracker and plan status before stopping.

## Definition of Done For A Lesson

A lesson is complete only if all of the following are true:

- folder path matches the curriculum convention
- `README.md` exists and includes at least 3 exercises
- `REFERENCE.md` exists and explains the lesson concepts clearly
- `ANSWERS.md` exists and explains the solution approach
- starter code exists
- solution code exists
- tests exist and pass with `go test`
- naming is consistent with surrounding lessons

## Current Status

Planning docs are complete, and the staged beginner path is now built end to end and passing tests.

Completed staged beginner modules:

- `01-setup-and-first-program`
- `02-basics`
- `03-conditionals`
- `04-loops`
- `05-functions`
- `06-scope-and-packages`
- `07-arrays-and-slices`
- `08-maps`
- `09-strings-runes-and-bytes`
- `10-structs`
- `11-methods`
- `12-interfaces-basics`
- `13-errors-basics`
- `14-intro-to-testing`

Current next file to act on:

- no pending phase plan; expand density or projects only if desired

The staged intermediate path is now built end to end and passes tests.

Completed staged intermediate modules:

- `01-pointers-and-semantics`
- `02-composition-and-embedding`
- `03-interfaces-in-practice`
- `04-custom-errors-and-error-wrapping`
- `05-standard-library-essentials`
- `06-file-io`
- `07-json-and-encoding`
- `08-time-randomness-and-utilities`
- `09-modules-and-package-design`
- `10-table-driven-tests`
- `11-benchmarks-and-fuzz-tests`
- `12-goroutines`
- `13-channels`
- `14-select-timeouts-and-cancellation`
- `15-context`
- `16-mutexes-atomics-and-race-detection`
- `17-http-clients`
- `18-http-servers`

The staged advanced path is now built end to end and passes tests.

Completed staged advanced modules:

- `01-generics-fundamentals`
- `02-generic-constraints`
- `03-reflection`
- `04-advanced-interface-design`
- `05-memory-layout-and-escape-analysis`
- `06-scheduler-and-runtime-concepts`
- `07-memory-model`
- `08-profiling-and-optimization`
- `09-build-tags-and-cross-compilation`
- `10-workspaces-and-larger-codebase-structure`
- `11-unsafe-package`
- `12-cgo-overview`
- `13-advanced-testing-strategy`
- `14-practical-mini-projects`

The advanced mini-project module currently includes:

- `01-generic-safe-cache`
- `02-context-aware-runner`
- `03-in-memory-status-service`

All planned phases are now implemented.
