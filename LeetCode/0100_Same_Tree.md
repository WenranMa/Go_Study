# 100. Same Tree

### 简单

Given the roots of two binary trees p and q, write a function to check if they are the same or not.

Two binary trees are considered the same if they are structurally identical, and the nodes have the same value.

### Example 1:

![bt1](/file/img/bt_ex1.jpg)

Input: p = [1,2,3], q = [1,2,3]

Output: true

### Example 2:

![bt2](/file/img/bt_ex2.jpg)

Input: p = [1,2], q = [1,null,2]

Output: false

### Example 3:

![bt3](/file/img/bt_ex3.jpg)

Input: p = [1,2,1], q = [1,1,2]

Output: false
 
### Constraints:

The number of nodes in both trees is in the range [0, 100].

-104 <= Node.val <= 104

### 解：

DFS

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil {
		return q == nil
	}
	if q == nil {
		return p == nil
	}
	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
```