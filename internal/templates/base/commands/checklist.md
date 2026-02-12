---
description: Quick pre-flight validation checklist (pass/fail).
subtask: true
---

# /checklist - Quick Pre-Flight Check (Pass/Fail)

$ARGUMENTS

## Purpose

**Quick validation** that produces a pass/fail checklist. Use before starting implementation.

**When to use /checklist vs /analyze:**
| Command | Purpose | Output | Time |
|---------|---------|--------|------|
| `/checklist` | Quick pre-flight validation | Pass/Fail checklist | Fast (1-2 min) |
| `/analyze` | Deep consistency analysis | Detailed analysis + remediation | Thorough (5-10 min) |

**Use /checklist when:**
- Quick sanity check before `/impl`
- Verifying basic completeness
- CI/CD gate validation
- Daily standup verification

## Output

- `docs/requirements/<feature>/CHECKLIST.md`

## Template

- `.opencode/templates/SDD-Checklist.md`

## Checklist Items

- No implementation details in spec
- Requirements are testable and unambiguous
- All acceptance criteria are measurable
- Edge cases identified
- Risks documented
- Plan includes data model and contracts (if applicable)
- Research captured for unknowns (if applicable)
- Quickstart documented for non-trivial setup (if applicable)
- Tasks trace to user stories
- Verification steps defined

## Workflow

1. Create checklist file if missing.
2. Mark pass/fail based on current artifacts.
3. Note required fixes with references.

## Rules

- Do not modify code.
- Use file references for evidence.
- Use the question tool if multiple fix paths exist.

## STOP POINT

After checklist:

```javascript
question({
  questions: [{
      question: "Checklist ready. Apply fixes now or proceed to /tasks?",
      header: "Next Step",
      options: [
        { label: "Apply fixes", description: "Revise spec/plan before tasks" },
        { label: "Proceed to /tasks", description: "Generate tasks now" }
      ]
    }]
})
```
