# 19. 删除链表的倒数第 N 个结点

### 中等

给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。

### 示例 1：
![remove](/file/img/remove_ex1.jpg)

    输入：head = [1,2,3,4,5], n = 2
    输出：[1,2,3,5]

### 示例 2：

    输入：head = [1], n = 1
    输出：[]

### 示例 3：

    输入：head = [1,2], n = 1
    输出：[1]

### 提示：
- 链表中结点的数目为 sz
- 1 <= sz <= 30
- 0 <= Node.val <= 100
- 1 <= n <= sz

### 进阶：
你能尝试使用一趟扫描实现吗？
    - 可以用数组（stack）存下每个node, 这样长度就知道了
    - 两个指针

### 解：
计算链表长度

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	l := 1
	cur := head
	for cur.Next != nil {
		l += 1
		cur = cur.Next
	}

	cur = head
	step := l - n - 1
	if step < 0 { //针对倒数第一的情况。
		return head.Next
	}

	for step > 0 {
		cur = cur.Next
		step -= 1
	}
	cur.Next = cur.Next.Next
	return head
}

// 优化，双指针，遍历一遍，一个指针先走N，再两个指针同时走。
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	p1, p2 := dummy, dummy
	pos := 0
	for pos < n {
		p1 = p1.Next
		pos++
	}
	for p1.Next != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	p2.Next = p2.Next.Next
	return dummy.Next
}
```