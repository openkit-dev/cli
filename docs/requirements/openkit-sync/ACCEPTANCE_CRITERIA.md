# Acceptance Criteria: Sync Engine + managed.json

- Running `openkit opencode sync` creates/updates only managed target files and writes `.openkit/managed.json`.
- Default behavior skips conflicts and never overwrites unmanaged files.
- Drift detection is checksum-based using `installed_sha256` from `.openkit/managed.json`.
- `--dry-run` makes no filesystem changes and prints a deterministic plan.
- Safe deletes only remove managed-and-unchanged files that are no longer in the target plan.

## Related

- [[docs/requirements/openkit-sync/HUB-OPENKIT-SYNC.md]]
- [[docs/requirements/openkit-sync/PROBLEM_STATEMENT.md]]
- [[docs/requirements/openkit-sync/PLAN.md]]
- [[docs/content-protocol/SYNC_SEMANTICS.md]]
