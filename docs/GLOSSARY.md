# Glossary

| Term | Definition |
|---|---|
| OpenKit | Go CLI toolkit for syncing SDD assets for multiple coding agents. |
| SDD | Spec-Driven Development workflow: `/discover -> /specify -> /create -> /verify -> /deploy`. |
| Managed State | `.openkit/managed.json` registry tracking installed artifacts and checksums. |
| Drift | A managed file whose current checksum differs from recorded `installed_sha256`. |
| Conflict | Existing unmanaged file where sync wants to write managed content. |
| Pack | Agent content bundle id/version recorded in managed state. |
| Doctor | Health command (`openkit <agent> doctor`) for install and drift diagnostics. |
| Overlay | Optional context-doc extension selected during discovery generation logic. |

## Related

- [[HUB-DOCS.md]]
- [[CONTEXT.md]]
