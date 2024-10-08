# 1464. 数组中两元素的最大乘积 Maximum Product of Two Elements in an Array

### Easy

Given the array of integers nums, you will choose two different indices i and j of that array. Return the maximum value of (nums[i]-1) * (nums[j]-1).

### Example 1:

Input: nums = [3,4,5,2]
Output: 12 
Explanation: If you choose the indices i=1 and j=2 (indexed from 0), you will get the maximum value, that is, (nums[1]-1)* (nums[2]-1) = (4-1)* (5-1) = 3* 4 = 12. 

### Example 2:

Input: nums = [1,5,4,5]
Output: 16
Explanation: Choosing the indices i=1 and j=3 (indexed from 0), you will get the maximum value of (5-1)* (5-1) = 16.

### Example 3:

Input: nums = [3,7]
Output: 12

Constraints:

2 <= nums.length <= 500
1 <= nums[i] <= 10^3

### 解：

leetcode 414
leetcode 1913
leetcode 747


```go
//找到最大，第二大两个数即可，O(n) time.
func maxProduct(nums []int) int {
	max := 0
	sec := 0
	for _, n := range nums {
		if max < n {
			sec = max
			max = n
		} else if sec < n {
			sec = n
		}
	}
	return (max - 1) * (sec - 1)
}
```