# Problem Statement: Gemini Target Sync (commands, skills, rules)

## Problem

`openkit gemini sync` currently installs only `GEMINI.md` and `.gemini/settings.json`.
This leaves the Gemini CLI target without:

- Project-local custom commands (`.gemini/commands/**/*.toml`)
- A consistent location for OpenKit rules and skills that can be referenced from `GEMINI.md`

As a result, Gemini users cannot reliably use the OpenKit workflow commands or discover the project rules/skills content.

## Why now

- We already have a canonical command library under `cli/internal/templates/base/commands/*.md`.
- We already have canonical rules and skills under `cli/internal/templates/base/rules/` and `cli/internal/templates/base/skills/`.
- Gemini CLI supports project-local commands via TOML files; syncing them completes the Gemini target.

## Non-goals

- Do not execute any Gemini commands during sync.
- Do not rely on Gemini auto-loading skills/rules directories beyond what `GEMINI.md` explicitly references.

## Related

- [[docs/requirements/gemini-sync/HUB-GEMINI-SYNC.md]]
- [[docs/requirements/gemini-sync/USER_STORIES.md]]
- [[docs/requirements/gemini-sync/ACCEPTANCE_CRITERIA.md]]
- [[docs/requirements/gemini-sync/PLAN.md]]
