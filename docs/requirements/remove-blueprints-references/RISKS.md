# Risks: Remove Blueprint References

- Removing `blueprints` from skip lists may increase scan noise if some workspaces include a user-created `blueprints/` directory.
  - Mitigation: keep skip lists focused on OpenKit-owned directories (e.g. `.opencode/`, internal tooling dirs); document how to exclude custom dirs if needed.
- Some occurrences of "blueprint" may be generic language (e.g. "mental blueprint") and not related to the product concept.
  - Mitigation: decide a standard alternative term (e.g. "mental model") and update prompts consistently.

## Related

- [[docs/requirements/remove-blueprints-references/HUB-REMOVE-BLUEPRINTS-REFERENCES.md]]
- [[docs/requirements/remove-blueprints-references/PROBLEM_STATEMENT.md]]
- [[docs/requirements/remove-blueprints-references/PLAN.md]]
- [[docs/sprint/Sprint-06/RISK_REGISTER.md]]
