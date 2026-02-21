package main

import (
	"encoding/json"
	"os"
	"strings"
)

func main() {

	// No input provided → BLOCKED
	if len(os.Args) < 2 {
		os.Exit(20)
	}

	file := os.Args[1]

	data, err := os.ReadFile(file)
	if err != nil {
		os.Exit(20)
	}

	var obj map[string]interface{}

	// Bad JSON format → FAIL
	err = json.Unmarshal(data, &obj)
	if err != nil {
		os.Exit(10)
	}

	// case_id must exist
	rawCaseID, ok := obj["case_id"]
	if !ok {
		os.Exit(20)
	}

	caseID, ok := rawCaseID.(string)
	if !ok {
		os.Exit(10)
	}

	// Enforce format: C-YYYYMMDD-###
	if len(caseID) != 14 {
		os.Exit(10)
	}
	if !strings.HasPrefix(caseID, "C-") {
		os.Exit(10)
	}
	if caseID[10] != '-' {
		os.Exit(10)
	}
	digitPositions := []int{2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 13}
	for _, pos := range digitPositions {
		if caseID[pos] < '0' || caseID[pos] > '9' {
			os.Exit(10)
		}
	}

	// action must exist and be a non-empty string → otherwise BLOCKED
	rawAction, ok := obj["action"]
	if !ok {
		os.Exit(20)
	}
	action, ok := rawAction.(string)
	if !ok {
		os.Exit(10)
	}
	if strings.TrimSpace(action) == "" {
		os.Exit(20)
	}

	// PASS
	os.Exit(0)
}