# RISK REGISTER

**Sprint**: Sprint-07
**Status**: Monitored (Sprint Complete)

## Risk Matrix

| ID | Risk | Probability | Impact | Owner | Mitigation | Linked Artifact |
|---|---|---|---|---|---|---|
| R07-001 | Rust migration causes command regressions | Medium | High | backend-specialist | Golden tests for each migrated command | [[requirements/memory-kernel-rust-cli/RISKS.md]] |
| R07-002 | Documentation standards adoption is inconsistent | High | Medium | documentation-writer | Enforce inline + related checks in doctor | [[requirements/memory-kernel-rust-cli/ACCEPTANCE_CRITERIA.md]] |
| R07-003 | Sprint scope grows beyond foundation | High | High | product-owner | Freeze scope to kernel + bootstrap deliverables | [[sprint/Sprint-07/BACKLOG.md]] |
| R07-004 | Tier policy misinterpreted by contributors | Medium | Medium | project-planner | Publish support matrix with explicit boundaries | [[requirements/memory-kernel-rust-cli/PLAN.md]] |
| R07-005 | Legacy memory deprecation causes user disruption | Medium | High | backend-specialist | Two-release sunset, warnings, migration guide | [[requirements/memory-kernel-rust-cli/RISKS.md]] |

## Escalation Rules

- Any `High Impact` risk not mitigated by mid-sprint must trigger backlog reprioritization.
- Any unresolved `R07-001` item blocks migration cutover tasks.
- Any unresolved `R07-005` item blocks legacy removal tasks.

## Sprint Exit Disposition

- No High Impact blocker remained open at sprint close.
- Residual risks moved to monitoring in the exit report.

## Related

- [[sprint/Sprint-07/TASKS.md]]
- [[requirements/memory-kernel-rust-cli/RISKS.md]]
- [[requirements/memory-kernel-rust-cli/ACCEPTANCE_CRITERIA.md]]
- [[sprint/Sprint-07/HUB-SPRINT-07.md]]
- [[sprint/Sprint-07/EXIT_REPORT.md]]
