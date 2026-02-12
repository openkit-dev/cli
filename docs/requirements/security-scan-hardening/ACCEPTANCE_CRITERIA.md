
## Acceptance Criteria

- Running `python3 .opencode/skills/vulnerability-scanner/scripts/security_scan.py .` exits with:
  - `0` when there are no critical/high findings
  - `1` when there is at least one high finding
  - `2` when there is at least one critical finding
- Running the repository template version `python3 internal/templates/base/skills/vulnerability-scanner/scripts/security_scan.py .` follows the same exit code rules.
- The scan does not report findings solely from scanning:
  - `.opencode/`
  - `.agents/`
  - `.cursor/`
  - `.gemini/`
  - `.tmp/`
- `.opencode/scripts/checklist.py .` reports FAIL when the security scan exits non-zero.

## Related

- [[docs/requirements/security-scan-hardening/HUB-SECURITY-SCAN-HARDENING.md]]
- [[docs/requirements/security-scan-hardening/PROBLEM_STATEMENT.md]]
- [[docs/requirements/security-scan-hardening/PLAN.md]]
- [[docs/sprint/Sprint-06/TASKS.md]]
