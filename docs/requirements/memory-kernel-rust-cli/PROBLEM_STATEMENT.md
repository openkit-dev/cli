# Problem Statement: Memory Kernel + Rust CLI Migration

## Problem

OpenKit currently delivers value with strong command orchestration, but persistent project memory is still fragmented across operational and planning files. At the same time, the current Go CLI architecture is stable but not optimized for the next phase of plugin-oriented extensibility and long-term multi-agent adapters.

The repository also still contains a legacy semantic memory plugin workflow (`--memory`, `openkit memory`, `.opencode/plugins/semantic-memory`) that conflicts with the new docs-first permanent memory direction. Maintaining both models in parallel increases user confusion, operational overhead, and migration risk.

The strategic goal discussed for this cycle is to turn OpenKit into a robust development CLI where project memory is first-class, with docs as source-of-truth and operational state clearly separated. This direction requires explicit kernel rules in documentation and deterministic validation, while also preparing the CLI foundation for Rust.

## Why Now

- The roadmap now prioritizes durable memory and docs governance as product differentiation.
- The team decided to adopt OpenCode-first depth while preserving multi-agent adapters.
- A deep architecture refactor is already planned, reducing incremental migration cost.

## Desired Outcome

1. Define and enforce OpenKit Memory Kernel v1 across `docs/` and `.openkit/`.
2. Establish inline wikilinking plus `## Related` indexing as documentation graph standard.
3. Start a safe, non-big-bang Rust migration with compatibility guarantees.
4. Sunset and remove legacy semantic memory plugin flows with a controlled deprecation path.
5. Keep delivery predictable via staged sprint execution in [[sprint/Sprint-07/TASKS.md]].

## Non-Goals (This Cycle)

- Full Go codebase replacement in one release.
- Full feature parity for every Tier-3 adapter.
- Immediate removal of existing sync/state mechanisms before parity checks.

## Legacy Sunset Scope

- Deprecation release (N): keep legacy commands with warning banners and migration guidance.
- Removal release (N+1): remove legacy plugin install/sync pathways and legacy `openkit memory` behavior.
- Migration guidance must be explicit from old artifacts to new docs-first memory model.

## Related

- [[requirements/memory-kernel-rust-cli/USER_STORIES.md]]
- [[requirements/memory-kernel-rust-cli/PLAN.md]]
- [[requirements/memory-kernel-rust-cli/RISKS.md]]
- [[requirements/memory-kernel-rust-cli/DATA_CONTRACTS.md]]
- [[requirements/memory-kernel-rust-cli/HUB-MEMORY-KERNEL-RUST-CLI.md]]
