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