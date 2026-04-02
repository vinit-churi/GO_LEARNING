# Go Study Guide Planning Docs

This folder is the control center for building a full beginner-to-advanced Go curriculum in this repository.

Start here:

1. Read [00-main-orchestrator.md](./00-main-orchestrator.md)
2. Check [06-progress-tracker.md](./06-progress-tracker.md)
3. Open the first subplan whose status is still `pending` or `in_progress`

Planning files:

- [00-main-orchestrator.md](./00-main-orchestrator.md): authoritative entrypoint for future work
- [01-curriculum-blueprint.md](./01-curriculum-blueprint.md): staged roadmap from beginner to advanced
- [02-content-spec.md](./02-content-spec.md): required folder layout, file conventions, and lesson format
- [03-phase-1-beginner-plan.md](./03-phase-1-beginner-plan.md): implementation plan for beginner lessons
- [04-phase-2-intermediate-plan.md](./04-phase-2-intermediate-plan.md): implementation plan for intermediate lessons
- [05-phase-3-advanced-plan.md](./05-phase-3-advanced-plan.md): implementation plan for advanced lessons
- [06-progress-tracker.md](./06-progress-tracker.md): current status and next actions

Rules for future contributors and LLMs:

- Keep [00-main-orchestrator.md](./00-main-orchestrator.md) in sync with real repo state.
- Update [06-progress-tracker.md](./06-progress-tracker.md) whenever a phase or milestone changes.
- Do not mark a lesson complete unless markdown, starter code, solution code, and tests all exist.
- Prefer small, coherent batches of lessons over partially built coverage across many topics.

## Reading Lesson Structure Correctly

- A folder like `lessons/02-intermediate/12-goroutines/01-concurrent-workers` is **one lesson unit** inside the module, not one exercise.
- Each lesson unit is expected to include multiple exercises in its `README.md` (content spec requires at least 3).
- If you want deeper practice, the curriculum can be expanded by adding additional lesson units per module (for example `02-*`, `03-*`) while still keeping 3+ exercises per unit.

Current implementation state:

- Beginner phase is complete across modules `01` through `14`.
- Intermediate phase is complete across modules `01` through `18`.
- Advanced phase is complete across modules `01` through `14`, including multiple mini-project lessons.
- `go test ./...` passes with the staged beginner, intermediate, and advanced lessons included.
