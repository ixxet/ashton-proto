# MCP

This directory holds shared MCP tool manifest definitions.

First-wave scope:

- [`athena.get_current_occupancy.json`](athena.get_current_occupancy.json)
  - source service: `athena`
  - route target: `GET /api/v1/presence/count`
  - required input: `facility_id`
  - read-only: `true`

Write-path manifests and broader gateway-facing tool expansion are deferred
until the first read slices are real.
