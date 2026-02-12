# Claude Code Target

## Discovery / Required Paths

- Primary instruction file (choose one):
  - `.claude/CLAUDE.md` (preferred when present)
  - `CLAUDE.md` (fallback)
- Rules: `.claude/rules/*.md`
- Skills: `.claude/skills/<name>/SKILL.md`
- Subagents/agents: `.claude/agents/*.md`
- Settings:
  - `.claude/settings.json` (project settings; syncable)
  - `.claude/settings.local.json` (user-local overrides; MUST NOT be managed)

## Supported Artifacts and File Types

| Canonical type | Target file types | Notes |
| --- | --- | --- |
| rules | `.md` | Multiple files are supported; define ordering explicitly (lexicographic filename) |
| commands | `.md` (or embedded in `CLAUDE.md`) | If the agent expects a single entrypoint file, map commands into `.claude/agents/` or `CLAUDE.md` |
| prompts | `.md` | Usually becomes `.claude/CLAUDE.md` or `.claude/agents/<id>.md` |
| skills | `SKILL.md` (+ optional assets) | Keep YAML frontmatter if used; otherwise treat as Markdown |
| templates | any | Only if Claude tooling supports loading from disk; otherwise keep as opaque assets for other targets |

## Precedence and Merge Rules

Recommended precedence for instruction entrypoint:
1) Use `.claude/CLAUDE.md` if it exists (managed by sync when configured).
2) Otherwise use `CLAUDE.md`.

Settings precedence (recommended):
1) `.claude/settings.local.json` overrides `.claude/settings.json`.
2) Sync MUST NOT overwrite or delete `.claude/settings.local.json`.

Rules precedence (recommended):
- Load all `.claude/rules/*.md` in lexicographic order; if a rule must override another, encode ordering via filename prefix (e.g. `00-...`, `90-...`).

## Trust/Safety Gotchas (Sync-Relevant)

- Local settings file: `.claude/settings.local.json` commonly contains machine/user-specific values; managing it risks leaks and breakage.
- Some environments may restrict tool execution regardless of instructions. Sync should not assume that adding rules enables capabilities.
- Large instruction files may be truncated by the agent runtime. Prefer splitting content into `.claude/rules/` and `.claude/agents/` where supported.

## Canonical OpenKit -> Claude Mapping

| Canonical artifact | Claude target path | Mapping notes |
| --- | --- | --- |
| `prompts/root.md` | `.claude/CLAUDE.md` | Prefer a single canonical `root` prompt for top-level behavior |
| `rules/<id>.md` | `.claude/rules/<id>.md` | Keep IDs filename-safe; consider numeric prefixes for ordering |
| `skills/<name>/SKILL.md` | `.claude/skills/<name>/SKILL.md` | Skill directory name is the canonical skill name |
| `agents/<id>.md` | `.claude/agents/<id>.md` | Use for subagent/tooling instructions when available |
| `settings/claude.settings.json` | `.claude/settings.json` | Syncable; do not touch `.claude/settings.local.json` |

## Related

- [[agent-compat/agents/README.md]]
- [[agent-compat/README.md]]
- [[content-protocol/PROTOCOL.md]]
