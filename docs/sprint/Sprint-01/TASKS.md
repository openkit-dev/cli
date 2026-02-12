# Sprint 01 Tasks

## Task 1: Managed State (read/write)

INPUT: project root, planned outputs
OUTPUT: `.openkit/managed.json` updated with installed entries
VERIFY: unit tests for deterministic JSON + sha256 drift detection

## Task 2: Sync Planner

INPUT: pack manifest + targets + current filesystem + managed.json
OUTPUT: deterministic plan buckets (create/update/skip/conflict/delete/move)
VERIFY: unit tests for each conflict class

## Task 3: Apply Engine

INPUT: plan + overwrite flags
OUTPUT: filesystem changes + backups (when overwriting/removing)
VERIFY: integration test writes backups and preserves unmanaged files

## Related

- [[sprint/Sprint-01/HUB-SPRINT-01.md]]
- [[sprint/Sprint-01/SPRINT_GOAL.md]]
- [[sprint/Sprint-01/BACKLOG.md]]
- [[sprint/Sprint-01/RISK_REGISTER.md]]
