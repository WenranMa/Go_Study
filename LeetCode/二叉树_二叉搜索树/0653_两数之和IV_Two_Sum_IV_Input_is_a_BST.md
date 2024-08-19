# 653_两数之和IV_Two_Sum_IV_Input_is_a_BST
Given a Binary Search Tree and a target number, return true if there exist two elements in the BST such that their sum is equal to the given target.

Example 1:

    Input:
        5
       / \
      3   6
     / \   \
    2   4   7

    Target = 9
    Output: True

Example 2:

    Input:
        5
       / \
      3   6
     / \   \
    2   4   7

    Target = 28
    Output: False

### 解：

inorder 遍历，然后双指针

```go
func findTarget(root *TreeNode, k int) bool {
	var nums []int
	var inOrder func(*TreeNode)
	inOrder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inOrder(root.Left)
		nums = append(nums, root.Val)
		inOrder(root.Right)
	}
    inOrder(root)
	l, r := 0, len(nums)-1
	for l < r {
		if nums[l]+nums[r] < k {
			l++
		} else if nums[l]+nums[r] > k {
			r--
		} else {
			return true
		}
	}
	return false
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
func findTarget(root *TreeNode, k int) bool {
    nums := inOrder(root, []int{})
    l, r := 0, len(nums)-1
    for l < r {
        if nums[l]+nums[r] < k {
            l++
        } else if nums[l]+nums[r] > k {
            r--
        } else {
            return true
        }
    }
    return false
}

func inOrder(root *TreeNode, nums []int) []int {
    if root == nil {
        return nums
    }
    nums = inOrder(root.Left, nums)
    nums = append(nums, root.Val)
    nums = inOrder(root.Right, nums)
    return nums
}
```