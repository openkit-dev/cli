# Acceptance Criteria: Port NPX scripts into CLI templates

- `openkit opencode sync` installs the following files under `.opencode/scripts/`:
  - `checklist.py`
  - `verify_all.py`
  - `auto_preview.py`
  - `session_manager.py`

- Scripts remain generic:
  - Use only Python standard library.
  - Default to `sys.executable` when spawning Python.
  - Avoid non-ASCII output.

- `cli/README.md` includes a short, user-facing section describing the main workflows:
  - Standard SDD flow: `/specify` -> `/clarify` -> `/plan` -> `/tasks` -> `/impl` -> `/test`
  - Orchestrated flow via `/engineer`
  - Verification scripts (`.opencode/scripts/checklist.py`, `.opencode/scripts/verify_all.py`)

## Related

- [[docs/requirements/opencode-scripts/HUB-OPENCODE-SCRIPTS.md]]
- [[docs/requirements/opencode-scripts/PROBLEM_STATEMENT.md]]
- [[docs/requirements/opencode-scripts/PLAN.md]]
- [[docs/sprint/Sprint-03/TASKS.md]]
