---
description: Run quality assurance checks - unit tests, E2E, security scans, performance audits
subtask: true
---

**TESTING MODE ACTIVATED**

**Target:** $ARGUMENTS (all, unit, integration, e2e, security, performance, coverage, ux)

**If $ARGUMENTS is empty:**
Use the question tool to ask:
```javascript
question({
  questions: [{
      question: "Which tests to run?",
      header: "Test Type",
      options: [
        { label: "All", description: "Full test suite" },
        { label: "Unit", description: "Unit tests only" },
        { label: "E2E", description: "End-to-end tests" },
        { label: "Security", description: "Security scans" }
      ]
    }]
})
```

**Testing Protocol (by priority):**

## P0 - Critical Checks

### 1. Lint & Type Check
```bash
npm run lint
npx tsc --noEmit
```
- Status: [PASS/FAIL]

### 2. Security Scan
!`python .opencode/skills/vulnerability-scanner/scripts/security_scan.py .`
- Status: [PASS/WARN/FAIL]
- Findings: [details]

## P1 - Quality Checks

### 3. Unit Tests
```bash
npm test
# ou
pytest
```
- Status: [PASS/FAIL]
- Coverage: [XX%]
- Failures: [if any]

### 4. UX/Accessibility Audit (if frontend)
!`python .opencode/skills/frontend-design/scripts/ux_audit.py .`
- Status: [PASS/WARN]
- Issues: [if any]

## P2 - Performance

### 5. Build Verification
```bash
npm run build
```
- Status: [PASS/FAIL]
- Warnings: [if any]

### 6. Lighthouse Audit (requires running server)
If a server is detected at http://localhost:3000:
!`python .opencode/skills/performance-profiling/scripts/lighthouse_audit.py http://localhost:3000`
- Score: [XX/100]
- Web Vitals: [LCP, FID, CLS]

## P3 - E2E Tests

### 7. Playwright E2E (requires server)
!`python .opencode/skills/webapp-testing/scripts/playwright_runner.py http://localhost:3000 --screenshot`
- Status: [PASS/FAIL]
- Screenshots: [path]

---

## Final Report

```markdown
## Test Results Summary

| Check | Status | Details |
|-------|--------|---------|
| Lint | PASS/FAIL | ... |
| Type Check | PASS/FAIL | ... |
| Security | PASS/WARN/FAIL | ... |
| Unit Tests | PASS/FAIL | XX% coverage |
| UX Audit | PASS/WARN | ... |
| Build | PASS/FAIL | ... |
| Lighthouse | PASS/WARN | XX/100 |
| E2E Tests | PASS/FAIL | ... |

### Action Items
- [ ] [If there are failures, list required fixes]
```

** IMPORTANT:** Do not mark checks as passing without actually running the commands!
