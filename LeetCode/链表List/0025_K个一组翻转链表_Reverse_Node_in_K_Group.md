# 25. K个一组翻转链表

### 困难

给你链表的头节点 head ，每 k 个节点一组进行翻转，请你返回修改后的链表。

k 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。

你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。

### 示例 1：

    输入：head = [1,2,3,4,5], k = 2
    输出：[2,1,4,3,5]

### 示例 2：

    输入：head = [1,2,3,4,5], k = 3
    输出：[3,2,1,4,5]

### 解：
```go
func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{-1, head}
	pre, cur := dummy, dummy.Next
	position := 1
	for cur != nil {
		if position == k {
			position = 0
			pre = reverse(pre, cur)
			cur = pre.Next
		} else {
			cur = cur.Next
		}
		position++
	}
	return dummy.Next
}

func reverse(pre, end *ListNode) *ListNode {
	cur := pre.Next
	for pre.Next != end {
		temp := pre.Next
		Nx := cur.Next
		pre.Next = cur.Next
		cur.Next = Nx.Next
		Nx.Next = temp
	}
	return cur
}
```