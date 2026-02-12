# Risks: Cursor + Codex Targets

## R1: Cursor Rules Format Uncertainty (Medium)

**Risk:** Cursor's `.cursor/rules/*.mdc` format may not be fully documented or may change.

**Impact:** Rules may not load correctly in Cursor.

**Mitigation:**
- Use simple markdown with minimal frontmatter
- Test with actual Cursor installation
- Fallback to `.cursorrules` for compatibility

**Status:** Mitigated (Verified by Task P1-3)

---

## R2: Codex Starlark Syntax Errors (Low)

**Risk:** Generated `.codex/rules/openkit.rules` may have syntax errors.

**Impact:** Codex may fail to load rules.

**Mitigation:**
- Keep rules simple (only prefix_rule)
- Add unit test that validates syntax
- Reference official examples from Codex docs

**Status:** Mitigated (Verified by Task P1-2)

---

## R3: AGENTS.md Size Limit (Low)

**Risk:** If we include too much content, AGENTS.md may exceed Codex's 32KB limit.

**Impact:** Instructions may be truncated.

**Mitigation:**
- Keep AGENTS.md concise (<10KB target)
- Reference skills by path instead of inlining
- Document limit in generated file

**Status:** Mitigated (Verified by Task P0-3)

---

## R4: Breaking Existing Installations (Medium)

**Risk:** Users with existing `.cursorrules` or `AGENTS.md` may have conflicts.

**Impact:** Sync may skip files or require `--overwrite`.

**Mitigation:**
- Safe-by-default behavior (skip unmanaged)
- Clear messaging when conflicts detected
- Doctor command shows conflict status

**Status:** Mitigated (Verified by Tasks P0-2, P0-4)

---

## R5: Cursor Skills Path Not Standard (Low)

**Risk:** Cursor may not have a standard skills path like other agents.

**Impact:** Skills may not be discoverable.

**Mitigation:**
- Use `.cursor/skills/` as reasonable convention
- Document in generated rules that skills are available
- Let users customize via their own rules

**Status:** Mitigated (Verified by Task P2-3)

## Related

- [[docs/requirements/cursor-codex-targets/HUB-CURSOR-CODEX-TARGETS.md]]
- [[docs/requirements/cursor-codex-targets/PROBLEM_STATEMENT.md]]
- [[docs/requirements/cursor-codex-targets/PLAN.md]]
- [[docs/sprint/Sprint-04/HUB-SPRINT-04.md]]
