# Contract Changes Runbook

## Purpose

Use this runbook when a tracer needs to change shared contracts.

## Rules

- lock envelope shape before detailed payloads
- avoid speculative schemas for repos that are not active in the tracer
- keep subject naming in `{service}.{entity}.{action}` form
- prefer additive changes inside `v1`
- create a new version only for true breaking changes
- when a tracer message becomes active across repos, add one shared runtime
  helper or validator in `ashton-proto` and have both producer and consumer
  import it instead of defining private wire structs

## Maintainer Decision Guide

| If the change is... | Update proto | Update event schema | Update Go helper | Update MCP manifest | Cut a tag now? |
| --- | --- | --- | --- | --- | --- |
| new HTTP or RPC request/response type shared across repos | Yes | No | Usually no | Only if the gateway consumes it | Only when a downstream repo now depends on the released contract |
| new event payload shared across producer and consumer repos | Maybe | Yes | Yes | No | Yes if producer or consumer expects the new contract from a released line |
| new enum or field added to an active proto surface | Yes | Maybe | Maybe | Maybe | Yes if a downstream released repo needs it; otherwise keep it additive on `main` |
| stricter event validation without changing the outer shape | No | Yes | Yes | No | Tag if downstream repos or hardening artifacts now depend on that stricter behavior |
| new routed read-only tool for the gateway | Maybe | No | No | Yes | Tag if the gateway or service repo consumes the manifest from a released line |
| docs-only clarification of current vs planned state | No | No | No | No | No |

## Verification

- `protoc` syntax validation passes
- event-specific JSON Schema validation tests cover one valid payload and the required rejection cases
- shared runtime helpers cover marshal, parse, and source-value validation for
  active event paths
- fixture bytes used by consuming repos come from `ashton-proto`, not
  hand-written JSON strings copied into service tests
- any generated config files stay coherent
- consuming repos update in the same tracer when required
