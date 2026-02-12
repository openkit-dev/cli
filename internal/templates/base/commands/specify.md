---
description: Create a feature specification focused on what and why.
subtask: false
---

# /specify - Feature Specification

$ARGUMENTS

## Purpose

Create a feature specification that captures user intent, scenarios, and acceptance criteria.
This command MUST focus on what and why. No implementation details.

## If $ARGUMENTS is empty

Use the question tool to ask for a feature description.

```javascript
question({
  questions: [{
      question: "Describe the feature you want to build (what and why, no tech stack).",
      header: "Feature Description"
    }]
})
```

## Outputs (required)

Create or update:

- `docs/requirements/<feature>/PROBLEM_STATEMENT.md`
- `docs/requirements/<feature>/USER_STORIES.md`
- `docs/requirements/<feature>/ACCEPTANCE_CRITERIA.md`
- `docs/requirements/<feature>/RISKS.md`

Ensure `docs/requirements/<feature>/` exists before writing.

## Templates

- `.opencode/templates/SDD-ProblemStatement.md`
- `.opencode/templates/SDD-UserStories.md`
- `.opencode/templates/SDD-AcceptanceCriteria.md`
- `.opencode/templates/SDD-Risks.md`

## Workflow

1. Confirm feature name and scope.
2. Create `docs/requirements/<feature>/` if missing.
3. Fill templates for Problem Statement, User Stories, Acceptance Criteria, and Risks.
4. Write prioritized user stories with independent acceptance scenarios.
5. Capture functional requirements and edge cases.
6. Record measurable success criteria (tech-agnostic).
7. Note risks and assumptions.
8. Add Obsidian-compatible links between generated artifacts and include `## Related`.
9. Follow canonical docs filenames from `.opencode/rules/DOCS_FILE_GLOSSARY.md`.

## Rules

- Do not include tech stack, APIs, or file structure.
- If ambiguity is critical, mark it as `NEEDS CLARIFICATION` and defer to `/clarify`.
- Use the question tool for any multi-option decision.

## STOP POINT

After writing the spec:

> Use the question tool to ask "Specification recorded in docs/requirements/. Proceed to /clarify or /plan?"
