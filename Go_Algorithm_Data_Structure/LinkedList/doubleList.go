package main

import (
	"fmt"
)

type HeroNode struct {
	No       int
	Name     string
	NickName string
	Pre      *HeroNode
	Next     *HeroNode
}

func InsertToTail(head *HeroNode, newHeroNode *HeroNode) {
	for head.Next != nil { // head is tail
		head = head.Next
	}
	head.Next = newHeroNode
	newHeroNode.Pre = head
}

func InsertOrder(head *HeroNode, newHeroNode *HeroNode) {
	cur := head
	for cur.Next != nil && cur.Next.No < newHeroNode.No {
		cur = cur.Next
	}
	if cur.Next != nil && cur.Next.No == newHeroNode.No {
		fmt.Println("对不起，已经存在No=", newHeroNode.No)
		return
	} else {
		newHeroNode.Next = cur.Next
		newHeroNode.Pre = cur
		if cur.Next != nil {
			cur.Next.Pre = newHeroNode
		}
		cur.Next = newHeroNode
	}
}

func DelHeroNode(head *HeroNode, id int) {
	cur := head
	for cur.Next != nil && cur.Next.No != id {
		cur = cur.Next
	}
	if cur.Next != nil && cur.Next.No == id {
		cur.Next = cur.Next.Next //ok
		if cur.Next != nil {
			cur.Next.Pre = cur
		}
	} else {
		fmt.Println("sorry, 要删除的id不存在")
	}
}

func ListForward(head *HeroNode) {
	if head.Next == nil {
		fmt.Println("空空如也。。。。")
		return
	}
	for head.Next != nil {
		fmt.Printf("[%d , %s , %s]==>", head.Next.No,
			head.Next.Name, head.Next.NickName)
		head = head.Next
	}
	fmt.Println()
}

func ListBackward(head *HeroNode) {
	if head.Next == nil {
		fmt.Println("空空如也。。。。")
		return
	}
	for head.Next != nil { // head is tail for now
		head = head.Next
	}
	for head.Pre != nil {
		fmt.Printf("[%d , %s , %s]==>", head.No,
			head.Name, head.NickName)
		head = head.Pre
	}
	fmt.Println()
}

func main() {
	head := &HeroNode{}
	hero1 := &HeroNode{
		No:       1,
		Name:     "宋江",
		NickName: "及时雨",
	}
	hero2 := &HeroNode{
		No:       2,
		Name:     "卢俊义",
		NickName: "玉麒麟",
	}
	hero3 := &HeroNode{
		No:       3,
		Name:     "吴用",
		NickName: "智多星",
	}
	hero4 := &HeroNode{
		No:       6,
		Name:     "林冲",
		NickName: "豹子头",
	}
	InsertOrder(head, hero1)
	InsertOrder(head, hero3)
	InsertOrder(head, hero2)
	InsertToTail(head, hero4)
	ListForward(head)
	ListBackward(head)
	DelHeroNode(head, 2)
	DelHeroNode(head, 5)
	ListForward(head)
}
