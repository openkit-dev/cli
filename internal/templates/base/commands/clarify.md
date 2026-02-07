---
description: Resolve ambiguities in the feature specification.
subtask: false
---

# /clarify - Specification Clarification

$ARGUMENTS

## Purpose

Resolve ambiguities and critical gaps in the feature spec before planning.

## Inputs

Read from `docs/requirements/<feature>/`:

- `PROBLEM_STATEMENT.md`
- `USER_STORIES.md`
- `ACCEPTANCE_CRITERIA.md`
- `RISKS.md`

## Workflow

1. Scan for `NEEDS CLARIFICATION` markers and unclear requirements.
2. Ask targeted questions using the question tool (max 5 at a time).
3. Update the spec files with explicit answers.
4. Re-check for remaining ambiguities.

## Rules

- Only ask what materially changes scope, security, or UX.
- Do not add implementation details.
- Use the question tool for all multi-option questions.

## STOP POINT

After clarifications are applied:

> Use the question tool to ask "Clarifications applied. Proceed to /plan?"
