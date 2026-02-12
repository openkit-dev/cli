---
name: iac-patterns
description: Infrastructure-as-Code patterns for Terraform/Docker/Kubernetes workflows with safe change management.
allowed-tools: Read, Write, Edit, Bash
---

# IaC Patterns

## Structure

- Organize infra by environment and domain boundaries.
- Reuse modules, avoid copy/paste stacks.
- Keep variable defaults safe and explicit.

## Change Safety

- Always plan before apply.
- Review destructive actions explicitly.
- Keep rollback/recovery procedure next to manifests.

## State and Secrets

- Use remote state with locking for shared stacks.
- Never commit plaintext secrets.
- Use secret managers and short-lived credentials.

## Kubernetes and Containers

- Pin image tags and base images intentionally.
- Define probes/resources explicitly.
- Keep least-privilege RBAC policies.

## Delivery

- Validate configs in CI before deployment.
- Promote with environment gates (dev -> stage -> prod).
- Capture drift and reconcile deliberately.
