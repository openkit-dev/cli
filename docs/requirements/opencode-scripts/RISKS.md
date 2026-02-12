# Risks: Port NPX scripts into CLI templates

- Python availability varies per environment.
  - Mitigation: scripts use `sys.executable` for subprocess.

- Some projects do not use Docker Compose.
  - Mitigation: `auto_preview.py` should fail with a clear message when compose files are missing.

- Scripts may reference skill paths that are not installed (if the user did not sync OpenCode).
  - Mitigation: keep scripts scoped under `.opencode/scripts/` and document they are for OpenCode target.

## Related

- [[requirements/opencode-scripts/HUB-OPENCODE-SCRIPTS.md]]
- [[requirements/opencode-scripts/PROBLEM_STATEMENT.md]]
- [[requirements/opencode-scripts/PLAN.md]]
- [[sprint/Sprint-03/RISK_REGISTER.md]]
