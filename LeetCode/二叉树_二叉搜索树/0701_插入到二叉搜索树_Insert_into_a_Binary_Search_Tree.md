# 701_插入到二叉搜索树_Insert_into_a_Binary_Search_Tree
Given the root node of a binary search tree (BST) and a value to be inserted into the tree, insert the value into the BST. Return the root node of the BST after the insertion. It is guaranteed that the new value does not exist in the original BST.

Note that there may exist multiple valid ways for the insertion, as long as the tree remains a BST after insertion. You can return any of them.

For example, 

    Given the tree:  
            4  
           / \  
          2   7  
         / \  
        1   3  
    And the value to insert: 5  
    You can return this binary search tree:  
            4    
           /  \  
          2    7  
         / \   /  
        1   3  5 

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
func insertIntoBST(root *TreeNode, val int) *TreeNode {
    if root == nil {
        return &TreeNode{val, nil, nil}
    }
    if root.Val < val {
        root.Right = insertIntoBST(root.Right, val)
    } else if root.Val > val {
        root.Left = insertIntoBST(root.Left, val)
    }
    return root
}
```