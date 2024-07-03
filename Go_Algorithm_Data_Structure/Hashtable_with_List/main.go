package main

import (
	"fmt"
	"os"
)

// 定义emp
type Emp struct {
	Id   int
	Name string
	Next *Emp
}

// show 单个雇员..
func (this *Emp) ShowMe() {
	fmt.Printf("链表%d 找到该雇员 %d:%s\n", this.Id%7, this.Id, this.Name)
}

// 定义EmpLink
type EmpLink struct {
	Head *Emp
}

// insert 添加员工的方法, 保证添加时，编号从小到大
func (this *EmpLink) Insert(emp *Emp) {
	pre := this.Head
	cur := pre.Next
	for cur != nil {
		if cur.Id > emp.Id {
			break
		}
		pre = cur
		cur = cur.Next
	}
	pre.Next = emp
	emp.Next = cur
}

// 显示链表的信息
func (this *EmpLink) ShowLink(no int) {
	cur := this.Head.Next
	fmt.Printf("链表%d:", no)
	for cur != nil {
		fmt.Printf("雇员id=%d 名字=%s ->", cur.Id, cur.Name)
		cur = cur.Next
	}
	fmt.Println()
}

// 根据id查找对应的雇员，如果没有就返回nil
func (this *EmpLink) FindById(id int) *Emp {
	cur := this.Head.Next
	for cur != nil {
		if cur.Id == id {
			return cur
		}
		cur = cur.Next
	}
	return nil
}

// 定义hashtable ,含有一个链表数组
type HashTable struct {
	Cap     int
	LinkArr []*EmpLink
}

// 给HashTable 编写Insert 雇员的方法.
func (this *HashTable) Insert(emp *Emp) {
	linkNo := this.HashFun(emp.Id)
	this.LinkArr[linkNo].Insert(emp) //
}

// 编写方法，显示hashtable的所有雇员
func (this *HashTable) ShowAll() {
	for i := 0; i < len(this.LinkArr); i++ {
		this.LinkArr[i].ShowLink(i)
	}
}

// 编写一个散列方法
func (this *HashTable) HashFun(id int) int {
	return id % this.Cap
}

func (this *HashTable) FindById(id int) *Emp {
	linkNo := this.HashFun(id)
	return this.LinkArr[linkNo].FindById(id)
}

func NewTable(cap int) *HashTable {
	t := &HashTable{
		Cap:     cap,
		LinkArr: make([]*EmpLink, cap),
	}
	for i := 0; i < cap; i++ {
		t.LinkArr[i] = &EmpLink{
			Head: &Emp{
				Id: -1,
			},
		}
	}
	return t
}

func main() {
	key := ""
	id := 0
	name := ""
	hashtable := NewTable(7)
	for {
		fmt.Println("===============雇员系统菜单============")
		fmt.Println("input 表示添加雇员")
		fmt.Println("show  表示显示雇员")
		fmt.Println("find  表示查找雇员")
		fmt.Println("exit  表示退出系统")
		fmt.Println("请输入你的选择")
		fmt.Scanln(&key)
		switch key {
		case "input":
			fmt.Println("输入雇员id")
			fmt.Scanln(&id)
			fmt.Println("输入雇员name")
			fmt.Scanln(&name)
			emp := &Emp{
				Id:   id,
				Name: name,
			}
			hashtable.Insert(emp)
		case "show":
			hashtable.ShowAll()
		case "find":
			fmt.Println("请输入id号:")
			fmt.Scanln(&id)
			emp := hashtable.FindById(id)
			if emp == nil {
				fmt.Printf("id=%d 的雇员不存在\n", id)
			} else {
				emp.ShowMe()
			}

		case "exit":
			os.Exit(0)
		default:
			fmt.Println("输入错误")
		}
	}

}
