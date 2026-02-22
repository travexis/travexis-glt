@"
# golden2 (PASS)

Purpose: simulate different working directories (cwd immunity test).

Run:
- from repo root
- from nested subdir

Expected:
- PASS
- exit code = 0
- verdict identical across cwd
"@ | Set-Content -Encoding utf8 .\tests\golden2\README.md