# 654_最大二叉树_Maximum_Binary_Tree
Given an integer array with no duplicates. A maximum tree building on this array is defined as follow:

The root is the maximum number in the array.
The left subtree is the maximum tree constructed from left part subarray divided by the maximum number.
The right subtree is the maximum tree constructed from right part subarray divided by the maximum number.
Construct the maximum tree by the given array and output the root node of this tree.

Example 1:  
    Input: [3,2,1,6,0,5]  
    Output: return the tree root node representing the following tree:  
        6  
      /   \  
     3     5  
      \    /   
       2  0   
        \  
         1  

6是最大值，左边3，2，1中最大值3位left. 右边0，5中最大值5位right.

Note:  
The size of the given array will be in the range [1,1000].


```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructMaximumBinaryTree(nums []int) *TreeNode {
    l := len(nums)
    if l == 0 {
        return nil
    }
    index := 0
    for i, n := range nums {
        if nums[index] < n { //find the max
            index = i
        }
    }
    root := &TreeNode{nums[index], nil, nil}
    root.Left = constructMaximumBinaryTree(nums[:index])
    root.Right = constructMaximumBinaryTree(nums[index+1:])
    return root
}
```