TGL v1.0.x — Minimal PRD (Locked)

Goal: Stabilize profiles and introduce a constitutional adjudication layer without breaking deterministic guarantees.

Core Principles

Determinism is mandatory.

Exit code SSOT (0 / 10 / 20) is immutable.

Validation must remain offline (no network, no time, no randomness).

All behavior must be reproducible given identical input.

Profiles

Introduce optional profile field in input JSON.

Rules:

Missing profile → treated as "core"

Allowed values:

"core"

"ledger"

If profile is present but not allowed → exit 10 (FAIL)

Core Profile

Default profile.

Core MUST support the following actions:

validate

verdict_core

Action: validate
Performs structural validation only.

Action: verdict_core
Defines the constitutional adjudication boundary.

Constitutional Adjudication Layer (v1.0.x)

For action verdict_core:

Required fields inside payload:

core_lock_sha256 (string)

case_id (string)

profile (string)

action (string)

payload (object)

proof_hash (string)

Enforcement rules:

core_lock_sha256:

Missing → 10 (FAIL)

Mismatch with repository root file core_lock.sha256 → 20 (BLOCKED)

proof_hash:

Must equal:

"sha256:" + sha256(
CanonicalizeV01({
case_id,
profile,
action,
payload
})
)

Missing → 10 (FAIL)

Mismatch → 20 (BLOCKED)

These rules define the adjudication boundary and are constitutional invariants.

Ledger Profile

profile="ledger"

Ledger profile extends validation rules but does not modify core field requirements.

Allowed action:

ledger_append

Ledger profile must:

Remain deterministic and offline

Enforce chain integrity rules

Preserve exit code SSOT

Ledger deterministic constraints:

Hard BLOCKED (20):

ledger_seq != ledger_head.ledger_seq + 1

prev_hash != ledger_head.line_hash

amount_usd < 0

cfu_impact < 0

Missing required field → 10 (FAIL)

Exit Code Contract (Immutable)

0 = PASS
10 = FAIL
20 = BLOCKED

These meanings MUST NOT change in v1.0.x.

Test Requirements

All core tests must pass.

All ledger tests must pass.

Golden vectors must enforce:

core_lock validation

proof_hash validation

CI must remain green before release.

Release Policy

v1.0.x is the first stable adjudication line.

Tags are immutable.

Constitutional invariants cannot be altered without MAJOR version increment.