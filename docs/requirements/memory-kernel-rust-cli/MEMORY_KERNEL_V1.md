# Memory Kernel v1 Specification

## Context

This document defines OpenKit Memory Kernel v1 primitives and their validation model for docs-first permanent memory.

Kernel scope and rollout are linked to [[requirements/memory-kernel-rust-cli/PLAN.md]] and execution in [[sprint/Sprint-07/TASKS.md]].

## Primitives

### 1) Repository Memory Source of Record

- Rule: `docs/` is the canonical durable memory surface.
- Validation: required hubs and canonical files must exist and be link-valid.

### 2) Entry Map Brevity

- Rule: top-level entry docs remain navigational and concise.
- Validation: hubs point to details rather than duplicating full operational manuals.

### 3) Inline Semantic Linking

- Rule: first relevant mention of internal artifacts must use wikilinks inline.
- Validation: memory doctor checks inline links in docs content sections.

### 4) Related Index Linking

- Rule: each canonical doc must include `## Related` with stable links.
- Validation: memory doctor validates presence on required files.

### 5) Documentation Contract Schema

- Rule: requirements and sprint artifacts follow canonical filenames and consistent sections.
- Validation: lint and review checks against artifact set in `docs/requirements/` and `docs/sprint/`.

### 6) Three-Space Separation

- Rule: durable knowledge in `docs/`; operational logs in `.openkit/ops/`; memory config in `.openkit/memory/`.
- Validation: command contracts generate expected paths and avoid legacy plugin paths.

### 7) Decision Traceability

- Rule: implementation and governance changes map to requirement and sprint artifacts.
- Validation: sprint tasks reference requirement artifacts with verifiable outputs.

### 8) Session Rhythm

- Rule: session-level context must be capturable as versioned snapshots.
- Validation: `memory capture` writes versioned JSON snapshots under `.openkit/ops/sessions/`.

### 9) Operational Learning Loop

- Rule: observations and tensions must be reviewable for remediation planning.
- Validation: `memory review` summarizes counts and deterministic recommendations.

### 10) Memory Health Gates

- Rule: memory quality contributes to delivery safety.
- Validation: `memory doctor` outputs health score/status and fails on broken wikilinks.

## Validation Matrix

| Primitive | Command/Check | Evidence |
|---|---|---|
| Source of record | `memory doctor` + docs lint | `docs/` hubs and link integrity |
| Inline + Related links | `memory doctor` | `inline_links`, `related_sections` checks |
| Three-space separation | `memory init` | `.openkit/memory/` and `.openkit/ops/` contracts |
| Session rhythm | `memory capture` | versioned session JSON snapshot |
| Learning loop | `memory review` | deterministic recommendations |
| Health gates | `memory doctor` | structured report + broken link failure |

## Related

- [[requirements/memory-kernel-rust-cli/PLAN.md]]
- [[requirements/memory-kernel-rust-cli/ACCEPTANCE_CRITERIA.md]]
- [[requirements/memory-kernel-rust-cli/DATA_CONTRACTS.md]]
- [[sprint/Sprint-07/TASKS.md]]
