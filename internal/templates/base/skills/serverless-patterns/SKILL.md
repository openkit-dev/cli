---
name: serverless-patterns
description: Serverless design patterns for function boundaries, cold start mitigation, and idempotent execution.
allowed-tools: Read, Write, Edit, Bash
---

# Serverless Patterns

## Function Design

- Keep handlers small and single-purpose.
- Extract shared logic into pure modules.
- Validate all inputs at the edge.

## Reliability

- Make operations idempotent by default.
- Use retries only for safe, repeatable operations.
- Prefer explicit timeouts and dead-letter paths.

## Performance

- Minimize startup paths to reduce cold start impact.
- Avoid heavyweight initialization in handler scope.
- Reuse clients across invocations when runtime allows.

## Security

- Least-privilege IAM/policies for each function.
- Keep secrets in managed stores, never in code.
- Sanitize logs to avoid leaking tokens/PII.

## Observability

- Emit structured logs with request IDs.
- Track latency, error rate, and throttling.
- Keep runbooks for common incident classes.
