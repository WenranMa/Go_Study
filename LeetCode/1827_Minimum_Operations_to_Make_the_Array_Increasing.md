# 1827. Minimum Operations to Make the Array Increasing

### Easy

You are given an integer array nums (0-indexed). In one operation, you can choose an element of the array and increment it by 1.

For example, if nums = [1,2,3], you can choose to increment nums[1] to make nums = [1,3,3].
Return the minimum number of operations needed to make nums strictly increasing.

An array nums is strictly increasing if nums[i] < nums[i+1] for all 0 <= i < nums.length - 1. An array of length 1 is trivially strictly increasing.

### Example 1:

Input: nums = [1,1,1]
Output: 3
Explanation: You can do the following operations:
1) Increment nums[2], so nums becomes [1,1,2].
2) Increment nums[1], so nums becomes [1,2,2].
3) Increment nums[2], so nums becomes [1,2,3].

### Example 2:

Input: nums = [1,5,2,4,1]
Output: 14
Example 3:

Input: nums = [8]
Output: 0
 

Constraints:

1 <= nums.length <= 5000
1 <= nums[i] <= 10^4

```go
//找到非上升的数字，然后+1再减去下一个数即可。
//改变了原有数组。
func minOperations(nums []int) int {
	ans := 0
	l := len(nums)
	for i := 0; i < l-1; i++ {
		if nums[i] >= nums[i+1] {
			ans += nums[i] + 1 - nums[i+1]
			nums[i+1] = nums[i] + 1
		}
	}
	return ans
}

//改进，不改变原数组
func minOperations(nums []int) int {
	ans := 0
	prev := 0
	for _, n := range nums {
		if n <= prev {
			prev += 1
			ans += prev - n
		} else {
			prev = n
		}
	}
	return ans
}
```