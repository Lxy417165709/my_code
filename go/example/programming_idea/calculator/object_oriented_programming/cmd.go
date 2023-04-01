package main

import "fmt"

type Cmd interface {
	Do()
}

type CmdCal struct {
	Op          *string
	Left, Right *float64
}

func (c *CmdCal) Do() {
	// 1. 检查参数。
	if c.Op == nil || c.Left == nil || c.Right == nil {
		fmt.Println("参数错误。")
		return
	}

	// 2. 执行业务逻辑。
	operator := OperatorMap[*c.Op]
	if operator == nil {
		operator = &OperatorUnknown{}
	}
	operator.Cal(*c.Op, *c.Left, *c.Right)
}

type CmdHelp struct {
	Op *string
}

func (c *CmdHelp) Do() {
	// 1. 检查参数。
	if c.Op == nil {
		fmt.Println("参数错误。")
		return
	}

	// 2. 执行业务逻辑。
	operator := OperatorMap[*c.Op]
	if operator == nil {
		operator = &OperatorUnknown{}
	}
	operator.Help(*c.Op)
}

type CmdUnknown struct {
	Cmd *string
}

func (c *CmdUnknown) Do() {
	fmt.Printf("错误命令(%s)。\n", *c.Cmd)
}
