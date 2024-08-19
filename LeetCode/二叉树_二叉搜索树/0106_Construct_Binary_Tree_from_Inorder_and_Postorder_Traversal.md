# 106. 从中序与后序遍历序列构造二叉树

### 中等

给定两个整数数组 inorder 和 postorder ，其中 inorder 是二叉树的中序遍历， postorder 是同一棵树的后序遍历，请你构造并返回这颗 二叉树 。

### 示例 1:

	输入：inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]
	输出：[3,9,20,null,null,15,7]

### 示例 2:

	输入：inorder = [-1], postorder = [-1]
	输出：[-1]

### 提示:
- 1 <= inorder.length <= 3000
- postorder.length == inorder.length
- -3000 <= inorder[i], postorder[i] <= 3000
- inorder 和 postorder 都由 不同 的值组成
- postorder 中每一个值都在 inorder 中
- inorder 保证是树的中序遍历
- postorder 保证是树的后序遍历

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
func buildTree(inorder []int, postorder []int) *TreeNode {
	m := make(map[int]int)
	for i, n := range inorder {
		m[n] = i
	}
	return build(postorder, m, len(postorder)-1, 0, len(postorder)-1)
}

func build(postorder []int, m map[int]int, index, start, end int) *TreeNode {
	root := &TreeNode{
		Val: postorder[index],
	}
	mid := m[postorder[index]]
	if mid > start {
		root.Left = build(postorder, m, index-1-(end-mid), start, mid-1)
	}
	if mid < end {
		root.Right = build(postorder, m, index-1, mid+1, end)
	}
	return root
}
```