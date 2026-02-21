package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/travexis/travexis-glt/internal/glt"
)

func main() {
	var inPath string
	flag.StringVar(&inPath, "in", "", "path to case json")
	flag.Parse()

	if inPath == "" {
		fmt.Fprintln(os.Stderr, "[FAIL] -in is required")
		os.Exit(10)
	}

	data, err := os.ReadFile(inPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "[FATAL] read file:", err)
		os.Exit(20)
	}

	var c glt.Case
	if err := json.Unmarshal(data, &c); err != nil {
		fmt.Fprintln(os.Stderr, "[FAIL] bad json:", err)
		os.Exit(10)
	}

	exit, msg := glt.ValidateCase(&c)

	// never silent on nonzero
	if msg == "" && exit != 0 {
		msg = fmt.Sprintf("[FAIL] exit=%d (no message)", exit)
	}

	if msg != "" {
		if exit == 0 {
			fmt.Fprintln(os.Stdout, msg)
		} else {
			fmt.Fprintln(os.Stderr, msg)
		}
	}

	os.Exit(exit)
}