# 108. 将有序数组转换为二叉搜索树

### 简单

给你一个整数数组 nums ，其中元素已经按 升序 排列，请你将其转换为一棵 
平衡
 二叉搜索树。

### 示例 1：
![bst1](/file/img/btree1.jpg)

输入：nums = [-10,-3,0,5,9]

输出：[0,-3,9,-10,null,5]

![bst2](/file/img/btree2.jpg)

解释：[0,-10,5,null,-3,null,9] 也将被视为正确答案：

### 示例 2：
![bst3](/file/img/btree3.jpg)

输入：nums = [1,3]

输出：[3,1]

解释：[1,null,3] 和 [3,1] 都是高度平衡二叉搜索树。
 

### 提示：

1 <= nums.length <= 104

-104 <= nums[i] <= 104

nums 按 严格递增 顺序排列

```go

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
func sortedArrayToBST(nums []int) *TreeNode {
	l, r := 0, len(nums)-1
	if l > r {
		return nil
	}
	m := (l + r) / 2
	return &TreeNode{
		Val:   nums[m],
		Left:  sortedArrayToBST(nums[l:m]),
		Right: sortedArrayToBST(nums[m+1 : r+1]),
	}
}
```