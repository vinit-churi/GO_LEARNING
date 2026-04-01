# Content Specification

This file defines the required lesson format and naming rules.

## Lesson Folder Convention

Use numeric prefixes to make the path self-ordering:

```text
lessons/
  01-beginner/
    01-setup-and-first-program/
      01-hello-world/
        README.md
        REFERENCE.md
        ANSWERS.md
        starter.go
        solution.go
        lesson_test.go
```

## Why Not Reuse The Existing Flat Format Exactly

The current repo uses a minimal `main.go` plus `main_test.go` layout. That is fine for small isolated drills, but it is too thin for a full learning system with explanations, multiple exercises, and answer separation.

The new curriculum should keep compatibility with Go tooling while adding the learning materials needed for structured study.

## Required Files

### `README.md`

Must include:

- lesson title
- difficulty
- prerequisites
- learning objectives
- concept summary
- run instructions
- test instructions
- at least 3 exercises
- problem statements for each exercise

### `REFERENCE.md`

Must include:

- concept explanations
- syntax patterns
- common mistakes
- why the feature exists
- when to use and when not to use it
- links or pointers to official Go documentation when appropriate

### `ANSWERS.md`

Must include:

- a short explanation for each exercise
- hints before full answers when useful
- discussion of tradeoffs or alternate approaches
- references to the relevant implementation in `solution.go`

`ANSWERS.md` should not replace the solution file. The code answer should live in `solution.go`; `ANSWERS.md` explains it.

### `starter.go`

Learner-facing scaffold. It may contain:

- TODOs
- stubs
- partial implementations
- comments that direct the exercise

### `solution.go`

Reference implementation. It should be clear and idiomatic, not code-golfed.

### `lesson_test.go`

Should:

- validate the exercises
- use clear test names
- prefer table-driven tests where they help readability
- remain maintainable for learners reading them

## Optional Files

Only add when the lesson benefits from them:

- `example.go`
- `example_test.go`
- `data/`
- `cmd/`
- `internal/`

## Test and Run Model

- Default validation path is `go test`.
- Some lessons should also include `go run` entry examples when execution flow matters.
- If a lesson needs both learner code and solution code to coexist cleanly, use packages or build tags carefully so test behavior remains clear.

## Authoring Standards

- Prefer small packages over giant files.
- Use ASCII unless the content requires otherwise.
- Keep terminology consistent across lessons.
- Avoid hidden requirements in exercises.
- Exercise prompts should state expected inputs, outputs, and constraints.
- Assume the learner is progressing linearly unless the lesson says otherwise.

## Recommendation For Coexisting With Existing Repo Content

Do not rewrite the old lessons first. Build the new staged curriculum beside the current content. Migration or consolidation can happen later after enough new lessons exist.
