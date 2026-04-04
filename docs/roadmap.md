# ashton-proto Roadmap

## Objective

Keep `ashton-proto` small, active, and tracer-driven so shared contracts only
expand when a real cross-repo slice needs them.

## Current Line

Current shipped line: `v0.3.0`

Current unreleased working line on `main`: `v0.3.1`

- common shared contract baseline is real
- ATHENA read contracts are real
- identified arrival and departure event schemas and runtime helpers are real
- the first real ATHENA occupancy manifest is real on `main`
- active downstream consumers reuse shared helpers instead of private event
  structs

## Planned Release Lines

| Planned tag | Intended purpose | Restrictions | What it should not do yet |
| --- | --- | --- | --- |
| `v0.4.0` | broader routed manifest expansion for later gateway lines | only expand when a second routed read actually lands | do not add speculative tool manifests for unreal services |
| `v0.5.0` | later cross-repo contract expansion | only add contracts that a real tracer needs | do not turn this repo into a speculative schema dump |

## Boundaries

- do not add broad speculative schemas for future repos just to feel complete
- do not expand manifests before the routed tool surface is actually real
- keep additive change inside the current version line when a breaking change is
  not truly needed
- keep runtime helpers on the active wire paths only

## Tracer / Workstream Ownership

- `Tracer 1`: first ATHENA read contract line
- `Tracer 2`: identified-arrival event line
- `Tracer 5`: identified-departure event line
- `Tracer 9`: first ATHENA MCP manifest line
- later gateway lines: broader routed manifest expansion only when those routes
  are real
