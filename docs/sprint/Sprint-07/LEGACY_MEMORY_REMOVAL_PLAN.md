# Legacy Memory Removal Plan (N+1)

## Context

This plan defines the N+1 removal patch set for legacy semantic memory plugin workflows after the deprecation window established in release N.

The deprecation and migration prerequisites are tracked in [[DEPRECATIONS.md]] and [[MEMORY_LEGACY_MIGRATION.md]].

## Gate Criteria (Must Pass Before Removal)

1. Deprecation warnings shipped in release N.
2. Migration guide published and linked from CLI notices.
3. No open `High Impact` migration blockers in [[sprint/Sprint-07/RISK_REGISTER.md]].
4. Acceptance criteria for legacy sunset are satisfied in [[requirements/memory-kernel-rust-cli/ACCEPTANCE_CRITERIA.md]].

## Removal Scope

### Code Surface

- Remove `--memory` flag behavior and legacy install flow from `internal/cli/init.go`.
- Remove `--memory` flag behavior and legacy sync flow from `internal/cli/agent_targets.go`.
- Remove legacy `openkit memory` command group in `internal/cli/memory.go`.
- Remove legacy extraction helpers in `internal/templates/embed.go`:
  - `ExtractMemoryPlugin`
  - `ExtractMemoryRules`
- Remove now-unused helper code paths linked only to legacy plugin setup.

### Docs Surface

- Update `README.md` to remove legacy memory command examples and `--memory` usage.
- Keep a short historical note in [[DEPRECATIONS.md]].
- Replace or remove old references in sprint/requirements artifacts where actionable guidance still points to legacy commands.

### Runtime Artifacts

- Stop generating or mutating `.opencode/plugins/semantic-memory/`.
- Stop generating or mutating `.opencode/memory/`.
- Continue supporting migration report contract at `.openkit/memory/legacy-migration-report.json` for audit trail.

## Patch Sequence

1. Remove legacy command/flag registration paths.
2. Remove legacy plugin template extraction APIs.
3. Remove dead code and update tests.
4. Update docs/help output and migration references.
5. Run verification and link lint.

## Verification

- `go test ./...` passes.
- CLI help contains no legacy `--memory` option.
- Root command tree contains no legacy `openkit memory` group.
- `docs/` link lint passes (excluding pre-existing unrelated broken links tracked separately).
- README migration section points only to docs-first model.

## Rollback Strategy

If N+1 removal blocks critical workflows:

1. Revert removal commit set for legacy memory surfaces.
2. Keep deprecation warnings active.
3. Re-open sunset risk in [[sprint/Sprint-07/RISK_REGISTER.md]].
4. Publish hotfix note with revised removal timeline in [[DEPRECATIONS.md]].

## Related

- [[DEPRECATIONS.md]]
- [[MEMORY_LEGACY_MIGRATION.md]]
- [[requirements/memory-kernel-rust-cli/ACCEPTANCE_CRITERIA.md]]
- [[requirements/memory-kernel-rust-cli/PLAN.md]]
- [[sprint/Sprint-07/TASKS.md]]
- [[sprint/Sprint-07/RISK_REGISTER.md]]
