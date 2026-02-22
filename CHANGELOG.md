Changelog

All notable changes to this project will be documented in this file.

This project follows semantic versioning for tags.

v1.0.3 — 2026-02-22
Added

Constitutional enforcement for verdict_core

core_lock_sha256 enforcement against repository root core_lock.sha256

proof_hash enforcement using CanonicalizeV01

Golden vector test coverage for:

core_lock mismatch

missing core_lock

proof_hash mismatch

missing proof_hash

Locked

Exit code SSOT (0 = PASS, 10 = FAIL, 20 = BLOCKED)

Deterministic validation requirement

CanonicalizeV01 algorithm invariance

core_lock enforcement boundary

proof_hash enforcement boundary

Notes

v1.0.x establishes the adjudication layer.

Tags from v1.0.0 onward are immutable.

v1.0.2 — 2026-02-22
Fixed

Clean build state for core profile wiring

Resolved validateCore signature mismatch

Ensured go test ./... passes deterministically

v1.0.1 — 2026-02-22
Changed

Stabilized core + ledger profile routing

Hardened deterministic exit behavior

v1.0.0 — 2026-02-22
Added

verdict_core action (constitutional boundary)

profile mechanism (default core; allowed: core|ledger)

Deterministic exit codes (SSOT):
0 = PASS
10 = FAIL
20 = BLOCKED

Notes

First immutable release line.

Tags from this version forward MUST NOT be moved.

v0.2.0 — 2026-02-21
Added

ledger profile v0.1 rules and ledger_append action

Ledger test cases under testdata/v0_2

Automated validation tests for ledger rules

Changed

CLI runner supports -in <case.json>

SPEC updated to document profile and ledger rules

v0.1.3 — 2026-02-21
Added

Baseline core validation rules:

JSON must be readable

case_id, action, payload required

case_id format C-YYYYMMDD-###

payload must be an object

action whitelist for core: validate

Deterministic exit code contract:

0 = PASS

10 = FAIL

20 = BLOCKED

Notes

Last pre-profile release.

v0.2.0 introduced profile expansion.