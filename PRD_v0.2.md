# TGL v0.2 — Minimal PRD (Locked)

Goal: Add optional validation profiles without breaking v0.1 core.

1) Keep v0.1 behavior as default (core profile).
2) Introduce `profile` (string) in input JSON:
   - missing `profile` => treated as `"core"`
   - allowed: `"core"`, `"ledger"`
3) `profile="core"`: identical to v0.1.4 rules and exit codes.
4) `profile="ledger"`: adds ledger-related requirements (no changes to core fields).
5) Ledger profile must remain deterministic and offline (no network/time/random).
6) Ledger requirements are enforced only when profile="ledger".
7) Exit codes stay fixed: 0 PASS, 10 FAIL, 20 BLOCKED.
8) Add two new tests: `pass_ledger/` and `fail_ledger_missing/`.
9) CI must stay green with all tests.
10) Release as v0.2.0 only after core+ledger tests both pass.