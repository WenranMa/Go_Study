# 669_修剪二叉搜索树_Trim_a_Binary_Search_Tree

Given a binary search tree and the lowest and highest boundaries as L and R, trim the tree so that all its elements lies in [L, R] (R >= L). You might need to change the root of the tree, so the result should return the new root of the trimmed binary search tree.

Example:  

    Input:   
        1
       / \
      0   2

    L = 1
    R = 2

    Output:
        1
         \
          2

    Input:
        3
       / \
      0   4
       \
        2
       /
      1

    L = 1
    R = 3

    Output:
        3
       /
      2
     /
    1

所有节点的值要在[L, R]之间。

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
func trimBST(root *TreeNode, L int, R int) *TreeNode {
    if root == nil {
        return root
    }
    if root.Val < L {
        return trimBST(root.Right, L, R)
    } else if root.Val > R {
        return trimBST(root.Left, L, R)
    }
    root.Left = trimBST(root.Left, L, R)
    root.Right = trimBST(root.Right, L, R)
    return root
}
```