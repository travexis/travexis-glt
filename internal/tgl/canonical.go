package tgl

import (
	"bytes"
	"encoding/json"
	"errors"
	"sort"
)

// CanonicalizeV01 produces the canonical bytes used by proof_hash.
// v0.1 rule: canonical JSON with:
// - UTF-8
// - object keys sorted lexicographically
// - no insignificant whitespace
// - arrays kept in given order
// - numbers: reject NaN/Inf; otherwise keep JSON number form as parsed
//
// NOTE: This function MUST be deterministic across implementations.
// Any change requires new version CanonicalizeV02 + new golden corpus.
func CanonicalizeV01(v any) ([]byte, error) {
	return canonicalizeJSONV01(v)
}

func canonicalizeJSONV01(v any) ([]byte, error) {
	// Canonicalize by walking decoded JSON types:
	// - map[string]any
	// - []any
	// - string, bool, nil
	// - json.Number
	b, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	dec := json.NewDecoder(bytes.NewReader(b))
	dec.UseNumber()

	var x any
	if err := dec.Decode(&x); err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	if err := writeCanonicalV01(&buf, x); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func writeCanonicalV01(buf *bytes.Buffer, v any) error {
	switch t := v.(type) {
	case nil:
		buf.WriteString("null")
		return nil
	case bool:
		if t {
			buf.WriteString("true")
		} else {
			buf.WriteString("false")
		}
		return nil
	case string:
		// json.Marshal on string gives correct escaping
		b, _ := json.Marshal(t)
		buf.Write(b)
		return nil
	case json.Number:
		s := t.String()
		// Reject NaN/Inf explicitly (shouldn't appear in JSON, but be strict)
		if s == "NaN" || s == "Infinity" || s == "-Infinity" {
			return errors.New("invalid number")
		}
		buf.WriteString(s)
		return nil
	case []any:
		buf.WriteByte('[')
		for i := range t {
			if i > 0 {
				buf.WriteByte(',')
			}
			if err := writeCanonicalV01(buf, t[i]); err != nil {
				return err
			}
		}
		buf.WriteByte(']')
		return nil
	case map[string]any:
		keys := make([]string, 0, len(t))
		for k := range t {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		buf.WriteByte('{')
		for i, k := range keys {
			if i > 0 {
				buf.WriteByte(',')
			}
			kb, _ := json.Marshal(k)
			buf.Write(kb)
			buf.WriteByte(':')
			if err := writeCanonicalV01(buf, t[k]); err != nil {
				return err
			}
		}
		buf.WriteByte('}')
		return nil
	default:
		return errors.New("unsupported json type")
	}
}