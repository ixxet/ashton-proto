# Events

This directory holds the shared event contracts that are already active in the
ASHTON stack.

Current inventory:

- `envelope.schema.json`
  - shared outer event shape
- `athena.identified_presence.arrived.schema.json`
  - active arrival payload schema
- `athena.identified_presence.departed.schema.json`
  - active departure payload schema
- `identified_presence_arrived.go`
  - shared marshal, parse, source mapping, and timestamp validation for arrival
- `identified_presence_departed.go`
  - shared marshal, parse, source mapping, and timestamp validation for departure
- fixture files and tests
  - valid bytes and rejection cases for the current active subjects

Rules for this directory:

- lock the envelope before widening payload detail
- keep subject naming in `{service}.{entity}.{action}` form
- when an event becomes active across repos, add one shared helper here and
  make both producer and consumer import it
- adding or breaking an externally consumed event/helper surface here is a
  pre-`1.0.0` `MINOR` change, not a `PATCH`
- docs, tests, and non-breaking validation/tooling fixes here are `PATCH` work
