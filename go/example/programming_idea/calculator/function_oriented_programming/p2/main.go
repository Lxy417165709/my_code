package main

import (
	"flag"
	"fmt"
)

func main() {
	// 1. 定义参数。
	cmd := flag.String("cmd", "", "命令")
	op := flag.String("op", "", "操作符")
	left := flag.Float64("left", 0.0, "左操作数")
	right := flag.Float64("right", 0.0, "右操作数")

	// 2. 解析。
	flag.Parse()

	// 3. 检查命令。
	if cmd == nil {
		fmt.Println("命令为空。")
		return
	}

	// 4. 根据不同的命令，执行不同的逻辑。
	if *cmd == "cal" {
		CmdCal(op, left, right)
	} else if *cmd == "help" {
		CmdHelp(op)
	} else {
		CmdUnknown(cmd)
	}
}
