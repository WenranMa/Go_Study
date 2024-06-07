# 94. 二叉树的中序遍历

### 简单
给定一个二叉树的根节点 root ，返回 它的 中序 遍历 。

### 示例 1：

    输入：root = [1,null,2,3]
    输出：[1,3,2]

### 示例 2：

    输入：root = []
    输出：[]

### 示例 3：

    输入：root = [1]
    输出：[1]

### 提示：
- 树中节点数目在范围 [0, 100] 内
- -100 <= Node.val <= 100

进阶: 递归算法很简单，你可以通过迭代算法完成吗？

### 解：
Left -- Root -- Right

```go
func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	var inOrder func(*TreeNode)
	inOrder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inOrder(root.Left)
		res = append(res, root.Val)
		inOrder(root.Right)
	}
	inOrder(root)
	return res
}

// stack 方式
func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	stack := []*TreeNode{}
	for root != nil || len(stack) != 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
            res = append(res, root.Val)
			root = root.Right
		}
	}
	return res
}
```