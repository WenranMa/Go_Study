# 1512. Number of Good Pairs

### Easy

Given an array of integers nums, return the number of good pairs.

A pair (i, j) is called good if nums[i] == nums[j] and i < j.

### Example 1:

Input: nums = [1,2,3,1,1,3]
Output: 4
Explanation: There are 4 good pairs (0,3), (0,4), (3,4), (2,5) 0-indexed.

### Example 2:

Input: nums = [1,1,1,1]
Output: 6
Explanation: Each pair in the array are good.

### Example 3:

Input: nums = [1,2,3]
Output: 0

Constraints:

1 <= nums.length <= 100
1 <= nums[i] <= 100

```go
//找规律
// 1:0
// 2:1
// 3:3
// 4:6
// 5:10
// 6:15
// 所以 n = 1 + 2 + 3 + ... + (n - 1)

func numIdenticalPairs(nums []int) int {
	res := 0
	m := make(map[int]int)
	for _, n := range nums {
		m[n] += 1
	}
	for _, v := range m {
		if v > 1 {
			res += calPairNum(v - 1) // res += v * (v - 1) / 2
		}
	}
	return res
}

func calPairNum(n int) int {
	return (1 + n) * n / 2
}
```