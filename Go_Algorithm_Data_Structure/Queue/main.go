package main

import (
	"errors"
	"fmt"
	"os"
)

// 使用一个结构体管理队列
type Queue struct {
	Array []int // Array=>模拟队列
}

// 添加数据到队尾
func (this *Queue) Push(val int) {
	this.Array = append(this.Array, val)
	return
}

// 从队列头部取出数据
func (this *Queue) Pop() (val int, err error) {
	if len(this.Array) == 0 {
		return -1, errors.New("queue empty")
	}
	val = this.Array[0]
	this.Array = this.Array[1:]
	return
}

// 显示队列, 找到队首，然后到遍历到队尾
func (this *Queue) ShowQueue() {
	fmt.Println("队列当前的情况是:")
	fmt.Println(this.Array)
}
func main() {
	queue := &Queue{}
	var key string
	var val int
	for {
		fmt.Println("1. 输入add 表示添加数据到队列")
		fmt.Println("2. 输入get 表示从队列获取数据")
		fmt.Println("3. 输入show 表示显示队列")
		fmt.Println("4. 输入exit 表示显示队列")

		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("输入你要入队列数")
			fmt.Scanln(&val)
			queue.Push(val)
			fmt.Println("加入队列ok")
		case "get":
			val, err := queue.Pop()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("从队列中取出了一个数=", val)
			}
		case "show":
			queue.ShowQueue()
		case "exit":
			os.Exit(0)
		}
	}
}
