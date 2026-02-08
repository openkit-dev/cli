# QUALITY GATES

**Created**: 2026-02-08

## Linters

- **Tool**: `golangci-lint`
- **Command**: `make lint`
- **Config**: `.golangci.yml` (standard Go linters enabled).
- **Status**: Configured. Requires `golangci-lint` binary in PATH.

## Testing

- **Tool**: Go Test
- **Command**: `make test` (`go test -v ./...`)
- **Coverage**: `make test-coverage` generates HTML report. Unit tests exist for `internal/ui`, `internal/cli`, `internal/selfupdate`, `internal/targets`.
- **Status**: Infrastructure exists.

## Build

- **Command**: `make build`
- **Artifacts**: `openkit` binary.
- **Multi-platform**: `make build-all` supports Darwin, Linux, Windows.

## CI/CD

- **Status**: Configured via GitHub Actions.
- **Workflow**: `.github/workflows/ci.yml` (Build, Test, Lint).
