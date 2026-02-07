---
trigger: always_on
---

# MASTER RULESET - OpenKit Agent System

> Single source of mandatory rules for all agents, skills, and commands.

---

## Context and Language

- Respond in the user's language; keep code/comments in English.
- Follow clean code principles: SRP, DRY, KISS, functions <= 20 lines, max 3 args, use guard clauses.
- Before changing any file, understand dependencies via CODEBASE.md or local references.

## Response Style (MANDATORY)

- Be objective, technical, consistent, and direct.
- Avoid embellishment, filler, and marketing language.
- Prefer short, information-dense sentences.

---

## SDD Gate (Mandatory)

Before any implementation, the following artifacts MUST exist:

- `docs/requirements/<feature>/PROBLEM_STATEMENT.md`
- `docs/requirements/<feature>/USER_STORIES.md`
- `docs/requirements/<feature>/ACCEPTANCE_CRITERIA.md`
- `docs/requirements/<feature>/RISKS.md`
- `docs/requirements/<feature>/PLAN.md`

If any are missing, STOP and direct the user to run /specify, /clarify, and /plan first.

---

## Sprint Documentation

1. Locate `docs/sprint/Sprint-XX/` and review SPRINT_GOAL.md, BACKLOG.md, TASKS.md.
2. Ask the user whether to use the latest sprint or create a new one.
3. Create requirements in `docs/requirements/<feature>/` and update Backlog/Tasks.
4. Mark tasks as `[x]` when complete; register changes in CHANGELOG.md when requested.

---

## Final Checklist

Execution order: Security -> Lint -> Schema -> Tests -> UX -> SEO -> Lighthouse/E2E.

A task only ends when all checks pass. If critical blockers exist, resolve them before proceeding.
