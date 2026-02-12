# Sprint 04 Tasks

## Task 1: Cursor Content Generator ✓

**ID:** P0-1
**Priority:** P0
**Status:** COMPLETED

INPUT: Research on Cursor rules format
OUTPUT: `cli/internal/targets/cursor_content.go` with:
- `generateCursorRules() []byte` - Enhanced .cursorrules content
- `generateCursorMDC() []byte` - .cursor/rules/openkit.mdc content

VERIFY:
```bash
go build ./...  # ✓ Passed
```

---

## Task 2: Update buildCursor() ✓

**ID:** P0-2
**Priority:** P0
**Depends:** P0-1
**Status:** COMPLETED

INPUT: `cli/internal/targets/targets.go` buildCursor function
OUTPUT: Enhanced function that generates:
- `.cursorrules` (improved)
- `.cursor/rules/openkit.mdc`
- `.cursor/skills/**`

VERIFY:
```bash
go test ./internal/targets/...  # ✓ Passed
```

---

## Task 3: Codex Content Generator ✓

**ID:** P0-3
**Priority:** P0
**Status:** COMPLETED

INPUT: Research on Codex AGENTS.md and rules format
OUTPUT: `cli/internal/targets/codex_content.go` with:
- `generateAgentsMD() []byte` - Comprehensive AGENTS.md
- `generateCodexRules() []byte` - .codex/rules/openkit.rules

VERIFY:
```bash
go build ./...  # ✓ Passed
```

---

## Task 4: Update buildCodex() ✓

**ID:** P0-4
**Priority:** P0
**Depends:** P0-3
**Status:** COMPLETED

INPUT: `cli/internal/targets/targets.go` buildCodex function
OUTPUT: Enhanced function that generates:
- `AGENTS.md` (comprehensive)
- `.codex/rules/openkit.rules`
- `.agents/skills/**`

VERIFY:
```bash
go test ./internal/targets/...  # ✓ Passed
```

---

## Task 5: Cursor Unit Tests ✓

**ID:** P1-1
**Priority:** P1
**Depends:** P0-1
**Status:** COMPLETED

INPUT: cursor_content.go
OUTPUT: `cli/internal/targets/cursor_content_test.go`

VERIFY:
```bash
go test ./internal/targets/... -run Cursor  # ✓ Passed (12 tests)
```

---

## Task 6: Codex Unit Tests ✓

**ID:** P1-2
**Priority:** P1
**Depends:** P0-3
**Status:** COMPLETED

INPUT: codex_content.go
OUTPUT: `cli/internal/targets/codex_content_test.go`

VERIFY:
```bash
go test ./internal/targets/... -run Codex  # ✓ Passed (10 tests)
```

---

## Task 7: Cursor Integration Test ✓

**ID:** P1-3
**Priority:** P1
**Depends:** P0-2
**Status:** COMPLETED

INPUT: Updated buildCursor
OUTPUT: Test in `cli/internal/cli/agent_targets_integration_test.go`

VERIFY:
```bash
go test ./internal/cli/... -run TestCursorSync  # ✓ Passed (147 files created)
```

---

## Task 8: Codex Integration Test ✓

**ID:** P1-4
**Priority:** P1
**Depends:** P0-4
**Status:** COMPLETED

INPUT: Updated buildCodex
OUTPUT: Test in `cli/internal/cli/agent_targets_integration_test.go`

VERIFY:
```bash
go test ./internal/cli/... -run TestCodexSync  # ✓ Passed (147 files created)
```

---

## Task 9: Doctor Enhancements ✓

**ID:** P2-1, P2-2
**Priority:** P2
**Depends:** P0-2, P0-4
**Status:** COMPLETED

INPUT: `cli/internal/cli/agent_targets.go` runAgentDoctor
OUTPUT: Enhanced checks for Cursor and Codex

VERIFY:
```bash
openkit cursor doctor  # ✓ Checks .cursorrules, .cursor/rules/, .cursor/skills/
openkit codex doctor   # ✓ Checks AGENTS.md, .codex/rules/, .agents/skills/
```

---

## Task 10: Documentation ✓

**ID:** P2-3, P2-4
**Priority:** P2
**Depends:** All P0, P1 tasks
**Status:** COMPLETED

INPUT: Implementation
OUTPUT:
- Updated `cli/docs/agent-compat/agents/cursor.md` ✓
- Updated `cli/docs/agent-compat/agents/codex.md` ✓

VERIFY: Review docs for accuracy ✓

---

## Task 11: Harden Security Scan

**ID:** P0-5
**Priority:** P0
**Status:** COMPLETED

INPUT:
- `internal/templates/base/skills/vulnerability-scanner/scripts/security_scan.py`
- `.opencode/skills/vulnerability-scanner/scripts/security_scan.py` (generated copy, ignored in git)
OUTPUT:
- Scanner skips internal tooling directories (reduces self-scan noise)
- Scanner exits non-zero on critical/high findings

VERIFY:
```bash
python3 .opencode/skills/vulnerability-scanner/scripts/security_scan.py . --output summary  # exit=0
python3 .opencode/scripts/checklist.py .  # PASS
```

---

## Task 12: Remove blueprint references

**ID:** P0-6
**Priority:** P0
**Status:** COMPLETED

INPUT:
- Repo content containing blueprint alias mentions
OUTPUT:
- Updated skills/templates/scripts that reference only shipped OpenKit artifacts

VERIFY:
```bash
rg -n "\\bblueprints\\b" .
python3 .opencode/scripts/checklist.py .
```

---

## Task 13: Guardrail against reintroduction

**ID:** P0-7
**Priority:** P0
**Depends:** P0-6
**Status:** COMPLETED

INPUT: Repo policies and CI/check scripts
OUTPUT: A deterministic check that fails if blueprint alias mentions are reintroduced in tracked sources

VERIFY:
```bash
python3 .opencode/scripts/checklist.py .
```

---

## Task 14: Requirements docs

**ID:** P2-5
**Priority:** P2
**Status:** COMPLETED

INPUT: Cleanup plan
OUTPUT:
- `docs/requirements/remove-blueprints-references/*`

VERIFY: Review docs for accuracy

## Related

- [[docs/sprint/Sprint-04/HUB-SPRINT-04.md]]
- [[docs/sprint/Sprint-04/SPRINT_GOAL.md]]
- [[docs/sprint/Sprint-04/BACKLOG.md]]
- [[docs/requirements/cursor-codex-targets/PLAN.md]]
