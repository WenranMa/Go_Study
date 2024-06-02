# 199. 二叉树的右视图

### 中等

给定一个二叉树的 根节点 root，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。

### 示例 1:

	输入: [1,2,3,null,5,null,4]
	输出: [1,3,4]

### 示例 2:

	输入: [1,null,3]
	输出: [1,3]

### 示例 3:

	输入: []
	输出: []
 
### 提示:
- 二叉树的节点个数的范围是 [0,100]
- -100 <= Node.val <= 100 

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
func rightSideView(root *TreeNode) []int {
	res := [][]int{}
	var order func(root *TreeNode, level int)
	order = func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		if level > len(res) {
			res = append(res, []int{})
		}
		res[level-1] = append(res[level-1], root.Val)
		order(root.Left, level+1)
		order(root.Right, level+1)
	}
	order(root, 1)
	r := []int{}
	for _, l := range res {
		r = append(r, l[len(l)-1])
	}
	return r
}

// 空间优化，不用每行都存，只要存每行最后一个。
func rightSideView(root *TreeNode) []int {
	rMap := make(map[int]int)
	var order func(root *TreeNode, level int)
	order = func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		rMap[level] = root.Val
		order(root.Left, level+1)
		order(root.Right, level+1)
	}
	order(root, 1)
	res := make([]int, len(rMap))
	for k, v := range rMap {
		res[k-1] = v
	}
	return res
}
```