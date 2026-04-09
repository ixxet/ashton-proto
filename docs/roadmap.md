# ashton-proto Roadmap

## Objective

Keep `ashton-proto` small, active, and tracer-driven so shared contracts only
expand when a real cross-repo slice needs them.

## Current Line

Current shipped line: `v0.4.0`

Current released line: `v0.4.0`

- common shared contract baseline is real
- ATHENA read contracts are real
- identified arrival and departure event schemas and runtime helpers are real
- two real ATHENA occupancy manifests are present in the current released line
- active downstream consumers reuse shared helpers instead of private event
  structs

`v0.4.0` stays intentionally narrow after release:

- manifest semantics are intentionally supported for that line
- `buf generate` stays reproducible and leaves the tree clean
- manifest tests cover the narrow supported surface, not only happy-path keys

## Planned Release Lines

| Planned tag | Intended purpose | Restrictions | What it should not do yet |
| --- | --- | --- | --- |
| `v0.5.0` | later cross-repo contract expansion | only add contracts that a real tracer needs | do not turn this repo into a speculative schema dump |

## Boundaries

- do not add broad speculative schemas for future repos just to feel complete
- do not expand manifests before the routed tool surface is actually real
- treat this repo as formal pre-`1.0.0` semver now: `PATCH` for non-breaking
  clarifications and tooling, `MINOR` for any externally consumed shared
  contract or manifest addition or breaking change
- keep runtime helpers on the active wire paths only

## Tracer / Workstream Ownership

- `Tracer 1`: first ATHENA read contract line
- `Tracer 2`: identified-arrival event line
- `Tracer 5`: identified-departure event line
- `Tracer 9`: first ATHENA MCP manifest line
- `Tracer 15`: second ATHENA MCP manifest line for caller-aware audited gateway routing
- later gateway lines: broader routed manifest expansion only when those routes are real
