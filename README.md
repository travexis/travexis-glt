Travexis TGL

TGL (Travexis Gate Language) is a deterministic validation standard.

Current stable line: v1.0.x

TGL defines:

A minimal specification

A fixed exit code contract (SSOT)

Deterministic validation rules

A reproducible test suite

A reference validator implementation

A constitutional adjudication boundary (verdict_core)

Exit Code Contract (SSOT)

TGL defines exactly three exit codes:

0 → PASS
10 → FAIL
20 → BLOCKED

Exit code is the single source of truth.

Stdout is informational only and MUST NOT override exit semantics.

Exit codes are immutable and MUST NOT be reinterpreted.

Constitutional Layer (v1.0.x)

Starting at v1.0.0, TGL introduces a constitutional adjudication boundary.

Core profile action:

verdict_core

This action enforces:

core_lock_sha256 validation against repository root core_lock.sha256

proof_hash validation using CanonicalizeV01

Deterministic invariant enforcement

These rules define the stability boundary of TGL.

Any implementation claiming v1.0.x compliance MUST enforce:

Deterministic validation

Exit code SSOT (0 / 10 / 20)

CanonicalizeV01 invariance

core_lock enforcement

proof_hash enforcement

Tags from v1.0.0 onward are immutable.

Profiles

profile is optional.
If missing or empty, default = core.

Allowed profiles:

core

ledger

Unknown profile → exit 10 (FAIL)

Repository Structure

engine_go/
Reference validator implementation.

runner_go/
Test runner for executing structured test cases.

internal/tgl/
Core validation logic.

test_cases/
Deterministic test suite (golden vectors).

SPEC.md
Formal specification.

TEST_SUITE.md
Test contract and golden vector definitions.

CHANGELOG.md
Release history.

RELEASE_POLICY.md
Tag immutability and versioning rules.