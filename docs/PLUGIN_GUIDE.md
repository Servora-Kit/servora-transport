# Plugin Guide

## Minimal Rules

1. Implement either `runtime.ServerPlugin` or `runtime.ClientPlugin`.
2. Expose a stable `const Type = "..."`.
3. Validate `Build` input types and return explicit errors.
4. Keep protocol-specific logic in plugin package; keep generic helpers in `shared/`.
5. Add tests for `Type()`, successful `Build`, and failure paths.
