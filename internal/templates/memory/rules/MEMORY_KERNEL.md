---
trigger: always_on
agent: opencode
---

# MEMORY_KERNEL

OpenKit memory is docs-first and runtime-native.

## Supported Commands

- `openkit memory init` (repair only)
- `openkit memory doctor --json`
- `openkit memory capture --summary "..." --action "..."`
- `openkit memory review --json`

## Required Artifacts

- `.openkit/memory/config.yaml`
- `.openkit/memory/derivation.yaml`
- `.openkit/ops/queue.yaml`
- `.openkit/ops/sessions/*.json`

## Policy

- Use docs (`openkit-memory/**`) as the primary memory graph.
- Keep `## Related` and wikilinks consistent.
- Do not introduce plugin-only memory tools in active flows.

## Related

- [[HUB-DOCS.md]]
- [[MEMORY_LEGACY_MIGRATION.md]]
- [[DEPRECATIONS.md]]
