# Implementation Plan: Sync Engine + managed.json

## Scope

Implement project content sync for multiple agent targets using `.openkit/managed.json` as the managed state.

## CLI Commands

- `openkit <agent> sync`
- `openkit <agent> upgrade`
- `openkit <agent> doctor`

Where `<agent>` is one of: `opencode|claude|gemini|codex`.

## State File

- Create/update `.openkit/managed.json` per `cli/docs/content-protocol/MANAGED_STATE_SCHEMA.md`.
- (Optional but recommended) If a legacy state file is detected, migrate into `.openkit/managed.json` once and report what happened.

## Conflict Rules

Follow `cli/docs/content-protocol/SYNC_SEMANTICS.md`.

## Verification

- Unit tests for: hashing, path safety, drift detection, plan determinism.
- Integration tests for: sync then upgrade with a modified file -> conflict bucket.

## Related

- [[requirements/openkit-sync/HUB-OPENKIT-SYNC.md]]
- [[content-protocol/PROTOCOL.md]]
- [[content-protocol/MANAGED_STATE_SCHEMA.md]]
- [[content-protocol/SYNC_SEMANTICS.md]]
