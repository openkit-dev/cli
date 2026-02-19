# SECURITY

## Threats

- Supply-chain compromise via malicious release artifact downloads.
- Documentation drift causing incorrect operator actions.
- Accidental destructive local operations (`uninstall` without dry-run/confirm awareness).
- Weak observability due to plain stdout logging (limited audit trail for automation).

## Controls

- Upgrade flow verifies SHA-256 checksum before binary replacement.
- Release pipeline generates and publishes `checksums.txt`.
- CLI uninstall requires explicit confirmation unless `--yes` is provided.
- CI enforces format/lint/build/test gates before merge.

## Gaps

- No CI dependency audit gate (`cargo audit`/equivalent) was found.
- No signature verification for release artifacts was found (checksum-only model).
- No structured logger/correlation-id pattern is implemented in runtime handlers.
- README references non-existent `docs/...` paths, increasing operational drift risk.

## Prioritized Actions

- P0: Add Rust dependency vulnerability scan in CI and fail on high severity findings.
- P1: Add artifact signing verification path (minisign/cosign) for installer and self-update.
- P1: Resolve README stale links from `docs/...` to `openkit-memory/...` canonical paths.
- P2: Introduce structured logging mode for machine-readable execution traces.

## Evidence

- `rust-cli/src/main.rs:424`: `verify_checksum` compares expected and actual SHA-256.
- `.github/workflows/release.yml:98`: release workflow generates checksums.
- `rust-cli/src/main.rs:287`: uninstall prompts user confirmation when `--yes` is not set.
- `.github/workflows/ci.yml:55`: quality gates enforced in CI.
- `grep("cargo audit|cargo-audit|cargo deny")`: not found in repository workflow files.
- `README.md:77`: stale reference to `docs/sprint/Sprint-09/PARITY_MATRIX.md`.

## Related

- [[CONTEXT.md]]
- [[QUALITY_GATES.md]]
- [[ACTION_ITEMS.md]]
- [[HUB-DOCS.md]]
