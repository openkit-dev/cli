---
description: Create a quality checklist for spec and plan readiness.
subtask: true
---

# /checklist - Spec/Plan Checklist

$ARGUMENTS

## Purpose

Generate a quality checklist to validate that specification and plan are ready for implementation.

## Output

- `docs/requirements/<feature>/checklists/requirements.md`

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
