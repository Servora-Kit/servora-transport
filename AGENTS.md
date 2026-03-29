# AGENTS.md - servora-transport

## Purpose

Independent repository for transport extension plugins used by Servora ecosystem projects.

## Scope

- Protocol plugins beyond built-in `grpc`/`http`
- Integration-specific transport adapters
- Reusable transport middleware extensions
- Repository is multi-module: one protocol plugin should be one independent Go module.

## Non-Goals

- Business service logic
- Re-implementing Servora core runtime contracts
