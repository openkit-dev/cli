# SECURITY

**Created**: 2026-02-10
**Last Updated**: 2026-02-12

## Threats & Risks

### High Priority

1. **Arbitrary Code Execution**: Agents execute commands via `bash` tool. `opencode.json` defines permissions (allow/ask/deny) but rigorous sandboxing is not enforced.
   - **Evidence**: `opencode.json` defines `permission` blocks for each agent.
   - **Impact**: Malicious agent or compromised configuration could execute arbitrary commands.

2. **Supply Chain Attacks**: `internal/selfupdate` verifies SHA256 checksums but lacks cryptographic signature verification (Cosign/GPG).
   - **Evidence**: `internal/selfupdate/upgrade.go` implements SHA256 verification only.
   - **Impact**: Compromised GitHub releases or man-in-the-middle attacks could install malicious binaries.

3. **File System Access**: Agents have `read`/`write`/`edit` access to project files. Scope limitation is critical but not enforced.
   - **Evidence**: `opencode.json` defines tools but no path-based restrictions.
   - **Impact**: Agents could modify sensitive files outside the project scope.

4. **Unverified Remote Content**: Some agents (security-auditor, penetration-tester) have `webfetch` tool access.
   - **Evidence**: `opencode.json` - security-auditor, penetration-tester agents have `webfetch: true`.
   - **Impact**: Could fetch and process malicious content.

### Medium Priority

5. **State File Tampering**: `.openkit/managed.json` tracks managed files but no integrity verification exists.
   - **Evidence**: `internal/managedstate/managedstate.go` reads/writes JSON without signatures.
   - **Impact**: Tampered state could lead to incorrect sync behavior.

6. **Dependency Vulnerabilities**: External Go dependencies could contain known vulnerabilities.
   - **Evidence**: `go.mod` lists 3 direct dependencies with transitive dependencies.
   - **Impact**: Vulnerable dependencies could be exploited.

7. **Template Injection**: Embedded templates could be compromised if build process is not secure.
   - **Evidence**: `internal/templates/embed.go` embeds files.
   - **Impact**: Malicious templates could inject code into user projects.

## Controls

### Existing Controls

1. **Permission System**: `opencode.json` defines `allow`/`ask`/`deny` for sensitive tools (`bash`, `write`, `edit`, `webfetch`).
   - **Evidence**: `opencode.json` - Each agent has `permission` block.
   - **Status**: Implemented but not enforced at runtime (depends on agent compliance).

2. **Dependency Management**: Dependencies managed via `go.mod` / `go.sum`.
   - **Evidence**: `go.mod` and `go.sum` files.
   - **Status**: Standard Go module security.

3. **Checksum Verification**: Self-update verifies SHA256 checksums.
   - **Evidence**: `internal/selfupdate/upgrade.go`.
   - **Status**: Implemented but insufficient (no signatures).

4. **Linter**: `golangci-lint` includes `staticcheck` for common issues.
   - **Evidence**: `.golangci.yml`.
   - **Status**: Configured and run in CI.

5. **CI/CD**: GitHub Actions runs tests and linting on PRs.
   - **Evidence**: `.github/workflows/ci.yml`.
   - **Status**: Basic quality checks only.

6. **Drift Detection**: Warns about manual changes to managed files.
   - **Evidence**: `internal/syncer/syncer.go` detects drift.
   - **Status**: Implemented but informational only.

## Gaps

### Critical Gaps

1. **No Cryptographic Signatures**: Self-update lacks GPG/Cosign signature verification.
   - **Current**: SHA256 checksums only.
   - **Needed**: Sign releases with Cosign or GPG and verify before install.
   - **Evidence**: `internal/selfupdate/upgrade.go` - No signature verification code.

2. **No SAST**: No static application security testing in CI.
   - **Current**: Basic linting only.
   - **Needed**: Add SAST scanner (e.g., golangci-lint with security linters, gosec, SonarQube).
   - **Evidence**: `.github/workflows/ci.yml` - No security scanners.

3. **No Secret Scanning**: No automated secret scanning detected.
   - **Current**: No secret scanning in CI or pre-commit hooks.
   - **Needed**: Add secret scanning (e.g., gitleaks, trufflehog).
   - **Evidence**: `.github/workflows/ci.yml` - No secret scanners.

4. **No Sandbox**: No explicit containerization or sandbox for agent execution.
   - **Current**: Agents run directly on host system.
   - **Needed**: Consider sandboxing for agent tool execution.
   - **Evidence**: No sandbox code in `internal/`.

5. **No SBOM**: No Software Bill of Materials generated.
   - **Current**: No SBOM in CI or releases.
   - **Needed**: Generate and sign SBOM (e.g., using Syft, CycloneDX).
   - **Evidence**: No SBOM files or generation in `.github/workflows/`.

### Medium Priority Gaps

6. **No Dependency Scanning**: No automated dependency vulnerability scanning.
   - **Current**: No `govulncheck`, `npm audit`, or similar in CI.
   - **Needed**: Add `govulncheck` to CI pipeline.
   - **Evidence**: `.github/workflows/ci.yml` - No dependency scanning.

7. **No Runtime Logging**: No audit logging for agent actions.
   - **Current**: Actions are not logged for audit trails.
   - **Needed**: Add structured logging for sensitive operations.
   - **Evidence**: No audit logging in codebase.

8. **No Input Validation**: Template parameters not validated before rendering.
   - **Current**: No input sanitization for user-provided values.
   - **Needed**: Add validation for template variables.
   - **Evidence**: No validation in template code.

9. **No Rate Limiting**: No rate limiting on external API calls (webfetch, self-update).
   - **Current**: Unbounded HTTP requests.
   - **Needed**: Add rate limiting and timeouts.
   - **Evidence**: `internal/selfupdate/upgrade.go` has 2s timeout but no rate limiting.

10. **No Integrity Verification**: No verification of embedded templates at build time.
    - **Current**: Templates embedded without integrity checks.
    - **Needed**: Hash templates and verify at runtime.
    - **Evidence**: `internal/templates/embed.go` - No integrity verification.

## Prioritized Actions

| Priority | Action | Impact | Effort | Owner | Notes |
|----------|--------|--------|--------|-------|-------|
| **P0** | Implement Update Signatures | Prevent supply chain attacks | High | Security | Add GPG/Cosign verification to `internal/selfupdate`. Checksum only is insufficient. |
| **P0** | Add SAST to CI | Catch security issues early | High | Security | Add `gosec` or enable security linters in golangci-lint. |
| **P0** | Add Secret Scanning | Prevent secrets in repo | High | Security | Add gitleaks to pre-commit hooks and CI. |
| **P1** | Add Dependency Scanning | Catch vulnerable deps | Medium | Security | Add `govulncheck` to CI pipeline. |
| **P1** | Generate and Sign SBOM | Supply chain transparency | Medium | Security | Generate CycloneDX SBOM with Syft. |
| **P2** | Add Audit Logging | Traceability | Medium | DevOps | Log sensitive operations to structured logs. |
| **P2** | Add Input Validation | Prevent injection | Medium | Backend | Validate template parameters. |
| **P3** | Explore Sandboxing | Containment | High | Security | Research sandbox options for agent execution. |
| **P3** | Add Rate Limiting | DoS protection | Low | Backend | Rate limit external API calls. |
| **P3** | Template Integrity | Verify templates | Low | Backend | Hash templates and verify at runtime. |

## Security Best Practices Not Implemented

1. **Principle of Least Privilege**: Agents have broad file system access without path restrictions.
2. **Defense in Depth**: Single layer of permission controls, no sandbox or containment.
3. **Secure by Default**: Some agents have `bash: allow` (devops-engineer) without additional controls.
4. **Supply Chain Security**: No SBOM, no signatures, no provenance tracking.
5. **Auditability**: No audit logs for security-relevant events.

## Related

- [[HUB-DOCS.md]]
- [[CONTEXT.md]]
- [[QUALITY_GATES.md]]
