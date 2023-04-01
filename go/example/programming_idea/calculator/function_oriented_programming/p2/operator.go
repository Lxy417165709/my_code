package main

import "fmt"

// -------------------------------------------------------------------------

type HandlerForCalOp = func(op string, left, right float64)

var HandlerForCalOpMap = map[string]HandlerForCalOp{
	"+": HandlerForCalOpAdd,
	"-": HandlerForCalOpSub,
	"*": HandlerForCalOpMul,
	"/": HandlerForCalOpDiv,
}

func HandlerForCalOpAdd(op string, left, right float64) {
	fmt.Printf("%f %s %f = %f\n", left, op, right, left+right)
}

func HandlerForCalOpSub(op string, left, right float64) {
	fmt.Printf("%f %s %f = %f\n", left, op, right, left-right)
}

func HandlerForCalOpMul(op string, left, right float64) {
	fmt.Printf("%f %s %f = %f\n", left, op, right, left*right)
}

func HandlerForCalOpDiv(op string, left, right float64) {
	if right == 0 {
		fmt.Println("除数为0。")
		return
	}
	fmt.Printf("%f %s %f = %f\n", left, op, right, left/right)
}

func HandlerForCalOpUnknown(op string, left, right float64) {
	fmt.Printf("不支持的符号(%s)。\n", op)
}

// -------------------------------------------------------------------------

// -------------------------------------------------------------------------

type HandlerForHelpOp = func(op string)

var HandlerForHelpOpMap = map[string]HandlerForHelpOp{
	"+": HandlerForHelpOpAdd,
	"-": HandlerForHelpOpSub,
	"*": HandlerForHelpOpMul,
	"/": HandlerForHelpOpDiv,
}

func HandlerForHelpOpAdd(op string) {
	fmt.Printf("%s: 加号。\n", op)
}

func HandlerForHelpOpSub(op string) {
	fmt.Printf("%s: 减号。\n", op)
}

func HandlerForHelpOpMul(op string) {
	fmt.Printf("%s: 乘号。\n", op)
}

func HandlerForHelpOpDiv(op string) {
	fmt.Printf("%s: 除号。\n", op)
}

func HandlerForHelpOpUnknown(op string) {
	fmt.Printf("不支持的符号(%s)。\n", op)
}

// -------------------------------------------------------------------------
