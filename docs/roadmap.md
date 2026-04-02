# ashton-proto Roadmap

## Objective

Create the smallest shared contract repo that lets ATHENA start cleanly without guessing interfaces later.

## First Implementation Slice

- define shared common types
- define the first ATHENA proto messages
- define the standard event envelope
- define one initial ATHENA MCP manifest
- keep inner event payloads flexible until the first producing adapters exist

## Boundaries

- do not add HERMES, APOLLO, or gateway contracts beyond what is needed to keep naming and versioning coherent
- do not create broad speculative schemas just to feel complete
- keep versioning simple and explicit from day one

## Exit Criteria

- ATHENA can import generated contract types from this repo
- the first ATHENA read path has a stable contract
- event naming is fixed for the first tracer bullet
- the repo has a clear path for later expansion without breaking the first slice

## Current State

Tracer 1 now has the minimum reproducible contract baseline:

- Buf lint and generation pass on the narrow health and ATHENA read contracts
- generated Go packages compile through `make check`
- the first package layout and RPC naming rules are locked without widening payload detail
- Tracer 2 extends the ATHENA contract surface with one identified-arrival event payload, one strict event schema, and one shared runtime helper so active producers and consumers do not drift on the JSON wire shape
- Tracer 5 extends that same lifecycle with one identified-departure event payload, one strict event schema, and one shared runtime helper so visit closing stays on the shared contract path too

## Tracer Ownership

- `Tracer 1`: owns the first active contract surface for ATHENA
- `Tracer 2`: owns the identified-arrival event contract used to turn physical presence into APOLLO visit history
- `Tracer 5`: owns the identified-departure event contract used to close APOLLO visit history deterministically
- later tracers should only expand this repo when a real cross-repo slice needs it
