# Changelog

All notable changes to this project will be documented in this file.

This project follows semantic versioning for tags.

---

## v0.2.0 — 2026-02-21

### Added
- `profile` mechanism (default `core`; allowed: `core|ledger`)
- `ledger` profile v0.1 rules and `ledger_append` action
- Deterministic exit codes (SSOT): `0=PASS`, `10=FAIL`, `20=BLOCKED`
- Ledger test cases under `testdata/v0_2` + automated validation tests

### Changed
- CLI runner supports `-in <case.json>`
- SPEC updated to document `profile` and `ledger` rules

---

## v0.1.3 — 2026-02-21

### Added
- Baseline core validation rules:
  - JSON must be readable
  - `case_id`, `action`, `payload` required
  - `case_id` format `C-YYYYMMDD-###`
  - `payload` must be an object
  - `action` whitelist for core: `validate`
- Deterministic exit code contract:
  - `0=PASS`, `10=FAIL`, `20=BLOCKED`

### Notes
- This was the last pre-profile release. `v0.2.0` introduces the profile expansion point.
