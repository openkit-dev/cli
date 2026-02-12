# User Stories: Sync Engine + managed.json

- As a developer, I can run `openkit <agent> sync` to install the OpenKit content pack into my project.
- As a developer, I can run `openkit <agent> upgrade` to move to a newer compatible pack version without losing customizations.
- As a CI pipeline, I can run `openkit <agent> upgrade --fail-on-changes` to detect drift/conflicts and fail the job deterministically.
- As a developer, I can run `openkit <agent> doctor` to see missing entrypoints, trust-mode issues, and drift summary.

## Related

- [[requirements/openkit-sync/HUB-OPENKIT-SYNC.md]]
- [[requirements/openkit-sync/PROBLEM_STATEMENT.md]]
- [[requirements/openkit-sync/ACCEPTANCE_CRITERIA.md]]
- [[requirements/openkit-sync/PLAN.md]]
