# 144. 二叉树的前序遍历

### 简单

给你二叉树的根节点 root ，返回它节点值的 前序 遍历。

### 示例 1：

    输入：root = [1,null,2,3]
    输出：[1,2,3]

### 示例 2：

    输入：root = []
    输出：[]

### 示例 3：

    输入：root = [1]
    输出：[1]

### 示例 4：

    输入：root = [1,2]
    输出：[1,2]

### 示例 5：

    输入：root = [1,null,2]
    输出：[1,2]
 

### 提示：
- 树中节点数目在范围 [0, 100] 内
- -100 <= Node.val <= 100

进阶：递归算法很简单，你可以通过迭代算法完成吗？

### 解：
root - left - right的遍历顺序

```go
func preorderTraversal(root *TreeNode) []int {
	res := []int{}
	var preOrder func(*TreeNode)
	preOrder = func(root *TreeNode) {
		if root == nil {
			return
		}
		res = append(res, root.Val)
		preOrder(root.Left)
		preOrder(root.Right)
	}
	preOrder(root)
	return res
}

// Stack 方式
func preorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	stack := []*TreeNode{}
	for root != nil || len(stack) != 0 {
		if root != nil {
			res = append(res, root.Val)
			stack = append(stack, root)
			root = root.Left
		} else {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			root = root.Right
		}
	}
	return res
}
```

