---
description: Deep cross-artifact consistency analysis with remediation actions.
subtask: true
---

# /analyze - Deep Analysis (Gaps & Remediation)

$ARGUMENTS

## Purpose

**Deep dive analysis** that validates consistency across ALL artifacts and produces actionable remediation.

**When to use /analyze vs /checklist:**
| Command | Purpose | Output | Time |
|---------|---------|--------|------|
| `/checklist` | Quick pre-flight validation | Pass/Fail checklist | Fast (1-2 min) |
| `/analyze` | Deep consistency analysis | Detailed analysis + remediation | Thorough (5-10 min) |

**Use /analyze when:**
- You need to understand WHY something is inconsistent
- You want specific remediation actions
- You're preparing for a complex implementation
- QA review before merging to main

## Inputs

Read from:

- `docs/requirements/<feature>/PROBLEM_STATEMENT.md`
- `docs/requirements/<feature>/USER_STORIES.md`
- `docs/requirements/<feature>/ACCEPTANCE_CRITERIA.md`
- `docs/requirements/<feature>/DATA_CONTRACTS.md` (if data is involved)
- `docs/requirements/<feature>/PLAN.md`
- `docs/sprint/Sprint-XX/TASKS.md`

## Output

Write analysis to:

- `docs/requirements/<feature>/ANALYSIS.md`

## Template

- `.opencode/templates/SDD-Analysis.md`

## Analysis Dimensions

| Dimension | What to Check |
|-----------|---------------|
| **Requirement Traceability** | Each requirement maps to plan sections |
| **Story Coverage** | Each user story has one or more tasks |
| **Contract Completeness** | DATA_CONTRACTS.md exists if data involved |
| **Test Coverage** | Each AC has corresponding test task |
| **Risk Mitigation** | Each risk has mitigation in plan |

## Workflow

1. Trace each requirement to plan sections.
2. Trace each user story to one or more tasks.
3. Validate auxiliary artifacts when applicable:
   - `RESEARCH.md` when unknowns or dependencies exist
   - `DATA_CONTRACTS.md` when data entities exist
   - `QUICKSTART.md` when setup or verification steps are non-trivial
4. Identify missing tests, missing contracts, or unclear verification steps.
5. **Produce remediation actions** with priority and responsible command.

## Rules

- No code changes.
- Use file references for evidence.
- Use the question tool if multiple fix options exist.

## STOP POINT

After analysis:

```javascript
question({
  questions: [{
      question: "Analysis complete. Apply fixes via /plan or /tasks?",
      header: "Next Step",
      options: [
        { label: "Apply via /plan", description: "Update spec/plan artifacts" },
        { label: "Apply via /tasks", description: "Update tasks only" }
      ]
    }]
})
```
