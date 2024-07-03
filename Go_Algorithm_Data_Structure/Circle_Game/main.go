package main

import (
	"fmt"
)

type Boy struct {
	No   int  // 编号
	Next *Boy // 指向下一个小孩的指针[默认值是nil]
}

// Constructor 构成单向的环形链表 num ：表示小孩的个数 返回该环形的链表的第一个小孩的指针
func AddBoy(num int) *Boy {
	var first, curBoy *Boy
	for i := 1; i <= num; i++ {
		boy := &Boy{
			No: i,
		}
		//1. 因为第一个小孩比较特殊
		if i == 1 { //第一个小孩
			first = boy //不要动
			curBoy = boy
			curBoy.Next = first //
		} else {
			curBoy.Next = boy
			curBoy = boy
			curBoy.Next = first //构造环形链表
		}
	}
	return first
}

// 显示单向的环形链表[遍历]
func ShowBoy(first *Boy) {
	if first == nil {
		fmt.Println("链表为空，没有小孩...")
		return
	}
	curBoy := first
	for {
		fmt.Printf("小孩编号=%d ->", curBoy.No)
		//退出的条件?curBoy.Next == first
		if curBoy.Next == first {
			break
		}
		curBoy = curBoy.Next
	}
	fmt.Println()
}

/*
设编号为1，2，… n的n个人围坐一圈，约定编号为k（1<=k<=n）
的人从1开始报数，数到m 的那个人出列，它的下一位又从1开始报数，
数到m的那个人又出列，依次类推，直到所有人出列为止，由此产生一个出队编号的序列
*/
func PlayGame(first *Boy, countNum int) {
	if first == nil {
		fmt.Println("空的链表，没有小孩")
		return
	}
	tail := first //tail 用于指向环形链表的最后一个小孩
	for tail.Next != first {
		tail = tail.Next
	}
	for tail != first {
		//开始数countNum-1次
		for i := 1; i <= countNum-1; i++ {
			first = first.Next
			tail = tail.Next
		}
		fmt.Printf("小孩编号为%d 出圈 \n", first.No)
		//删除first执行的小孩
		first = first.Next
		tail.Next = first
	}
	fmt.Printf("小孩编号为%d 出圈 \n", first.No)
}

func main() {
	first := AddBoy(10)
	ShowBoy(first)
	PlayGame(first, 13)
}
