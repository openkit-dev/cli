# Problem Statement: Cursor + Codex Targets

## Current State

The OpenKit CLI currently has minimal implementations for Cursor and Codex targets:

### Cursor Target (buildCursor)
- Generates only `.cursorrules` (6 lines of generic text)
- Does not leverage Cursor's `.cursor/rules/*.mdc` system
- No skills, no commands, no detailed guidance

### Codex Target (buildCodex)
- Generates `AGENTS.md` (basic 8 lines)
- Copies skills to `.agents/skills/`
- Does not include rules or detailed instructions
- Does not follow Codex's `AGENTS.md` + `AGENTS.override.md` hierarchy

## Research Summary

### Cursor Capabilities
- `.cursorrules` at project root (legacy, still supported)
- `.cursor/rules/*.mdc` for modular rules (modern approach)
- No native command system (uses prompts/chat)
- Rules are markdown with frontmatter metadata

### Codex Capabilities
- `AGENTS.md` at project root (primary entrypoint)
- `AGENTS.override.md` for temporary overrides
- Hierarchical discovery: walks from root to current directory
- `config.toml` for advanced settings
- `.codex/rules/*.rules` for Starlark command policies
- Skills support via `/codex/skills`

## Target Outcome

| Agent | Current Files | Improved Files |
|-------|--------------|----------------|
| Cursor | `.cursorrules` (6 lines) | `.cursorrules` + `.cursor/rules/*.mdc` + skills reference |
| Codex | `AGENTS.md` + `.agents/skills/**` | `AGENTS.md` (rich) + `.codex/rules/openkit.rules` + skills |

## Success Criteria

1. Cursor sync installs modular rules that reference OpenKit skills
2. Codex sync installs AGENTS.md with proper structure and skill references
3. Both targets pass existing tests and maintain managed state compatibility
4. Doctor commands provide useful diagnostics for each target

## Related

- [[docs/requirements/cursor-codex-targets/HUB-CURSOR-CODEX-TARGETS.md]]
- [[docs/requirements/cursor-codex-targets/USER_STORIES.md]]
- [[docs/requirements/cursor-codex-targets/ACCEPTANCE_CRITERIA.md]]
- [[docs/requirements/cursor-codex-targets/PLAN.md]]
