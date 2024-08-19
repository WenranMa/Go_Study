# 747.至少是其他数字两倍的最大数 Largest Number At Least Twice of Others

### Easy

You are given an integer array nums where the largest integer is unique.

Determine whether the largest element in the array is at least twice as much as every other number in the array. If it is, return the index of the largest element, or return -1 otherwise.

### Example 1:

Input: nums = [3,6,1,0]
Output: 1
Explanation: 6 is the largest integer.
For every other number in the array x, 6 is at least twice as big as x.
The index of value 6 is 1, so we return 1.

### Example 2:

Input: nums = [1,2,3,4]
Output: -1
Explanation: 4 is less than twice the value of 3, so we return -1.

### Example 3:

Input: nums = [1]
Output: 0
Explanation: 1 is trivially at least twice the value as any other number because there are no other numbers.

Constraints:

1 <= nums.length <= 50
0 <= nums[i] <= 100
The largest element in nums is unique.

### 解：

找最大，第二大两个数。

leetcode 414
leetcode 1464.
leetcode 1913.

```go
// find max and second max.
// 注意边界情况。
func dominantIndex(nums []int) int {
	first := 0
	second := -1
	res := 0
	if len(nums) == 1 {
		return 0
	}
	for i, c := range nums {
		if first < c {
			second = first
			first = c
			res = i
		} else if second < c {
			second = c
		}
	}
	if second == 0 || first/second >= 2 {
		return res
	}
	return -1
}

```