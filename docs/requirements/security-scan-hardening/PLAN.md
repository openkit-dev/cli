
## Plan

1. Update `security_scan.py` (template + generated copy) to:
   - Skip internal tooling directories by default.
   - Skip `internal/templates/**` without skipping `internal/**` as a whole.
   - Exit non-zero on critical/high findings.
2. Verify behavior locally:
   - Run `python3 internal/templates/base/skills/vulnerability-scanner/scripts/security_scan.py . --output summary` and confirm exit code.
   - Run `python3 .opencode/skills/vulnerability-scanner/scripts/security_scan.py . --output summary` (if present) and confirm exit code.
   - Run `python3 .opencode/scripts/checklist.py .` and confirm Security Scan FAILs on non-zero exit.

## Related

- [[requirements/security-scan-hardening/HUB-SECURITY-SCAN-HARDENING.md]]
- [[requirements/security-scan-hardening/PROBLEM_STATEMENT.md]]
- [[requirements/security-scan-hardening/USER_STORIES.md]]
- [[sprint/Sprint-06/TASKS.md]]
