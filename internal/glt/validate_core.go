package glt

func validateCore(action string, payload map[string]any) (int, string) {
	// v0.1 core semantics locked
	if action != "validate" {
		return 10, "[FAIL] action not allowed for core (allowed: validate)"
	}
	_ = payload // reserved
	return 0, "[PASS] core"
}