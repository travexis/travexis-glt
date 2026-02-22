package tgl

import (
	"os"
	"path/filepath"
	"strings"
)

func validateCore(action string, payload map[string]any) (int, string) {
	// Only support verdict_core in v1.0.x core.
	if action != "verdict_core" {
		return 10, "[FAIL] unknown core action (allowed: verdict_core)"
	}

	// --- core_lock enforcement ---
	wantLock, err := readCoreLock()
	if err != nil {
		return 20, "[BLOCK] core_lock read error"
	}
	gotLock, ok := payload["core_lock_sha256"].(string)
	if !ok || strings.TrimSpace(gotLock) == "" {
		return 10, "[FAIL] missing core_lock_sha256"
	}
	if strings.TrimSpace(gotLock) != wantLock {
		return 20, "[BLOCK] core_lock_sha256 mismatch"
	}

	// --- proof_hash enforcement ---
	// required fields for verdict_core canonicalization
	caseID, ok := payload["case_id"].(string)
	if !ok || strings.TrimSpace(caseID) == "" {
		return 10, "[FAIL] missing case_id"
	}
	profile, ok := payload["profile"].(string)
	if !ok || strings.TrimSpace(profile) == "" {
		return 10, "[FAIL] missing profile"
	}
	act2, ok := payload["action"].(string)
	if !ok || strings.TrimSpace(act2) == "" {
		return 10, "[FAIL] missing action"
	}
	pl2, ok := payload["payload"].(map[string]any)
	if !ok {
		return 10, "[FAIL] missing payload (object)"
	}

	canon, err := CanonicalizeV01(map[string]any{
		"case_id": caseID,
		"profile": profile,
		"action":  act2,
		"payload": pl2,
	})
	if err != nil {
		return 20, "[BLOCK] canonicalize error"
	}
	wantProof := "sha256:" + sha256Hex(canon)

	gotProof, ok := payload["proof_hash"].(string)
	if !ok || strings.TrimSpace(gotProof) == "" {
		return 10, "[FAIL] missing proof_hash"
	}
	if strings.TrimSpace(gotProof) != wantProof {
		return 20, "[BLOCK] proof_hash mismatch"
	}

	return 0, "[PASS] verdict_core"
}

func readCoreLock() (string, error) {
	// validate_core.go lives in internal/tgl, so repo root is ../..
	p := filepath.Join("..", "..", "core_lock.sha256")
	b, err := os.ReadFile(p)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(b)), nil
}