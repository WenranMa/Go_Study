# 103. 二叉树的锯齿形层序遍历

### 中等

给你二叉树的根节点 root ，返回其节点值的 锯齿形层序遍历 。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。


### 示例 1：

    输入：root = [3,9,20,null,null,15,7]
    输出：[[3],[20,9],[15,7]]

### 示例 2：

    输入：root = [1]
    输出：[[1]]

### 示例 3：

    输入：root = []
    输出：[]
 
### 提示：
- 树中节点数目在范围 [0, 2000] 内
- -100 <= Node.val <= 100

### 解：

加入一个boolean变量判断是奇数行或者偶数行，每一层翻转一次。通过if语句判断。
Recursion

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	var order func(root *TreeNode, level int, flag bool)
	order = func(root *TreeNode, level int, flag bool) {
		if root == nil {
			return
		}
		if level > len(res) {
			res = append(res, []int{})
		}
		if flag {
			res[level-1] = append(res[level-1], root.Val)
		} else {
			res[level-1] = append([]int{root.Val}, res[level-1]...)
		}
		order(root.Left, level+1, !flag)
		order(root.Right, level+1, !flag)
	}
	order(root, 1, true)
	return res
}
```

方法2：BFS，用队列
```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	q := []*TreeNode{}
	q = append(q, root)
	q = append(q, nil)
	flag := true
	for len(q) != 0 {
		level := []int{}
		cur := q[0]
		q = q[1:len(q)]

		for cur != nil {
			if flag {
				level = append(level, cur.Val)
			} else {
				level = append([]int{cur.Val}, level...)
			}
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
			cur = q[0]
			q = q[1:len(q)]
		}
		res = append(res, level)
		if len(q) != 0 {
			q = append(q, nil)
		}
		flag = !flag
	}
	return res
}
```