# Codex Target

This document targets OpenAI Codex CLI that discovers agents via `AGENTS.md` and supports project-local skills under `.agents/`.

## Discovery / Required Paths

- Agent discovery chain (project root):
  - `AGENTS.override.md` (highest precedence when present)
  - `AGENTS.md` (base)
- Skills: `.agents/skills/<name>/SKILL.md`
- Rules: `.codex/rules/*.rules` (Starlark command policies)
- Config (trust-gated by some runtimes):
  - `.codex/config.toml` (project-local)
  - `~/.codex/config.toml` (user-global)

## Supported Artifacts and File Types

| Canonical type | Target file types | Notes |
| --- | --- | --- |
| agents/subagents | `.md` | Render into `AGENTS.md`/`AGENTS.override.md` using the runtime's expected structure |
| skills | `SKILL.md` (+ optional assets) | Size cap: 32KB default for `AGENTS.md`; keep skills small and modular |
| rules | `.rules` (Starlark) | Project-local rules under `.codex/rules/` for command policies |
| config | `.toml` | Project config may load only when repo is trusted |

## Precedence and Merge Rules

- `AGENTS.override.md` overrides `AGENTS.md` completely or by documented sections (runtime-specific). Sync should treat it as a single managed file if used.
- Project config `.codex/config.toml` (if loaded) typically augments or overrides parts of `~/.codex/config.toml`.
- Rules under `.codex/rules/*.rules` are loaded in lexicographic order; use numeric prefixes for ordering.

## Trust/Safety Gotchas (Sync-Relevant)

- **Trusted projects**: Codex CLI may ignore project commands unless the repo is trusted.
- **Home-directory files**: writing to `~/.codex/*` is risky (affects all projects, may contain secrets). Default sync behavior SHOULD be project-local only.
- **Size caps**:
  - `AGENTS.md` default limit: 32KB (`project_doc_max_bytes`)
  - If exceeded, runtime truncates or ignores content
  - **Mitigation**: keep compact top-level index, move detail into skills
- **Rules format**: must use valid Starlark `prefix_rule()` syntax

## Canonical OpenKit -> Codex Mapping

| Canonical artifact | Codex target path | Mapping notes |
| --- | --- | --- |
| `agents/root.md` | `AGENTS.md` | Comprehensive agent configuration (< 32KB) |
| `agents/override.md` | `AGENTS.override.md` | Optional; use only to override existing project file |
| `skills/<name>/SKILL.md` | `.agents/skills/<name>/SKILL.md` | Skill directory name is canonical skill name |
| `rules/openkit.rules` | `.codex/rules/openkit.rules` | Starlark command policies with `prefix_rule()` |
| `config/codex.config.toml` | `.codex/config.toml` | Only effective in trusted projects |

## Files Installed by OpenKit

When you run `openkit codex sync`, the following files are created:

1. **`AGENTS.md`** - Comprehensive agent configuration with SDD workflow (< 32KB)
2. **`.codex/rules/openkit.rules`** - Starlark command policies (safe commands pre-approved)
3. **`.agents/skills/**`** - Complete skills library

All files are tracked in `.openkit/managed.json`.

## Starlark Rules Format

Rules use `prefix_rule()` function:

```starlark
prefix_rule(
    pattern = ["git", "status"],
    decision = "allow",  # allow | prompt | forbidden
    justification = "Read-only git status is safe",
    match = ["git status", "git status --short"],
)
```

See `.codex/rules/openkit.rules` for complete examples.

## Related

- [[agent-compat/agents/README.md]]
- [[agent-compat/README.md]]
- [[content-protocol/PROTOCOL.md]]
