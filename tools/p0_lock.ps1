New-Item -ItemType Directory -Force -Path .\tools | Out-Null

@'
# tools/p0_lock.ps1  (Final Locked v0.1, PS5.1-safe)
Set-StrictMode -Version Latest
$ErrorActionPreference = "Stop"

function Fail([string]$msg) { Write-Host "[FAIL] $msg" -ForegroundColor Red; exit 2 }
function Ok([string]$msg) { Write-Host "[OK] $msg" -ForegroundColor Green }

try { $repoRoot = (git rev-parse --show-toplevel).Trim() } catch { Fail "not in a git repo" }
if (-not $repoRoot) { Fail "not in a git repo" }

Push-Location $repoRoot
try {
  if (-not (Test-Path ".\out")) { New-Item -ItemType Directory -Force -Path ".\out" | Out-Null }

  Ok "build: glt + runner"
  & go build -o .\out\glt.exe .\cmd\glt
  if ($LASTEXITCODE -ne 0) { Fail "go build glt failed (exit=$LASTEXITCODE)" }

  & go build -o .\out\runner.exe .\runner_go
  if ($LASTEXITCODE -ne 0) { Fail "go build runner failed (exit=$LASTEXITCODE)" }

  Ok "run: conformance suite"
  & .\out\runner.exe .\out\glt.exe .\tests 2>&1 | Tee-Object -FilePath .\out\runner_last.txt
  if ($LASTEXITCODE -ne 0) { Fail "runner suite failed. See .\out\runner_last.txt" }
  Ok "conformance: PASS"

  $porcelain = (git status --porcelain)
  if ($porcelain) {
    Write-Host ""
    Write-Host "[WARN] working tree dirty:" -ForegroundColor Yellow
    $porcelain | ForEach-Object { Write-Host "  $_" -ForegroundColor Yellow }
    Write-Host ""
    Fail "dirty working tree: commit+push required"
  }
  Ok "git: working tree clean"

  & git fetch --quiet
  if ($LASTEXITCODE -ne 0) { Fail "git fetch failed" }

  $sb = (git status -sb)
  if ($sb -match "\[ahead\s+\d+\]") { Fail "branch ahead of origin: git push required" }
  Ok "git: pushed (not ahead of origin)"

  Write-Host ""
  Ok "P0-LOCAL-LOCK complete. Safe to close this window."
  exit 0
}
finally { Pop-Location }
'@ | Set-Content -Encoding utf8 -LiteralPath .\tools\p0_lock.ps1