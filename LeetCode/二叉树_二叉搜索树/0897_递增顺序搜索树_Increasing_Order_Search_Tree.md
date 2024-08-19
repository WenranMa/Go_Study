# 897_递增顺序搜索树_Increasing_Order_Search_Tree

Given a tree, rearrange the tree in in-order so that the leftmost node in the tree is now the root of the tree, and every node has no left child and only 1 right child.

Example 1:

    Input: [5,3,6,2,4,null,8,1,null,null,null,7,9]  
          5   
         / \  
        3   6  
       / \   \  
      2   4   8  
     /       / \   
    1       7   9  
    
    Output: [1,null,2,null,3,null,4,null,5,null,6,null,7,null,8,null,9]   
    1  
     \  
      2  
       \  
        3  
         \  
          4  
           \  
            5  
             \  
              6  
               \  
                7   
                 \  
                  8  
                   \  
                    9   

Note:
- The number of nodes in the given tree will be between 1 and 100.
- Each node will have a unique integer value from 0 to 1000.

### 解：

中序遍历

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func increasingBST(root *TreeNode) *TreeNode {
	var vals []int
	var inorder func(*TreeNode)
	inorder = func(n *TreeNode) {
		if n == nil {
			return
		}
		inorder(n.Left)
		vals = append(vals, n.Val)
		inorder(n.Right)
	}
	inorder(root)
	head := &TreeNode{vals[0], nil, nil}
	node := head
	for i := 1; i < len(vals); i++ {
		node.Right = &TreeNode{vals[i], nil, nil}
		node = node.Right
	}
	return head
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
func increasingBST(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    values := dfs(root, []int{})
    l := len(values)
    r := &TreeNode{}
    r.Val = values[0]
    for i, node := 1, r; i < l; i++ {
        node.Right = &TreeNode{values[i], nil, nil}
        node = node.Right
    }
    return r
}

func dfs(node *TreeNode, values []int) []int {
    if node == nil {
        return values
    }
    values = dfs(node.Left, values)
    values = append(values, node.Val)
    values = dfs(node.Right, values)
    return values
}
```