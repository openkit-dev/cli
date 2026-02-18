# Legacy Memory Migration Guide

## Context

OpenKit is replacing the legacy semantic memory plugin workflow (`--memory`, `openkit memory`, `.opencode/plugins/semantic-memory`) with a docs-first permanent memory model based on `docs/`, `.openkit/ops/`, and `.openkit/memory/`.

This guide defines a safe migration path with rollback instructions.

## When to Migrate

- Migrate during release N (deprecation window), before release N+1 removal.
- If your project still references `openkit init --memory`, `openkit opencode sync --memory`, or `openkit memory`, migrate now.

## Artifact Mapping (Old -> New)

| Legacy Artifact | New Artifact | Purpose |
|---|---|---|
| `.opencode/plugins/semantic-memory/` | `.openkit/memory/` + `.openkit/ops/` | Runtime plugin logic replaced by native memory state + operations |
| `.opencode/memory/config.json` | `.openkit/memory/config.yaml` | Memory configuration contract |
| `.opencode/memory/index.lance/` | `docs/` + `.openkit/ops/sessions/` + `.openkit/ops/observations/` | Persistent context moves to docs graph and operational snapshots |
| `.opencode/rules/SEMANTIC_MEMORY.md` | Memory Kernel docs policy + `docs/DEPRECATIONS.md` | Rules shift from plugin-specific to platform policy |
| `openkit memory stats/list/search/...` | `openkit memory doctor/capture/review` (new model) | Legacy plugin commands replaced by kernel lifecycle commands |

## Migration Steps

1. Snapshot current memory state.
   - Export legacy memory data if still needed.
   - Record active workflows depending on legacy commands.
2. Initialize new kernel structure.
   - Ensure `.openkit/memory/` and `.openkit/ops/` exist.
   - Ensure `docs/` hubs and requirement/sprint links are valid.
3. Move durable context into docs.
   - Convert important legacy insights into requirement, ADR, or sprint artifacts.
   - Add inline wikilinks at first relevant mention and keep `## Related` sections.
4. Record migration state.
   - Create `.openkit/memory/legacy-migration-report.json` using the contract in [[requirements/memory-kernel-rust-cli/DATA_CONTRACTS.md]].
5. Disable legacy operations.
   - Stop using `--memory` flags and `openkit memory` legacy subcommands.
   - Follow deprecation policy in [[DEPRECATIONS.md]].

## Verification Checklist

- [ ] Legacy paths inventoried and backed up.
- [ ] New memory directories exist and are versioned as needed.
- [ ] Critical memory content is represented in docs artifacts.
- [ ] Wikilink graph remains valid.
- [ ] Deprecation warnings no longer appear in daily workflows.

## Rollback (Interrupted Migration)

If migration is interrupted or new workflows fail:

1. Keep backup copies of `.opencode/plugins/semantic-memory/` and `.opencode/memory/`.
2. Restore legacy directories from backup.
3. Continue using legacy commands temporarily during release N only.
4. Register rollback event in `.openkit/memory/legacy-migration-report.json` with `status: "rolled_back"`.
5. Open a migration issue with affected commands and paths before retrying.

## Support Boundaries

- Release N: legacy flows are supported with warnings.
- Release N+1: legacy flows are removed; only docs-first memory model is supported.

## Related

- [[DEPRECATIONS.md]]
- [[requirements/memory-kernel-rust-cli/PROBLEM_STATEMENT.md]]
- [[requirements/memory-kernel-rust-cli/DATA_CONTRACTS.md]]
- [[requirements/memory-kernel-rust-cli/ACCEPTANCE_CRITERIA.md]]
- [[sprint/Sprint-07/TASKS.md]]
