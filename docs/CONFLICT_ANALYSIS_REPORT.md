# Conflict Analysis Report - internal/templates/base/

**Date:** 2026-02-09
**Scope:** Analysis of agent orchestration conflicts in `internal/templates/base/`
**Status:** All conflicts resolved (P0-P3)

---

## Executive Summary

Analysis of the `internal/templates/base/` directory identified **7 conflicts** between commands, agents, skills, and orchestration patterns. All conflicts have been resolved through systematic fixes across 14 files.

---

## Conflicts Identified and Resolved

### CONFLICT 1: `/plan` vs `/tasks` (P0 - RESOLVED)

**Problem:** Both commands were creating `TASKS.md`, causing confusion and potential overwrites.

| Aspect | Before | After |
|--------|--------|-------|
| `/plan` creates TASKS.md | YES | NO |
| `/tasks` creates TASKS.md | YES | YES (PRIMARY) |
| Ownership clarity | Ambiguous | Clear |

**Solution:**
- Removed TASKS.md creation from `/plan`
- Clarified `/tasks` as the PRIMARY and ONLY command for TASKS.md
- Updated documentation to reflect ownership

**Files Modified:**
- `commands/plan.md`
- `commands/tasks.md`
- `prompts/project-planner.md`

---

### CONFLICT 2: `/specify` Redundancy (P0 - RESOLVED)

**Problem:** `/plan` was creating all specification artifacts, making `/specify` redundant.

**Solution:**
- Added SDD Gate Check (Phase 1.5) to `/plan` that verifies if `/specify` was executed
- If specification is missing, `/plan` prompts user to run `/specify` first
- Clear separation: `/specify` creates spec artifacts, `/plan` creates plan artifacts

**Files Modified:**
- `commands/plan.md`
- `rules/MASTER.md`

---

### CONFLICT 3: TodoList IDs Inconsistent (P1 - RESOLVED)

**Problem:** Each command defined its own TodoList ID scheme, causing confusion when commands run in sequence.

| Command | Before | After |
|---------|--------|-------|
| `/plan` | `plan-01-stack` | `sdd-plan-01-gate` |
| `/tasks` | (none) | `sdd-tasks-01-read` |
| `/impl` | `impl-01-precheck` | `sdd-impl-01-precheck` |
| `/engineer` | `phase1-analysis` | `orch-01-analysis` |

**Solution:**
- Created standard ID schema: `{workflow}-{phase}-{step}`
- Added "Standard TodoList ID Schema" section to `TODOLIST_PROTOCOL.md`
- Updated all commands to use standardized IDs

**Files Modified:**
- `commands/plan.md`
- `commands/tasks.md`
- `commands/impl.md`
- `commands/engineer.md`
- `commands/brainstorm.md`
- `prompts/orchestrator.md`
- `rules/TODOLIST_PROTOCOL.md`

---

### CONFLICT 4: STOP Points Inconsistent (P1 - RESOLVED)

**Problem:** Some commands used plain text STOP points, others used `question` tool inconsistently.

**Solution:**
- Standardized all STOP points to use `question` tool with proper structure
- Consistent options: "Proceed", "Review", "Adjust"
- Clear header naming: "Phase X Complete"

**Files Modified:**
- `commands/plan.md`
- `commands/tasks.md`
- `commands/impl.md`
- `commands/engineer.md`

---

### CONFLICT 5: `/analyze` vs `/checklist` (P2 - RESOLVED)

**Problem:** Both commands validated artifacts with overlapping purposes, confusing users.

| Aspect | `/checklist` | `/analyze` |
|--------|--------------|------------|
| Purpose | Quick pre-flight | Deep analysis |
| Output | Pass/Fail checklist | Detailed analysis + remediation |
| Time | Fast (1-2 min) | Thorough (5-10 min) |
| Use Case | Before `/impl` | QA review, complex prep |

**Solution:**
- Added clear differentiation tables to both commands
- Updated descriptions to clarify when to use each

**Files Modified:**
- `commands/analyze.md`
- `commands/checklist.md`

---

### CONFLICT 6: `/context` vs `/brainstorm` (P2 - RESOLVED)

**Problem:** Discovery Gate accepted both commands interchangeably, but they serve different purposes.

| Aspect | `/context` | `/brainstorm` |
|--------|------------|---------------|
| Required? | MANDATORY | OPTIONAL |
| Output | Technical docs (CONTEXT.md, SECURITY.md) | Options analysis, trade-offs |
| Use Case | Always before /specify | When scope unclear |

**Solution:**
- Marked `/context` as MANDATORY in documentation
- Marked `/brainstorm` as OPTIONAL
- Clarified that `/brainstorm` does NOT replace `/context`

**Files Modified:**
- `commands/context.md`
- `commands/brainstorm.md`
- `rules/MASTER.md`
- `prompts/orchestrator.md`

---

### CONFLICT 7: Phase Naming Inconsistent (P3 - RESOLVED)

**Problem:** Different files used different phase naming conventions.

| File | Before | After |
|------|--------|-------|
| `project-planner.md` | Phase 1-4 + X | Phase 0-5 |
| `orchestrator.md` | Phase 1-3 | Phase 0-5 |
| `engineer.md` | Phase 1-2 | Phase 0-5 |

**Solution:**
- Created standard 6-phase workflow (Phase 0-5)
- Added "Standard Phase Workflow" table to MASTER.md
- Updated all agent/command files to use consistent naming

**Files Modified:**
- `prompts/project-planner.md`
- `prompts/orchestrator.md`
- `commands/engineer.md`
- `rules/MASTER.md`

---

## New Standard Phase Workflow

| Phase | Name | Command(s) | Output | Code? |
|-------|------|------------|--------|-------|
| **0** | Discovery | `/context` (MANDATORY), `/brainstorm` (optional) | `docs/CONTEXT.md`, decisions | NO |
| **1** | Specification | `/specify`, `/clarify` | `docs/requirements/<feature>/` | NO |
| **2** | Planning | `/plan` | `PLAN.md`, `SPRINT_GOAL.md`, `BACKLOG.md` | NO |
| **3** | Task Breakdown | `/tasks` | `TASKS.md` | NO |
| **4** | Implementation | `/impl` | Working code | YES |
| **5** | Verification | `/test`, scripts | Verified project | Scripts |

**Flow Diagram:**
```
Phase 0: /context (MANDATORY) --> /brainstorm (optional)
                |
                v
Phase 1: /specify --> /clarify (optional)
                |
                v
Phase 2: /plan --> STOP (approval required)
                |
                v
Phase 3: /tasks --> STOP (approval required)
                |
                v
Phase 4: /impl --> parallel agents (P0->P1->P2->P3)
                |
                v
Phase 5: /test, /checklist scripts --> /doc
```

---

## New Standard TodoList ID Schema

| Prefix | Command | Examples |
|--------|---------|----------|
| `sdd-spec-` | `/specify` | `sdd-spec-01-problem`, `sdd-spec-02-stories` |
| `sdd-plan-` | `/plan` | `sdd-plan-01-gate`, `sdd-plan-02-stack` |
| `sdd-tasks-` | `/tasks` | `sdd-tasks-01-read`, `sdd-tasks-02-breakdown` |
| `sdd-impl-` | `/impl` | `sdd-impl-01-precheck`, `sdd-impl-02-p0-foundation` |
| `orch-` | `/engineer` | `orch-01-analysis`, `orch-02-planning` |
| `brainstorm-` | `/brainstorm` | `brainstorm-01-context`, `brainstorm-02-options` |

---

## Files Modified Summary

### Commands (8 files)
1. `internal/templates/base/commands/plan.md`
2. `internal/templates/base/commands/tasks.md`
3. `internal/templates/base/commands/impl.md`
4. `internal/templates/base/commands/engineer.md`
5. `internal/templates/base/commands/brainstorm.md`
6. `internal/templates/base/commands/context.md`
7. `internal/templates/base/commands/analyze.md`
8. `internal/templates/base/commands/checklist.md`

### Prompts (2 files)
9. `internal/templates/base/prompts/project-planner.md`
10. `internal/templates/base/prompts/orchestrator.md`

### Rules (2 files)
11. `internal/templates/base/rules/MASTER.md`
12. `internal/templates/base/rules/TODOLIST_PROTOCOL.md`

---

## Recommendations for Future Development

1. **New Commands:** Follow the Standard Phase Workflow and TodoList ID Schema
2. **New Agents:** Use the standardized phase naming (Phase 0-5)
3. **STOP Points:** Always use `question` tool with consistent options
4. **Artifact Ownership:** Document which command creates which artifact
5. **Testing:** After adding new commands, verify they don't conflict with existing ones

---

## Appendix: Fix Priority Reference

| Priority | Description | Criteria |
|----------|-------------|----------|
| **P0** | Critical | Causes workflow failures, data loss, or confusion |
| **P1** | High | Inconsistency that affects user experience |
| **P2** | Medium | Clarity and documentation issues |
| **P3** | Low | Cosmetic or minor alignment issues |

---

---

## Additional Artifacts Created

### Glossary ([[GLOSSARY.md]])

A comprehensive glossary was created defining standard terminology for:
- Agent system concepts (Agent, Orchestrator, Skill)
- Workflow phases (Discovery, Specification, Planning, etc.)
- Gates and checkpoints (SDD Gate, Discovery Gate, EXIT GATE)
- Command reference
- TodoList ID schema
- Acronym definitions

### Templates Updated

The following templates were updated to include Glossary sections:
- `templates/DOCS-CONTEXT.md` - Added Terminology section
- `templates/SDD-ProblemStatement.md` - Added Glossary section
- `templates/SDD-Plan.md` - Added Glossary and References sections
- `commands/doc.md` - Added Glossary Updates workflow

---

**Report Generated:** 2026-02-09
**Total Fixes Applied:** 7 conflicts across 12 files + Glossary system

## Related

- [[HUB-DOCS.md]]
- [[GLOSSARY.md]]
- [[CONTEXT.md]]
