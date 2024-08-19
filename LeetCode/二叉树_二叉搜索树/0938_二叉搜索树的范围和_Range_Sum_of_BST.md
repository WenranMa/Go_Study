# 938_二叉搜索树的范围和_Range_Sum_of_BST

Given the root node of a binary search tree, return the sum of values of all nodes with value between L and R (inclusive). The binary search tree is guaranteed to have unique values.

Example 1:

Input: root = [10,5,15,3,7,null,18], L = 7, R = 15
Output: 32

Example 2:

Input: root = [10,5,15,3,7,13,18,1,null,6], L = 6, R = 10
Output: 23
 
Note:

The number of nodes in the tree is at most 10000.
The final answer is guaranteed to be less than 2^31.

### 解：

中序遍历，判断在范围内的加。
```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rangeSumBST(root *TreeNode, low int, high int) int {
    var sum int
    var inOrder func(*TreeNode)
    inOrder = func(root *TreeNode) {
        if root == nil {
            return
        }
        inOrder(root.Left)
        if root.Val >= low && root.Val <= high {
            sum += root.Val
        }
        inOrder(root.Right)
    }
    inOrder(root)
    return sum
}
```

方法2：DFS

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rangeSumBST(root *TreeNode, low int, high int) int {
    var sum int
    var dfs func(*TreeNode)
    dfs = func(root *TreeNode) {
        if root == nil {
            return
        }
        if root.Val < low {
            dfs(root.Right)
        } else if root.Val > high {
            dfs(root.Left)
        } else {
            sum += root.Val
            dfs(root.Left)
            dfs(root.Right)
        }
    }
    dfs(root)
    return sum
}
```



```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rangeSumBST(root *TreeNode, L int, R int) int {
    return helper(root, L, R, 0)
}

func helper(root *TreeNode, L, R, n int) int {
    if root == nil {
        return n
    }
    if root.Val < L {
        n = helper(root.Right, L, R, n)
    } else if root.Val > R {
        n = helper(root.Left, L, R, n)
    } else {
        n += root.Val
        n = helper(root.Left, L, R, n)
        n = helper(root.Right, L, R, n)
    }
    return n
}
```