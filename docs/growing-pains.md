# Growing Pains

Use this document to record real contract mistakes, schema mismatches, breaking
change scares, and the fixes that made the shared contracts more durable.

## 2026-04-01

- Early planning tried to lock detailed event payloads before any producer existed.
  We corrected this by locking the event envelope and subject naming first, while
  leaving `data` intentionally flexible for the first wave.
