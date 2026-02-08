# Acceptance Criteria: Cursor + Codex Targets

## AC-1: Cursor Target Sync

### AC-1.1: File Generation
- [ ] `openkit cursor sync` creates `.cursorrules` (backward compat)
- [ ] `openkit cursor sync` creates `.cursor/rules/openkit.mdc`
- [ ] `openkit cursor sync` copies skills to `.cursor/skills/`
- [ ] All files tracked in `.openkit/managed.json`

### AC-1.2: Content Quality
- [ ] `.cursorrules` contains meaningful project guidance (not just 6 lines)
- [ ] `.cursor/rules/openkit.mdc` has frontmatter + markdown body
- [ ] Skills copied maintain directory structure

### AC-1.3: Idempotency
- [ ] Running sync twice produces no changes on second run
- [ ] `--dry-run` shows plan without writing

---

## AC-2: Codex Target Sync

### AC-2.1: File Generation
- [ ] `openkit codex sync` creates `AGENTS.md` at project root
- [ ] `openkit codex sync` creates `.codex/rules/openkit.rules`
- [ ] `openkit codex sync` copies skills to `.agents/skills/`
- [ ] All files tracked in `.openkit/managed.json`

### AC-2.2: AGENTS.md Content
- [ ] Contains project expectations section
- [ ] Contains skills reference with paths
- [ ] Contains SDD workflow summary
- [ ] File size < 32KB (Codex default limit)

### AC-2.3: Rules File Content
- [ ] Uses valid Starlark syntax
- [ ] Includes prefix_rule() for common commands
- [ ] Has comments explaining each rule

### AC-2.4: Idempotency
- [ ] Running sync twice produces no changes on second run
- [ ] `--dry-run` shows plan without writing

---

## AC-3: Doctor Commands

### AC-3.1: Cursor Doctor
- [ ] Checks for `.cursorrules` presence
- [ ] Checks for `.cursor/rules/` directory
- [ ] Reports managed file count
- [ ] Reports drift status

### AC-3.2: Codex Doctor
- [ ] Checks for `AGENTS.md` presence
- [ ] Checks for `.codex/rules/` directory
- [ ] Checks for `.agents/skills/` directory
- [ ] Reports managed file count
- [ ] Reports drift status

---

## AC-4: Tests

### AC-4.1: Unit Tests
- [ ] Test Cursor file generation functions
- [ ] Test Codex AGENTS.md generation
- [ ] Test Codex rules file syntax

### AC-4.2: Integration Tests
- [ ] Test `openkit cursor sync` end-to-end
- [ ] Test `openkit codex sync` end-to-end
- [ ] Verify managed.json entries
- [ ] Verify idempotency

---

## AC-5: Documentation

### AC-5.1: Agent Compat Docs
- [ ] Update `cli/docs/agent-compat/agents/cursor.md`
- [ ] Update `cli/docs/agent-compat/agents/codex.md`

### AC-5.2: README
- [ ] Update README supported agents table
