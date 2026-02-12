# Implementation Plan: Gemini Target Sync

## Scope

Complete the Gemini target so `openkit gemini sync` installs:

- Project commands as TOML: `.gemini/commands/openkit/*.toml`
- Rules: `.gemini/rules/*.md`
- Skills: `.gemini/skills/**`
- A `GEMINI.md` entrypoint that references the synced rules/skills

## Design

1. Add a small internal generator that:
   - Loads `base/commands/*.md` from the embedded FS
   - Splits YAML frontmatter (if present) from markdown body
   - Extracts `description` (optional)
   - Replaces `$ARGUMENTS` -> `{{args}}`
   - Emits TOML with `prompt` and optional `description`

2. Update `buildGemini` in `cli/internal/targets/targets.go` to:
   - Copy base rules into `.gemini/rules/`
   - Copy base skills into `.gemini/skills/`
   - Generate TOML files into `.gemini/commands/openkit/`
   - Update `GEMINI.md` content to point to `.gemini/rules/` and `.gemini/skills/`

3. Update docs:
   - `cli/docs/agent-compat/agents/gemini.md` to reflect rules and skills directories.

## Verification

- Run `go test ./...` in `cli/`.
- Add focused unit tests if feasible for:
  - frontmatter parsing
  - TOML string escaping
  - `$ARGUMENTS` replacement

## Related

- [[requirements/gemini-sync/HUB-GEMINI-SYNC.md]]
- [[requirements/gemini-sync/PROBLEM_STATEMENT.md]]
- [[content-protocol/PROTOCOL.md]]
- [[sprint/Sprint-02/TASKS.md]]
