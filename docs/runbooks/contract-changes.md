# Contract Changes Runbook

## Purpose

Use this runbook when a tracer needs to change shared contracts.

## Rules

- lock envelope shape before detailed payloads
- avoid speculative schemas for repos that are not active in the tracer
- keep subject naming in `{service}.{entity}.{action}` form
- prefer additive changes inside `v1`
- create a new version only for true breaking changes

## Verification

- `protoc` syntax validation passes
- any generated config files stay coherent
- consuming repos update in the same tracer when required
