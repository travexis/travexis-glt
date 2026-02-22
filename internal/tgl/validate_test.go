package tgl

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func mustReadJSON(t *testing.T, p string) *Case {
	t.Helper()
	b, err := os.ReadFile(p)
	if err != nil {
		t.Fatal(err)
	}
	var c Case
	if err := json.Unmarshal(b, &c); err != nil {
		t.Fatal(err)
	}
	return &c
}

func TestLedgerV02(t *testing.T) {
	type tc struct {
		file string
		want int
	}
	cases := []tc{
		{"pass_ledger_normal.json", 0},
		{"fail_ledger_seq_jump.json", 20},
		{"fail_ledger_prev_hash_wrong.json", 20},
		{"fail_ledger_negative_amount.json", 20},
		{"fail_ledger_missing_evidence_ref.json", 10},
	}

	base := filepath.Join("..", "..", "testdata", "v0_2")
	for _, x := range cases {
		t.Run(x.file, func(t *testing.T) {
			c := mustReadJSON(t, filepath.Join(base, x.file))
			got, _ := ValidateCase(c)
			if got != x.want {
				t.Fatalf("exit=%d want=%d", got, x.want)
			}
		})
	}
}

func TestVerdictCoreV01_Golden1(t *testing.T) {
	// Minimal golden vector to lock canonical + proof_hash rule.
	vc := &Case{
		CaseID:  "C-20260222-001",
		Action:  "verdict_core",
		Profile: "core",
		Payload: map[string]any{
			"case_id": "C-20260222-001",
			"profile": "ledger",
			"action":  "ledger_append",
			"payload": map[string]any{
				"ledger_seq":        float64(42),
				"prev_hash":         "abc123",
				"entry_type":        "append",
				"amount_usd":        float64(10),
				"cfu_impact":        float64(1),
				"responsible_party": "operator",
				"evidence_ref":      "sha256:0000000000000000000000000000000000000000000000000000000000000000",
				"ledger_head": map[string]any{
					"ledger_seq": float64(41),
					"line_hash":  "abc123",
				},
			},
		},
	}

	// Compute expected proof_hash via canonicalizer (same as spec).
	p := vc.Payload.(map[string]any)
	canon, err := CanonicalizeV01(map[string]any{
		"case_id": p["case_id"],
		"profile": p["profile"],
		"action":  p["action"],
		"payload": p["payload"],
	})
	if err != nil {
		t.Fatal(err)
	}
	want := "sha256:" + sha256Hex(canon)
	p["proof_hash"] = want

	got, _ := ValidateCase(vc)
	if got != 0 {
		t.Fatalf("exit=%d want=0", got)
	}

	// Mismatch must BLOCK (20)
	p["proof_hash"] = "sha256:ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"
	got2, _ := ValidateCase(vc)
	if got2 != 20 {
		t.Fatalf("exit=%d want=20", got2)
	}
}