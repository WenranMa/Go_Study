
## 两个链表相加，高位在前。

```go
// 先反转，再相加。
type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	h1 := &ListNode{8, &ListNode{7, &ListNode{9, &ListNode{6, nil}}}}
	h5 := &ListNode{7, &ListNode{8, &ListNode{7, &ListNode{5, nil}}}}
	h:= sum(h1, h5)
	for h != nil {
		fmt.Println(h.Val)
		h = h.Next
	}
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

func sum(h1, h2 *ListNode) *ListNode {
	h1 = reverseList(h1)
	h2 = reverseList(h2)
	dummy:= &ListNode{-1, nil}
	cur := dummy
	carry := 0
	for h1 != nil || h2 != nil {	
		var n1, n2 int
		if h1 != nil {
			n1 = h1.Val
		}
		if h2 != nil {
			n2 = h2.Val
		}
		sum := (n1 + n2 + carry)%10
		carry = (n1 + n2 + carry)/10
		cur.Next = &ListNode{sum, nil}
		if h1 != nil {
			h1 = h1.Next
		}
		if h2 != nil {
			h2 = h2.Next
		}
		cur = cur.Next
	}
	if carry > 0 {
		cur.Next = &ListNode{carry, nil}
	}
	return	reverseList(dummy.Next)
}
```


// 164, 200 再看




// 130 , 140, 149, 154, 174, 188, 204, 212,216, 218, 220, 221. 240