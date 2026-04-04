# Proto

This directory holds the shared protobuf definitions that are already active in
the ASHTON stack.

Current inventory:

- `ashton/common/v1/health.proto`
  - common health response baseline
- `ashton/athena/v1/athena.proto`
  - occupancy read request and response
  - presence source enums
  - shared ATHENA event-related message shapes and enums

Rules for this directory:

- keep package layout rooted under `ashton/...` so Buf lint and generated
  imports stay aligned
- add proto only when a real tracer needs a shared contract
- do not turn this directory into a speculative schema catalog for unreal
  services
