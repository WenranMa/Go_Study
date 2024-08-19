# 872_叶子相似的树_Leaf-Similar_Trees
Consider all the leaves of a binary tree.  From left to right order, the values of those leaves form a leaf value sequence. For example, in the given tree above, the leaf value sequence is (6, 7, 4, 9, 8). Two binary trees are considered leaf-similar if their leaf value sequence is the same. Return true if and only if the two given trees with head nodes root1 and root2 are leaf-similar.

Note: Both of the given trees will have between 1 and 100 nodes.

### 解：

两个数的叶子节点序列相同即可。

递归遍历，叶子节点直接加入数组。

前序遍历即可

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
    v1 := dfs(root1, []int{})
    v2 := dfs(root2, []int{})
    l1 := len(v1)
    l2 := len(v2)
    if l1 != l2 {
        return false
    }
    for i := 0; i < l1; i++ {
        if v1[i] != v2[i] {
            return false
        }
    }
    return true
}

func dfs(node *TreeNode, values []int) []int {
    if node == nil {
        return values
    }
    if node.Left == nil && node.Right == nil {
        values = append(values, node.Val)
    }
    values = dfs(node.Left, values)
    values = dfs(node.Right, values)
    return values
}
```

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	var val []int
	var preorder func(*TreeNode)
	preorder = func(r *TreeNode) {
		if r == nil {
			return
		}
		if r.Left == nil && r.Right == nil {
			val = append(val, r.Val)
		}
		preorder(r.Left)
		preorder(r.Right)
	}
	preorder(root1)
	v1 := make([]int, len(val))
	copy(v1, val)
	val = make([]int, 0)
	preorder(root2)
	if len(v1) != len(val) {
		return false
	}
	for i := 0; i < len(v1); i++ {
		if v1[i] != val[i] {
			return false
		}
	}
	return true
}
```