# Gemini CLI Target

## Discovery / Required Paths

- Primary instruction file: `GEMINI.md`
- Commands: `.gemini/commands/**/*.toml`
- Settings: `.gemini/settings.json`
- Rules (content): `.gemini/rules/*.md`
- Skills (content): `.gemini/skills/**`

## Supported Artifacts and File Types

| Canonical type | Target file types | Notes |
| --- | --- | --- |
| prompts | `.md` | Typically maps to `GEMINI.md` |
| commands | `.toml` | Command definitions are TOML; schema is agent-specific and must be validated by the CLI/agent |
| rules | `.md` | Prefer a dedicated directory and reference from `GEMINI.md` |
| skills | `.md` + assets | Gemini CLI may not natively load skills-as-folders; treat as content to be referenced from `GEMINI.md` |
| templates | any | Usually stored but not automatically loaded unless referenced |

## Precedence and Merge Rules

Recommended precedence:
1) `GEMINI.md` defines top-level behavior.
2) `.gemini/settings.json` configures runtime behavior.
3) `.gemini/commands/**/*.toml` provides command catalog.

Sync should treat TOML files as authoritative, file-level managed artifacts (no semantic merges).

## Trust/Safety Gotchas (Sync-Relevant)

- Trust mode: the runtime may ignore project commands and/or user commands when a repository is not trusted.
  - Implication: syncing files successfully does not guarantee they will be loaded.
  - Doctor checks should detect and report whether project commands are currently enabled.
- Command execution: commands can trigger tools. Sync MUST NOT execute commands during install/upgrade.
- TOML schema drift: keep per-command schema versioning in the command itself (or via manifest metadata) to support migrations.

## Canonical OpenKit -> Gemini Mapping

| Canonical artifact | Gemini target path | Mapping notes |
| --- | --- | --- |
| `prompts/root.md` | `GEMINI.md` | Keep as single entrypoint; link to other docs as needed |
| `commands/<id>.toml` | `.gemini/commands/<id>.toml` | Use filename-safe IDs; nested directories allowed for grouping |
| `settings/gemini.settings.json` | `.gemini/settings.json` | Syncable; treat as managed config |
| `rules/<id>.md` | `.gemini/rules/<id>.md` | Prefer referencing files from `GEMINI.md` |
| `skills/<name>/SKILL.md` | `.gemini/skills/<name>/SKILL.md` | Store skills as content for reference |

## Related

- [[agent-compat/agents/README.md]]
- [[agent-compat/README.md]]
- [[content-protocol/PROTOCOL.md]]
