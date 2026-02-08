# SECURITY

**Created**: 2026-02-08

## Threats & Risks

- **Arbitrary Code Execution**: Agents execute commands. `opencode.json` defines permissions, but rigorous sandboxing is needed.
- **Supply Chain**: `internal/selfupdate` verifies SHA256 checksums but lacks cryptographic signature verification (Cosign/GPG).
- **File System Access**: Agents have `read`/`write` access. Scope limitation is critical.

## Controls

- **Permission System**: `opencode.json` defines `allow`/`ask`/`deny` for sensitive tools (`bash`, `write`).
- **Dependencies**: Managed via `go.mod` / `go.sum`.

## Gaps

- **SAST**: No static analysis for security vulnerabilities in CI.
- **Secret Scanning**: No automated secret scanning detected.
- **Sandbox**: No explicit containerization or sandbox for agent execution visible in root.
