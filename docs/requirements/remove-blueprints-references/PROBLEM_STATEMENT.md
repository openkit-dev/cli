# Problem Statement: Remove Blueprint References

OpenKit currently contains scattered mentions of "blueprints" (including @-prefixed alias references) across skills, templates, and audit scripts.

These references imply a supported concept and/or a repository location that is not actually shipped with the CLI, which creates confusion for contributors and users.

We need to remove product-level references to "blueprints" and standardize on the terminology and structure that OpenKit actually ships (templates under `internal/templates/` and generated packs under `.opencode/`).

## Related

- [[docs/requirements/remove-blueprints-references/HUB-REMOVE-BLUEPRINTS-REFERENCES.md]]
- [[docs/requirements/remove-blueprints-references/USER_STORIES.md]]
- [[docs/requirements/remove-blueprints-references/ACCEPTANCE_CRITERIA.md]]
- [[docs/requirements/remove-blueprints-references/PLAN.md]]
