package main

import (
	"errors"
	"fmt"
	"strconv"
)

type Stack struct {
	arr []int // slice模拟栈
}

func (this *Stack) Push(val int) {
	this.arr = append(this.arr, val)
}

func (this *Stack) Pop() (int, bool) {
	var top int
	if len(this.arr) > 0 {
		top = this.arr[len(this.arr)-1]
		this.arr = this.arr[:len(this.arr)-1]
		return top, true
	}
	return -1, false
}

func (this *Stack) IsEmpty() bool {
	return len(this.arr) == 0
}

func (this *Stack) Top() (int, bool) {
	if len(this.arr) > 0 {
		return this.arr[len(this.arr)-1], true
	}
	return -1, false
}

// 遍历栈，注意需要从栈顶开始遍历
func (this *Stack) List() {
	fmt.Println("栈的情况如下：")
	for i, v := range this.arr {
		fmt.Printf("arr[%d]=%d\n", i, v)
	}
}

// 判断一个字符是不是一个运算符[+, - , * , /]
func (this *Stack) IsOper(val int) bool {
	if val == 42 || val == 43 || val == 45 || val == 47 {
		return true
	}
	return false
}

func (this *Stack) Cal(oper int) (int, error) {
	res := 0
	num1, ok := this.Pop()
	num2, ok := this.Pop()
	if !ok {
		return 0, errors.New("Calcuate failed")
	}
	switch oper {
	case 42:
		res = num2 * num1
	case 43:
		res = num2 + num1
	case 45:
		res = num2 - num1
	case 47:
		res = num2 / num1
	}
	return res, nil
}

// [* / => 1,  + - => 0]
func (this *Stack) Priority(oper int) int {
	if oper == 42 || oper == 47 {
		return 1
	}
	return 0
}

func main() {
	numStack := &Stack{}
	operStack := &Stack{}
	exp := "20+30*3-4-6*3"
	keepNum := []byte{}
	for index := 0; index < len(exp); index += 1 {
		//处理多位数的问题
		ch := exp[index]            // 字符 ch == '+' ===> 43
		temp := int(ch)             // 就是字符对应的ASCiI码
		if operStack.IsOper(temp) { // 说明是符号
			if operStack.IsEmpty() {
				operStack.Push(temp)
			} else {
				top, _ := operStack.Top()
				if operStack.Priority(top) >= operStack.Priority(temp) {
					oper, _ := operStack.Pop()
					result, _ := numStack.Cal(oper)
					numStack.Push(result)
				}
				operStack.Push(temp)
			}
		} else { //说明是数 //处理多位数
			keepNum = append(keepNum, ch)
			if index == len(exp)-1 || operStack.IsOper(int(exp[index+1])) {
				val, _ := strconv.ParseInt(string(keepNum), 10, 64)
				numStack.Push(int(val))
				keepNum = []byte{}
			}
		}
	}
	for !operStack.IsEmpty() {
		oper, _ := operStack.Pop()
		result, _ := numStack.Cal(oper)
		numStack.Push(result)
	}
	res, _ := numStack.Pop()
	fmt.Printf("表达式%s = %v", exp, res)
}
