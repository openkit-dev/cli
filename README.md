# OpenKit CLI

OpenKit is now a Rust-only CLI runtime.

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

## Commands (Current)

```bash
openkit --help
openkit --version
openkit upgrade --check
openkit upgrade --dry-run
openkit check
openkit check --json
openkit init my-app --ai opencode --no-git
openkit opencode sync --overwrite
openkit opencode doctor --json

openkit memory init
openkit memory doctor --json --write
openkit memory capture --session-id s01 --summary "Sprint work" --action check
openkit memory review --json
```

## From Source

```bash
cargo build --release --manifest-path rust-cli/Cargo.toml
cargo test --manifest-path rust-cli/Cargo.toml
```

## Migration Status

- Runtime is Rust-only.
- Legacy Go CLI code was removed.
- Core parity migration continues in Sprint-09 (agent lifecycle commands).

See:

- `docs/sprint/Sprint-09/PARITY_MATRIX.md`
- `docs/requirements/memory-kernel-rust-cli/PLAN.md`
