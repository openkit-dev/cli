# CONTEXT

**Created**: 2026-02-10
**Last Updated**: 2026-02-17
**Scope**: Full (CLI-focused)

## Executive Summary (10 bullets)

- Project is a Go CLI, not a web app (`cmd/openkit/main.go`, `internal/cli/root.go`).
- Core workflow is agent sync/upgrade/doctor via `openkit <agent> <command>`.
- Sync engine uses managed state and checksum drift detection (`internal/syncer/syncer.go`).
- Managed state is file-based JSON at `.openkit/managed.json` (`internal/managedstate/managedstate.go`).
- CI enforces lint, test, build; security scanning is not in CI yet (`.github/workflows/ci.yml`).
- Release uses GoReleaser on git tags (`.github/workflows/release.yml`).
- Self-update verifies SHA256 checksums but not signatures (`internal/selfupdate/upgrade.go`).
- `make test` and `make build` pass on this discovery run (2026-02-17).
- Discovery command naming is drifted: code still exposes `openkit context` (`internal/cli/context.go`).
- No frontend routes, HTTP backend endpoints, or DB migrations found in repository Go sources.

## Repository Map

| Area | Path(s) | Notes |
|---|---|---|
| CLI entry | `cmd/openkit/main.go` | Boots CLI and executes Cobra root command. |
| CLI commands | `internal/cli/` | Root commands plus `context`, `init`, `upgrade`, `memory`, agent subcommands. |
| Sync engine | `internal/syncer/syncer.go` | Plans create/update/conflict/delete with safe path checks. |
| Managed state | `internal/managedstate/managedstate.go` | Persists per-agent installed artifact hashes. |
| Targets | `internal/targets/` | Generates per-agent content packs. |
| Self-update | `internal/selfupdate/` | Fetches latest tag and installs release artifacts. |
| Build and release | `Makefile`, `.github/workflows/` | Local build/test/lint and CI/release automation. |
| OpenKit config | `opencode.json` | Agent registry, tools, permissions. |
| Docs graph | `docs/` | Hubs, sprint/requirements artifacts, context pack. |

## Key Flows

1. CLI command path: `main()` -> `cli.Execute()` -> `rootCmd` dispatch.
2. Agent sync path: `openkit <agent> sync` -> `runAgentSync` -> `syncer.Apply` -> `managedstate.Save`.
3. Upgrade path: `openkit upgrade` -> GitHub release fetch -> artifact download -> checksum verify -> atomic replace.
4. Verification path: CI runs `make test` and `make build`; local scripts exist in `.opencode/scripts/`.

## Discovery Notes (Not Found)

- Frontend routing/data fetching: not found.
- HTTP backend endpoints/models/migrations: not found.
- Database migrations/schema files: not found.

## Evidence

- `cmd/openkit/main.go`: `cli.SetVersionInfo(version, commit, date)` then `cli.Execute()`.
- `internal/cli/root.go`: `Use: "openkit"` and update check against GitHub latest release URL.
- `internal/cli/agent_targets.go`: defines `sync`, `upgrade`, `doctor` subcommands per agent.
- `internal/syncer/syncer.go`: `ActionConflict` on unmanaged or drifted files unless `--overwrite`.
- `internal/managedstate/managedstate.go`: state saved under `.openkit/managed.json` with SHA256.
- `internal/selfupdate/upgrade.go`: verifies checksums from `checksums.txt` before replace.
- `go.mod`: module `github.com/openkit-devtools/openkit`, Go `1.25.7`, Cobra + color.
- `Makefile`: targets `test`, `build`, `build-all`, `lint`, `test-coverage`.
- `.github/workflows/ci.yml`: CI steps are lint, test, build.
- `.github/workflows/release.yml`: tag-triggered GoReleaser publish job.
- `internal/cli/context.go`: command still registered as `Use: "context"`.

## Terminology

> For standard terminology definitions, see [[GLOSSARY.md]]

| Term | Definition (project-specific) |
|------|-------------------------------|
| Managed state | JSON registry of installed artifacts and hashes for each agent. |
| Drift | Managed file hash no longer matches recorded installed hash. |
| Conflict | Existing unmanaged file blocks safe sync unless overwrite is enabled. |
| Pack | Embedded content version for an agent target distribution. |

## Related

- [[HUB-DOCS.md]]
- [[QUALITY_GATES.md]]
- [[SECURITY.md]]
- [[ACTION_ITEMS.md]]
