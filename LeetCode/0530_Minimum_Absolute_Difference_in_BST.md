# 530. 二叉搜索树的最小绝对差

### 简单

给你一个二叉搜索树的根节点 root ，返回 树中任意两不同节点值之间的最小差值 。

差值是一个正数，其数值等于两值之差的绝对值。

### 示例 1：
![bst1](/file/img/bst1.jpg)

输入：root = [4,2,6,1,3]

输出：1

### 示例 2：
![bst2](/file/img/bst2.jpg)

输入：root = [1,0,48,null,null,12,49]

输出：1
 
### 提示：

树中节点的数目范围是 [2, 10^4]

0 <= Node.val <= 10^5

### 解：

因为是BST, 所以中序遍历（left, root, right）后得到一定是个升序数组。
最小值一定是`相邻`的某两个元素的差。

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var vals []int

func getMinimumDifference(root *TreeNode) int {
	vals = []int{}
	inorder(root)
	res := 100001
	for i := len(vals) - 1; i > 0; i-- {
		res = min(res, vals[i]-vals[i-1])
	}
	return res
}

func inorder(root *TreeNode) {
	if root == nil {
		return
	}
	inorder(root.Left)
	vals = append(vals, root.Val)
	inorder(root.Right)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
```

优化：不需要用数组存，一个变量即可~

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func getMinimumDifference(root *TreeNode) int {
	res, pre:= math.MaxInt64, -1
    var dfs func(node *TreeNode)
    dfs = func(node *TreeNode) {
        if node == nil {
            return
        }
        dfs(node.Left)
        if pre != -1 && node.Val - pre < res{
            res = node.Val - pre
        }
        pre = node.Val
        dfs(node.Right)
    }
    dfs(root)
    return res
}
```