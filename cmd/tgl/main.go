package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/travexis/travexis-tgl/internal/tgl"
)

// 定义标准退出码
const (
	EXIT_OK      = 0  // 验证成功
	EXIT_INVALID = 10 // JSON 格式错误或严重验证失败
	EXIT_USAGE   = 20 // 用法错误或缺少必填字段 (action/payload)
)

func main() {
	// 初始化命令行标志解析器，错误输出到 stderr
	fs := flag.NewFlagSet("validator", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)

	var inPath string
	fs.StringVar(&inPath, "in", "", "path to case json")

	// 解析命令行参数
	if err := fs.Parse(os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "[FAIL] usage error: %v\n", err)
		os.Exit(EXIT_USAGE)
	}

	// 兼容性处理：如果 -in 未提供，尝试使用位置参数 <file>
	if inPath == "" {
		if rest := fs.Args(); len(rest) >= 1 {
			inPath = rest[0]
		}
	}

	// 最终检查：必须提供文件路径
	if inPath == "" {
		fmt.Fprintln(os.Stderr, "[FAIL] -in is required (or provide <file> as positional arg)")
		os.Exit(EXIT_USAGE)
	}

	// 读取文件内容
	data, err := os.ReadFile(inPath)
	if err != nil {
		// 统一使用 [FAIL] 前缀
		fmt.Fprintf(os.Stderr, "[FAIL] read file: %v\n", err)
		os.Exit(EXIT_USAGE)
	}

	// --- 第一阶段：JSON 预检 (Pre-check) ---
	// 目的：在类型化解析之前，先检查合同规定的必填字段 (action, payload)。
	// 这样可以将“缺少必填项”归类为 EXIT_USAGE (20)，而不是通用的 EXIT_INVALID (10)。
	var raw map[string]any
	if err := json.Unmarshal(data, &raw); err != nil {
		fmt.Fprintf(os.Stderr, "[FAIL] bad json: %v\n", err)
		os.Exit(EXIT_INVALID)
	}

	// 检查 action 字段：必须存在且为非空字符串
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

	// 检查 payload 字段：必须存在且不为 null
	if v, ok := raw["payload"]; !ok || v == nil {
		fmt.Fprintln(os.Stderr, "[FAIL] missing payload")
		os.Exit(EXIT_USAGE)
	}

	// --- 第二阶段：类型化解析与业务验证 ---
	// 将数据解码为具体的 tgl.Case 结构体
	var c tgl.Case
	if err := json.Unmarshal(data, &c); err != nil {
		fmt.Fprintf(os.Stderr, "[FAIL] bad json: %v\n", err)
		os.Exit(EXIT_INVALID)
	}

	// 执行核心验证逻辑
	exit, msg := tgl.ValidateCase(&c)

	// --- 退出码调整逻辑 ---
	// 如果验证返回 EXIT_INVALID，但错误消息包含特定关键字（如 blocked_usage 或缺少字段），
	// 则将其升级为 EXIT_USAGE，以符合调用方合约。
	if exit == EXIT_INVALID && msg != "" {
		m := strings.ToLower(msg)
		if strings.Contains(m, "blocked_usage") ||
			strings.Contains(m, "blocked usage") ||
			strings.Contains(m, "missing action") ||
			strings.Contains(m, "missing payload") {
			exit = EXIT_USAGE
		}
	}

	// 确保非零退出码时有对应的错误消息
	if msg == "" && exit != 0 {
		msg = fmt.Sprintf("[FAIL] exit=%d (no message)", exit)
	}

	// 输出消息
	if msg != "" {
		if exit == EXIT_OK {
			fmt.Fprintln(os.Stdout, msg)
		} else {
			fmt.Fprintln(os.Stderr, msg)
		}
	}

	// --- 最终退出处理 ---
	// 钳制退出码以确保符合合约定义
	switch exit {
	case EXIT_OK, EXIT_INVALID, EXIT_USAGE:
		os.Exit(exit)
	default:
		// 处理意外的退出码值，强制转换为 EXIT_USAGE
		// 注意：移除了原代码中 unreachable 的 "if exit == 0" 分支
		fmt.Fprintf(os.Stderr, "[FAIL] unexpected exit=%d (clamped to %d)\n", exit, EXIT_USAGE)
		os.Exit(EXIT_USAGE)
	}
}