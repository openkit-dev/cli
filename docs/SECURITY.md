# SECURITY

## Threats

- Supply-chain tampering during binary upgrade flow.
- Over-privileged agent execution (`bash`, `edit`, `write`) in misconfigured usage.
- Drift/conflict bypass risk when users force `--overwrite`.
- Missing CI security checks allows vulnerable dependencies or leaked secrets.

## Controls

- SHA256 verification for downloaded upgrade artifacts.
- Optional minisign verification for `checksums.txt` when `OPENKIT_MINISIGN_PUBKEY` is configured.
- Safe path guard blocks writes outside project root in sync engine.
- Managed-state checksum tracking to detect drift.
- CI lint/test/build pipeline on `main` PR/push.

## Gaps

- Signature verification is optional and depends on environment key configuration.
- No dependency vulnerability scanner in CI (`govulncheck`/equivalent absent).
- No secret scanning in CI.
- No dedicated structured audit log/correlation ID pattern for security operations.

## Prioritized Actions

| Priority | Action | Impact | Effort | Owner | Notes |
|---|---|---|---|---|---|
| P0 | Enforce signed release verification by default | High | Medium | Security/DevOps | Require public key configuration in managed environments. |
| P0 | Add secret scanning in CI | High | Low | Security | Add gitleaks or equivalent in PR pipeline. |
| P1 | Add dependency scanning in CI | Medium | Low | Security | Add `govulncheck ./...` gate. |
| P1 | Add security-specific static analysis | Medium | Medium | Security | Add `gosec` or security linters. |
| P2 | Add audit-oriented structured logs | Medium | Medium | Backend | Capture sync/upgrade actions and outcomes. |

## Evidence

- `rust-cli/src/main.rs`: `run_self_update_unix()` validates SHA256 checksums and optionally verifies minisign signatures.
- `internal/syncer/syncer.go`: `SafeAbsPath()` rejects writes outside project root.
- `internal/managedstate/managedstate.go`: file hash tracking (`InstalledSHA256`) and schema checks.
- `opencode.json`: some agents have broad permissions (example: `devops-engineer` sets `bash: allow`).
- `.github/workflows/release.yml`: optional checksums signing when `MINISIGN_SECRET_KEY` is configured.
- `.github/workflows/release.yml`: signing uses non-interactive minisign mode (`-W`) for CI compatibility.
- `.github/workflows/ci.yml`: includes lint/test/build only; no security or secret scan jobs.

## Related

- [[CONTEXT.md]]
- [[QUALITY_GATES.md]]
- [[ACTION_ITEMS.md]]
