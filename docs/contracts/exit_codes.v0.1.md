# Exit Codes (v0.1)

This repo uses a strict exit-code contract for all CLI validation commands.

- 0  = OK (validation passed)
- 10 = INVALID (schema/validation failed)
- 20 = USAGE/SYSTEM (missing args, wrong invocation, or other usage error)