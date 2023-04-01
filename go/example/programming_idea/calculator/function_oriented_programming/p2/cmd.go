package main

import "fmt"

func CmdCal(op *string, left, right *float64) {
	// 1. 检查参数。
	if op == nil || left == nil || right == nil {
		fmt.Println("参数错误。")
		return
	}

	// 2. 执行业务逻辑。
	handlerForCalOp := HandlerForCalOpMap[*op]
	if handlerForCalOp == nil {
		handlerForCalOp = HandlerForCalOpUnknown
	}
	handlerForCalOp(*op, *left, *right)
}

func CmdHelp(op *string) {
	// 1. 检查参数。
	if op == nil {
		fmt.Println("参数错误。")
		return
	}

	// 2. 执行业务逻辑。
	handlerForHelpOp := HandlerForHelpOpMap[*op]
	if handlerForHelpOp == nil {
		handlerForHelpOp = HandlerForHelpOpUnknown
	}
	handlerForHelpOp(*op)
}

func CmdUnknown(cmd *string) {
	fmt.Printf("错误命令(%s)。\n", *cmd)
}
