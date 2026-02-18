# Deprecations

## Context

This file tracks active and planned deprecations with expected removal windows.

## Memory Legacy Plugin

- Component: legacy semantic memory plugin workflow
- Legacy surfaces:
  - `openkit init --memory`
  - `openkit opencode sync --memory`
  - `openkit memory` legacy command group
  - `.opencode/plugins/semantic-memory/`
  - `.opencode/memory/`
- Release N: Deprecated (warnings enabled)
- Release N+1: Removed
- Migration guide: [[MEMORY_LEGACY_MIGRATION.md]]
- Removal plan: [[sprint/Sprint-07/LEGACY_MEMORY_REMOVAL_PLAN.md]]

## Policy

- Every deprecation must define scope, migration path, and removal release.
- Deprecation notices must be visible in CLI output for affected commands.
- Removal requires verification against acceptance criteria in linked requirement artifacts.

## Related

- [[MEMORY_LEGACY_MIGRATION.md]]
- [[requirements/memory-kernel-rust-cli/ACCEPTANCE_CRITERIA.md]]
- [[requirements/memory-kernel-rust-cli/DATA_CONTRACTS.md]]
- [[sprint/Sprint-07/TASKS.md]]
