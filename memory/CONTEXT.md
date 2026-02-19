# CONTEXT

**Created**: 2026-02-18
**Scope**: Full (CLI-centric)

## Executive Summary (10 bullets)

- Repository is a Rust-first CLI project with runtime in `rust-cli/`.
- Main command surface is implemented in one entrypoint file: `rust-cli/src/main.rs`.
- CI currently enforces format, lint, build, and contract tests for Rust.
- Release pipeline builds multi-platform binaries and publishes checksums.
- Installers fetch artifacts from GitHub Releases for Unix and Windows.
- Memory Kernel contract files exist under `.openkit/memory/` and `.openkit/ops/`.
- No web frontend application routing/data fetching was found.
- No backend HTTP API server/routes were found.
- No database schema/migrations were found.
- Documentation drift exists: README still references `docs/...` paths not present in repo.

## Repository Map

| Area | Path(s) | Notes |
|---|---|---|
| CLI Runtime | `rust-cli/src/main.rs` | Clap-based command parser and command handlers. |
| CLI Tests | `rust-cli/tests/command_contracts.rs` | Contract tests for init/sync/check/memory/upgrade/uninstall. |
| Installer Scripts | `scripts/install.sh`, `scripts/install.ps1` | Release download + install flows. |
| CI/CD | `.github/workflows/ci.yml`, `.github/workflows/release.yml` | Quality gates + release packaging matrix. |
| Memory Kernel | `.openkit/memory/config.yaml`, `.openkit/memory/derivation.yaml`, `.openkit/ops/queue.yaml` | Required memory contracts present. |
| Frontend | not found | No app routes/components outside template assets. |
| Backend API | not found | No HTTP framework/server endpoints found in runtime. |
| Database | not found | No migration folders, SQL files, or ORM schema files found. |

## Key Flows

1. Command dispatch: `openkit <subcommand>` -> `Cli::parse()` -> `match Commands` in `rust-cli/src/main.rs` -> handler result.
2. Upgrade flow (Unix): release lookup -> asset download -> checksum verification -> atomic binary replacement.
3. Memory health flow: `openkit memory doctor` -> wikilink/related/broken-link checks -> scored report output.

## Evidence

- `README.md:3`: `OpenKit is now a Rust-only CLI runtime.`
- `rust-cli/src/main.rs:28`: command enum defines CLI-only surface (`Check`, `Init`, `Sync`, `Doctor`, `Upgrade`, `Uninstall`, `Memory`).
- `rust-cli/src/main.rs:333`: `run_upgrade` performs release discovery and update execution.
- `rust-cli/src/main.rs:424`: `verify_checksum` validates SHA-256 before replacement.
- `.github/workflows/ci.yml:55`: CI enforces `cargo fmt`, `cargo clippy`, `cargo build`, `cargo test`.
- `.github/workflows/release.yml:18`: release matrix covers Linux/macOS/Windows targets.
- `rust-cli/tests/command_contracts.rs:52`: tests assert Memory Kernel contract files and command behavior.
- `README.md:77`: references `docs/sprint/Sprint-09/PARITY_MATRIX.md` (path not found in repository).
- `README.md:78`: references `docs/requirements/memory-kernel-rust-cli/PLAN.md` (path not found in repository).
- `glob("**/*.go")`: not found (legacy Go runtime removed as documented).
- `glob("docs/**")`: not found.
- `glob("**/*.sql")`: not found.

## Terminology

> For standard terminology definitions, see [[GLOSSARY.md]]

| Term | Definition (project-specific) |
|------|-------------------------------|
| Memory Kernel | Docs-first persistence layer under `.openkit/` and `openkit-memory/`. |
| Contract Tests | Rust integration tests validating command IO and generated artifacts. |
| Parity | Progress of Rust runtime feature coverage replacing legacy Go runtime. |

## Related

- [[HUB-DOCS.md]]
- [[API.md]]
- [[QUALITY_GATES.md]]
- [[SECURITY.md]]
- [[ACTION_ITEMS.md]]
