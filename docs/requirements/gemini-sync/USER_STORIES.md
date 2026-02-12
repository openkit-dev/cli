# User Stories: Gemini Target Sync

## Story 1: Project commands are available

As a developer using Gemini CLI,
I want OpenKit commands installed as `.gemini/commands/openkit/*.toml`,
so I can run consistent project workflows (specify/plan/tasks/impl/etc.).

## Story 2: Rules are discoverable

As a developer,
I want `GEMINI.md` to reference the OpenKit rules files synced into the repo,
so I can inspect and follow them during sessions.

## Story 3: Skills are available as content

As a developer,
I want OpenKit skills synced into `.gemini/skills/`,
so the content is present in the repo and can be referenced when needed.

## Story 4: Sync stays safe by default

As a repo maintainer,
I want `openkit gemini sync` to avoid overwriting unmanaged files by default,
so adopting OpenKit does not clobber local customization.

## Related

- [[requirements/gemini-sync/HUB-GEMINI-SYNC.md]]
- [[requirements/gemini-sync/PROBLEM_STATEMENT.md]]
- [[requirements/gemini-sync/ACCEPTANCE_CRITERIA.md]]
- [[requirements/gemini-sync/PLAN.md]]
