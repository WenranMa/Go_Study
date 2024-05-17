# 226. Invert Binary Tree

### 简单

Given the root of a binary tree, invert the tree, and return its root.

### Example 1:
![invert-tree1](/file/img/invert1-tree.jpg)

Input: root = [4,2,7,1,3,6,9]

Output: [4,7,2,9,6,3,1]

### Example 2:
![invert-tree2](/file/img/invert2-tree.jpg)

Input: root = [2,1,3]

Output: [2,3,1]

### Example 3:

Input: root = []

Output: []
 
### Constraints:

The number of nodes in the tree is in the range [0, 100].

-100 <= Node.val <= 100

### 解：

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
    if root == nil {
        return root
    }
    root.Left, root.Right = root.Right, root.Left
    invertTree(root.Left)
    invertTree(root.Right)
    return root
}
```