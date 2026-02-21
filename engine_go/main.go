package main

import (
	"encoding/json"
	"os"
	"strings"
)

func main() {

	if len(os.Args) < 2 {
		os.Exit(20)
	}

	file := os.Args[1]

	data, err := os.ReadFile(file)
	if err != nil {
		os.Exit(20)
	}

	var obj map[string]interface{}

	// Bad JSON
	err = json.Unmarshal(data, &obj)
	if err != nil {
		os.Exit(10)
	}

	// case_id
	rawCaseID, ok := obj["case_id"]
	if !ok {
		os.Exit(20)
	}
	caseID, ok := rawCaseID.(string)
	if !ok {
		os.Exit(10)
	}

	if len(caseID) != 14 ||
		!strings.HasPrefix(caseID, "C-") ||
		caseID[10] != '-' {
		os.Exit(10)
	}

	digitPositions := []int{2,3,4,5,6,7,8,9,11,12,13}
	for _, pos := range digitPositions {
		if caseID[pos] < '0' || caseID[pos] > '9' {
			os.Exit(10)
		}
	}

	// action required
	rawAction, ok := obj["action"]
	if !ok {
		os.Exit(20)
	}
	action, ok := rawAction.(string)
	if !ok {
		os.Exit(10)
	}
	action = strings.TrimSpace(action)
	if action == "" {
		os.Exit(20)
	}

	// whitelist: only "validate"
	if action != "validate" {
		os.Exit(10)
	}

	// payload required and must be object
	rawPayload, ok := obj["payload"]
	if !ok || rawPayload == nil {
		os.Exit(20)
	}
	_, ok = rawPayload.(map[string]interface{})
	if !ok {
		os.Exit(10)
	}

	os.Exit(0)
}