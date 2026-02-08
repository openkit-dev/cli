# Cursor Target

## Discovery / Required Paths

- Primary rules file: `.cursorrules` (backward compatibility)
- Modular rules: `.cursor/rules/*.mdc`
- Skills: `.cursor/skills/<name>/SKILL.md`

## Supported Artifacts and File Types

| Canonical type | Target file types | Notes |
| --- | --- | --- |
| rules | `.md` (`.cursorrules`) | Legacy single-file format |
| rules | `.mdc` (`.cursor/rules/*.mdc`) | Modern modular format with YAML frontmatter |
| skills | `SKILL.md` | Installed under `.cursor/skills/` |

## Precedence and Merge Rules

- `.cursorrules` is loaded first (backward compatibility)
- `.cursor/rules/*.mdc` files are loaded in lexicographic order
- Skills are referenced by path in rules files

## Trust/Safety Gotchas (Sync-Relevant)

- User workflows vary by IDE version and extensions; treat `.cursorrules` as best-effort.
- Keep `.cursorrules` concise to avoid truncation.
- `.mdc` files should include YAML frontmatter with `name` and `description`.

## Canonical OpenKit -> Cursor Mapping

| Canonical artifact | Cursor target path | Mapping notes |
| --- | --- | --- |
| `rules/master.md` | `.cursorrules` | Render essentials into single file for backward compatibility |
| `rules/openkit.mdc` | `.cursor/rules/openkit.mdc` | Modern modular rule with frontmatter |
| `skills/<name>/SKILL.md` | `.cursor/skills/<name>/SKILL.md` | Skill directory name is canonical skill name |

## Files Installed by OpenKit

When you run `openkit cursor sync`, the following files are created:

1. **`.cursorrules`** - Enhanced project rules (backward compatible)
2. **`.cursor/rules/openkit.mdc`** - Modular SDD workflow rules
3. **`.cursor/skills/**`** - Complete skills library

All files are tracked in `.openkit/managed.json`.
