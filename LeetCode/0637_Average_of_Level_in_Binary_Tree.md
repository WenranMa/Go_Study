# 637. 二叉树的层平均值

### 简单

给定一个非空二叉树的根节点 root , 以数组的形式返回每一层节点的平均值。与实际答案相差 10^-5 以内的答案可以被接受。

### 示例 1：
![average1](/file/img/avg1-tree.jpg)

输入：root = [3,9,20,null,null,15,7]

输出：[3.00000,14.50000,11.00000]

解释：第 0 层的平均值为 3,第 1 层的平均值为 14.5,第 2 层的平均值为 11 。因此返回 [3, 14.5, 11] 。

### 示例 2:
![average2](/file/img/avg2-tree.jpg)

输入：root = [3,9,20,15,7]

输出：[3.00000,14.50000,11.00000]
 
### 提示：

树中节点数量在 [1, 10^4] 范围内

-2^31 <= Node.val <= 2^31 - 1

### 解：

结合 102 题，Binary Tree Level Order Traversal.

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
	res := []float64{}
	if root == nil {
		return res
	}
	q := []*TreeNode{}
	q = append(q, root)
	q = append(q, nil)
	for len(q) != 0 {
		level := 0
		n := 0
		cur := q[0]
		q = q[1:len(q)]
		for cur != nil {
			level += cur.Val
			n += 1
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
			cur = q[0]
			q = q[1:len(q)]
		}
		res = append(res, float64(level)/float64(n))
		if len(q) != 0 {
			q = append(q, nil)
		}
	}
	return res
}
```