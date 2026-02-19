# OpenKit CLI

OpenKit is a Rust-first CLI for project bootstrap, agent pack sync, environment checks, upgrade/uninstall, and Memory Kernel maintenance.

## Install

### macOS / Linux / WSL

```bash
curl -fsSL https://raw.githubusercontent.com/orionlabz/openkit/main/scripts/install.sh | bash
```

### Windows (PowerShell)

```powershell
irm https://raw.githubusercontent.com/orionlabz/openkit/main/scripts/install.ps1 | iex
```

### Manual Download

Download from [latest release](https://github.com/orionlabz/openkit/releases/latest):

- `openkit_Darwin_x86_64.tar.gz`
- `openkit_Darwin_arm64.tar.gz`
- `openkit_Linux_x86_64.tar.gz`
- `openkit_Linux_arm64.tar.gz`
- `openkit_Windows_x86_64.zip`

## Command Surface

```bash
# General
openkit --help
openkit --version

# Environment/system checks
openkit check
openkit check --json

# Project bootstrap
openkit init [project-name] --agent opencode --no-git
openkit init [project-name] --agent codex --no-git
openkit init --overwrite --no-git

# Agent pack lifecycle
openkit sync --agent opencode --overwrite
openkit doctor --agent opencode --json

# Binary lifecycle
openkit upgrade --check
openkit upgrade --dry-run
openkit upgrade
openkit uninstall --dry-run
openkit uninstall --yes

# Memory Kernel maintenance / repair
openkit memory init
openkit memory doctor --json --write
openkit memory capture --session-id s01 --summary "Sprint work" --action check
openkit memory review --json
```

## Supported Agents (`--agent`)

- `opencode`
- `claude`
- `cursor`
- `gemini`
- `codex`
- `antigravity`

These are supported in `openkit init`, `openkit sync`, and `openkit doctor`.

## Platform Support

- macOS: `x86_64`, `arm64`
- Linux: `x86_64`, `arm64`
- Windows: `x86_64`

## Upgrade Behavior

- `openkit upgrade --check`: queries latest release tag from GitHub.
- `openkit upgrade`:
  - Linux/macOS: Rust-native self-update (download artifact, verify `checksums.txt` SHA-256, replace binary with rollback path).
  - Windows: executes the official PowerShell installer flow.
- `openkit upgrade --dry-run`: prints planned update source/asset without changing binaries.
- `openkit uninstall --dry-run`: prints candidate install paths that would be removed.

## From Source

```bash
cargo fmt --manifest-path rust-cli/Cargo.toml --all --check
cargo clippy --manifest-path rust-cli/Cargo.toml --all-targets -- -D warnings
cargo build --release --manifest-path rust-cli/Cargo.toml
cargo test --manifest-path rust-cli/Cargo.toml
```

## Project Documentation

Discovery/specification/sprint artifacts are maintained in `openkit-memory/`.

- `openkit-memory/HUB-DOCS.md`
- `openkit-memory/CONTEXT.md`
- `openkit-memory/requirements/HUB-REQUIREMENTS.md`
- `openkit-memory/sprint/HUB-SPRINTS.md`
