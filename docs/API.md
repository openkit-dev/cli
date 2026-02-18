# API

## Surface

| Type | Name | Location | Notes |
|---|---|---|---|
| CLI | `openkit` root command | `internal/cli/root.go` | Main command entrypoint. |
| CLI | `openkit init` | `internal/cli/init.go` | Initialize project scaffolding. |
| CLI | `openkit check` | `internal/cli/check.go` | Agent/system availability checks. |
| CLI | `openkit version` | `internal/cli/version.go` | Prints build metadata. |
| CLI | `openkit upgrade` | `internal/cli/upgrade.go` | Upgrades CLI binary. |
| CLI | `openkit uninstall` | `internal/cli/uninstall.go` | Removes CLI from system. |
| CLI | `openkit context` | `internal/cli/context.go` | Generates context documentation (drift vs `/discover` naming). |
| CLI | `openkit <agent> sync|upgrade|doctor` | `internal/cli/agent_targets.go` | Agent content lifecycle commands. |
| CLI | `openkit memory init|doctor|capture|review` | `internal/cli/memory_runtime.go` | Go bridge to docs-first memory runtime. |
| CLI | `openkit-rs memory init|doctor|capture|review` | `rust-cli/src/main.rs` | Docs-first Memory Kernel command bootstrap. |

## Contracts

- Command interface is Cobra-based (`*cobra.Command`) with `Use`, `Short`, `Run`/`RunE` handlers.
- Sync contract writes tracked artifacts and updates `.openkit/managed.json` state.
- Upgrade contract downloads release artifact + `checksums.txt`, verifies SHA256, then replaces binary.

## Gaps

- No HTTP/REST/GraphQL API surface found.
- No versioned machine-readable command schema exported.

## Evidence

- `internal/cli/root.go`: root command registration and execution flow.
- `internal/cli/agent_targets.go`: per-agent `sync`, `upgrade`, `doctor` command setup.
- `internal/cli/context.go`: command `Use: "context"`.
- `internal/selfupdate/upgrade.go`: binary upgrade checksum validation path.

## Related

- [[CONTEXT.md]]
- [[BACKEND.md]]
- [[SECURITY.md]]
