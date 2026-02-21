# Release Policy

This repo uses semantic versioning for tags: `vMAJOR.MINOR.PATCH`.

## What changes the version
- PATCH: docs / CI / refactor only. No validator behavior change.
- MINOR: backward-compatible behavior change (new profile/action/rule/test).
- MAJOR: breaking change to input contract, semantics, or exit code meanings.

## Tag immutability
- During incubation (pre-`v1.0.0`), version tags (`v0.x.y`) may be force-moved if needed to keep the release coherent.
- Starting at `v1.0.0`, version tags are immutable (no force updates).

## Immutable anchor tags (required during incubation)
- If a `v0.x.y` tag is force-moved, create an immutable anchor tag `r-YYYYMMDD-HHMM` pointing to the previous release commit.
- Anchor tags are never force-moved.

### Standard procedure for moving a version tag (incubation only)

1) Create an immutable anchor tag that points to the current version tag target:

```powershell
$old = (git rev-parse v0.2.0).Trim()
$stamp = Get-Date -Format "yyyyMMdd-HHmm"
git tag "r-$stamp" $old
git push origin "r-$stamp"