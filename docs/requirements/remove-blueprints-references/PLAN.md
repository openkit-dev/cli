# Implementation Plan: Remove Blueprint References

## Scope

Remove references and mentions of "blueprints" across the OpenKit CLI repository, and standardize wording to match shipped artifacts (templates and generated packs).

## Work Items

1. **Inventory**
   - Search for `blueprints` and `blueprint` occurrences.
   - Classify matches: product reference vs. generic language vs. skip-list handling.

2. **Standardize terminology**
   - Adopt "templates" as the canonical term in docs/skills.
   - Replace blueprint alias path references with a real, shipped path (prefer `internal/templates/base/` or another concrete example).

3. **Apply edits in source + template base**
   - Update `.opencode/**` sources.
   - Mirror the same edits into `internal/templates/base/**` so future generations stay aligned.

4. **Adjust audit scripts**
   - Remove "templates/blueprints" wording from comments.
   - Remove `blueprints` directory special-casing unless it is still required for performance/compat.

5. **Sprint docs updates**
   - Add backlog + tasks for the cleanup.
   - Record risks in the sprint risk register.

6. **Verification**
   - Repo-wide search confirms no blueprint alias path references remain.
   - Run:
     - `python .opencode/scripts/checklist.py .`
     - `python .opencode/skills/vulnerability-scanner/scripts/security_scan.py .`

## Out of Scope

- Introducing a new external repository for content examples.
- Changing CLI behavior unrelated to documentation/content cleanup.

## Verification Record

- 2026-02-08: `python3 .opencode/scripts/checklist.py .` PASS
- 2026-02-08: `python3 .opencode/scripts/verify_all.py . --url http://localhost:3000` PASS

## Related

- [[docs/requirements/remove-blueprints-references/HUB-REMOVE-BLUEPRINTS-REFERENCES.md]]
- [[docs/requirements/remove-blueprints-references/PROBLEM_STATEMENT.md]]
- [[docs/requirements/remove-blueprints-references/USER_STORIES.md]]
- [[docs/sprint/Sprint-06/TASKS.md]]
