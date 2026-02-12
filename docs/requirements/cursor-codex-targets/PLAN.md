# Implementation Plan: Cursor + Codex Targets

## Overview

Improve the Cursor and Codex targets to provide richer configuration files, following each agent's conventions.

## Phase 1: Cursor Target Enhancement

### Task 1.1: Update buildCursor() in targets.go

**Input:** Current minimal implementation
**Output:** Enhanced function that generates:
- `.cursorrules` (improved content)
- `.cursor/rules/openkit.mdc`
- `.cursor/skills/**` (copied from base)

**Implementation:**
```go
func buildCursor(agent *agents.Agent, cliVersion string) (DesiredResult, error) {
    var files []syncer.DesiredFile
    
    // 1. Enhanced .cursorrules
    cursorrules := generateCursorRules()
    files = append(files, syncer.DesiredFile{...})
    
    // 2. Modular rule file
    mdcContent := generateCursorMDC()
    files = append(files, syncer.DesiredFile{
        OutputPath: ".cursor/rules/openkit.mdc",
        ...
    })
    
    // 3. Skills
    skills, _ := syncer.DesiredFromEmbeddedSubdir(templates.BaseFS(), "base/skills", ".cursor/skills")
    files = append(files, skills...)
    
    return DesiredResult{...}, nil
}
```

### Task 1.2: Create cursor_content.go

**Input:** Content templates
**Output:** Functions to generate Cursor-specific content

```go
// cli/internal/targets/cursor_content.go
func generateCursorRules() []byte
func generateCursorMDC() []byte
```

### Task 1.3: Unit Tests

**Input:** cursor_content.go functions
**Output:** `cursor_content_test.go` with tests

---

## Phase 2: Codex Target Enhancement

### Task 2.1: Update buildCodex() in targets.go

**Input:** Current minimal implementation
**Output:** Enhanced function that generates:
- `AGENTS.md` (comprehensive)
- `.codex/rules/openkit.rules`
- `.agents/skills/**` (already implemented)

### Task 2.2: Create codex_content.go

**Input:** Content templates
**Output:** Functions to generate Codex-specific content

```go
// cli/internal/targets/codex_content.go
func generateAgentsMD() []byte
func generateCodexRules() []byte
```

### Task 2.3: Unit Tests

**Input:** codex_content.go functions
**Output:** `codex_content_test.go` with tests

---

## Phase 3: Doctor Command Enhancements

### Task 3.1: Update agent_targets.go doctor checks

**Input:** Current doctor checks
**Output:** Enhanced checks for Cursor and Codex

---

## Phase 4: Integration Tests

### Task 4.1: Cursor Integration Test

**Input:** Test framework
**Output:** `TestCursorSync_WritesFilesAndManagedState`

### Task 4.2: Codex Integration Test

**Input:** Test framework
**Output:** `TestCodexSync_WritesFilesAndManagedState`

---

## Phase 5: Documentation

### Task 5.1: Update agent-compat docs

**Input:** Current docs
**Output:** Updated `cursor.md` and `codex.md`

---

## Verification

After implementation:
```bash
# Run unit tests
go test ./internal/targets/...

# Run integration tests
go test ./internal/cli/... -run "TestCursorSync|TestCodexSync"

# Manual verification
cd /tmp && mkdir test-cursor && cd test-cursor
openkit cursor sync
ls -la .cursor/
cat .cursorrules

cd /tmp && mkdir test-codex && cd test-codex
openkit codex sync
ls -la .codex/
cat AGENTS.md
```

## Related

- [[requirements/cursor-codex-targets/HUB-CURSOR-CODEX-TARGETS.md]]
- [[requirements/cursor-codex-targets/PROBLEM_STATEMENT.md]]
- [[requirements/cursor-codex-targets/USER_STORIES.md]]
- [[sprint/Sprint-04/TASKS.md]]
