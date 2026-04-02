# ashton-proto

Shared contracts repo for the ASHTON platform.

This repo will hold the protobuf definitions, event schemas, and MCP tool manifests that every service depends on. Nothing here should become a runnable app. Its job is to make the rest of the platform type-safe, consistent, and easier to evolve without accidental drift.

This repo now holds the active shared contract surface for the first cross-repo tracers. The implementation brief lives in [ashton-platform/planning/repo-briefs/ashton-proto.md](https://github.com/ixxet/ashton-platform/blob/main/planning/repo-briefs/ashton-proto.md).

## Role In The Platform

- first repo to build
- source of truth for shared contracts
- dependency anchor for `athena`, `hermes`, `apollo`, and `ashton-mcp-gateway`

## First Execution Goal

Ship the minimum contract surface needed for the first ATHENA tracer bullet:

- common shared types
- initial ATHENA messages
- one event envelope definition
- one ATHENA MCP tool manifest

The event envelope gets locked first. Detailed payload schemas stay intentionally light until the first adapter-backed slices exist.

## Current State

Tracer 1 contract hardening complete:

- `buf.yaml` and `buf.gen.yaml` are active and lint-clean
- `make generate`, `make lint`, and `make check` cover the narrow contract path
- health and ATHENA presence contracts now live under the `ashton/...` package layout Buf expects
- generated Go packages are tracked and compile through a consumer-style import test
- Tracer 2 adds one narrow identified-arrival event contract and JSON schema for `athena.identified_presence.arrived`
- Tracer 2 closure hardening adds a shared Go helper for `athena.identified_presence.arrived` so producers and consumers reuse one subject constant, one schema-backed marshal/parse path, and one fixture set
- this repo is ready for a `v0.2.1` tracer-close tag

See:

- `docs/roadmap.md`
- `docs/runbooks/contract-changes.md`
- `docs/growing-pains.md`
