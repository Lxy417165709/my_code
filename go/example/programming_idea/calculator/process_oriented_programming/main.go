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
		// 4.1 检查参数。
		if op == nil || left == nil || right == nil {
			fmt.Println("参数错误。")
			return
		}
		leftNum, rightNum := *left, *right

		// 4.2 执行业务逻辑。
		switch *op {
		case "+":
			fmt.Printf("%f\n", leftNum+rightNum)
		case "-":
			fmt.Printf("%f\n", leftNum-rightNum)
		case "*":
			fmt.Printf("%f\n", leftNum*rightNum)
		case "/":
			if rightNum == 0 {
				fmt.Println("除数为0。")
				return
			}
			fmt.Printf("%f\n", leftNum/rightNum)
		default:
			fmt.Printf("不支持的符号(%s)。\n", *op)
			return
		}
	} else if *cmd == "help" {
		// 4.1 检查参数。
		if op == nil {
			fmt.Println("参数错误。")
			return
		}

		// 4.2 执行业务逻辑。
		switch *op {
		case "+":
			fmt.Println("我是加号。")
		case "-":
			fmt.Println("我是减号。")
		case "*":
			fmt.Println("我是乘号。")
		case "/":
			fmt.Println("我是除号。")
		default:
			fmt.Printf("不支持的符号(%s)。\n", *op)
			return
		}
	} else {
		fmt.Printf("错误命令(%s)。\n", *cmd)
	}
}
