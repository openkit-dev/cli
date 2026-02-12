# Agent Compatibility (Content Sync Targets)

This directory documents how a canonical OpenKit content pack is synced into different agent-specific project layouts. It is intentionally independent of the OpenKit CLI binary and focuses on file-system artifacts.

Canonical protocol definitions:
- [[docs/content-protocol/PROTOCOL.md]]
- [[docs/content-protocol/MANIFEST_SCHEMA.md]]
- [[docs/content-protocol/SYNC_SEMANTICS.md]]

Agent-specific profiles:
- [[docs/agent-compat/agents/README.md]]

## At A Glance

| Agent target | Primary entrypoint | Project directories (managed) | Key safety/trust gotcha |
| --- | --- | --- | --- |
| OpenCode | `opencode.json` | `.opencode/{commands,prompts,rules,skills,templates}` | Rule packs can change tool behavior; treat `rules/` as sensitive |
| Claude Code | `CLAUDE.md` or `.claude/CLAUDE.md` | `.claude/{rules,skills,agents}` | Never manage `.claude/settings.local.json` (user-local) |
| Gemini CLI | `GEMINI.md` | `.gemini/commands/**.toml` + `.gemini/settings.json` | Trust mode may ignore project/user commands entirely |
| Codex | `AGENTS.md` (+ `AGENTS.override.md`) | `.agents/skills/<name>/SKILL.md` | Many configs/rules are outside repo and/or only in trusted projects |
| Cursor | `.cursorrules` | `.cursorrules` | Single file; keep concise |

## What These Docs Cover

Each per-agent document includes:
- Supported artifact types (rules, commands, prompts, skills, templates, agents/subagents)
- Supported file types and structural constraints
- Required paths and discovery/precedence rules
- Trust/safety gotchas that affect syncing behavior
- Explicit mapping notes: canonical OpenKit protocol -> that agent's layout

## Canonical-to-Target Mapping (High Level)

OpenKit content packs are agent-neutral. A pack provides canonical artifacts (IDs + type + bytes) and a renderer target mapping (where each artifact should land for a specific agent).

The sync engine is responsible for:
- Selecting a pack version
- Rendering/copying artifacts into the agent layout
- Tracking managed files under `.openkit/managed.json`
- Avoiding destructive behavior by default (see sync semantics)

## Related

- [[docs/HUB-DOCS.md]]
- [[docs/agent-compat/agents/README.md]]
- [[docs/content-protocol/README.md]]
