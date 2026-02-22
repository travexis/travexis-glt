TGL Test Suite (v1.0.x)

This repository provides a deterministic test suite for TGL.

All test cases live under:

test_cases/

Each test case is a directory.
The directory name is the test ID.

Test Case Contract

Each test directory MUST contain:

input.*
The input file passed to the validator (typically input.json).

expected_exit.txt
A single integer representing the expected exit code.

Exit codes (fixed and authoritative):

0 PASS
10 FAIL
20 BLOCKED

Stdout is informational only.
Exit code is the single source of truth (SSOT).

How to Run

2.1 Build (Windows)

go build -o validator.exe ./engine_go
go build -o runner.exe ./runner_go

2.2 Run

.\runner.exe .\test_cases\your_test_case_directory

Golden Vectors (v1.0.x)

These tests lock constitutional invariants.

golden1 — Verdict Core Enforcement (PASS)

Requirements:

core_lock_sha256 exists

core_lock_sha256 matches repository root file core_lock.sha256

proof_hash matches canonicalized payload

Expected exit: 0

golden2 — core_lock_sha256 mismatch (BLOCKED)

If:

core_lock_sha256 does not equal repository core_lock.sha256

Expected exit: 20

golden3 — Missing core_lock_sha256 (FAIL)

If:

core_lock_sha256 is missing

Expected exit: 10

golden4 — proof_hash mismatch (BLOCKED)

If:

proof_hash does not match:

"sha256:" + sha256(
CanonicalizeV01({
case_id,
profile,
action,
payload
})
)

Expected exit: 20

golden5 — Missing proof_hash (FAIL)

If:

proof_hash is missing

Expected exit: 10

Example Test Case Layout

Each test case directory must follow this structure:

test_case_id/
input.json
expected_exit.txt

Example:

golden1/
input.json
expected_exit.txt

Conformance Notes (v1.0.x)

Any TGL implementation claiming v1.0.x compliance MUST enforce:

Deterministic validation

Exit code SSOT (0 / 10 / 20)

core_lock_sha256 enforcement

proof_hash enforcement

CanonicalizeV01 algorithm invariance

Violation handling must follow SSOT rules:

Structural invariant violation → 20 (BLOCKED)

Missing required field → 10 (FAIL)

Invalid but readable input → 10 (FAIL)

Implementations that do not enforce these invariants are NOT v1.0.x compliant.