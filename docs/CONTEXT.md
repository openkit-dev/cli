# CONTEXT

**Created**: 2026-02-10
**Last Updated**: 2026-02-10
**Scope**: Backend/CLI

## Executive Summary

- **Project Type**: Go-based CLI tool (`openkit`) for agent orchestration and Spec-Driven Development (SDD).
- **Primary Purpose**: Universal toolkit that configures SDD environments for multiple AI coding agents (OpenCode, Claude, Cursor, Gemini, Codex).
- **Framework**: Uses Cobra (`github.com/spf13/cobra`) for CLI command handling.
- **State Management**: Local state tracking via `internal/managedstate` with conflict detection and drift detection.
- **Agent Registry**: 6 supported agents registered in `internal/agents/registry.go` (opencode, claude, cursor, gemini, codex, windsurf).
- **Template System**: Embedded templates in `internal/templates/embed.go` with 150+ files per agent.
- **Build System**: Makefile-based build process with multi-platform support (Darwin, Linux, Windows, AMD64/ARM64).
- **Linter**: `golangci-lint` configured via `.golangci.yml` with standard Go linters enabled.
- **UI**: `internal/ui` provides colored console output helpers (success, error, info, warning).
- **Configuration**: Agent behaviors and permissions defined in `opencode.json` with 16 specialized agents.
- **Self-Update**: Binary auto-update mechanism via `internal/selfupdate` with GitHub API integration.
- **Verification Scripts**: Python scripts in `.opencode/scripts/` for verification and quality checks.
- **Release**: GoReleaser configured via `.goreleaser.yaml` for automated releases.

## Repository Map

| Area | Path(s) | Notes |
|---|---|---|
| **Entry Point** | `cmd/openkit/` | Main CLI entry point. |
| **CLI Logic** | `internal/cli/` | Command definitions, flags, and handlers. |
| **Agent Core** | `internal/agents/` | Agent registry and execution logic. |
| **State Management** | `internal/managedstate/` | Persistence layer for tracking managed files. |
| **Synchronization** | `internal/syncer/` | Sync logic for agent configuration deployment. |
| **Target Systems** | `internal/targets/` | Agent-specific content generators (Claude, Cursor, Codex, Gemini). |
| **Self-Update** | `internal/selfupdate/` | Binary auto-update mechanism with GitHub API. |
| **Templates** | `internal/templates/` | Embedded project scaffolding templates. |
| **Platform** | `internal/platform/` | OS-specific abstractions. |
| **UI** | `internal/ui/` | Console output helpers. |
| **Config** | `opencode.json` | Agent system configuration with 16 agents. |
| **Build** | `Makefile` | Build automation with multi-platform support. |
| **CI/CD** | `.github/workflows/` | GitHub Actions for CI and releases. |
| **Docs** | `docs/` | Project documentation and SDD artifacts. |
| **Embedded Assets** | `.opencode/` | Commands, prompts, rules, skills, scripts for agents. |

## Key Flows

1. **CLI Execution**: `openkit <command>` -> `cmd/openkit` -> `internal/cli` -> `internal/agents` -> Output.
2. **Agent Sync**: `openkit <agent> sync` -> `internal/cli/agent_targets.go` -> `internal/syncer` -> `internal/templates` -> Files written to project.
3. **Build Process**: `make build` -> `go build ./cmd/openkit` -> Binary with version info embedded via LDFLAGS.
4. **Self-Update**: `openkit upgrade` -> `internal/selfupdate` -> GitHub API -> Checksum verification -> Binary replacement.
5. **Verification**: `.opencode/scripts/verify_all.py` -> Security scan, lint, UX audit, Lighthouse, Playwright E2E.

## Evidence

- `go.mod`: Go 1.25.7, dependencies (cobra, color, mod).
- `Makefile`: Build targets (`build`, `test`, `lint`, `build-all`, `install`).
- `opencode.json`: Agent definitions (orchestrator, backend-specialist, frontend-specialist, etc.) with permissions.
- `internal/agents/registry.go`: Registry of 6 supported agents (opencode, claude, cursor, gemini, codex, windsurf).
- `internal/syncer/syncer.go`: Synchronization logic with conflict detection.
- `internal/selfupdate/upgrade.go`: Binary update mechanism with SHA256 verification.
- `.golangci.yml`: Linter configuration with standard Go linters.
- `.github/workflows/ci.yml`: CI pipeline with Go setup, lint, test, build.
- `.github/workflows/release.yml`: Release automation with GoReleaser.
- `internal/cli/root.go`: Root command with banner, version info, and update check.
- `.opencode/scripts/verify_all.py`: Comprehensive verification script.
- `.opencode/scripts/checklist.py`: Quality checklist runner.
- `.goreleaser.yaml`: Release automation configuration.
- `internal/templates/embed.go`: Embedded file system for templates.

## Terminology

> For standard terminology definitions, see [[GLOSSARY.md]].

| Term | Definition (project-specific) |
|------|-------------------------------|
| **Agent** | An AI coding agent (OpenCode, Claude, Cursor, etc.) that OpenKit configures. |
| **SDD** | Spec-Driven Development - a workflow that emphasizes specification before implementation. |
| **Managed File** | A file tracked by OpenKit in `.openkit/managed.json` for conflict detection. |
| **Drift** | When a managed file has been manually modified outside of OpenKit sync. |
| **Conflict** | When an unmanaged file exists at a path where OpenKit wants to create a managed file. |
| **Embed** | Files embedded in the Go binary at build time via `internal/templates/embed.go`. |
| **Sync** | The process of installing or updating agent configuration files. |
| **Orphaned File** | A managed file that no longer exists in the embedded templates. |
| **Permission** | Access control for agent tools (allow, ask, deny) defined in `opencode.json`. |

## Dependencies

### Go Dependencies

| Package | Version | Purpose |
|---------|---------|---------|
| github.com/spf13/cobra | v1.10.2 | CLI command framework |
| github.com/fatih/color | v1.18.0 | Colored console output |
| golang.org/x/mod | v0.32.0 | Go module versioning |

### Development Dependencies

| Tool | Purpose |
|------|---------|
| golangci-lint | Linting (configured in `.golangci.yml`) |
| GoReleaser | Release automation |
| Go 1.25.x | Build toolchain |

## Project Stats

- **Go Files**: 33
- **Test Files**: 8
- **Supported Agents**: 6 (opencode, claude, cursor, gemini, codex, windsurf)
- **Specialized Agents**: 16 (orchestrator, chat, backend-specialist, frontend-specialist, database-architect, security-auditor, test-engineer, devops-engineer, mobile-developer, debugger, explorer-agent, performance-optimizer, seo-specialist, product-owner, project-planner, penetration-tester, documentation-writer)
- **Skills**: 33+ domain skills embedded in `.opencode/skills/`
- **Commands**: 18 slash commands embedded in `.opencode/commands/`
- **Platforms**: 5 (Darwin amd64/arm64, Linux amd64/arm64, Windows amd64)

## Related

- [[HUB-DOCS.md]]
- [[GLOSSARY.md]]
- [[SECURITY.md]]
- [[QUALITY_GATES.md]]
