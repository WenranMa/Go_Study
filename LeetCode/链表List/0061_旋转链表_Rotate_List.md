# 61. 旋转链表

### 中等

给你一个链表的头节点 head ，旋转链表，将链表每个节点向右移动 k 个位置。

### 示例 1：
![r1](/file/img/rotate1.jpg)

    输入：head = [1,2,3,4,5], k = 2
    输出：[4,5,1,2,3]

### 示例 2：
![r2](/file/img/rotate2.jpg)

    输入：head = [0,1,2], k = 4
    输出：[2,0,1]

### 提示：
-链表中节点的数目在范围 [0, 500] 内
--100 <= Node.val <= 100
- 0 <= k <= 2 * 109

### 解：

首尾相连。

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	l := 1
	cur := head
	for cur.Next != nil {
		cur = cur.Next
		l += 1
	}
	cur.Next = head
	step := l - k%l
	for step > 0 {
		cur = cur.Next
		step -= 1
	}
	res := cur.Next
	cur.Next = nil
	return res
}
```