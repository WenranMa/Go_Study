package main

import (
	"fmt"
)

type HNode struct {
	No       int
	Name     string
	NickName string
	Next     *HNode //这个表示指向下一个结点
}

// 编写第一种插入方法，在单链表的最后加入.[简单]
func InsertHNodeToTail(head *HNode, newHNode *HNode) {
	for head.Next != nil {
		head = head.Next
	}
	head.Next = newHNode
}

// 编写第2种插入方法，根据No 的编号从小到大插入..【实用】
func InsertHNodeOrder(head *HNode, newHeroNode *HNode) {
	for head.Next != nil && head.Next.No < newHeroNode.No {
		head = head.Next
	}
	if head.Next != nil && head.Next.No == newHeroNode.No {
		fmt.Println("对不起，已经存在No=", newHeroNode.No)
		return
	} else {
		newHeroNode.Next = head.Next
		head.Next = newHeroNode
	}
}

func DelHNode(head *HNode, id int) {
	for head.Next != nil && head.Next.No != id {
		head = head.Next
	}
	if head.Next != nil && head.Next.No == id {
		temp := head.Next
		head.Next = head.Next.Next
		temp.Next = nil
	} else {
		fmt.Println("sorry, 要删除的id不存在")
	}
}

// 显示链表的所有结点信息
func ListHNode(head *HNode) {
	if head.Next == nil {
		fmt.Println("空空如也。。。。")
		return
	}
	for head.Next != nil {
		fmt.Printf("[%d , %s , %s]==>", head.Next.No, head.Next.Name, head.Next.NickName)
		head = head.Next
	}
}

func main() {
	head := &HNode{}
	hero1 := &HNode{
		No:       1,
		Name:     "宋江",
		NickName: "及时雨",
	}
	hero2 := &HNode{
		No:       2,
		Name:     "卢俊义",
		NickName: "玉麒麟",
	}
	hero3 := &HNode{
		No:       6,
		Name:     "林冲",
		NickName: "豹子头",
	}
	hero4 := &HNode{
		No:       3,
		Name:     "吴用",
		NickName: "智多星",
	}

	InsertHNodeToTail(head, hero2)
	InsertHNodeOrder(head, hero4)
	InsertHNodeOrder(head, hero1)
	InsertHNodeToTail(head, hero3)

	ListHNode(head)

	fmt.Println()
	DelHNode(head, 1)
	DelHNode(head, 3)
	ListHNode(head)
}
