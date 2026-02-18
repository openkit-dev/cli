# TASKS

**Sprint**: Sprint-07
**Title**: Memory Kernel v1 + Rust Migration Bootstrap
**Status**: Completed
**Total Tasks**: 15

## Task Breakdown

### P0: Foundation

#### Task 1: Publish Memory Kernel v1 Specification

**ID**: S07-T001
**Story**: S07-001
**Priority**: P0

**INPUT**:
- [[requirements/memory-kernel-rust-cli/PROBLEM_STATEMENT.md]]
- [[requirements/memory-kernel-rust-cli/ACCEPTANCE_CRITERIA.md]]

**OUTPUT**:
- Kernel specification doc with primitives and validation strategy

**VERIFY**:
- [x] Primitive list is explicit and versioned
- [x] Validation rules map to each primitive

**Execution Note**:
- Kernel primitives and validation matrix documented in [[requirements/memory-kernel-rust-cli/MEMORY_KERNEL_V1.md]].

---

#### Task 2: Define Docs Graph Policy (Inline + Related)

**ID**: S07-T002
**Story**: S07-002
**Priority**: P0

**INPUT**:
- [[requirements/memory-kernel-rust-cli/PLAN.md]]
- [[requirements/memory-kernel-rust-cli/DATA_CONTRACTS.md]]

**OUTPUT**:
- Policy doc and checklist for inline wikilinks + `## Related`

**VERIFY**:
- [x] Policy explains first-relevant-mention inline linking
- [x] Policy keeps `## Related` as mandatory index

**Execution Note**:
- Policy published in [[requirements/memory-kernel-rust-cli/DOCS_GRAPH_POLICY.md]].

---

#### Task 3: Define Memory Doctor Checks

**ID**: S07-T003
**Story**: S07-001
**Priority**: P0

**INPUT**:
- [[requirements/memory-kernel-rust-cli/DATA_CONTRACTS.md]]
- [[requirements/memory-kernel-rust-cli/ACCEPTANCE_CRITERIA.md]]

**OUTPUT**:
- Check list and score formula for memory health

**VERIFY**:
- [x] Score thresholds are documented
- [x] Checks include links, related, stale docs, traceability

**Execution Note**:
- Score thresholds and doctor checks are documented in [[requirements/memory-kernel-rust-cli/DATA_CONTRACTS.md]] and implemented in `rust-cli/src/main.rs`.

---

### P1: Rust Bootstrap

#### Task 4: Initialize Rust CLI Workspace

**ID**: S07-T004
**Story**: S07-003
**Priority**: P1

**INPUT**:
- [[requirements/memory-kernel-rust-cli/PLAN.md]]

**OUTPUT**:
- Rust workspace scaffold and module boundaries for CLI core

**VERIFY**:
- [x] Workspace builds successfully
- [x] Command contract module exists

---

#### Task 5: Implement `memory init` Command

**ID**: S07-T005
**Story**: S07-003
**Priority**: P1

**INPUT**:
- [[requirements/memory-kernel-rust-cli/DATA_CONTRACTS.md]]

**OUTPUT**:
- Rust command creating `.openkit/memory/` and `.openkit/ops/` baseline

**VERIFY**:
- [x] Creates contract-compliant config files
- [x] Is idempotent on second run

---

#### Task 6: Implement `memory doctor` Command

**ID**: S07-T006
**Story**: S07-004
**Priority**: P1

**INPUT**:
- [[requirements/memory-kernel-rust-cli/ACCEPTANCE_CRITERIA.md]]
- [[requirements/memory-kernel-rust-cli/DATA_CONTRACTS.md]]

**OUTPUT**:
- Rust command producing memory health status and score

**VERIFY**:
- [x] Emits structured output matching doctor contract
- [x] Fails with actionable diagnostics when links are broken

---

#### Task 7: Implement `memory capture` Command

**ID**: S07-T007
**Story**: S07-005
**Priority**: P1

**INPUT**:
- [[requirements/memory-kernel-rust-cli/DATA_CONTRACTS.md]]

**OUTPUT**:
- Session snapshot writer for `.openkit/ops/sessions/`

**VERIFY**:
- [x] Writes versioned JSON snapshot
- [x] Includes summary and action list

---

#### Task 8: Implement `memory review` Command

**ID**: S07-T008
**Story**: S07-005
**Priority**: P1

**INPUT**:
- `.openkit/ops/observations/`
- `.openkit/ops/tensions/`

**OUTPUT**:
- Review summary and queue suggestions for next sprint

**VERIFY**:
- [x] Aggregates pending observations and tensions
- [x] Produces deterministic recommendation list

**Execution Note**:
- Rust scaffold and command implementations were added under `rust-cli/`.
- Runtime verification executed with `cargo check`, `cargo test`, `memory init`, `memory doctor`, `memory capture`, and `memory review`.

---

### P2: Migration Safety and Governance

#### Task 9: Add Golden Compatibility Tests

**ID**: S07-T009
**Story**: S07-004
**Priority**: P2

**INPUT**:
- Current CLI behavior reference
- Rust command behavior

**OUTPUT**:
- Golden test fixtures and parity checks

**VERIFY**:
- [x] Golden tests run in CI
- [x] Command output contract differences are reported clearly

**Execution Note**:
- Golden contract tests implemented in `rust-cli/tests/command_contracts.rs` with fixtures in `rust-cli/tests/golden/`.
- CI updated in `.github/workflows/ci.yml` to run `cargo test --manifest-path rust-cli/Cargo.toml`.

---

#### Task 10: Publish Tier Support Matrix

**ID**: S07-T010
**Story**: S07-005
**Priority**: P2

**INPUT**:
- Product decisions documented in [[requirements/memory-kernel-rust-cli/PLAN.md]]

**OUTPUT**:
- Support matrix with Tier 1/2/3 scope and SLA (`[[SUPPORT_MATRIX.md]]`)

**VERIFY**:
- [x] OpenCode listed as Tier 1
- [x] Claude Code, Codex, Antigravity listed as Tier 2

**Execution Note**:
- Official support policy published in `docs/SUPPORT_MATRIX.md`.
- Matrix includes capability scope, SLA, and promotion/demotion criteria.

---

#### Task 11: Link Policy Retrofit for New Artifacts

**ID**: S07-T011
**Story**: S07-002
**Priority**: P2

**INPUT**:
- New Sprint-07 and requirement files

**OUTPUT**:
- Inline links and `## Related` sections normalized

**VERIFY**:
- [x] Every new artifact includes inline references
- [x] Every new artifact includes `## Related`

**Execution Note**:
- Sprint-07 and related requirement artifacts were normalized with inline wikilinks and `## Related` sections.

---

#### Task 12: Sprint Exit Verification

**ID**: S07-T012
**Story**: S07-004
**Priority**: P2

**INPUT**:
- All Sprint-07 outputs

**OUTPUT**:
- Exit report with acceptance criteria coverage

**VERIFY**:
- [x] Each acceptance criterion is mapped to completed tasks
- [x] Open risks are explicitly documented

**Execution Note**:
- Sprint exit mapping and risk disposition are documented in [[sprint/Sprint-07/EXIT_REPORT.md]].

---

### P2: Legacy Sunset (Release N -> N+1)

#### Task 13: Add Legacy Deprecation Warnings

**ID**: S07-T013
**Story**: S07-006
**Priority**: P2

**INPUT**:
- Legacy memory pathways in CLI (`--memory`, `openkit memory`)
- [[requirements/memory-kernel-rust-cli/ACCEPTANCE_CRITERIA.md]]

**OUTPUT**:
- Deprecation warnings with migration guidance in legacy command/flag surfaces

**VERIFY**:
- [x] Warnings appear consistently in release N pathways
- [x] Warnings reference docs-first migration target

---

#### Task 14: Publish Legacy-to-Kernel Migration Guide

**ID**: S07-T014
**Story**: S07-006
**Priority**: P2

**INPUT**:
- Legacy plugin path inventory
- [[requirements/memory-kernel-rust-cli/DATA_CONTRACTS.md]]

**OUTPUT**:
- Migration guide and mapping table for old -> new artifacts

**VERIFY**:
- [x] Includes path mapping from `.opencode/plugins/semantic-memory` and `.opencode/memory`
- [x] Includes rollback notes for interrupted migration

---

#### Task 15: Remove Legacy Memory Paths (N+1 Gate)

**ID**: S07-T015
**Story**: S07-006
**Priority**: P2

**INPUT**:
- Completion of S07-T013 and S07-T014
- [[requirements/memory-kernel-rust-cli/PLAN.md]]

**OUTPUT**:
- Removal patch plan for legacy plugin install/sync/command code paths (`[[sprint/Sprint-07/LEGACY_MEMORY_REMOVAL_PLAN.md]]`)

**VERIFY**:
- [x] Removal gated by migration criteria
- [x] No orphan references remain in docs and command help
- [x] Compatibility tests pass for new memory command set

**Execution Note**:
- Active command help and main docs were updated to remove legacy memory command surfaces.
- Historical sprint docs remain as archived historical records.

## Related

- [[sprint/Sprint-07/HUB-SPRINT-07.md]]
- [[sprint/Sprint-07/BACKLOG.md]]
- [[sprint/Sprint-07/RISK_REGISTER.md]]
- [[sprint/Sprint-07/LEGACY_MEMORY_REMOVAL_PLAN.md]]
- [[sprint/Sprint-07/EXIT_REPORT.md]]
- [[requirements/memory-kernel-rust-cli/ACCEPTANCE_CRITERIA.md]]
- [[requirements/memory-kernel-rust-cli/PLAN.md]]
