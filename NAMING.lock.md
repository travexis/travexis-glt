# NAMING LOCK (TGL v1.0)

This repository is the reference implementation + corpus for:

- Public protocol name: **TGL**
- Internal machine name: **tgl** (lowercase only)

Hard rules (non-negotiable):
1) The string `glt` MUST NOT appear in any path, package name, module path, import path, schema id, CLI flag, env var, or docs.
2) No mixed-case machine names. Only `tgl`.
3) Any future compatibility aliases (if ever) must live in an explicit, isolated shim folder and MUST NOT affect core_lock or proof_hash inputs.