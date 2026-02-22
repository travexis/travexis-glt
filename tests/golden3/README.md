# golden3 (PASS)

Purpose:
Simulate multi-profile execution to verify profile consistency.

How to run:
- Run with profile=core.
- Run with profile=ledger (or secondary profile).

Expected:
- PASS
- Exit code = 0
- Verdict must remain consistent across profiles.

Why it matters:
Profiles may vary, but core_lock and proof boundaries must remain stable.