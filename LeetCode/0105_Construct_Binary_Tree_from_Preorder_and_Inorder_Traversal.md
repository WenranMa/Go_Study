# 105. 从前序与中序遍历序列构造二叉树

### 中等

给定两个整数数组 preorder 和 inorder ，其中 preorder 是二叉树的先序遍历， inorder 是同一棵树的中序遍历，请构造二叉树并返回其根节点。

### 示例 1:
![pre_in_tree](/file/img/pre_in_tree.jpg)

输入: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]

输出: [3,9,20,null,null,15,7]

### 示例 2:

输入: preorder = [-1], inorder = [-1]

输出: [-1]

### 实例 3：
preorder = [1,2,4,5,3,6,7]

inorder = [4,2,5,1,6,3,7]

输出: [1,2,3,4,5,6,7]


### 提示:

- 1 <= preorder.length <= 3000
- inorder.length == preorder.length
- -3000 <= preorder[i], inorder[i] <= 3000
- preorder 和 inorder 均 无重复 元素
- inorder 均出现在 preorder
- preorder 保证 为二叉树的前序遍历序列
- inorder 保证 为二叉树的中序遍历序列

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
func buildTree(preorder []int, inorder []int) *TreeNode {
	m := make(map[int]int)
	for i, n := range inorder {
		m[n] = i
	}
	return build(preorder, m, 0, 0, len(preorder)-1)
}

func build(preorder []int, m map[int]int, index, start, end int) *TreeNode {
	root := &TreeNode{
		Val: preorder[index],
	}
	mid := m[preorder[index]]
	if mid > start {
		root.Left = build(preorder, m, index+1, start, mid-1)
	}
	if mid < end {
		root.Right = build(preorder, m, index+1+(mid-start), mid+1, end)
	}
	return root
}
```