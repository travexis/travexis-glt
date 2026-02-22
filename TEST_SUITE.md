# TGL Test Suite (v0.1.x)

This repository provides a deterministic test suite for TGL.](image.png)
All tests live under:

Each test case is a directory. The directory name is the test id.

---

## 1. Test Case Contract

Each test directory MUST contain:

- `input.*`  
  The input file passed to the validator (typically `input.json`).

- `expected_exit.txt`  
  A single integer representing the expected exit code.

Exit codes (fixed):
- `0`  PASS
- `10` FAIL
- `20` BLOCKED

Stdout is informational only. Exit code is the source of truth.

---

## 2. How to Run

### 2.1 Build and run locally (Windows)

Build:

```powershell
go build -o validator.exe ./engine_go
go build -o runner.exe ./runner_go