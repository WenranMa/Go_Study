# 445. 两数相加 II

### 中等

给你两个 非空 链表来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储一位数字。将这两数相加会返回一个新的链表。

你可以假设除了数字 0 之外，这两个数字都不会以零开头。

### 示例1：

输入：l1 = [7,2,4,3], l2 = [5,6,4]
输出：[7,8,0,7]

### 示例2：

输入：l1 = [2,4,3], l2 = [5,6,4]
输出：[8,0,7]

### 示例3：

输入：l1 = [0], l2 = [0]
输出：[0]

### 提示：

链表的长度范围为 [1, 100]
0 <= node.val <= 9
输入数据保证链表代表的数字无前导 0

进阶：如果输入链表不能翻转该如何解决？

### 解：

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(h1 *ListNode, h2 *ListNode) *ListNode {
	h1 = reverseList(h1)
	h2 = reverseList(h2)
	dummy := &ListNode{-1, nil}
	cur := dummy
	carry := 0
	for h1 != nil || h2 != nil {
		var n1, n2 int
		if h1 != nil {
			n1 = h1.Val
			h1 = h1.Next
		}
		if h2 != nil {
			n2 = h2.Val
			h2 = h2.Next
		}
		sum := (n1 + n2 + carry) % 10
		carry = (n1 + n2 + carry) / 10
		cur.Next = &ListNode{sum, nil}
		cur = cur.Next
	}
	if carry > 0 {
		cur.Next = &ListNode{carry, nil}
	}
	return reverseList(dummy.Next)
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{-1, head}
	for head.Next != nil {
		nx := head.Next
		head.Next = nx.Next
		nx.Next = dummy.Next
		dummy.Next = nx
	}
	return dummy.Next
}
```

如果不让反转链表，可以用栈

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(h1 *ListNode, h2 *ListNode) *ListNode {
	var s1, s2, s3 []int
	for h1 != nil {
		s1 = append(s1, h1.Val)
		h1 = h1.Next
	}
	for h2 != nil {
		s2 = append(s2, h2.Val)
		h2 = h2.Next
	}
	carry := 0
	for len(s1) != 0 || len(s2) != 0 {
		var n1, n2 int
		if len(s1) != 0 {
			n1 = s1[len(s1)-1]
			s1 = s1[:len(s1)-1]
		}
		if len(s2) != 0 {
			n2 = s2[len(s2)-1]
			s2 = s2[:len(s2)-1]
		}
		sum := (n1 + n2 + carry) % 10
		carry = (n1 + n2 + carry) / 10
		s3 = append(s3, sum)
	}
	if carry > 0 {
		s3 = append(s3, carry)
	}
	dummy := &ListNode{-1, nil}
	var temp = dummy
	for len(s3) != 0 {
		n := s3[len(s3)-1]
		s3 = s3[:len(s3)-1]
		temp.Next = &ListNode{n, nil}
		temp = temp.Next
	}
	return dummy.Next
}
```