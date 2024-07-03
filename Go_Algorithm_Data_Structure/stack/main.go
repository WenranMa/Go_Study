package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	arr []int
}

func (this *Stack) Push(val int) {
	this.arr = append(this.arr, val)
}

func (this *Stack) Pop() (val int, err error) {
	if len(this.arr) == 0 {
		return -1, errors.New("stack empty")
	}
	val = this.arr[len(this.arr)-1]
	this.arr = this.arr[:len(this.arr)-1]
	return val, nil
}

func (this *Stack) List() {
	fmt.Println("栈的情况如下：")
	fmt.Println(this.arr)
}

func main() {
	stack := &Stack{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)

	stack.List()
	val, _ := stack.Pop()
	fmt.Println("出栈val=", val) // 5
	stack.List()               //

	fmt.Println()
	val, _ = stack.Pop()
	val, _ = stack.Pop()
	val, _ = stack.Pop()
	val, _ = stack.Pop()
	val, _ = stack.Pop()       // 出错
	fmt.Println("出栈val=", val) // 5
	stack.List()               //
}
