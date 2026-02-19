# QUALITY_GATES

## Required

- Lint
- Type check
- Tests
- Security scan

## Gate Status

| Gate | Command/Source | Status | Evidence |
|---|---|---|---|
| Lint | `cargo clippy --manifest-path rust-cli/Cargo.toml --all-targets -- -D warnings` | Configured | `.github/workflows/ci.yml` step `Rust Lint`. |
| Format | `cargo fmt --manifest-path rust-cli/Cargo.toml --all --check` | Configured | `.github/workflows/ci.yml` step `Rust Format Check`. |
| Tests | `cargo test --manifest-path rust-cli/Cargo.toml` | Configured | `.github/workflows/ci.yml` step `Rust Contract Tests`. |
| Build | `cargo build --release --manifest-path rust-cli/Cargo.toml` | Configured | `.github/workflows/ci.yml` step `Rust Build`. |
| Security scan | CI | Missing | `.github/workflows/ci.yml` has no security scan step. |
| Dependency scan | CI | Missing | No Rust dependency audit step in CI workflow. |
| Coverage gate | CI | Missing | No coverage reporting gate configured for `rust-cli` tests. |

## Commands

```bash
# Project quality commands
make lint
make test
make build
python .opencode/scripts/checklist.py .
```

## CI Notes

- `.github/workflows/ci.yml` triggers on push and pull request to `main`.
- CI steps are checkout, setup-rust, release matrix validation, fmt, lint, test, build.
- `.github/workflows/release.yml` publishes `openkit_<OS>_<ARCH>` assets and `checksums.txt` on `v*` tags.

## Gaps

- Security and dependency scanning are not part of CI.
- Test coverage is not enforced as a gate.
- Python verification scripts exist but are not wired into CI.

## Evidence

- `.github/workflows/ci.yml`: Rust fmt/clippy/build/test pipeline.
- `Makefile`: Rust targets `lint`, `test`, `build`.
- `.opencode/scripts/checklist.py`: local checklist script available.

## Related

- [[CONTEXT.md]]
- [[SECURITY.md]]
- [[ACTION_ITEMS.md]]
