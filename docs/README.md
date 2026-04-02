# Go Study Guide Planning Docs

This folder is the control center for building a full beginner-to-advanced Go curriculum in this repository.

Start here:

1. Read [00-main-orchestrator.md](/Users/vinitchuri/projects/go_learning/docs/00-main-orchestrator.md)
2. Check [06-progress-tracker.md](/Users/vinitchuri/projects/go_learning/docs/06-progress-tracker.md)
3. Open the first subplan whose status is still `pending` or `in_progress`

Planning files:

- [00-main-orchestrator.md](/Users/vinitchuri/projects/go_learning/docs/00-main-orchestrator.md): authoritative entrypoint for future work
- [01-curriculum-blueprint.md](/Users/vinitchuri/projects/go_learning/docs/01-curriculum-blueprint.md): staged roadmap from beginner to advanced
- [02-content-spec.md](/Users/vinitchuri/projects/go_learning/docs/02-content-spec.md): required folder layout, file conventions, and lesson format
- [03-phase-1-beginner-plan.md](/Users/vinitchuri/projects/go_learning/docs/03-phase-1-beginner-plan.md): implementation plan for beginner lessons
- [04-phase-2-intermediate-plan.md](/Users/vinitchuri/projects/go_learning/docs/04-phase-2-intermediate-plan.md): implementation plan for intermediate lessons
- [05-phase-3-advanced-plan.md](/Users/vinitchuri/projects/go_learning/docs/05-phase-3-advanced-plan.md): implementation plan for advanced lessons
- [06-progress-tracker.md](/Users/vinitchuri/projects/go_learning/docs/06-progress-tracker.md): current status and next actions

Rules for future contributors and LLMs:

- Keep [00-main-orchestrator.md](/Users/vinitchuri/projects/go_learning/docs/00-main-orchestrator.md) in sync with real repo state.
- Update [06-progress-tracker.md](/Users/vinitchuri/projects/go_learning/docs/06-progress-tracker.md) whenever a phase or milestone changes.
- Do not mark a lesson complete unless markdown, starter code, solution code, and tests all exist.
- Prefer small, coherent batches of lessons over partially built coverage across many topics.

Current implementation state:

- Beginner phase is complete across modules `01` through `14`.
- Intermediate phase is complete across modules `01` through `18`.
- Advanced phase is complete across modules `01` through `14`, including multiple mini-project lessons.
- `go test ./...` passes with the staged beginner, intermediate, and advanced lessons included.
