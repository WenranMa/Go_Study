# 414. 第三大的数 Third Maximum Number

### Easy

Given an integer array nums, return the third distinct maximum number in this array. If the third maximum does not exist, return the maximum number.

### Example 1:

Input: nums = [3,2,1]
Output: 1
Explanation:
The first distinct maximum is 3.
The second distinct maximum is 2.
The third distinct maximum is 1.

### Example 2:

Input: nums = [1,2]
Output: 2
Explanation:
The first distinct maximum is 2.
The second distinct maximum is 1.
The third distinct maximum does not exist, so the maximum (2) is returned instead.

### Example 3:

Input: nums = [2,2,3,1]
Output: 1
Explanation:
The first distinct maximum is 3.
The second distinct maximum is 2 (both 2's are counted together since they have the same value).
The third distinct maximum is 1.

Constraints:

1 <= nums.length <= 104
-231 <= nums[i] <= 231 - 1

Follow up: Can you find an O(n) solution?

### 解：

和找第二大的数类似。多一层

leetcode 1464.
leetcode 1913.
leetcode 747

```go
// O(n) time, O(1) space. 定义三个数即可。
// 注意处理重复数据，以及系统的int类型位数，默认64位。
func thirdMax(nums []int) int {
	f := math.MinInt64
	s := math.MinInt64
	t := math.MinInt64
	for _, n := range nums {
		if f < n {
			t = s
			s = f
			f = n

		} else if s < n && f != n { // 不等于用于去重
			t = s
			s = n
		} else if t < n && s != n && f != n {  // 不等于用于去重
			t = n
		}
	}
	if t == math.MinInt64 {
		return f
	}
	return t
}
```