# ADR 0001: Managed State File (.openkit/managed.json)

**Status:** Accepted
**Date:** 2026-02-08

## Context

OpenKit needs a reliable, tool-owned source of truth for which files it is allowed to update, move, or delete during sync/upgrade.

Historically, Node-based OpenKit tracked baseline state in a legacy file under `.opencode/`.

## Decision

The CLI will use a repository-agnostic managed state file:

- Path: `.openkit/managed.json` (relative to project root)
- Purpose: track installed pack identity and per-file install metadata
- Ownership: OpenKit only (user edits are permitted but treated as potentially unsafe)

The CLI SHOULD provide a one-time migration path from the legacy state file (if present) into `.openkit/managed.json`.

## Consequences

- Sync/upgrade can be safe-by-default (skip conflicts, never touch unmanaged files).
- Agent targets (.opencode/.claude/.gemini/.codex) remain implementation details; state tracking is consistent.
- The legacy file location is deprecated.

## Related

- [[docs/adr/HUB-ADR.md]]
- [[docs/content-protocol/MANAGED_STATE_SCHEMA.md]]
- [[docs/requirements/openkit-sync/PLAN.md]]
