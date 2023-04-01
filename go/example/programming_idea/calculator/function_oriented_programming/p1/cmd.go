package main

import "fmt"

func CmdCal(op *string, left, right *float64) {
	// 1. 检查参数。
	if op == nil || left == nil || right == nil {
		fmt.Println("参数错误。")
		return
	}
	leftNum, rightNum := *left, *right

	// 2. 执行业务逻辑。
	switch *op {
	case "+":
		fmt.Printf("%f %s %f = %f\n", leftNum, *op, rightNum, leftNum+rightNum)
	case "-":
		fmt.Printf("%f %s %f = %f\n", leftNum, *op, rightNum, leftNum-rightNum)
	case "*":
		fmt.Printf("%f %s %f = %f\n", leftNum, *op, rightNum, leftNum*rightNum)
	case "/":
		if rightNum == 0 {
			fmt.Println("除数为0。")
			return
		}
		fmt.Printf("%f %s %f = %f\n", leftNum, *op, rightNum, leftNum/rightNum)
	default:
		fmt.Printf("不支持的符号(%s)。\n", *op)
		return
	}
}

func CmdHelp(op *string) {
	// 1. 检查参数。
	if op == nil {
		fmt.Println("参数错误。")
		return
	}

	// 2. 执行业务逻辑。
	switch *op {
	case "+":
		fmt.Printf("%s: 加号。\n", *op)
	case "-":
		fmt.Printf("%s: 减号。\n", *op)
	case "*":
		fmt.Printf("%s: 乘号。\n", *op)
	case "/":
		fmt.Printf("%s: 除号。\n", *op)
	default:
		fmt.Printf("不支持的符号(%s)。\n", *op)
		return
	}
}

func CmdUnknown(cmd *string) {
	fmt.Printf("错误命令(%s)。\n", *cmd)
}
