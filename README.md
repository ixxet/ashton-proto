# ashton-proto

Shared contracts repo for the ASHTON platform.

This repo will hold the protobuf definitions, event schemas, and MCP tool manifests that every service depends on. Nothing here should become a runnable app. Its job is to make the rest of the platform type-safe, consistent, and easier to evolve without accidental drift.

This folder is intentionally docs-first right now. The implementation brief lives in [ashton-platform/planning/repo-briefs/ashton-proto.md](https://github.com/ixxet/ashton-platform/blob/main/planning/repo-briefs/ashton-proto.md).

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

Docs-first stub only. No proto layout or code generation files have been created yet.
