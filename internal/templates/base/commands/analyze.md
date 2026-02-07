---
description: Cross-artifact consistency analysis for spec, plan, and tasks.
subtask: true
---

# /analyze - Spec/Plan/Tasks Analysis

$ARGUMENTS

## Purpose

Validate consistency across requirements, plan, and tasks. Identify gaps, contradictions, or missing coverage.

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

- `docs/requirements/<feature>/analysis.md`

## Workflow

1. Trace each requirement to plan sections.
2. Trace each user story to one or more tasks.
3. Validate auxiliary artifacts when applicable:
   - `research.md` when unknowns or dependencies exist
   - `DATA_CONTRACTS.md` when data entities exist
   - `contracts/` when APIs/events/integrations exist
   - `quickstart.md` when setup or verification steps are non-trivial
4. Identify missing tests, missing contracts, or unclear verification steps.
4. List remediation actions.

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
