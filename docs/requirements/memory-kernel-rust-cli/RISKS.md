# Risks: Memory Kernel + Rust CLI Migration

## R01: Behavior Regression During Migration

- Probability: Medium
- Impact: High
- Description: Command behavior differences between Go and Rust versions can break user workflows.
- Mitigation: Golden contract tests before and after each migrated command; phased cutover.

## R02: Documentation Debt Migration

- Probability: High
- Impact: Medium
- Description: Existing docs may not follow inline-link + related standards immediately.
- Mitigation: Prioritize hub and requirements docs first; run iterative memory doctor remediation.

## R03: Scope Creep (Kernel + Runtime + Adapters)

- Probability: High
- Impact: High
- Description: Combining architecture redesign with language migration can overload one sprint.
- Mitigation: Lock sprint scope to foundation artifacts and migration bootstrap only.

## R04: Tier Confusion

- Probability: Medium
- Impact: Medium
- Description: Teams may expect Tier-2 adapters to receive Tier-1 features immediately.
- Mitigation: Publish explicit support matrix and parity rules in sprint deliverables.

## R05: Operational Noise

- Probability: Medium
- Impact: Medium
- Description: New ops logs (observations/tensions/sessions) can become noisy without thresholds.
- Mitigation: Add threshold-based review triggers and queue compaction rules.

## R06: Legacy Memory Removal Backlash

- Probability: Medium
- Impact: High
- Description: Users depending on `--memory` and legacy `openkit memory` behavior may face workflow disruption after removal.
- Mitigation: Two-release deprecation policy, migration report, and compatibility notes in README/docs.

## R07: Partial Migration State

- Probability: Medium
- Impact: Medium
- Description: Repositories can end up with mixed old/new memory artifacts if migration is interrupted.
- Mitigation: Idempotent migration command, explicit `legacy-migration-report.json`, and doctor checks for mixed state.

## Related

- [[requirements/memory-kernel-rust-cli/PLAN.md]]
- [[requirements/memory-kernel-rust-cli/ACCEPTANCE_CRITERIA.md]]
- [[sprint/Sprint-07/RISK_REGISTER.md]]
- [[requirements/memory-kernel-rust-cli/PROBLEM_STATEMENT.md]]
- [[requirements/memory-kernel-rust-cli/HUB-MEMORY-KERNEL-RUST-CLI.md]]
