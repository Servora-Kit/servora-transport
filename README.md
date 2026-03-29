# servora-transport

Multi-module repository for Servora transport extension plugins.

## Principles

- One protocol, one module.
- Users import only the protocol module they need.
- Core built-in protocols remain in `servora/transport`.

## Modules

- `server/tcp` -> `github.com/Servora-Kit/servora-transport/server/tcp`

## Layout

```text
servora-transport/
├── server/
│   └── tcp/           # independent Go module
├── client/            # protocol modules (to be added)
├── middleware/        # shared optional helpers
├── shared/            # shared optional utilities
├── examples/
└── docs/
```
