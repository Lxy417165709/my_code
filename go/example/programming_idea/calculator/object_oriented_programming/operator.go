package main

import "fmt"

var OperatorMap = map[string]Operator{
	"+": &OperatorAdd{},
	"-": &OperatorSub{},
	"*": &OperatorMul{},
	"/": &OperatorDiv{},
}

type Operator interface {
	Cal(op string, left, right float64)
	Help(op string)
}

type OperatorAdd struct{}

func (o *OperatorAdd) Cal(op string, left, right float64) {
	fmt.Printf("%f %s %f = %f\n", left, op, right, left+right)
}

func (o *OperatorAdd) Help(op string) {
	fmt.Printf("%s: 加号。\n", op)
}

type OperatorSub struct{}

func (o *OperatorSub) Cal(op string, left, right float64) {
	fmt.Printf("%f %s %f = %f\n", left, op, right, left-right)
}

func (o *OperatorSub) Help(op string) {
	fmt.Printf("%s: 减号。\n", op)
}

type OperatorMul struct{}

func (o *OperatorMul) Cal(op string, left, right float64) {
	fmt.Printf("%f %s %f = %f\n", left, op, right, left*right)
}

func (o *OperatorMul) Help(op string) {
	fmt.Printf("%s: 乘号。\n", op)
}

type OperatorDiv struct{}

func (o *OperatorDiv) Cal(op string, left, right float64) {
	if right == 0 {
		fmt.Println("除数为0。")
		return
	}
	fmt.Printf("%f %s %f = %f\n", left, op, right, left/right)
}

func (o *OperatorDiv) Help(op string) {
	fmt.Printf("%s: 除号。\n", op)
}

type OperatorUnknown struct{}

func (o *OperatorUnknown) Cal(op string, left, right float64) {
	fmt.Printf("不支持的符号(%s)。\n", op)
}

func (o *OperatorUnknown) Help(op string) {
	fmt.Printf("不支持的符号(%s)。\n", op)
}
