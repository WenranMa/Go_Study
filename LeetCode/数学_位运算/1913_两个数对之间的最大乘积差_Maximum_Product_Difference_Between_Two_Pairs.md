# 1913. 两个数对之间的最大乘积差 Maximum Product Difference Between Two Pairs

### Easy

The product difference between two pairs (a, b) and (c, d) is defined as (a * b) - (c * d).

For example, the product difference between (5, 6) and (2, 7) is (5 * 6) - (2 * 7) = 16.
Given an integer array nums, choose four distinct indices w, x, y, and z such that the product difference between pairs (nums[w], nums[x]) and (nums[y], nums[z]) is maximized.

Return the maximum such product difference.

### Example 1:

Input: nums = [5,6,2,7,4]
Output: 34
Explanation: We can choose indices 1 and 3 for the first pair (6, 7) and indices 2 and 4 for the second pair (2, 4).
The product difference is (6 * 7) - (2 * 4) = 34.

### Example 2:

Input: nums = [4,2,5,9,7,4,8]
Output: 64
Explanation: We can choose indices 3 and 6 for the first pair (9, 8) and indices 1 and 5 for the second pair (2, 4).
The product difference is (9 * 8) - (2 * 4) = 64.

Constraints:

4 <= nums.length <= 104
1 <= nums[i] <= 104

### 解：

找到最大两个数和最小两个数。 一次遍历即可。

leetcode 1464
leetcode 414
leetcode 747


```go
// find the max 2 numbers and the minimum 2 numbers.
func maxProductDifference(nums []int) int {
	m1 := 0 // 最大
	m2 := 0 // 第二大
	s1 := 10001 // 最小
	s2 := 10001 // 第二小
	for _, n := range nums {
		if m1 < n {
			m2 = m1
			m1 = n
		} else if m2 < n {
			m2 = n
		}
		if s1 > n {
			s2 = s1
			s1 = n
		} else if s2 > n {
			s2 = n
		}
	}
	return m1*m2 - s1*s2
}
```