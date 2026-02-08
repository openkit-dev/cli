# CONTEXT

**Created**: 2026-02-08
**Scope**: Backend/CLI

## Executive Summary

- **Project Type**: Go-based CLI tool (`openkit`) for agent orchestration.
- **Framework**: Uses Cobra for CLI command handling.
- **State Management**: Local state management via `internal/managedstate`.
- **Agents**: Agent logic encapsulated in `internal/agents`.
- **Build System**: Makefile based build process (`make build`, `make test`).
- **Linter**: `golangci-lint` configured via `.golangci.yml`.
- **UI**: `internal/ui` provides colored console output helpers.
- **Configuration**: `opencode.json` defines agent behaviors and permissions.
- **Dependencies**: Minimal external deps (`cobra`, `color`, `mod`).
- **Deployment**: Local binary installation (`make install`).

## Repository Map

| Area | Path(s) | Notes |
|---|---|---|
| **Entry Point** | `cmd/openkit/` | Main CLI entry point. |
| **CLI Logic** | `internal/cli/` | Command definitions and flags. |
| **Agent Core** | `internal/agents/` | Agent registry and execution logic. |
| **State** | `internal/managedstate/` | Persistence layer. |
| **Config** | `opencode.json` | Agent system configuration. |
| **Build** | `Makefile` | Automation scripts. |
| **Docs** | `docs/` | Project documentation. |

## Key Flows

1. **CLI Execution**: `openkit <command>` -> `cmd/openkit` -> `internal/cli` -> `internal/agents` -> Output.
2. **Build Process**: `make build` -> `go build ./cmd/openkit`.

## Evidence

- `go.mod`: Go 1.25.7, dependencies.
- `Makefile`: Build targets (`build`, `test`, `lint`).
- `opencode.json`: Agent definitions (`orchestrator`, `backend-specialist`, etc.).
- `internal/ui/`: Console output helpers (Fixed).
