# 124. 二叉树中的最大路径和

### 困难

二叉树中的 路径 被定义为一条节点序列，序列中每对相邻节点之间都存在一条边。同一个节点在一条路径序列中 至多出现一次 。该路径 至少包含一个 节点，且不一定经过根节点。

路径和 是路径中各节点值的总和。

给你一个二叉树的根节点 root ，返回其 最大路径和 。

### 示例 1：
![path1](/file/img/bt_path_sum_1.jpg)

输入：root = [1,2,3]
输出：6
解释：最优路径是 2 -> 1 -> 3 ，路径和为 2 + 1 + 3 = 6

### 示例 2：
![path2](/file/img/bt_path_sum_2.jpg)

输入：root = [-10,9,20,null,null,15,7]
输出：42
解释：最优路径是 15 -> 20 -> 7 ，路径和为 15 + 20 + 7 = 42

### 提示：

树中节点数目范围是 [1, 3 * 104]
-1000 <= Node.val <= 1000

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
func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := root.Val
	var path func(node *TreeNode) int
	path = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l := path(node.Left)
		r := path(node.Right)
		onePath := node.Val
		if l > 0 {
			onePath += l
		}
		if r > 0 {
			onePath += r
		}
		if onePath > res {
			res = onePath
		}
		return max(node.Val, node.Val+max(l, r))
	}
	path(root)
	return res
}
```