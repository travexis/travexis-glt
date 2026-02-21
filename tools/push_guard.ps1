@'
# tools/push_guard.ps1 (Final Locked v0.1, PS5.1-safe)
Set-StrictMode -Version Latest
$ErrorActionPreference = "Stop"

function Fail([string]$msg) { Write-Host "[FAIL] $msg" -ForegroundColor Red; exit 2 }
function Ok([string]$msg) { Write-Host "[OK] $msg" -ForegroundColor Green }

$repoRoot = (git rev-parse --show-toplevel).Trim()
if (-not $repoRoot) { Fail "not in a git repo" }

Push-Location $repoRoot
try {
  $dirty = (git status --porcelain)
  if ($dirty) { Fail "dirty working tree: commit required" }

  & git fetch --quiet
  if ($LASTEXITCODE -ne 0) { Fail "git fetch failed" }

  $sb = (git status -sb)
  if ($sb -match "\[ahead\s+\d+\]") { Fail "branch ahead of origin: push required" }

  Ok "clean + pushed"
  exit 0
}
finally { Pop-Location }
'@ | Set-Content -Encoding utf8 -LiteralPath .\tools\push_guard.ps1