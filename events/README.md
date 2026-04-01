# Events

This directory holds shared event schemas for the ASHTON platform.

First rule: lock the envelope before locking detailed payloads.

For the first implementation wave:

- lock the standard envelope
- lock subject naming convention
- keep the `data` payload flexible until real producing slices exist
