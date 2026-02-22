# Travexis TGL

**TGL (Travexis Gate Language)** is a deterministic validation standard.

## Current Stable Line: **v0.2-golden1**

TGL defines:

- A minimal specification
- A fixed exit code contract (SSOT)
- Deterministic validation rules
- A reproducible test suite
- A reference validator implementation
- A constitutional adjudication boundary (**verdict_core**)

### **Exit Code Contract (SSOT)**

TGL defines exactly three exit codes:

- **0 → PASS**
- **10 → FAIL**
- **20 → BLOCKED**

Exit code is the single source of truth.

Stdout is informational only and **MUST NOT override exit semantics**.

Exit codes are immutable and **MUST NOT be reinterpreted**.

---

### **Constitutional Layer (v1.0.x)**

Starting at **v1.0.0**, TGL introduces a constitutional adjudication boundary.

#### **Core Profile Action**:
- **verdict_core**: Enforces `core_lock_sha256` validation, `proof_hash` validation, and deterministic invariant enforcement.

---

### **New Features in v0.2-golden1**:

- **Golden Vector Test Coverage** for `core_lock` and `proof_hash` enforcement.
- Ensured that all files have consistent **LF line endings** to prevent cross-platform issues.
- Updated **`input.txt`** and **`expected_exit.txt`** files for **golden2/3** to match expected test outcomes.

---

### **Profiles**

- **profile** is optional.
  - If missing or empty, default = **core**.
- **Allowed profiles**:
  - **core**
  - **ledger**
  - **Unknown profile** → exit **10** (FAIL)

---

### **Repository Structure**

- **engine_go/**: Reference validator implementation.
- **runner_go/**: Test runner for executing structured test cases.
- **internal/tgl/**: Core validation logic.
- **test_cases/**: Deterministic test suite (golden vectors).
- **SPEC.md**: Formal specification.
- **TEST_SUITE.md**: Test contract and golden vector definitions.
- **CHANGELOG.md**: Release history.
- **RELEASE_POLICY.md**: Tag immutability and versioning rules.