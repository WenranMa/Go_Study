# 876_链表中间节点_Middle_of_the_Linked_List

Given a non-empty, singly linked list with head node head, return a middle node of linked list. If there are two middle nodes, return the second middle node. 

Example:

    Input: [1,2,3,4,5]  
    Output: Node 3 from this list (Serialization: [3,4,5])  
    The returned node has value 3.  (The judge's serialization of this node is [3,4,5]).  
    Note that we returned a ListNode object ans, such that:  
    ans.val = 3, ans.next.val = 4, ans.next.next.val = 5, and ans.next.next.next = NULL.  
    
    Input: [1,2,3,4,5,6]  
    Output: Node 4 from this list (Serialization: [4,5,6])  
    Since the list has two middle nodes with values 3 and 4, we return the second one.

### 解：

先遍历得到链表的长度，然后遍历到链表的中间节点。

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    l := 1
    h := head
    for h.Next != nil {
        h = h.Next
        l++
    }
    m := l / 2
    for i := 0; i < m; i++ {
        head = head.Next
    }
    return head
}
```