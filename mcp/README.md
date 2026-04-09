# MCP

This directory holds the shared MCP-style tool manifest layer for routed gateway
tools.

Current inventory:

- [`athena.get_current_occupancy.json`](athena.get_current_occupancy.json)
  - source service: `athena`
  - route target: `GET /api/v1/presence/count`
  - required input: `facility_id`
  - read-only: `true`
- [`athena.get_current_zone_occupancy.json`](athena.get_current_zone_occupancy.json)
  - source service: `athena`
  - route target: `GET /api/v1/presence/count`
  - required inputs: `facility_id`, `zone_id`
  - read-only: `true`

Current status:

- present in the current working line
- still intentionally smaller than a broad shared tool catalog
- intentionally limited to two read-only ATHENA routes

Rules for this directory:

- add a manifest only when a real routed tool exists
- do not add speculative write manifests before approval runtime exists
- keep manifest expansion tracer-driven, just like proto and event growth
- a new externally consumed manifest surface here is a pre-`1.0.0` `MINOR`
  change, not a `PATCH`
- docs, tests, and non-breaking manifest clarifications here are `PATCH` work
