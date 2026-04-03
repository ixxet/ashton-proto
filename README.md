# ashton-proto

Shared contracts repo for the ASHTON platform.

> Nothing in this repo should become a runnable app. Its job is to define the
> shared wire contracts, validation rules, and runtime helpers that keep the
> service repos from drifting.

This repo is already part of the real executable stack. `athena` publishes the
current identified-arrival and identified-departure events through the shared
helpers here, and `apollo` consumes those same helpers instead of maintaining
its own private JSON shape.

## Why This Repo Exists

`ashton-proto` exists so the platform can stay contract-first without stopping
at documentation theater. The current goal is not to create a giant speculative
schema catalog. The goal is to keep the real active surface small, versioned,
and reusable across repos.

## Architecture

The standalone Mermaid source for this flow lives at
[`docs/diagrams/ashton-proto-contract-flow.mmd`](docs/diagrams/ashton-proto-contract-flow.mmd).

```mermaid
flowchart LR
  briefs["ashton-platform<br/>repo briefs and tracer rules"]
  proto["proto/ashton/...<br/>shared protobuf contracts"]
  schema["events/*.schema.json<br/>event envelope and payload rules"]
  helper["events/*.go<br/>runtime marshal, parse, validate"]
  gen["gen/go/...<br/>generated Go packages"]
  fixtures["shared fixtures and contract tests"]
  mcp["mcp/<br/>shared tool manifests<br/>planned expansion"]
  athena["athena<br/>producer and contract consumer"]
  apollo["apollo<br/>consumer and contract consumer"]
  hermes["hermes<br/>future consumer"]
  gateway["ashton-mcp-gateway<br/>future manifest and contract consumer"]

  briefs --> proto
  briefs --> schema
  proto --> gen
  proto --> helper
  schema --> helper
  helper --> fixtures
  gen --> athena
  gen --> apollo
  helper --> athena
  helper --> apollo
  proto -. future surface .-> hermes
  mcp -. future manifests .-> gateway
  proto -. future contracts .-> gateway
```

## Current Contract Surface

| Surface | Path | Status | Purpose |
| --- | --- | --- | --- |
| Common health proto | [`proto/ashton/common/v1/health.proto`](proto/ashton/common/v1/health.proto) | Real | Shared health contract baseline |
| ATHENA read proto | [`proto/ashton/athena/v1/athena.proto`](proto/ashton/athena/v1/athena.proto) | Real | Presence source enums, occupancy types, and the first read RPC |
| Event envelope schema | [`events/envelope.schema.json`](events/envelope.schema.json) | Real | Shared outer event shape and subject naming discipline |
| Identified-arrival schema | [`events/athena.identified_presence.arrived.schema.json`](events/athena.identified_presence.arrived.schema.json) | Real | Active arrival event payload for visit opening |
| Identified-departure schema | [`events/athena.identified_presence.departed.schema.json`](events/athena.identified_presence.departed.schema.json) | Real | Active departure event payload for visit closing |
| Runtime helpers | [`events/identified_presence_arrived.go`](events/identified_presence_arrived.go), [`events/identified_presence_departed.go`](events/identified_presence_departed.go) | Real | Shared marshal, parse, source mapping, and timestamp validation |
| Generated Go packages | `gen/go/...` | Real | Consumer import path for Go services |
| MCP manifests | [`mcp/`](mcp/) | Real, narrow | Shared manifest layer now exists for one ATHENA occupancy route |
| SQL naming guidance | [`sql/naming.md`](sql/naming.md) | Real | Cross-repo relational naming conventions |

## Tech Stack

| Layer | Technology | Status | Line | Notes |
| --- | --- | --- | --- | --- |
| Contract definition | Protobuf + Buf | Instituted | `v0.0.x` -> `v0.3.0` | The package layout is now Buf-clean and generation is reproducible |
| Event validation | JSON Schema 2020-12 | Instituted | `v0.2.x` -> `v0.3.0` | Active event schemas validate the first cross-repo subjects |
| Runtime enforcement | Go helpers + explicit timestamp parsing | Instituted | `v0.2.x` -> `v0.3.0` | Schema validation alone is not trusted for contract-critical semantics |
| Generated consumers | Go generated code | Instituted | `v0.0.x` -> `v0.3.0` | `athena` and `apollo` import generated packages from this repo |
| Test discipline | Go tests + shared fixtures | Instituted | `v0.0.x` -> `v0.3.0` | Repos should reuse shared fixture bytes instead of copying JSON strings |
| Tool manifest layer | MCP manifests | Instituted | `v0.3.1` | The first real manifest line exists and stays ATHENA occupancy only at first |
| Broader routed manifest coverage | MCP manifests + tracer-owned expansion | Planned | `v0.4.0` | Expand only when a second routed read actually lands |
| Later contract expansion | additive proto, schema, and helper growth | Deferred | `v0.5.0` | Only widen when a real cross-repo tracer requires it |

## Ownership Rules

| Rule | Why It Exists |
| --- | --- |
| Lock the event envelope before widening payload detail | Keeps the first slices stable without speculating ahead of real producers |
| Keep subject names in `{service}.{entity}.{action}` form | Makes event routing and ownership obvious |
| Prefer additive change inside `v1` | Avoids unnecessary version sprawl while the surface is still small |
| Publish a runtime helper when a cross-repo message becomes active | Producers and consumers should not maintain private copies of the same wire contract |
| Keep repo expansion tracer-driven | New contracts should exist because a real slice needs them, not because a future repo might someday want them |

## Current State Block

### Already real in this repo

- `buf.yaml` and `buf.gen.yaml` are active and lint-clean
- generated Go code is tracked and compile-checked through consumer-style tests
- the first shared event envelope is locked
- `athena.identified_presence.arrived` and
  `athena.identified_presence.departed` are defined as both schema and shared
  runtime helpers
- shared fixture bytes and validation tests exist for the active visit
  lifecycle event paths
- `mcp/athena.get_current_occupancy.json` now defines the first real shared
  manifest-backed gateway tool line

### Real and active across repos

- `athena` publishes identified-arrival and identified-departure events
  through this repo's helpers
- `apollo` parses those same events through this repo's helpers
- runtime timestamp and source validation now happen in one place instead of
  being recopied in each service

### Planned next

The planned release lines below are the authoritative expansion path. These
bullets are only the short summary.

- expand contract surfaces only when a real tracer requires them
- add broader proto and manifest coverage for `apollo`, `hermes`, and later
  gateway routes only when those executable slices are real

### Deferred on purpose

- broad speculative schemas for features that do not yet have a tracer
- gateway-wide manifest expansion before any routed tool is real
- version churn for changes that are still additive inside the current surface

## Release History

| Release line | Exact tags | Status | What became real | What stayed deferred |
| --- | --- | --- | --- | --- |
| `v0.0.x` | `v0.0.1` | Shipped | common scaffold, Buf baseline, generated Go path | active cross-repo event lines and manifests |
| `v0.1.x` | `v0.1.0` | Shipped | first ATHENA read contract line | lifecycle events and manifests |
| `v0.2.x` | `v0.2.0`, `v0.2.1` | Shipped | identified-arrival event schema and shared helper line | departure contract and manifests |
| `v0.3.0` | `v0.3.0` | Shipped | identified-departure event schema and shared helper line | MCP manifest runtime remains deferred |

## Planned Release Lines

| Planned tag | Intended purpose | Restrictions | What it should not do yet |
| --- | --- | --- | --- |
| `v0.4.0` | broader routed manifest expansion for later gateway lines | expand only when a second routed read actually exists | do not add speculative manifests for unreal service routes |
| `v0.5.0` | later cross-repo contract expansion only when a real tracer needs it | stay tracer-driven and additive where possible | do not turn this repo into a speculative schema dump |

## Versioning And Drift Prevention

| Concern | Current Decision |
| --- | --- |
| Breaking changes | Avoid them while the active surface is still narrow; add new versions only when a real incompatibility exists |
| Producer and consumer drift | Use shared runtime helpers, not repo-local JSON structs |
| Timestamp and enum validation | Keep schema validation, then parse contract-critical values explicitly in runtime code |
| Test fixtures | Generate or reuse shared bytes from `ashton-proto` instead of duplicating hand-written payloads downstream |
| Contract expansion | Tie it to tracer scope so the repo stays small and defensible |

## Project Structure

| Path | Purpose |
| --- | --- |
| `proto/` | shared protobuf contracts |
| `events/` | event schemas, helper code, and fixtures |
| `gen/` | generated language bindings |
| `mcp/` | shared manifest layer for routed gateway tools |
| `tests/` | contract import and schema validation checks |
| `docs/` | roadmap, runbook, ADR index, growing-pains log, and diagrams |

## Docs Map

- [Contract flow diagram](docs/diagrams/ashton-proto-contract-flow.mmd)
- [Roadmap](docs/roadmap.md)
- [Growing pains](docs/growing-pains.md)
- [Contract changes runbook](docs/runbooks/contract-changes.md)
- [ADR index](docs/adr/README.md)
- [Events overview](events/README.md)
- [Proto overview](proto/README.md)
- [MCP overview](mcp/README.md)
- [SQL naming conventions](sql/naming.md)

## Why This Repo Matters

For the platform itself, this repo keeps the first real cross-repo flow honest.
For the engineering narrative, it shows a stronger habit than "we wrote some
services and hoped the payloads matched": contracts are authored once, enforced
once, and reused everywhere that matters.
