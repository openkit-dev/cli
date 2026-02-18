# Support Matrix

## Context

This matrix defines OpenKit support tiers for agent adapters during the OpenCode-first roadmap phase. The policy keeps memory and governance features moving fast while preserving multi-agent compatibility.

The Tier model aligns with the migration plan in [[requirements/memory-kernel-rust-cli/PLAN.md]] and sprint scope in [[sprint/Sprint-07/TASKS.md]].

## Tier Assignment

- Tier 1 (Primary): OpenCode
- Tier 2 (Supported): Claude Code, Codex, Antigravity
- Tier 3 (Experimental): Cursor, Gemini, Windsurf, others

## Capability Scope by Tier

| Capability | Tier 1 | Tier 2 | Tier 3 |
|---|---|---|---|
| Core workflow (`discover/specify/create/verify`) | Full | Full | Partial |
| Memory model | Docs + `.openkit/ops/` + `.openkit/memory/` with advanced automation | Docs-first + essential operational support | Best-effort |
| Adapter maintenance | Highest priority | Regular priority | Opportunistic |
| Contract tests | Full suite + regression gates | Essential regression suite | Smoke checks |
| Breaking-change tolerance | Very low | Moderate with migration notes | Higher tolerance |

## SLA and Quality Targets

| Metric | Tier 1 | Tier 2 | Tier 3 |
|---|---|---|---|
| Critical bug response | <= 72h | <= 10 days | Best-effort |
| Adapter regression threshold | <= 1.0% | <= 2.5% | <= 4.0% |
| Minimum test coverage target | >= 85% (core surfaces) | >= 70% (adapter surfaces) | Smoke only |

## Promotion Criteria

### Tier 3 -> Tier 2

- Adoption threshold met for two consecutive review cycles.
- Adapter error rate <= 2.5%.
- Essential regression tests available and passing.
- Documentation and migration guidance complete.

### Tier 2 -> Tier 1

- Adoption threshold met for three consecutive review cycles.
- Adapter error rate <= 1.0%.
- Full feature parity on mandatory workflow surfaces.
- Full regression and quality gate coverage.

## Demotion Criteria

- Tier 1 -> Tier 2: sustained error or support-cost breach across two review cycles.
- Tier 2 -> Tier 3: sustained reliability or adoption drop across two review cycles.

## Review Cadence

- Tier assignments are reviewed each major release.
- Changes require updates to this file and sprint backlog linkage.

## Related

- [[requirements/memory-kernel-rust-cli/PLAN.md]]
- [[requirements/memory-kernel-rust-cli/ACCEPTANCE_CRITERIA.md]]
- [[sprint/Sprint-07/BACKLOG.md]]
- [[sprint/Sprint-07/TASKS.md]]
- [[DEPRECATIONS.md]]
