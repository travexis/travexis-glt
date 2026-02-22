package tgl

import (
	"crypto/sha256"
	"encoding/hex"
	"regexp"
)

var reCaseID = regexp.MustCompile(`^C-\d{8}-\d{3}$`)

func isCaseID(s string) bool { return reCaseID.MatchString(s) }

func asString(v any) (string, bool) {
	s, ok := v.(string)
	return s, ok
}

func asFloat(v any) (float64, bool) {
	f, ok := v.(float64) // json.Unmarshal numbers -> float64
	return f, ok
}

func asInt(v any) (int64, bool) {
	f, ok := v.(float64)
	if !ok {
		return 0, false
	}
	if f != float64(int64(f)) {
		return 0, false
	}
	return int64(f), true
}

// sha256Hex returns lowercase hex SHA256 of bytes.
func sha256Hex(b []byte) string {
	h := sha256.Sum256(b)
	return hex.EncodeToString(h[:])
}