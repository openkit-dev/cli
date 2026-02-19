# Migration Guide

## Scope

This guide covers migration from legacy OpenKit flows to the current Rust runtime and docs-first Memory Kernel.

## Current Runtime Baseline

- Runtime: Rust-only CLI (`rust-cli/src/main.rs`)
- Release assets: `openkit_<OS>_<ARCH>` + `checksums.txt`
- Memory model: docs-first (`docs/` + `.openkit/memory/` + `.openkit/ops/`)

## Removed Legacy Surfaces

- `openkit init --memory`
- `openkit opencode sync --memory`
- plugin directory `.opencode/plugins/semantic-memory/`
- plugin-era commands (`openkit memory list/search/stats/prune/export/config/debug`)

## Supported Commands

- `openkit --version`
- `openkit check`
- `openkit init`
- `openkit upgrade`
- `openkit uninstall`
- `openkit memory init|doctor|capture|review`

## Migration Steps

1. Update to latest OpenKit release.
2. Initialize/refresh project scaffold with `openkit init` if needed.
3. Initialize memory kernel artifacts with `openkit memory init`.
4. Validate docs graph and memory health with `openkit memory doctor --json`.
5. Replace any legacy memory workflow references in local docs/scripts.

## Verification

- CLI command surface matches `docs/API.md`.
- `openkit check` succeeds.
- `openkit memory doctor --json` reports healthy status.

## Related

- [[API.md]]
- [[DEPRECATIONS.md]]
- [[MEMORY_LEGACY_MIGRATION.md]]
- [[audit/HUB-AUDIT.md]]
