# Data Contracts: Memory Kernel + Rust CLI Migration

## Contract 1: Memory Config

Path: `.openkit/memory/config.yaml`

```yaml
version: 1
mode: assisted # manual | assisted | auto
health_thresholds:
  healthy: 85
  warning: 70
linking:
  require_inline_links: true
  require_related_section: true
```

## Contract 2: Derivation State

Path: `.openkit/memory/derivation.yaml`

```yaml
version: 1
feature_slug: memory-kernel-rust-cli
decisions:
  runtime_language: rust
  migration_strategy: strangler
  tier_policy:
    tier1: [opencode]
    tier2: [claude-code, codex, antigravity]
```

## Contract 3: Operational Queue

Path: `.openkit/ops/queue.yaml`

```yaml
version: 1
items:
  - id: MK-001
    type: maintenance
    status: pending
    title: "Resolve stale links in requirements docs"
```

## Contract 4: Session Snapshot

Path: `.openkit/ops/sessions/<timestamp>.json`

```json
{
  "version": 1,
  "session_id": "mk-2026-02-18-001",
  "started_at": "2026-02-18T12:00:00Z",
  "ended_at": "2026-02-18T13:00:00Z",
  "summary": "Defined kernel artifacts and sprint plan",
  "actions": ["requirements-created", "sprint-created"]
}
```

## Contract 5: Memory Doctor Output

Path: stdout + optional `.openkit/ops/health/memory-health.json`

```json
{
  "version": 1,
  "score": 88,
  "status": "healthy",
  "checks": {
    "inline_links": "pass",
    "related_sections": "pass",
    "broken_wikilinks": "pass",
    "stale_docs": "warn"
  }
}
```

## Contract 6: Legacy Migration Report

Path: `.openkit/memory/legacy-migration-report.json`

```json
{
  "version": 1,
  "legacy_detected": true,
  "legacy_paths": [
    ".opencode/plugins/semantic-memory",
    ".opencode/memory"
  ],
  "migrated_at": "2026-02-18T14:00:00Z",
  "status": "completed",
  "notes": "Legacy plugin sunset migration completed"
}
```

## Contract 7: Deprecation Policy

Path: `docs/DEPRECATIONS.md` (or equivalent policy file)

```yaml
memory_legacy:
  release_n:
    status: deprecated
    warnings_enabled: true
  release_n_plus_1:
    status: removed
    legacy_commands_available: false
```

## Related

- [[requirements/memory-kernel-rust-cli/PLAN.md]]
- [[requirements/memory-kernel-rust-cli/ACCEPTANCE_CRITERIA.md]]
- [[requirements/memory-kernel-rust-cli/PROBLEM_STATEMENT.md]]
- [[sprint/Sprint-07/TASKS.md]]
- [[content-protocol/MANAGED_STATE_SCHEMA.md]]
