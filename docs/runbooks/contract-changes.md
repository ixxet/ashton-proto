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

## Verification

- `protoc` syntax validation passes
- event-specific JSON Schema validation tests cover one valid payload and the required rejection cases
- shared runtime helpers cover marshal, parse, and source-value validation for
  active event paths
- fixture bytes used by consuming repos come from `ashton-proto`, not
  hand-written JSON strings copied into service tests
- any generated config files stay coherent
- consuming repos update in the same tracer when required
