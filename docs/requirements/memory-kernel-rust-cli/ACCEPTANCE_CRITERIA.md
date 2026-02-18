# Acceptance Criteria: Memory Kernel + Rust CLI Migration

## Memory Kernel v1

- [x] Kernel primitive specification exists and is versioned in project docs.
- [x] `docs/` and `.openkit/ops/` responsibilities are explicitly documented.
- [x] Memory health checks are executable and produce actionable output.

## Documentation Graph Rules

- [x] New and updated docs include inline wikilinks for first relevant mention of internal artifacts.
- [x] New and updated docs include `## Related` with stable links.
- [x] Link linting reports no broken wikilinks in mandatory artifacts.

## Rust Migration Foundation

- [x] Rust CLI workspace is initialized with command contract definitions.
- [x] Compatibility (golden) tests exist for migrated commands.
- [x] Migration strategy is staged (no big-bang), with parity gates per phase.

## Legacy Memory Sunset

- [x] Release N prints deprecation warnings for legacy `--memory` and legacy `openkit memory` workflows.
- [x] Migration guide maps legacy plugin artifacts to new docs-first memory artifacts.
- [x] Release N+1 removes legacy semantic memory plugin install/sync code paths.
- [x] Release N+1 removes or repurposes legacy memory CLI surfaces without orphaning docs.

## Product/Tier Governance

- [x] Tier matrix is codified in docs with OpenCode as Tier 1 and Claude/Codex/Antigravity as Tier 2.
- [x] Adapter boundaries are documented as core-agnostic + adapter-specific layers.

## Delivery Controls

- [x] Sprint backlog and task artifacts map directly to this requirement set.
- [x] Risks and mitigations are tracked and linked to execution tasks.
- [x] Legacy sunset tasks are tracked with explicit rollback criteria.

## Related

- [[requirements/memory-kernel-rust-cli/PLAN.md]]
- [[requirements/memory-kernel-rust-cli/RISKS.md]]
- [[sprint/Sprint-07/TASKS.md]]
- [[sprint/Sprint-07/RISK_REGISTER.md]]
- [[requirements/memory-kernel-rust-cli/HUB-MEMORY-KERNEL-RUST-CLI.md]]
