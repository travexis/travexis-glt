package tgl

type Case struct {
	CaseID  any `json:"case_id"`
	Action  any `json:"action"`
	Profile any `json:"profile"`
	Payload any `json:"payload"`
}

func ValidateCase(c *Case) (exit int, msg string) {
	// required: case_id/action/payload
	if c.CaseID == nil || c.Action == nil || c.Payload == nil {
		return 10, "[FAIL] missing required fields: case_id/action/payload"
	}

	// case_id format
	caseID, ok := c.CaseID.(string)
	if !ok || !isCaseID(caseID) {
		return 10, "[FAIL] bad case_id (expect C-YYYYMMDD-###)"
	}

	// action string
	action, ok := c.Action.(string)
	if !ok {
		return 10, "[FAIL] action must be string"
	}

	// payload object
	payload, ok := c.Payload.(map[string]any)
	if !ok {
		return 10, "[FAIL] payload must be object"
	}

	// profile optional
	profile := "core"
	if c.Profile != nil {
		s, ok := c.Profile.(string)
		if !ok {
			return 10, "[FAIL] profile must be string"
		}
		if s != "" {
			profile = s
		}
	}

	// profile whitelist
	switch profile {
	case "core":
		return validateCore(action, payload)
	case "ledger":
		return validateLedger(action, payload)
	default:
		return 10, "[FAIL] unknown profile (allowed: core, ledger)"
	}
}