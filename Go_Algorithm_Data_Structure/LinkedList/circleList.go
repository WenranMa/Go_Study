package main

import (
	"fmt"
)

// 定义猫的结构体结点
type CatNode struct {
	No   int //猫猫的编号
	Name string
	Next *CatNode
}

func InsertCatNode(head *CatNode, newCatNode *CatNode) {
	//判断是不是添加第一只猫
	if head.Next == nil {
		head.No = newCatNode.No
		head.Name = newCatNode.Name
		head.Next = head //构成一个环形
		fmt.Println(newCatNode, "加入到环形的链表")
		return
	}

	//定义一个临时变量，帮忙,找到环形的最后结点
	cur := head
	for cur.Next != head {
		cur = cur.Next
	}
	//加入到链表中
	cur.Next = newCatNode
	newCatNode.Next = head
}

// 输出这个环形的链表
func ListCircleLink(head *CatNode) {
	fmt.Println("环形链表的情况如下：")
	cur := head
	if cur.Next == nil {
		fmt.Println("环形链表为空...")
		return
	}

	// 先打印一下head.
	fmt.Printf("猫的信息为=[id=%d Name=%s] ->\n", cur.No, cur.Name)
	for cur.Next != head {
		cur = cur.Next
		fmt.Printf("猫的信息为=[id=%d Name=%s] ->\n", cur.No, cur.Name)
	}
}

// 删除一只猫
func DelCatNode(head *CatNode, id int) *CatNode {
	cur := head
	//空链表
	if cur.Next == nil {
		fmt.Println("空的环形链表，不能删除")
		return head
	}
	//如果只有一个结点
	if cur.Next == head { //只有一个结点
		if cur.No == id {
			cur.Next = nil
		}
		return head
	}
	//将helper 定位到链表最后
	helper := head
	for {
		if helper.Next == head {
			break
		}
		helper = helper.Next
	}

	//如果有两个包含两个以上结点
	flag := true
	for {
		if cur.Next == head { //如果到这来，说明我比较到最后一个【最后一个还没比较】
			break
		}
		if cur.No == id {
			if cur == head { //说明删除的是头结点
				head = head.Next
			}
			//恭喜找到., 我们也可以在直接删除
			helper.Next = cur.Next
			fmt.Printf("猫猫=%d\n", id)
			flag = false
			break
		}
		cur = cur.Next       //移动 【比较】
		helper = helper.Next //移动 【一旦找到要删除的结点 helper】
	}
	//这里还有比较一次
	if flag { //如果flag 为真，则我们上面没有删除
		if cur.No == id {
			helper.Next = cur.Next
			fmt.Printf("猫猫=%d\n", id)
		} else {
			fmt.Printf("对不起，没有No=%d\n", id)
		}
	}
	return head
}

func main() {
	head := &CatNode{}
	cat1 := &CatNode{
		No:   1,
		Name: "tom",
	}
	cat2 := &CatNode{
		No:   2,
		Name: "tom2",
	}
	cat3 := &CatNode{
		No:   3,
		Name: "tom3",
	}
	InsertCatNode(head, cat1)
	ListCircleLink(head)

	InsertCatNode(head, cat2)
	InsertCatNode(head, cat3)
	ListCircleLink(head)

	head = DelCatNode(head, 2)

	ListCircleLink(head)
}
