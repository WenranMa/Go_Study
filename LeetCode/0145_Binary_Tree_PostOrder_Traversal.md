# 145. 二叉树的后序遍历

### 简单

给你一棵二叉树的根节点 root ，返回其节点值的 后序遍历 。

### 示例 1：

    输入：root = [1,null,2,3]
    输出：[3,2,1]

### 示例 2：

    输入：root = []
    输出：[]

### 示例 3：

    输入：root = [1]
    输出：[1]

### 提示：
- 树中节点的数目在范围 [0, 100] 内
- -100 <= Node.val <= 100

进阶：递归算法很简单，你可以通过迭代算法完成吗？

### 解：
Left -> Right -> Root

```go
func postorderTraversal(root *TreeNode) []int {
	res := []int{}
	var postOrder func(*TreeNode)
	postOrder = func(root *TreeNode) {
		if root == nil {
			return
		}
		postOrder(root.Left)
		postOrder(root.Right)
		res = append(res, root.Val)
	}
	postOrder(root)
	return res
}

// stack方式 相当于 Root -> Right -> Left 取反。
func postorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	stack := []*TreeNode{}
	for root != nil || len(stack) != 0 {
		if root != nil {
			res = append(res, root.Val)
			stack = append(stack, root)
			root = root.Right
		} else {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			root = root.Left
		}
	}
	for i, j := 0, len(res)-1; i <= j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}
```