# BACKEND

## Endpoints

| Method | Route | Handler | Auth | Notes |
|---|---|---|---|---|
| N/A | not found | not found | not found | Project is a CLI application, not an HTTP server. |

## Jobs / Async

- Update check executes during CLI pre-run via HTTP call to GitHub release API.
- Agent sync applies file operations synchronously with backup support when overwriting.

## Backend Components

| Component | Location | Responsibility |
|---|---|---|
| Command dispatch | `internal/cli/root.go` | Cobra root setup, update notification, command registration. |
| Agent operations | `internal/cli/agent_targets.go` | `sync`, `upgrade`, `doctor` per agent. |
| State model | `internal/managedstate/managedstate.go` | Persist/validate managed file metadata and hashes. |
| Sync planner/apply | `internal/syncer/syncer.go` | Build plan, detect conflict/drift, apply create/update/delete safely. |
| Self-update | `internal/selfupdate/upgrade.go` | Download release artifact, verify checksum, replace binary. |

## Evidence

- `internal/cli/root.go`: `rootCmd` with `PersistentPreRunE` invoking update check.
- `internal/cli/agent_targets.go`: command definitions `sync`, `upgrade`, `doctor` and `runAgentSync`.
- `internal/syncer/syncer.go`: `ActionConflict`/`ActionOverwrite` decision paths.
- `internal/managedstate/managedstate.go`: `State`, `AgentState`, `FileEntry` persisted in JSON.
- `internal/selfupdate/upgrade.go`: checksum verification before atomic replacement.

## Related

- [[CONTEXT.md]]
- [[API.md]]
- [[SECURITY.md]]
