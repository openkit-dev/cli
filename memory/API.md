# API

## Surface

| Type | Name | Location | Notes |
|---|---|---|---|
| CLI | `openkit check [--json]` | `rust-cli/src/main.rs` | Local environment/agent tooling health report. |
| CLI | `openkit init [project] [--agent ...]` | `rust-cli/src/main.rs` | Bootstraps project, docs graph, and Memory Kernel. |
| CLI | `openkit sync --agent <name>` | `rust-cli/src/main.rs` | Syncs agent pack/templates into project. |
| CLI | `openkit doctor --agent <name> [--json]` | `rust-cli/src/main.rs` | Reports agent configuration status. |
| CLI | `openkit upgrade [--check|--dry-run]` | `rust-cli/src/main.rs` | Self-update (Unix) or installer handoff (Windows). |
| CLI | `openkit uninstall [--dry-run|--yes]` | `rust-cli/src/main.rs` | Removes installed binary from known locations. |
| CLI | `openkit memory init` | `rust-cli/src/main.rs` | Repairs/creates required Memory Kernel files. |
| CLI | `openkit memory doctor [--json --write]` | `rust-cli/src/main.rs` | Validates docs links/related sections and scores health. |
| CLI | `openkit memory capture` | `rust-cli/src/main.rs` | Captures session snapshot into `.openkit/ops/sessions`. |
| CLI | `openkit memory review [--json]` | `rust-cli/src/main.rs` | Aggregates memory operational counters/recommendations. |

## Contracts

- JSON output contracts are validated in `rust-cli/tests/command_contracts.rs` for `check --json`, `memory doctor --json`, and `memory review --json`.
- Memory Kernel file contracts are validated against golden fixtures in `rust-cli/tests/golden/`.
- HTTP API endpoints: not found.

## Evidence

- `rust-cli/src/main.rs:28`: full command set declared in `enum Commands`.
- `rust-cli/src/main.rs:110`: memory subcommands declared in `enum MemorySubcommand`.
- `rust-cli/tests/command_contracts.rs:209`: verifies `check --json` schema shape.
- `rust-cli/tests/command_contracts.rs:96`: verifies `memory doctor --json` output.
- `rust-cli/tests/command_contracts.rs:139`: verifies `memory capture` + `memory review` flow.

## Related

- [[CONTEXT.md]]
- [[HUB-DOCS.md]]
- [[QUALITY_GATES.md]]
- [[SECURITY.md]]
