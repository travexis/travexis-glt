package main

import (
	"encoding/json"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		os.Exit(20) // BLOCKED
	}

	file := os.Args[1]

	data, err := os.ReadFile(file)
	if err != nil {
		os.Exit(20) // BLOCKED
	}

	var obj map[string]interface{}

	err = json.Unmarshal(data, &obj)
	if err != nil {
		os.Exit(10) // FAIL (bad format)
	}

	if _, ok := obj["case_id"]; !ok {
		os.Exit(20) // BLOCKED (missing required field)
	}

	os.Exit(0) // PASS
}