# Sync Semantics (Install / Upgrade)

This document defines how OpenKit syncs a canonical content pack into a project for a specific agent target.

## Target Selection

Inputs:
- `agent`: `opencode | claude | gemini | codex`
- `pack id`: default pack for that agent, or user-specified
- `pack version`: default `latest` (highest compatible SemVer)
- `pack url` (optional): explicit source override

Selection rules:
1) Determine agent target.
2) Resolve pack source:
   - If `--pack-url` provided, use it.
   - Else use configured registry/default.
3) Resolve version:
   - If `--pack-version` provided, use it.
   - Else choose latest compatible with `compat` constraints.
4) Download/extract to local cache.

## Local Cache

Recommended cache structure (implementation-defined, but stable):
- `~/.cache/openkit/packs/<pack-id>/<version>/...` (or OS equivalent)

Cache MUST be content-addressed or checksum-validated to prevent tampering.

## Managed Files Tracking

Projects SHOULD store sync state in:
- `.openkit/managed.json`

Recommended contents:
- Pack identity/version used for the last sync per agent
- For each managed file:
  - `output_path`
  - `artifact_id`
  - `installed_sha256` (bytes written)
  - `installed_at` (timestamp)
  - `mode` (`copy|render|template`)

The state file is the source of truth for determining what the tool is allowed to update/delete.

## Conflict Detection

Given a planned write to `output_path`, classify:

1) Unmanaged existing file
- Condition: file exists AND not present in `.openkit/managed.json`.
- Result: `CONFLICT_UNMANAGED_EXISTS`.

2) User-edited managed file (checksum drift)
- Condition: file exists AND is managed AND current file hash != `installed_sha256`.
- Result: `CONFLICT_MANAGED_DRIFT`.

3) Missing managed file
- Condition: managed entry exists but file is missing.
- Result: `MISSING_MANAGED_FILE`.

4) No conflict
- Condition: file absent OR managed and unchanged.
- Result: safe to create/update.

## Resolution Rules

Default behavior (no flags):
- Create missing files.
- Update managed-and-unchanged files.
- Skip and report all conflicts.
- Never delete unmanaged files.

`--overwrite` behavior:
- Overwrite unmanaged existing files that are in the target set.
- Overwrite managed files even if drifted.
- Still never write outside project root.

Safe deletes (only when explicitly implied by pack changes):
- A file may be deleted only if:
  - It is managed, and
  - It is no longer present in the new target plan (or is deprecated with `remove_after` satisfied), and
  - Its current hash matches the last installed hash (i.e., not user-edited).
- Drifted managed files are never auto-deleted.

Deprecations and renames:
- If `migrations.renames` maps `from_output_path -> to_output_path`, sync SHOULD:
  - Move the file if it is managed and unchanged.
  - Otherwise, create `to_output_path` and report `from_output_path` as a manual migration.

## Dry Run Output

`--dry-run` MUST make no filesystem changes and SHOULD output a deterministic plan.

Recommended per-file plan format:
- `ACTION` (`create|update|skip|conflict|delete|move`)
- `PATH`
- `REASON` (e.g., `managed-unchanged`, `unmanaged-exists`, `checksum-drift`)
- `ARTIFACT_ID`

## CLI UX Requirements

Commands (agent is the target):

1) Upgrade (resolve newer compatible version and apply)
- `openkit <agent> upgrade [--dry-run] [--overwrite] [--pack-version <semver>] [--pack-url <url>]`

2) Sync (apply a specific resolved pack version; idempotent)
- `openkit <agent> sync [--dry-run] [--overwrite] [--pack-version <semver>] [--pack-url <url>]`

3) Doctor (detect trust issues, missing entrypoints, drift, and compatibility)
- `openkit <agent> doctor [--dry-run] [--pack-version <semver>] [--pack-url <url>]`

Doctor checks SHOULD include:
- Agent entrypoint presence (e.g., `opencode.json`, `CLAUDE.md`, `GEMINI.md`, `AGENTS.md`)
- Trust-gated features (Gemini trust mode; Codex trusted-project config)
- `.openkit/managed.json` validity and checksum drift summary
- Orphaned managed entries (state references missing files)

## Related

- [[content-protocol/README.md]]
- [[content-protocol/PROTOCOL.md]]
- [[content-protocol/MANAGED_STATE_SCHEMA.md]]
