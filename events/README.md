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
