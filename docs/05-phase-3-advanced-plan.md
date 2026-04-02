# Phase 3 Plan: Advanced

Status: `completed`

## Goal

Teach advanced Go features and internals responsibly, with a strong emphasis on tradeoffs and real-world use.

## Scope

This phase should cover:

1. generics fundamentals
2. generic constraints
3. reflection
4. advanced interface design
5. memory layout and escape analysis
6. scheduler and runtime concepts
7. memory model
8. profiling and optimization
9. build tags and cross-compilation
10. workspaces and larger codebase structure
11. unsafe package
12. cgo overview
13. advanced testing strategy
14. practical mini-projects

## Guardrails

- Advanced topics must explain why not to overuse them.
- Unsafe or low-level topics should include safety warnings.
- Performance claims should be measurable, not vague.
- Runtime and memory topics should prefer concrete examples over abstract theory alone.

## Completion Criteria

Phase 3 is complete when the advanced path exists end-to-end and the repo contains at least a few mini-projects tying together multiple prior topics.

## Practice Depth Notes

- The current advanced path also starts with one primary lesson unit (`01-*`) per module, with mini-projects as additional integrated practice.
- Each primary unit should provide multiple exercises (minimum 3 in each lesson `README.md`, per content spec).
- To practice advanced concepts more thoroughly, add extra lesson units (`02-*`, `03-*`) for high-leverage modules like generics, reflection, profiling, memory model, and testing strategy.

Implemented at:

- `lessons/03-advanced/01-generics-fundamentals`
- `lessons/03-advanced/02-generic-constraints`
- `lessons/03-advanced/03-reflection`
- `lessons/03-advanced/04-advanced-interface-design`
- `lessons/03-advanced/05-memory-layout-and-escape-analysis`
- `lessons/03-advanced/06-scheduler-and-runtime-concepts`
- `lessons/03-advanced/07-memory-model`
- `lessons/03-advanced/08-profiling-and-optimization`
- `lessons/03-advanced/09-build-tags-and-cross-compilation`
- `lessons/03-advanced/10-workspaces-and-larger-codebase-structure`
- `lessons/03-advanced/11-unsafe-package`
- `lessons/03-advanced/12-cgo-overview`
- `lessons/03-advanced/13-advanced-testing-strategy`
- `lessons/03-advanced/14-practical-mini-projects`

Mini-project lessons included:

- `lessons/03-advanced/14-practical-mini-projects/01-generic-safe-cache`
- `lessons/03-advanced/14-practical-mini-projects/02-context-aware-runner`
- `lessons/03-advanced/14-practical-mini-projects/03-in-memory-status-service`
