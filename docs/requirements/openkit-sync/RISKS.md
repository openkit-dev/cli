# Risks: Sync Engine + managed.json

- Incorrect path normalization could write outside project root.
- Mistaken ownership of user-local files (e.g. Claude settings) could break agent tooling.
- Weak checksum handling could miss drift or cause false positives.
- Renames/migrations could be destructive if applied to drifted files.

## Related

- [[requirements/openkit-sync/HUB-OPENKIT-SYNC.md]]
- [[requirements/openkit-sync/PROBLEM_STATEMENT.md]]
- [[requirements/openkit-sync/PLAN.md]]
- [[content-protocol/MANAGED_STATE_SCHEMA.md]]
