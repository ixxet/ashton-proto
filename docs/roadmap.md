# ashton-proto Roadmap

## Objective

Create the smallest shared contract repo that lets ATHENA start cleanly without guessing interfaces later.

## First Implementation Slice

- define shared common types
- define the first ATHENA proto messages
- define the standard event envelope
- define one initial ATHENA MCP manifest

## Boundaries

- do not add HERMES, APOLLO, or gateway contracts beyond what is needed to keep naming and versioning coherent
- do not create broad speculative schemas just to feel complete
- keep versioning simple and explicit from day one

## Exit Criteria

- ATHENA can import generated contract types from this repo
- the first ATHENA read path has a stable contract
- event naming is fixed for the first tracer bullet
- the repo has a clear path for later expansion without breaking the first slice
