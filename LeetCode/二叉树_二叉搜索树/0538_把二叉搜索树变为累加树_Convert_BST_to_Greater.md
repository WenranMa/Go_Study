# 538_把二叉搜索树变为累加树_Convert_BST_to_Greater_Tree
Given a Binary Search Tree (BST), convert it to a Greater Tree such that every key of the original BST is changed to the original key plus sum of all keys greater than the original key in BST.

Example:

    Input: The root of a Binary Search Tree like this:  
             5  
            /  \  
           2    13  

    Output: The root of a Greater Tree like this:  
             18  
            /  \  
          20    13

这题于leetcode 1038重复。

### 解：

就是反向中序遍历。从右边开始遍历，然后累加。

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func convertBST(root *TreeNode) *TreeNode {
	var value int
	var inOrderReverse func(*TreeNode)
	inOrderReverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		inOrderReverse(root.Right)
		value += root.Val
		root.Val = value
		inOrderReverse(root.Left)
	}
	inOrderReverse(root)
	return root
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
func convertBST(root *TreeNode) *TreeNode {
    inOrderReverse(root, 0)
    return root
}
func inOrderReverse(root *TreeNode, n int) int {
    if root == nil {
        return n
    }
    n = inOrderReverse(root.Right, n)
    n += root.Val
    root.Val = n
    n = inOrderReverse(root.Left, n)
    return n
}
```