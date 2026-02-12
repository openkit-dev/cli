# Acceptance Criteria: Gemini Target Sync

- `openkit gemini sync` installs/updates:
  - `GEMINI.md`
  - `.gemini/settings.json`
  - `.gemini/commands/openkit/*.toml` generated from `cli/internal/templates/base/commands/*.md` (excluding `README.md`)
  - `.gemini/skills/**` copied from `cli/internal/templates/base/skills/**`
  - `.gemini/rules/*.md` copied from `cli/internal/templates/base/rules/*.md`

- TOML command generation:
  - Reads YAML frontmatter `description` when present and maps it to TOML `description`.
  - Uses the markdown body as TOML `prompt` (multiline).
  - Replaces `$ARGUMENTS` with `{{args}}` in the generated prompt.

- `GEMINI.md` references the synced rules directory paths (no embedded full rule content).

- Sync remains safe-by-default:
  - Unmanaged conflicts are skipped unless `--overwrite`.
  - Managed state is tracked in `.openkit/managed.json`.

## Related

- [[docs/requirements/gemini-sync/HUB-GEMINI-SYNC.md]]
- [[docs/requirements/gemini-sync/PROBLEM_STATEMENT.md]]
- [[docs/requirements/gemini-sync/USER_STORIES.md]]
- [[docs/sprint/Sprint-02/TASKS.md]]
