# 86. 分隔链表

### 中等

给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。

你应当 保留 两个分区中每个节点的初始相对位置。

### 示例 1：
![partition](/file/img/partition.jpg)

    输入：head = [1,4,3,2,5,2], x = 3
    输出：[1,2,2,4,3,5]

### 示例 2：

    输入：head = [2,1], x = 2
    输出：[1,2]
 
### 提示：
- 链表中节点的数目在范围 [0, 200] 内
- -100 <= Node.val <= 100
- -200 <= x <= 200

### 解：

用两个虚拟头部，一个存小于x的链表，一个存大于等于x的链表，然后小的尾部接大的头部。

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	s, l := &ListNode{}, &ListNode{}
	s1, l1 := s, l
	for cur := head; cur != nil; cur = cur.Next {
		if cur.Val < x {
			s1.Next = cur
			s1 = s1.Next
		} else if cur.Val >= x {
			l1.Next = cur
			l1 = l1.Next
		}
	}
	l1.Next = nil
	s1.Next = l.Next
	return s.Next
}
```