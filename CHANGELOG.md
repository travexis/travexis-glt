# Changelog

All notable changes to this project will be documented in this file.

This project follows semantic versioning for tags.

---

## v0.1.3 — 2026-02-21

### Added
- GitHub Actions CI workflow to run deterministic builds and the full test suite.
- Strict validation rules:
  - `case_id` required and must match `C-YYYYMMDD-###`
  - `action` required, non-empty, and whitelisted (`validate` only)
  - `payload` required and must be a JSON object
- Expanded test suite to cover:
  - invalid JSON format
  - invalid `case_id` format
  - missing `action`
  - missing `payload`
  - unknown `action`
  - usage blocked scenarios

### Changed
- Documented fixed exit code semantics in `SPEC.md`:
  - `0` PASS
  - `10` FAIL
  - `20` BLOCKED
- Documented test suite contract and runner usage in `TEST_SUITE.md`.

### Removed
- Removed build artifacts (`*.exe`) from version control and added `.gitignore`.

---

## v0.1.2 — 2026-02-21

### Added
- Initial CI-ready structure: spec + tests + Go reference validator + Go runner.

---

## v0.1.0 — 2026-02-21

### Added
- Initial repository with minimal deterministic validator behavior and seed tests.