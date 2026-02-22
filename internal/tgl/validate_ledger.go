package tgl

import "fmt"

func validateLedger(action string, payload map[string]any) (int, string) {
	// v0.2 ledger semantics: allow ledger_append only
	if action != "ledger_append" {
		return 10, "[FAIL] action not allowed for ledger (allowed: ledger_append)"
	}

	ledgerSeq, ok := asInt(payload["ledger_seq"])
	if !ok {
		return 10, "[FAIL] missing/invalid payload.ledger_seq (int)"
	}
	prevHash, ok := asString(payload["prev_hash"])
	if !ok || prevHash == "" {
		return 10, "[FAIL] missing/invalid payload.prev_hash (string)"
	}
	entryType, ok := asString(payload["entry_type"])
	if !ok || (entryType != "append" && entryType != "correction" && entryType != "snapshot") {
		return 10, "[FAIL] bad payload.entry_type (append|correction|snapshot)"
	}
	amountUSD, ok := asFloat(payload["amount_usd"])
	if !ok {
		return 10, "[FAIL] missing/invalid payload.amount_usd (number)"
	}
	cfuImpact, ok := asFloat(payload["cfu_impact"])
	if !ok {
		return 10, "[FAIL] missing/invalid payload.cfu_impact (number)"
	}
	responsible, ok := asString(payload["responsible_party"])
	if !ok || (responsible != "issuer" && responsible != "operator" && responsible != "holder") {
		return 10, "[FAIL] bad payload.responsible_party (issuer|operator|holder)"
	}
	evidenceRef, ok := asString(payload["evidence_ref"])
	if !ok || evidenceRef == "" {
		return 10, "[FAIL] missing/invalid payload.evidence_ref (string)"
	}

	// HARD BLOCK
	if amountUSD < 0 {
		return 20, "[BLOCKED] amount_usd must be >= 0"
	}
	if cfuImpact < 0 {
		return 20, "[BLOCKED] cfu_impact must be >= 0"
	}

	// ledger_head (deterministic previous record injected)
	headAny, ok := payload["ledger_head"]
	if !ok {
		return 10, "[FAIL] missing payload.ledger_head (object)"
	}
	head, ok := headAny.(map[string]any)
	if !ok {
		return 10, "[FAIL] payload.ledger_head must be object"
	}
	headSeq, ok := asInt(head["ledger_seq"])
	if !ok {
		return 10, "[FAIL] payload.ledger_head.ledger_seq must be int"
	}
	headHash, ok := asString(head["line_hash"])
	if !ok || headHash == "" {
		return 10, "[FAIL] payload.ledger_head.line_hash must be string"
	}

	if ledgerSeq != headSeq+1 {
		return 20, fmt.Sprintf("[BLOCKED] ledger_seq must be head+1 (got %d want %d)", ledgerSeq, headSeq+1)
	}
	if prevHash != headHash {
		return 20, "[BLOCKED] prev_hash must match ledger_head.line_hash"
	}

	return 0, "[PASS] ledger"
}