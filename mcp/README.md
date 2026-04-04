# MCP

This directory holds the shared MCP-style tool manifest layer for routed gateway
tools.

Current inventory:

- [`athena.get_current_occupancy.json`](athena.get_current_occupancy.json)
  - source service: `athena`
  - route target: `GET /api/v1/presence/count`
  - required input: `facility_id`
  - read-only: `true`

Current status:

- real on `main`
- not yet part of a shipped tag
- intentionally limited to one read-only ATHENA route

Rules for this directory:

- add a manifest only when a real routed tool exists
- do not add speculative write manifests before approval runtime exists
- keep manifest expansion tracer-driven, just like proto and event growth
