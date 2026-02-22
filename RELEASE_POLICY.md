## Release Policy

This repo uses semantic versioning for tags: vMAJOR.MINOR.PATCH.

### Current stable line: v0.2.x

### What changes the version:

#### PATCH
- Documentation / CI / refactor only.
- No validator behavior change.

#### MINOR
- Backward-compatible behavior change.
- Examples:
  - New profile
  - New action
  - Additional validation rule
  - Additional test vectors

#### MAJOR
- Breaking change to:
  - Input contract
  - Exit code semantics
  - Deterministic rules
  - Constitutional invariants
  - Tag immutability

### Starting at v1.0.0:
- Version tags are immutable.
- No force updates.
- No retagging.
- No moving tags.
- Tags represent permanent validator behavior.

### Immutable anchor tags (pre-v1.0 only):
- During incubation (v0.x.y only):
  - If a version tag must be moved, an immutable anchor tag MUST be created:
    - Format: `r-YYYYMMDD-HHMM`
    - Anchor tags:
      - Point to the previous release commit
      - Are never force-moved
      - Remain permanent historical anchors

This rule does NOT apply to v1.0.0 and later.

### Validator stability requirement (v1.0.x):
For any v1.0.x release:
- The following MUST be true before tagging:
  - `go test ./...` passes
  - `SPEC.md` reflects current constitutional invariants
  - `core_lock.sha256` matches the enforced rules
  - Exit code SSOT remains unchanged (0 / 10 / 20)
- If validator behavior changes:
  - Version MUST increment
  - Tags MUST NOT be rewritten