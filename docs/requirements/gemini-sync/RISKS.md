# Risks: Gemini Target Sync

## Command TOML formatting

- Risk: malformed TOML due to unescaped sequences in markdown.
- Mitigation: implement a dedicated TOML string encoder for multiline `prompt`.

## Repo trust and command loading

- Risk: Gemini CLI may ignore project commands unless the repo is trusted.
- Mitigation: document this in `GEMINI.md`; keep `doctor` checks as future work.

## Argument placeholder mismatch

- Risk: `$ARGUMENTS` mapping differs from Gemini expectations.
- Mitigation: standardize on `{{args}}` and document usage in generated prompts.

## Large content footprint

- Risk: syncing all skills/rules increases file count.
- Mitigation: rely on managed state + safe prune; keep directories scoped under `.gemini/`.

## Related

- [[requirements/gemini-sync/HUB-GEMINI-SYNC.md]]
- [[requirements/gemini-sync/PROBLEM_STATEMENT.md]]
- [[requirements/gemini-sync/PLAN.md]]
- [[sprint/Sprint-02/RISK_REGISTER.md]]
