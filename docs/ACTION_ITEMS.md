# ACTION_ITEMS

| Priority | Item | Impact | Effort | Owner | Notes |
|---|---|---|---|---|---|
| P0 | Add release signature verification | High | Medium | Security/DevOps | `internal/selfupdate/upgrade.go` verifies checksum only. |
| P0 | Add secret scanning in CI | High | Low | Security | Missing in `.github/workflows/ci.yml`. |
| P1 | Add dependency vulnerability scan | Medium | Low | Security | Add `govulncheck` gate in CI. |
| P1 | Add coverage gate in CI | Medium | Low | Testing | `make test-coverage` exists but not enforced in CI. |
| P1 | Align command docs with CLI command surface | Medium | Medium | Docs/Backend | Code exposes `openkit context`; docs emphasize `/discover`. |
| P2 | Add structured audit logging for sync/upgrade | Medium | Medium | Backend | No centralized audit/correlation-id logging pattern. |
| P2 | Lower complexity threshold or justify exceptions | Medium | Medium | Backend | `.golangci.yml` currently sets `gocyclo` threshold to 30. |

## Cross-Repo Impact

| Severity | Owner | Impact | Action |
|---|---|---|---|
| Medium | Product/Docs | Command naming drift can confuse downstream agent packs and user docs consumers. | Standardize on one naming path and update generated docs templates accordingly. |

## Blockers

- External scanners/tools (e.g., `govulncheck`, gitleaks) may require CI environment updates and policy approval.

## Related

- [[CONTEXT.md]]
- [[SECURITY.md]]
- [[QUALITY_GATES.md]]
