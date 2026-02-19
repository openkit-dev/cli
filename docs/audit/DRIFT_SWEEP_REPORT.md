# DRIFT_SWEEP_REPORT

## Snapshot

- Date: 2026-02-18
- Scope: `docs/**` + `internal/templates/**`
- Focus patterns: legacy Go runtime paths, legacy memory plugin commands/flags, plugin-era tool references.

## Result Summary

### Internal templates

- `internal/templates/**`: 0 drift hits in active templates after cleanup.
- Deprecated memory plugin files removed and replaced by Memory Kernel resources.

### Active top-level docs

- Active top-level docs (excluding legacy migration/deprecation guides): 0 drift hits.
- `docs/BACKEND.md`, `docs/QUALITY_GATES.md`, `docs/SECURITY.md`, `docs/GLOSSARY.md` aligned with Rust runtime.

### Historical/intentional legacy references

- `docs/requirements/**`: 77 hits (historical decision/migration context).
- `docs/sprint/**`: 297 hits (execution history from pre-Rust and migration sprints).
- `docs/audit/**`: 22 hits (expected baseline evidence references).
- Legacy references in these areas are mostly intentional historical records.

## Completed Fixes In This Sweep

- Replaced deprecated plugin memory templates with docs-first Memory Kernel templates in `internal/templates/memory/**`.
- Removed plugin-era files (`index.ts`, `lib/*`, `scripts/*`, `rules/SEMANTIC_MEMORY.md`).
- Added replacement files (`config.yaml`, `derivation.yaml`, `queue.yaml`, `rules/MEMORY_KERNEL.md`, `README.md`).
- Fixed stale path/runtime assumptions in internal prompt/skill templates.
- Rewrote `docs/MIGRATION_GUIDE.md` to current command surface and model.
- Updated `docs/DEPRECATIONS.md` and `docs/MEMORY_LEGACY_MIGRATION.md` to post-removal state.

## Remaining Drift Backlog

1. Archive plugin-era sprint docs (`docs/sprint/Sprint-05/**`) into explicit legacy namespace.
2. Mark legacy requirements packs as historical where migration is complete.
3. Add CI drift gate to block reintroduction of banned patterns in active docs/templates.

## Verification

- `obsidian-link-lint` (root `docs`): 0 broken links.
- Drift scan on `internal/templates/**`: 0 active legacy-pattern hits.
- Drift scan on active top-level docs: 0 hits.

## Related

- [[audit/HUB-AUDIT.md]]
- [[audit/SWEEP_PLAN.md]]
- [[audit/INTERNAL_SWEEP_PLAN.md]]
- [[audit/INTERNAL_BASELINE.md]]
