# ACTION_ITEMS

| Priority | Severity | Item | Impact | Effort | Owner | Notes |
|---|---|---|---|---|---|---|
| P0 | High | Add dependency vulnerability scan (`cargo audit` or equivalent) to CI | Reduces supply-chain exposure in release path | Low | DevOps | Evidence: `.github/workflows/ci.yml` has no dependency audit step |
| P1 | High | Add signed release verification path for installers/self-update | Improves integrity beyond checksum-only validation | Medium | Security + Release Eng | Evidence: `rust-cli/src/main.rs` verifies checksum only |
| P2 | Medium | Introduce structured logging option for CLI operations | Improves diagnostics and automated observability | Medium | Runtime Team | Evidence: widespread `println!` usage in `rust-cli/src/main.rs` |
| P2 | Low | Review legacy empty directories (`cmd/openkit`) for cleanup policy | Reduces repository ambiguity and drift | Low | Maintainers | Evidence: `cmd/openkit` exists but is empty |

## Cross-Repo Impacts

| Severity | Owner | Impact | Reference |
|---|---|---|---|
| Medium | Docs Maintainer | Public release documentation may point to non-existent paths for users and downstream templates | `README.md:77` |
| Medium | Template Maintainer | Agent/template consumers may inherit outdated navigation assumptions | `.opencode/templates/DOCS-README.md:19` |

## Related

- [[CONTEXT.md]]
- [[SECURITY.md]]
- [[QUALITY_GATES.md]]
- [[HUB-DOCS.md]]
- [[sprint/HUB-SPRINTS.md]]
