# QUALITY_GATES

## Required

- Lint
- Type check
- Tests
- Security scan

## Gate Status

| Gate | Command/Source | Status | Evidence |
|---|---|---|---|
| Lint | `make lint` | Configured | `Makefile` target `lint`; CI uses `golangci-lint-action@v6`. |
| Type check | `golangci-lint` (`typecheck`) | Configured | `.golangci.yml` enables `typecheck`. |
| Tests | `make test` | Passing | Discovery run on 2026-02-17 passed all Go tests. |
| Build | `make build` | Passing | Discovery run on 2026-02-17 produced `openkit` binary. |
| Security scan | CI | Missing | `.github/workflows/ci.yml` has no security scan step. |
| Dependency scan | CI | Missing | No `govulncheck` or equivalent in CI workflow. |
| Coverage gate | CI | Missing | `test-coverage` exists in `Makefile` but not in CI. |

## Commands

```bash
# Project quality commands
make lint
make test
make build
make test-coverage
python .opencode/scripts/checklist.py .
```

## CI Notes

- `.github/workflows/ci.yml` triggers on push and pull request to `main`.
- CI steps are checkout, setup-go, lint, test, build.
- `.github/workflows/release.yml` publishes on `v*` tags through GoReleaser.

## Gaps

- Security and dependency scanning are not part of CI.
- Test coverage is not enforced as a gate.
- Python verification scripts exist but are not wired into CI.

## Evidence

- `.github/workflows/ci.yml`: `Lint`, `Test`, `Build` steps only.
- `.golangci.yml`: enabled linters include `typecheck`, `gocyclo`, `staticcheck`.
- `Makefile`: targets `lint`, `test`, `build`, `test-coverage`.
- `.opencode/scripts/checklist.py`: local checklist script available.

## Related

- [[CONTEXT.md]]
- [[SECURITY.md]]
- [[ACTION_ITEMS.md]]
