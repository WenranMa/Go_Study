# 1991. 找到数组的中间位置 Find the Middle Index in Array

### Easy

Given a 0-indexed integer array nums, find the leftmost middleIndex (i.e., the smallest amongst all the possible ones).

A middleIndex is an index where nums[0] + nums[1] + ... + nums[middleIndex-1] == nums[middleIndex+1] + nums[middleIndex+2] + ... + nums[nums.length-1].

If middleIndex == 0, the left side sum is considered to be 0. Similarly, if middleIndex == nums.length - 1, the right side sum is considered to be 0.

Return the leftmost middleIndex that satisfies the condition, or -1 if there is no such index.

### Example 1:

Input: nums = [2,3,-1,8,4]
Output: 3
Explanation: The sum of the numbers before index 3 is: 2 + 3 + -1 = 4
The sum of the numbers after index 3 is: 4 = 4

### Example 2:

Input: nums = [1,-1,4]
Output: 2
Explanation: The sum of the numbers before index 2 is: 1 + -1 = 0
The sum of the numbers after index 2 is: 0

### Example 3:

Input: nums = [2,5]
Output: -1
Explanation: There is no valid middleIndex.

Constraints:

1 <= nums.length <= 100
-1000 <= nums[i] <= 1000

Note: This question is the same as 724

### 解：

遍历数组，计算前缀和，后缀和，如果前缀和等于后缀和，则返回当前索引。

后缀和初始值是 1 到 l-1的和。
前缀和初始值为 0

时间复杂度：O(n)

遍历数组，前缀和 += nums[i-1]，后缀和 -= nums[i], i 从 1 开始。

```go
//Prefix sum.
func pivotIndex(nums []int) int {
	sl := 0
	sr := 0
	l := len(nums)
	for i := l - 1; i >= 1; i-- {
		sr += nums[i]
	}
	if sl == sr {
		return 0
	}
	for i := 1; i < l; i++ {
		sl += nums[i-1]
		sr -= nums[i]
		if sl == sr {
			return i
		}
	}
	return -1
}
```