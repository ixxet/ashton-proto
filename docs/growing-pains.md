# Growing Pains

Use this document to record real contract mistakes, schema mismatches, breaking
change scares, and the fixes that made the shared contracts more durable.

## 2026-04-01

- Early planning tried to lock detailed event payloads before any producer existed.
  We corrected this by locking the event envelope and subject naming first, while
  leaving `data` intentionally flexible for the first wave.

- The first proto layout looked tidy to humans but failed Buf standard lint
  because the package path and RPC request/response names were too loose. The
  fix was to move the contracts under `proto/ashton/...` and use method-scoped
  request and response names before calling the repo reproducible.

- The first Tracer 2 schema test tried to follow the event schema `$id` over
  HTTPS instead of loading the local shared envelope resource. The fix was to
  preload both schema resources in the Go validation test so local contract
  checks stay deterministic.

## 2026-04-02

- Symptom: `athena` and `apollo` both spoke the identified-arrival event, but
  each repo still owned its own private JSON struct.
  Cause: the proto and schema existed, but nothing in `ashton-proto` gave the
  runtime code one shared marshal and parse surface to import.
  Fix: add a shared Go helper in `events/` that owns the subject constant,
  source-value mapping, schema-backed marshal and parse paths, and reusable
  fixture bytes.
  Rule: when a tracer message becomes active across repos, `ashton-proto` must
  publish a runtime helper or validator so producer and consumer do not hand-roll
  the same wire contract twice.

- Symptom: the shared schema allowed an invalid `recorded_at` string through the
  runtime helper tests.
  Cause: JSON Schema `format: date-time` was not sufficient runtime enforcement
  by itself in this stack.
  Fix: keep schema validation, then explicitly parse envelope and payload
  timestamps in the shared helper.
  Rule: use schema validation for structure, but still parse contract-critical
  timestamps and enums explicitly in runtime code.
