# ACTION_ITEMS

**Created**: 2026-02-10
**Last Updated**: 2026-02-10

| Priority | Item | Impact | Effort | Owner | Notes |
|---|---|---|---|---|---|
| **P0** | Implement Update Signatures | Prevent supply chain attacks | High | Security | Add GPG/Cosign verification to `internal/selfupdate`. Checksum only is insufficient. Blocker for secure releases. |
| **P0** | Add SAST to CI | Catch security issues early | High | Security | Add `gosec` or enable security linters in golangci-lint. Run on every PR. |
| **P0** | Add Secret Scanning | Prevent secrets in repo | High | Security | Add gitleaks to pre-commit hooks and CI. Detect leaked credentials/API keys. |
| **P0** | Add Coverage Reporting | Enforce test quality | Medium | Testing | Add coverage reporting to CI and set minimum threshold (e.g., 70%). |
| **P1** | Add Dependency Scanning | Catch vulnerable deps | Medium | Security | Add `govulncheck` to CI pipeline. Scan for known CVEs in Go dependencies. |
| **P1** | Generate and Sign SBOM | Supply chain transparency | Medium | Security | Generate CycloneDX SBOM with Syft. Sign with Cosign. |
| **P1** | Reduce Cyclomatic Complexity | Improve code maintainability | Medium | Backend | Change gocyclo threshold from 30 to 15. Refactor complex functions. |
| **P1** | Add Integration Tests | Improve test coverage | High | Testing | Add end-to-end tests for CLI commands (sync, upgrade, doctor). |
| **P1** | Integrate Verification Scripts | Ensure project quality | Medium | DevOps | Migrate Python verification scripts to Go or integrate via CI. |
| **P2** | Add Audit Logging | Traceability | Medium | DevOps | Log sensitive operations (sync, upgrade, install) to structured logs. |
| **P2** | Add Input Validation | Prevent injection | Medium | Backend | Validate template parameters before rendering. Sanitize user input. |
| **P2** | Add Code Signing | Verify binary authenticity | High | DevOps | Sign macOS (codesign) and Windows (signtool) binaries in release workflow. |
| **P2** | Add Performance Testing | Detect regressions | Medium | Performance | Add benchmarks for critical paths (sync, template rendering). |
| **P2** | Add Rate Limiting | DoS protection | Low | Backend | Rate limit external API calls (webfetch, self-update, GitHub API). |
| **P3** | Explore Sandboxing | Containment | High | Security | Research sandbox options for agent execution (gVisor, Firecracker, containers). |
| **P3** | Add Fuzzing | Find security bugs | High | Security | Add fuzz testing for input parsing and template rendering. |
| **P3** | Template Integrity | Verify templates | Low | Backend | Hash embedded templates and verify integrity at runtime. |
| **P3** | Add Chaos Engineering | Test resilience | Low | DevOps | Add chaos testing for distributed operations (if applicable). |
| **P3** | Automate CHANGELOG.md | Improve release process | Low | Documentation | Auto-generate changelog from git commits or PRs. |

## Status Summary

| Priority | Count | Total Effort |
|----------|-------|--------------|
| **P0** | 4 | High+High+High+Medium = Very High |
| **P1** | 5 | Medium+Medium+Medium+High+Medium = High |
| **P2** | 5 | Medium+Medium+High+Medium+Low = Medium-High |
| **P3** | 5 | High+High+Low+Low+Low = Medium |

**Total**: 19 items

## Quick Wins (Low Effort, High Impact)

1. **Add Secret Scanning** (P0, High Impact, High Effort): Critical for security, relatively easy to implement with gitleaks.
2. **Add SAST to CI** (P0, High Impact, High Effort): Easy to add, immediate security benefits.
3. **Add Dependency Scanning** (P1, Medium Impact, Medium Effort): One-line `govulncheck` command in CI.

## Long-Term Strategic Items

1. **Explore Sandboxing** (P3, High Impact, High Effort): Major architectural change, requires significant research and development.
2. **Add Fuzzing** (P3, High Impact, High Effort): Requires continuous integration and bug fixing effort.

## Blocked Items

| Item | Blocked By | Reason |
|------|------------|--------|
| Generate and Sign SBOM | Code Signing | Need code signing infrastructure before SBOM signing. |
| Release binary authenticity | Code Signing | Need codesign (macOS) and signtool (Windows) setup. |
| Supply chain transparency | Update Signatures | Need cryptographic signatures before full SBOM implementation. |

## Next Steps (Recommended Order)

1. **Week 1**: Add Secret Scanning and SAST to CI (P0 items)
2. **Week 2**: Add Coverage Reporting and Dependency Scanning (P0/P1)
3. **Week 3**: Reduce Complexity and Add Integration Tests (P1)
4. **Week 4**: Add Code Signing and SBOM Generation (P2)
5. **Month 2**: Explore Sandboxing, Add Fuzzing, Add Performance Testing (P2/P3)

## Metrics to Track

| Metric | Current | Target | Status |
|--------|---------|--------|--------|
| **Security Gates Passing** | 0/5 | 5/5 | Blocked |
| **Quality Gates Passing** | 3/8 | 8/8 | In Progress |
| **Test Coverage** | Unknown | 70%+ | Not Measured |
| **Cyclomatic Complexity** | < 30 | < 15 | Exceeds Target |
| **Known Vulnerabilities** | Unknown | 0 | Not Scanned |
| **Secrets in Repo** | Unknown | 0 | Not Scanned |

## Related Documentation

- **SECURITY.md** - Detailed security threat analysis and controls
- **QUALITY_GATES.md** - Current quality gate status
- **CONTEXT.md** - Project architecture and overview

## Related

- [[docs/HUB-DOCS.md]]
- [[docs/CONTEXT.md]]
- [[docs/QUALITY_GATES.md]]
- [[docs/SECURITY.md]]
