# NAMING LOCK (TGL v1.0)
Naming Lock (v1.0.x)

This repository is the reference implementation and corpus for:

Public protocol name: TGL
Internal machine name: tgl (lowercase only)

Current stable line: v1.0.x

Constitutional Naming Rules (Non-Negotiable)

The string glt MUST NOT appear anywhere in:

File paths

Folder names

Package names

Module paths

Import paths

Schema identifiers

CLI flags

Environment variables

Documentation

Test fixtures

Machine identifier MUST be lowercase only:

Allowed:

tgl

Disallowed:

glt

Tgl

TGL (as machine identifier)

Public display name MAY use uppercase:

Public protocol name: TGL
Machine identifier: tgl

Compatibility Aliases (Strict Isolation Rule)

If a compatibility alias is ever introduced:

It MUST live inside an isolated shim directory.

It MUST NOT affect:

CanonicalizeV01 inputs

core_lock enforcement

proof_hash computation

Exit code semantics

Deterministic behavior

Constitutional Status (v1.0.x)

Naming rules are part of the v1.0.x compliance boundary.

If violated, the implementation is NOT v1.0.x compliant.