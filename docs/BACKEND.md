# BACKEND / INTERNAL ARCHITECTURE

**Created**: 2026-02-08

## Architecture

Monolithic Go application organized by functional packages in `internal/`.

## Packages

| Package | Responsibility |
|---|---|
| `internal/agents` | Core agent logic and registry. |
| `internal/cli` | Cobra command definitions. |
| `internal/managedstate` | State persistence (file-based or local DB). |
| `internal/platform` | OS-specific abstractions. |
| `internal/selfupdate` | Binary auto-update mechanism. |
| `internal/syncer` | Synchronization logic (git/remote). |
| `internal/targets` | Target systems? |
| `internal/templates` | Project scaffolding templates. |
| `internal/ui` | Console output helpers (Success, Error, Info, Warning). |

## Data Models

- **State**: Managed by `managedstate` package.
- **Configuration**: Defined in `opencode.json` and parsed at runtime.
