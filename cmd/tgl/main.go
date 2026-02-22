package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/travexis/travexis-tgl/internal/tgl"
)

const (
	EXIT_OK      = 0
	EXIT_INVALID = 10
	EXIT_USAGE   = 20
)

func main() {
	fs := flag.NewFlagSet("validator", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)

	var inPath string
	fs.StringVar(&inPath, "in", "", "path to case json")

	if err := fs.Parse(os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "[FAIL] usage error: %v\n", err)
		os.Exit(EXIT_USAGE)
	}

	// Compatibility: allow positional arg: validator <file>
	if inPath == "" {
		if rest := fs.Args(); len(rest) >= 1 {
			inPath = rest[0]
		}
	}

	if inPath == "" {
		fmt.Fprintln(os.Stderr, "[FAIL] -in is required (or provide <file> as positional arg)")
		os.Exit(EXIT_USAGE)
	}

	data, err := os.ReadFile(inPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "[FATAL] read file:", err)
		os.Exit(EXIT_USAGE)
	}

	// --- Pre-check required fields at JSON level (runner contract expects USAGE=20) ---
	var raw map[string]any
	if err := json.Unmarshal(data, &raw); err != nil {
		fmt.Fprintln(os.Stderr, "[FAIL] bad json:", err)
		os.Exit(EXIT_INVALID)
	}

	// action: must exist and be non-empty string
	act, ok := raw["action"]
	if !ok {
		fmt.Fprintln(os.Stderr, "[FAIL] missing action")
		os.Exit(EXIT_USAGE)
	}
	actStr, ok := act.(string)
	if !ok || strings.TrimSpace(actStr) == "" {
		fmt.Fprintln(os.Stderr, "[FAIL] missing action")
		os.Exit(EXIT_USAGE)
	}

	// payload: must exist and not be null
	if v, ok := raw["payload"]; !ok || v == nil {
		fmt.Fprintln(os.Stderr, "[FAIL] missing payload")
		os.Exit(EXIT_USAGE)
	}

	// --- Now parse into typed Case and validate ---
	var c tgl.Case
	if err := json.Unmarshal(data, &c); err != nil {
		fmt.Fprintln(os.Stderr, "[FAIL] bad json:", err)
		os.Exit(EXIT_INVALID)
	}

	exit, msg := tgl.ValidateCase(&c)

	// Upgrade specific validation messages to USAGE=20 (blocked usage contract)
	// (kept narrow to avoid touching real INVALID cases)
	if exit == EXIT_INVALID && msg != "" {
		m := strings.ToLower(msg)
		if strings.Contains(m, "blocked_usage") ||
			strings.Contains(m, "blocked usage") ||
			strings.Contains(m, "missing action") ||
			strings.Contains(m, "missing payload") {
			exit = EXIT_USAGE
		}
	}

	// never silent on nonzero
	if msg == "" && exit != 0 {
		msg = fmt.Sprintf("[FAIL] exit=%d (no message)", exit)
	}

	if msg != "" {
		if exit == EXIT_OK {
			fmt.Fprintln(os.Stdout, msg)
		} else {
			fmt.Fprintln(os.Stderr, msg)
		}
	}

	// Clamp to contract
	switch exit {
	case EXIT_OK, EXIT_INVALID, EXIT_USAGE:
		os.Exit(exit)
	default:
		if exit == 0 {
			os.Exit(EXIT_OK)
		}
		fmt.Fprintf(os.Stderr, "[FAIL] unexpected exit=%d (clamped to %d)\n", exit, EXIT_USAGE)
		os.Exit(EXIT_USAGE)
	}
}