
## Problem

The validation flow reports the security scan as passing even when it detects critical/high issues.

Root causes:

- `.opencode/skills/vulnerability-scanner/scripts/security_scan.py` prints findings but always exits with status code 0.
- In this repository, the source-of-truth for shipped scripts lives under `internal/templates/base/**` and is later materialized into `.opencode/**` by sync/build steps.
- The scanner can also scan internal tooling packs (e.g. `.opencode/`, `.agents/`, `.cursor/`, `.gemini/`, `.tmp/`) and match its own regex/pattern definitions, producing false positives.

This creates a false sense of safety in `.opencode/scripts/checklist.py` because it treats the scan as PASS based only on exit code.

## Goal

- Make `security_scan.py` fail (non-zero exit code) when critical/high findings exist.
- Reduce self-scan noise by skipping internal tooling directories and template packs.

## Non-goals

- Do not redesign the security ruleset or add external dependencies.
- Do not change other checklist steps.

## Related

- [[docs/requirements/security-scan-hardening/HUB-SECURITY-SCAN-HARDENING.md]]
- [[docs/requirements/security-scan-hardening/USER_STORIES.md]]
- [[docs/requirements/security-scan-hardening/ACCEPTANCE_CRITERIA.md]]
- [[docs/requirements/security-scan-hardening/PLAN.md]]
