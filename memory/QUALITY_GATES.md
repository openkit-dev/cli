# QUALITY_GATES

## Required

- Lint: required
- Type/compile safety: required
- Tests: required
- Security scan: required

## Commands

```bash
# Rust runtime
cargo fmt --manifest-path rust-cli/Cargo.toml --all --check
cargo clippy --manifest-path rust-cli/Cargo.toml --all-targets -- -D warnings
cargo build --release --manifest-path rust-cli/Cargo.toml
cargo test --manifest-path rust-cli/Cargo.toml

# OpenKit validation pipeline
python .opencode/scripts/checklist.py .
python .opencode/scripts/verify_all.py . --url http://localhost:3000
```

## CI Notes

- CI workflow currently runs format/lint/build/tests in `.github/workflows/ci.yml`.
- Release workflow builds and packages binaries per target matrix in `.github/workflows/release.yml`.
- Dependency vulnerability audit (for example `cargo audit`) is not configured in CI.
- Verification status for this discovery run: static review only; commands not executed.

## Evidence

- `.github/workflows/ci.yml:55`: `cargo fmt` check step exists.
- `.github/workflows/ci.yml:58`: `cargo clippy` step exists with `-D warnings`.
- `.github/workflows/ci.yml:61`: release build step runs in CI.
- `.github/workflows/ci.yml:64`: contract tests run via `cargo test`.
- `.opencode/scripts/checklist.py:63`: checklist includes security, lint, tests, and audits.
- `grep("cargo audit|cargo-audit|cargo deny")`: not found in workflows.

## Related

- [[CONTEXT.md]]
- [[SECURITY.md]]
- [[ACTION_ITEMS.md]]
- [[HUB-DOCS.md]]
