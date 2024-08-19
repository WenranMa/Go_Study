# 965_单值二叉树_Univalued_Binary_Tree
A binary tree is univalued if every node in the tree has the same value. Return true if and only if the given tree is univalued.

Example:

    Input: [1,1,1,1,1,null,1]
    Output: true
    Input: [2,2,2,5,2]
    Output: false

Note:
- The number of nodes in the given tree will be in the range [1, 100].
- Each node's value will be an integer in the range [0, 99].

### 解：

递归，DFS。

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isUnivalTree(root *TreeNode) bool {
    if root == nil {
        return true
    }
    l, r := true, true
    if root.Left != nil {
        l = root.Left.Val == root.Val && isUnivalTree(root.Left)
    }
    if root.Right != nil {
        r = root.Right.Val == root.Val && isUnivalTree(root.Right)
    }
    return l && r
}
```