TGL Specification (v1.0.x)

TGL (Travexis Gate Language) is a deterministic validation protocol.

The validator MUST return a single exit code.
Exit code is the only authoritative result.

Stdout is informational only.

1. Input Requirements

Input MUST be a readable JSON file.

If:

no input argument is provided

file cannot be read

→ exit 20 (BLOCKED)

If:

JSON is invalid

→ exit 10 (FAIL)

2. Required Fields

A valid TGL input MUST contain:

case_id (string)

action (string)

payload (object)

2.1 case_id

Format MUST be:

C-YYYYMMDD-###

Example:

C-20260221-001

Rules:

Must start with C-

Date portion must be numeric

Sequence must be numeric

If format invalid → exit 10 (FAIL)
If missing → exit 20 (BLOCKED)

2.2 action

Must exist

Must be a non-empty string

If missing → exit 20 (BLOCKED)
If wrong type → exit 10 (FAIL)

Allowed actions are profile-dependent.

2.3 payload

Must exist

Must not be null

Must be a JSON object

If missing → exit 20 (BLOCKED)
If wrong type → exit 10 (FAIL)

3. Exit Code Semantics (SSOT)

Exit codes are fixed and MUST NOT be overloaded:

0 = PASS

10 = FAIL

20 = BLOCKED

BLOCKED (20)

Structural or deterministic policy violation:

Missing required file

Chain mismatch

Hash mismatch

Core lock mismatch

Negative forbidden values

Deterministic invariant violation

FAIL (10)

Readable but invalid:

Invalid format

Unknown profile

Unknown action

Wrong type

Missing required field inside valid structure

4. Determinism

Given the same input file,
a conforming TGL validator MUST always return the same exit code.

No network calls.
No time-dependent logic.
No randomness.

Validation MUST be deterministic.

Profiles (v0.2+)

profile is optional.
If missing or empty, default = core.

Allowed profiles:

core

ledger

If profile is present but not allowed → exit 10 (FAIL).

Core Profile (v1.0.x)

Allowed actions:

validate

verdict_core

Core Action: validate

Basic structural validation only.

Core Action: verdict_core (Constitutional Invariant)

This action defines the adjudication boundary of TGL.

Required fields inside payload:

core_lock_sha256 (string)

case_id (string)

profile (string)

action (string)

payload (object)

proof_hash (string)

5. Core Lock Enforcement (Locked)

core_lock_sha256 MUST exist
→ missing = 10 (FAIL)

core_lock_sha256 MUST equal the content of repository root file:

core_lock.sha256

→ mismatch = 20 (BLOCKED)

This prevents unauthorized rule modification.

6. Proof Hash Enforcement (Locked)

For action verdict_core:

proof_hash MUST equal:

"sha256:" + sha256(
CanonicalizeV01({
case_id,
profile,
action,
payload
})
)

Canonicalization algorithm = CanonicalizeV01

Requirements:

Deterministic UTF-8 encoding

Stable key sorting

No whitespace ambiguity

Rules:

Missing proof_hash → 10 (FAIL)

Mismatch → 20 (BLOCKED)

This establishes a cryptographic adjudication boundary.

Ledger Profile (v0.2)

Allowed action:

ledger_append

Ledger Payload Requirements

Inside payload:

ledger_seq (int)

prev_hash (string)

entry_type (enum): append | correction | snapshot

amount_usd (number, must be >= 0)

cfu_impact (number, must be >= 0)

responsible_party (enum): issuer | operator | holder

evidence_ref (string)

ledger_head (object):

ledger_seq (int)

line_hash (string)

Ledger Deterministic Constraints

Hard BLOCKED (20):

ledger_seq != ledger_head.ledger_seq + 1

prev_hash != ledger_head.line_hash

amount_usd < 0

cfu_impact < 0

Missing required field → 10 (FAIL)

Constitutional Status (v1.0.x)

The following are constitutional invariants:

Exit code SSOT

Deterministic validation

CanonicalizeV01 algorithm

core_lock enforcement

proof_hash enforcement

Any implementation claiming v1.0.x compliance MUST enforce these invariants.