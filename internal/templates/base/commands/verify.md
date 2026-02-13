---
description: Quality verification - tests, lint, security, performance
subtask: true
---

# /verify - Quality Verification

$ARGUMENTS

## Overview

Run comprehensive quality checks on the codebase. This command unifies (old commands):
- `/test` (old) - Test execution
- `/checklist` (old) - Quality checklist
- `/analyze` (old) - Cross-artifact analysis

## If $ARGUMENTS is empty

Use the question tool:

```javascript
question({
  questions: [{
      question: "Which verification scope?",
      header: "Verification Scope",
      options: [
        { label: "All", description: "Full verification suite" },
        { label: "Quick", description: "Lint + Security only" },
        { label: "Custom", description: "Specific checks" }
      ]
    }]
})
```

---

## Verification Protocol (by priority)

### P0 - Critical Checks

#### 1. Lint & Type Check
```bash
npm run lint
npx tsc --noEmit
```
- Status: [PASS/FAIL]

#### 2. Security Scan
!`python .opencode/skills/vulnerability-scanner/scripts/security_scan.py .`
- Status: [PASS/WARN/FAIL]
- Findings: [details]

### P1 - Quality Checks

#### 3. Unit Tests
```bash
npm test
# or
pytest
```
- Status: [PASS/FAIL]
- Coverage: [XX%]
- Failures: [if any]

#### 4. UX/Accessibility Audit (if frontend)
!`python .opencode/skills/frontend-design/scripts/ux_audit.py .`
- Status: [PASS/WARN]
- Issues: [if any]

### P2 - Performance

#### 5. Build Verification
```bash
npm run build
```
- Status: [PASS/FAIL]
- Warnings: [if any]

#### 6. Lighthouse Audit (requires running server)
If a server is detected at http://localhost:3000:
!`python .opencode/skills/performance-profiling/scripts/lighthouse_audit.py http://localhost:3000`
- Score: [XX/100]
- Web Vitals: [LCP, FID, CLS]

### P3 - E2E Tests

#### 7. Playwright E2E (requires server)
!`python .opencode/skills/webapp-testing/scripts/playwright_runner.py http://localhost:3000 --screenshot`
- Status: [PASS/FAIL]
- Screenshots: [path]

---

## All-in-One Command

```bash
python .opencode/scripts/verify_all.py . --url http://localhost:3000
```

---

## Final Report

```markdown
## Verification Results Summary

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

---

## STOP Point

```javascript
question({
  questions: [{
      header: "Verification Complete",
      question: "Results: [PASS/FAIL]. Proceed to deploy (/deploy)?",
      options: [
        { label: "Yes, proceed to /deploy", description: "Deploy to target environment" },
        { label: "Fix issues first", description: "Address verification failures" }
      ]
    }]
})
```

**IMPORTANT:** Do not mark checks as passing without actually running the commands!
