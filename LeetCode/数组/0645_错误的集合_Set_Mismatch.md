# 645. 错误的集合 Set Mismatch

### Easy

You have a set of integers s, which originally contains all the numbers from 1 to n. Unfortunately, due to some error, one of the numbers in s got duplicated to another number in the set, which results in repetition of one number and loss of another number.

You are given an integer array nums representing the data status of this set after the error.

Find the number that occurs twice and the number that is missing and return them in the form of an array.

### Example 1:

Input: nums = [1,2,2,4]
Output: [2,3]

### Example 2:

Input: nums = [1,1]
Output: [1,2]

Constraints:

2 <= nums.length <= 10^4
1 <= nums[i] <= 10^4

### 解：

1. 可以用map.
```go
func findErrorNums(nums []int) []int {
	m := make(map[int]int)
	res := make([]int, 2)
	for _, n := range nums {
		m[n] += 1
	}
	for i := 1; i <= len(nums); i++ {
		if v, ok := m[i]; ok {
			if v == 2 {
				res[0] = i
			}
		} else {
			res[1] = i
		}
	}
	return res
}
```

2. 也可以类似 leetcode 268, 442, 448

```go
func findErrorNums(nums []int) []int {
	res := make([]int, 2)
	for i := 0; i < len(nums); {
		if nums[i] != i+1 && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		} else {
			i += 1
		}
	}
	for i, n := range nums {
		if n != i+1 {
			res[0] = n
			res[1] = i + 1
		}
	}
	return res
}
```