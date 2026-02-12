# OpenCode Target (OpenKit Native Layout)

## Discovery / Required Paths

- Root marker: `opencode.json`
- Project content root: `.opencode/`
- Managed directories (by convention):
  - `.opencode/commands/*.md`
  - `.opencode/prompts/*.md`
  - `.opencode/rules/*.md`
  - `.opencode/skills/<name>/SKILL.md`
  - `.opencode/templates/**` (optional; pack-defined)

## Supported Artifacts and File Types

| Canonical type | Typical file types | Notes |
| --- | --- | --- |
| rules | `.md` | May include frontmatter-like headers used by some runtimes; keep deterministic ordering |
| commands | `.md` | One command per file; filename is the command name |
| prompts | `.md` | Agent prompt/instructions; referenced by runtime configuration |
| skills | `SKILL.md` (+ optional adjacent assets) | `SKILL.md` MUST contain YAML frontmatter; adjacent `scripts/`, `templates/` allowed |
| templates | any text/binary | Treated as opaque assets; pack decides target subpath |
| agents/subagents | `.md` | Only if the runtime supports subagents; otherwise mapped to prompts |

## Precedence and Merge Rules

Recommended sync precedence (non-destructive by default):
1) If target file is managed and unchanged, update in-place.
2) If target file is managed but checksum drifted, treat as user-edited conflict (skip unless `--overwrite`).
3) If target file is unmanaged and exists, do not overwrite by default.

No semantic merges are assumed for Markdown. Sync is file-level (copy/render) with explicit conflict behavior (see sync semantics).

## Trust/Safety Gotchas (Sync-Relevant)

- `.opencode/rules/` can materially change safety posture (tool usage constraints, execution permissions). Treat updates to rules as high-risk; prefer explicit `--dry-run` review.
- Skills can ship runnable scripts/assets. Sync MUST NOT execute any scripts. Only copy/render bytes.
- YAML frontmatter in `.opencode/skills/<name>/SKILL.md` is structurally significant. Sync must preserve it exactly.

## Canonical OpenKit -> OpenCode Mapping

| Canonical artifact | OpenCode target path | Mapping notes |
| --- | --- | --- |
| `rules/<id>.md` | `.opencode/rules/<id>.md` | Prefer stable IDs that also work as filenames |
| `commands/<id>.md` | `.opencode/commands/<id>.md` | Command IDs MUST be filename-safe |
| `prompts/<id>.md` | `.opencode/prompts/<id>.md` | If prompt is agent-specific, keep ID namespaced (e.g. `opencode.<id>`) |
| `skills/<name>/SKILL.md` | `.opencode/skills/<name>/SKILL.md` | Skill directory name is the canonical skill name |
| `skills/<name>/**` | `.opencode/skills/<name>/**` | Optional scripts/templates/assets; copied verbatim |
| `templates/**` | `.opencode/templates/**` (or pack-defined) | Prefer pack-defined destinations to avoid collisions |

## Related

- [[agent-compat/agents/README.md]]
- [[agent-compat/README.md]]
- [[content-protocol/PROTOCOL.md]]
