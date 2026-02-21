# GLT Specification (v0.1)

GLT (Gate Language for Travexis) is a deterministic validation protocol.

The validator MUST return a single exit code.
Exit code is the only authoritative result.

Stdout is informational only.

---

## 1. Input Requirements

Input MUST be a readable JSON file.

If:
- no input argument is provided
- file cannot be read

→ exit `20` (BLOCKED)

If:
- JSON is invalid

→ exit `10` (FAIL)

---

## 2. Required Fields

A valid GLT input MUST contain:

- `case_id` (string)
- `action` (string)
- `payload` (object)

### 2.1 case_id

Format MUST be:

C-YYYYMMDD-###

Example:

C-20260221-001

Rules:
- Total length = 14
- Must start with `C-`
- Character at position 10 must be `-`
- All other specified positions must be numeric

If format invalid → exit `10` (FAIL)

If missing → exit `20` (BLOCKED)

---

### 2.2 action

- Must exist
- Must be a non-empty string
- MUST match whitelist

v0.1 whitelist:

validate

If missing → exit `20` (BLOCKED)

If wrong type or unknown action → exit `10` (FAIL)

---

### 2.3 payload

- Must exist
- Must not be null
- Must be a JSON object

If missing → exit `20` (BLOCKED)

If wrong type → exit `10` (FAIL)

---

## 3. Exit Code Semantics

GLT exit codes have fixed meanings:

- `20` (BLOCKED): structural preconditions are not met.  
  Missing required fields, unreadable input file, empty required string.

- `10` (FAIL): input is readable but invalid by GLT rules.  
  Invalid JSON, wrong field type, format mismatch, unknown action.

- `0` (PASS): input is valid under GLT rules.

Exit code MUST NOT be overloaded.

---

## 4. Determinism

Given the same input file,
a conforming GLT validator MUST always return the same exit code.

No network calls.
No time-dependent logic.
No randomness.

GLT validation MUST be deterministic.