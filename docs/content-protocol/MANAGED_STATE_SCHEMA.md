# Managed State Schema (.openkit/managed.json)

This document defines the JSON structure stored at `.openkit/managed.json` in a project.

## Goals

- Determine what files are managed by OpenKit
- Detect drift (user edits) via checksum
- Support safe upgrades and pack version tracking

## Minimal Shape

```json
{
  "schema_version": "1",
  "agents": {
    "opencode": {
      "pack": {"id": "openkit-core", "version": "1.2.3"},
      "files": {
        ".opencode/rules/MASTER.md": {
          "artifact_id": "rules.master",
          "installed_sha256": "0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef",
          "installed_at": "2026-02-08T00:00:00Z",
          "mode": "copy"
        }
      }
    }
  }
}
```

## Field Definition (Normative)

- `schema_version` (string): currently `"1"`.
- `agents` (object): keyed by agent id (`opencode|claude|gemini|codex`).
  - `pack` (object): last synced pack identity.
    - `id` (string)
    - `version` (string)
  - `files` (object): map of `output_path -> entry`.
    - `output_path` keys are relative to project root and MUST NOT contain `..`.
    - entry:
      - `artifact_id` (string)
      - `installed_sha256` (string): hex sha256 of bytes written
      - `installed_at` (string): RFC3339 timestamp
      - `mode` (string enum): `copy|render|template`

## Notes

- The CLI MUST treat this file as the sole authority for safe deletes/moves.
- Unmanaged files are never deleted.

## Related

- [[content-protocol/README.md]]
- [[content-protocol/PROTOCOL.md]]
- [[content-protocol/SYNC_SEMANTICS.md]]
