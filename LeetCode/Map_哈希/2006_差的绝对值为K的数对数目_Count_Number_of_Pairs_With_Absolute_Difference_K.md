# 2006. 差的绝对值为K的数对数目 Count Number of Pairs With Absolute Difference K

### Easy

Given an integer array nums and an integer k, return the number of pairs (i, j) where i < j such that |nums[i] - nums[j]| == k.

The value of |x| is defined as:

x if x >= 0.
-x if x < 0.

### Example 1:

Input: nums = [1,2,2,1], k = 1
Output: 4
Explanation: The pairs with an absolute difference of 1 are:
- [1,2],[1,2],[2,1],[2,1]

### Example 2:

Input: nums = [1,3], k = 3
Output: 0
Explanation: There are no pairs with an absolute difference of 3.

### Example 3:

Input: nums = [3,2,1,5,4], k = 2
Output: 3
Explanation: The pairs with an absolute difference of 2 are:
- [3,1],[3,5],[2,4]

Constraints:

1 <= nums.length <= 200
1 <= nums[i] <= 100
1 <= k <= 99

### 解：

map

```go
func countKDifference(nums []int, k int) int {
	m := make(map[int]int)
	var res int
	for _, n := range nums {
		res += m[n+k] + m[n-k]
		m[n] += 1
	}
	return res
}
```


```go
//counting sort. use num as index.
//index + k is the diff nums. 
func countKDifference(nums []int, k int) int {
	set := make([]int, 101)
	res := 0
	for _, n := range nums {
		set[n] += 1
	}
	for i, _ := range set {
		if i+k < 101 {
			res += set[i] * set[i+k]
		}
	}
	return res
}
```