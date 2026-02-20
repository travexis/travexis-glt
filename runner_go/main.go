package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

func main() {
	validator := "./validator.exe"
	testsDir := "./tests"

	// allow overrides: runner.exe <validator> <testsDir>
	if len(os.Args) >= 2 {
		validator = os.Args[1]
	}
	if len(os.Args) >= 3 {
		testsDir = os.Args[2]
	}

	caseDirs, err := os.ReadDir(testsDir)
	if err != nil {
		fmt.Println("[FATAL] cannot read tests dir:", err)
		os.Exit(2)
	}

	var cases []string
	for _, e := range caseDirs {
		if e.IsDir() {
			cases = append(cases, e.Name())
		}
	}
	sort.Strings(cases)

	total, pass, fail := 0, 0, 0
	var failed []string

	for _, name := range cases {
		total++
		casePath := filepath.Join(testsDir, name)

		expPath := filepath.Join(casePath, "expected_exit.txt")
		expBytes, err := os.ReadFile(expPath)
		if err != nil {
			fail++
			failed = append(failed, fmt.Sprintf("%s missing expected_exit.txt", name))
			fmt.Printf("[FAIL] %s missing expected_exit.txt\n", name)
			continue
		}
		expStr := strings.TrimSpace(string(expBytes))
		expected, err := strconv.Atoi(expStr)
		if err != nil {
			fail++
			failed = append(failed, fmt.Sprintf("%s bad expected_exit.txt=%q", name, expStr))
			fmt.Printf("[FAIL] %s bad expected_exit.txt=%q\n", name, expStr)
			continue
		}

		// find input.*
		input := ""
		entries, _ := os.ReadDir(casePath)
		for _, f := range entries {
			if !f.IsDir() && strings.HasPrefix(strings.ToLower(f.Name()), "input.") {
				input = filepath.Join(casePath, f.Name())
				break
			}
		}
		if input == "" {
			fail++
			failed = append(failed, fmt.Sprintf("%s missing input.*", name))
			fmt.Printf("[FAIL] %s missing input.*\n", name)
			continue
		}

		cmd := exec.Command(validator, input)
		err = cmd.Run()
		actual := cmd.ProcessState.ExitCode()

		if actual == expected {
			pass++
			fmt.Printf("[PASS] %s exit=%d\n", name, actual)
		} else {
			fail++
			failed = append(failed, fmt.Sprintf("%s expected=%d actual=%d", name, expected, actual))
			fmt.Printf("[FAIL] %s expected=%d actual=%d\n", name, expected, actual)
		}
	}

	fmt.Println()
	fmt.Printf("[SUMMARY] total=%d pass=%d fail=%d\n", total, pass, fail)

	if fail > 0 {
		fmt.Println("[FAILED CASES]")
		for _, s := range failed {
			fmt.Println(" -", s)
		}
		os.Exit(1)
	}
	os.Exit(0)
}