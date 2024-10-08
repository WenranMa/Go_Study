# 617_合并二叉树_Merge Two Binary Trees

Given two binary trees and imagine that when you put one of them to cover the other, some nodes of the two trees are overlapped while the others are not. You need to merge them into a new binary tree. The merge rule is that if two nodes overlap, then sum node values up as the new value of the merged node. Otherwise, the NOT null node will be used as the node of new tree.

Example 1:

    Input: 
    Tree 1                     Tree 2                  
          1                         2                             
         / \                       / \                            
        3   2                     1   3                        
       /                           \   \                      
      5                             4   7                  
    Output:  
    Merged tree:    
            3  
            / \  
           4   5  
          / \   \   
         5   4   7  
 
Note: The merging process must start from the root nodes of both trees.

### 解：

递归

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
    if t1 == nil && t2 == nil {
        return nil
    } else if t1 == nil {
        return t2
    } else if t2 == nil {
        return t1
    } else {
        t := &TreeNode{0, nil, nil}
        t.Val = t1.Val + t2.Val
        t.Left = mergeTrees(t1.Left, t2.Left)
        t.Right = mergeTrees(t1.Right, t2.Right)
        return t
    }
    return nil
}
```

```go
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil && t2 == nil {
		return nil
	} else if t1 == nil {
		return t2
	} else if t2 == nil {
		return t1
	} else {
		return &TreeNode{t1.Val + t2.Val, mergeTrees(t1.Left, t2.Left), mergeTrees(t1.Right, t2.Right)}
	}
}
```