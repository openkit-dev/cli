---
name: desktop-patterns
description: Desktop app architecture patterns for Electron/Tauri/native wrappers. Use for IPC boundaries, update channels, and local permissions.
allowed-tools: Read, Write, Edit, Bash
---

# Desktop Patterns

## Process Boundaries

- Separate UI process from privileged host process.
- Use explicit IPC contracts and validate every payload.
- Keep privileged operations in minimal modules.

## Update Strategy

- Use staged rollout channels (alpha, beta, stable).
- Verify update integrity before applying.
- Always provide rollback instructions.

## Security

- Disable unnecessary host permissions by default.
- Restrict filesystem access to approved roots.
- Audit shell/process spawn entry points.

## Packaging

- Keep reproducible build configs.
- Sign binaries when supported.
- Document platform-specific constraints clearly.

## UX for Desktop

- Respect offline operation where possible.
- Handle app restarts and crash recovery gracefully.
- Provide concise diagnostics for support cases.
