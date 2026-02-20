# GLT Test Suite (v0.1.0)

Tests are located under `tests/`.

Each test case folder MUST contain:
- `input.*` (e.g. input.json)
- `expected_exit.txt` (single integer)

Exit codes:
- 0  PASS
- 10 FAIL
- 20 BLOCKED

The runner is the reference way to execute the suite.