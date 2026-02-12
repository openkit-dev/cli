
## Risks

- Tightening exit codes can break existing workflows that assume the scan is informational only.
- Skipping internal tooling directories may hide real issues if a project stores runtime code or secrets there.

## Mitigations

- Only skip directories that are clearly tooling/config for agent frameworks.
- Keep medium/low findings informational (exit 0) to avoid excessive false failures.

## Related

- [[docs/requirements/security-scan-hardening/HUB-SECURITY-SCAN-HARDENING.md]]
- [[docs/requirements/security-scan-hardening/PROBLEM_STATEMENT.md]]
- [[docs/requirements/security-scan-hardening/PLAN.md]]
- [[docs/sprint/Sprint-06/RISK_REGISTER.md]]
