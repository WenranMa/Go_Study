# 101. 对称二叉树

### 简单

给你一个二叉树的根节点 root ， 检查它是否轴对称。

### 示例 1：
![symmtric1](/file/img/symtree1.jpg)

输入：root = [1,2,2,3,4,4,3]

输出：true

### 示例 2：
![symmtric2](/file/img/symtree2.jpg)

输入：root = [1,2,2,null,3,null,3]

输出：false
 

### 提示：
树中节点数目在范围 [1, 1000] 内

-100 <= Node.val <= 100
 
### 进阶：
你可以运用递归和迭代两种方法解决这个问题吗？

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
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return helper(root.Left, root.Right)
}

func helper(l, r *TreeNode) bool {
	if l != nil && r != nil {
		return l.Val == r.Val && helper(l.Right, r.Left) && helper(l.Left, r.Right)
	}
	return l == nil && r == nil
}

// 迭代方式，用Queue. 先进先出。

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	lq := []*TreeNode{}
	rq := []*TreeNode{}
	lq = append(lq, root.Left)
	rq = append(rq, root.Right)
	for len(lq) != 0 && len(rq) != 0 {
		l := lq[0]
		lq = lq[1:len(lq)]
		r := rq[0]
		rq = rq[1:len(rq)]
		if l == nil && r == nil {
			continue
		}
		if l == nil && r != nil || l != nil && r == nil {
			return false
		}
		if l.Val != r.Val {
			return false
		}
		lq = append(lq, l.Left)
		lq = append(lq, l.Right)
		rq = append(rq, r.Right)
		rq = append(rq, r.Left)
	}
	return true
}
```