# 148. 排序链表

### 中等

给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。

### 示例 1：

    输入：head = [4,2,1,3]
    输出：[1,2,3,4]

### 示例 2：

    输入：head = [-1,5,3,4,0]
    输出：[-1,0,3,4,5]

### 示例 3：

    输入：head = []
    输出：[]

### 提示：
链表中节点的数目在范围 [0, 5 * 10^4] 内
-105 <= Node.val <= 10^5

进阶：你可以在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序吗？

### 解:
merge sort

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	fast = slow.Next
	slow.Next = nil
	slow = sortList(head)
	fast = sortList(fast)
	return merge(slow, fast)
}

func merge(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var head *ListNode
	if l1.Val <= l2.Val {
		head = l1
		l1 = l1.Next
	} else {
		head = l2
		l2 = l2.Next
	}
	cur := head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}
	return head
}
```