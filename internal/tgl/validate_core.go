package tgl

func validateCore(action string, payload map[string]any) (int, string) {
	// v1.0 core semantics (baseline)
	switch action {
	case "validate":
		_ = payload // reserved
		return 0, "[PASS] core"
		case "verdict_core":
		// Required fields for v0.1 verdict_core
		// We compute: proof_hash = "sha256:" + SHA256(canonical(verdict_core))
		vc := map[string]any{
			"case_id":  payload["case_id"],
			"profile":  payload["profile"],
			"action":   payload["action"],
			"payload":  payload["payload"],
		}

		// Field presence checks (strict)
		if vc["case_id"] == nil {
			return 10, "[FAIL] missing payload.case_id"
		}
		if vc["profile"] == nil {
			return 10, "[FAIL] missing payload.profile"
		}
		if vc["action"] == nil {
			return 10, "[FAIL] missing payload.action"
		}
		if vc["payload"] == nil {
			return 10, "[FAIL] missing payload.payload"
		}

		canon, err := CanonicalizeV01(vc)
		if err != nil {
			return 10, "[FAIL] canonicalize_v0_1 error"
		}
				want := "sha256:" + sha256Hex(canon)

		gotAny, ok := payload["proof_hash"]
		if !ok {
			return 10, "[FAIL] missing payload.proof_hash"
		}
		got, ok := gotAny.(string)
		if !ok || got == "" {
			return 10, "[FAIL] payload.proof_hash must be string"
		}
		if got != want {
			return 20, "[BLOCKED] proof_hash mismatch"
		}

		return 0, "[PASS] verdict_core"
	default:
		return 10, "[FAIL] action not allowed for core (allowed: validate, verdict_core)"
	}
}