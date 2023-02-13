# 2089. Find Target Indices After Sorting Array

### Easy

You are given a 0-indexed integer array nums and a target element target.

A target index is an index i such that nums[i] == target.

Return a list of the target indices of nums after sorting nums in non-decreasing order. If there are no target indices, return an empty list. The returned list must be sorted in increasing order.

### Example 1:

Input: nums = [1,2,5,2,3], target = 2
Output: [1,2]
Explanation: After sorting, nums is [1,2,2,3,5].
The indices where nums[i] == 2 are 1 and 2.

### Example 2:

Input: nums = [1,2,5,2,3], target = 3
Output: [3]
Explanation: After sorting, nums is [1,2,2,3,5].
The index where nums[i] == 3 is 3.

### Example 3:

Input: nums = [1,2,5,2,3], target = 5
Output: [4]
Explanation: After sorting, nums is [1,2,2,3,5].
The index where nums[i] == 5 is 4.

Constraints:

1 <= nums.length <= 100
1 <= nums[i], target <= 100

```go
func targetIndices(nums []int, target int) []int {
	res := []int{}
	ind := make([]int, 101)
	for _, n := range nums {
		ind[n] += 1
	}
	for i := 1; i < 101; i++ {
		ind[i] += ind[i-1]
	}
	for i := ind[target-1]; i < ind[target]; i++ {
		res = append(res, i)
	}
	return res
}

// better solution: less memory
func targetIndices(nums []int, target int) []int {
	res := []int{}
	tn := 0
	less := 0
	for _, n := range nums {
		if n < target {
			less += 1
		}
		if n == target {
			tn += 1
		}
	}
	for i := less; i < less+tn; i++ {
		res = append(res, i)
	}
	return res
}
```