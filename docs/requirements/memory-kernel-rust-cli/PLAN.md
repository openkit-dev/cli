# Implementation Plan: Memory Kernel + Rust CLI Migration

## Scope

Deliver a foundation release that defines OpenKit Memory Kernel v1 and starts Rust migration without disrupting active workflows. This plan depends on requirement constraints in [[requirements/memory-kernel-rust-cli/ACCEPTANCE_CRITERIA.md]] and execution control in [[sprint/Sprint-07/TASKS.md]].

## Architecture Decisions

1. Memory source-of-truth stays in `docs/` with explicit governance.
2. Operational state moves to `.openkit/ops/` (sessions, queue, observations, tensions).
3. Rust migration uses strangler strategy, preserving compatibility until parity gates pass.
4. Documentation graph policy enforces inline wikilinks plus `## Related` index.
5. Legacy semantic memory plugin flows are deprecated in release N and removed in release N+1.

## Execution Phases

### Phase A: Kernel Specification and Governance

- Define kernel primitives and validation strategy.
- Publish support matrix and adapter tier policy.
- Add docs policy for inline links and related sections.

### Phase B: Rust Foundation

- Create Rust CLI workspace and command contracts.
- Implement baseline commands for memory lifecycle (`init`, `doctor`, `capture`, `review`).
- Add golden tests comparing expected command output semantics.

### Phase C: Integration and Verification

- Integrate memory doctor checks into quality gates.
- Connect sprint artifacts and requirement artifacts to parity checklist.
- Validate that no mandatory workflows regress.

### Phase D: Legacy Sunset Execution

- Add warnings to legacy memory flags/commands during release N.
- Provide migration guide from legacy plugin paths to docs-first memory paths.
- Remove legacy plugin install/sync/command paths in release N+1 once criteria pass.

## Deliverables

- Full requirement artifact pack in `docs/requirements/memory-kernel-rust-cli/`.
- Sprint-07 planning pack in `docs/sprint/Sprint-07/`.
- Execution-ready tasks with INPUT -> OUTPUT -> VERIFY.
- Legacy sunset checklist with rollback and support notes.

## Verification Strategy

- Link integrity: run wikilink lint for docs graph.
- Scope integrity: ensure tasks map one-to-one with acceptance criteria.
- Migration integrity: pass contract tests for each migrated command before cutover.
- Sunset integrity: verify no orphan references to legacy memory plugin flows after removal phase.

## Related

- [[requirements/memory-kernel-rust-cli/PROBLEM_STATEMENT.md]]
- [[requirements/memory-kernel-rust-cli/DATA_CONTRACTS.md]]
- [[requirements/memory-kernel-rust-cli/RISKS.md]]
- [[requirements/memory-kernel-rust-cli/ACCEPTANCE_CRITERIA.md]]
- [[sprint/Sprint-07/TASKS.md]]
