# 1207. Unique Number of Occurrences

### Easy

Given an array of integers arr, return true if the number of occurrences of each value in the array is unique, or false otherwise.

### Example 1:

Input: arr = [1,2,2,1,1,3]
Output: true
Explanation: The value 1 has 3 occurrences, 2 has 2 and 3 has 1. No two values have the same number of occurrences.

### Example 2:

Input: arr = [1,2]
Output: false

### Example 3:

Input: arr = [-3,0,1,-3,1,1,1,-3,10,0]
Output: true

Constraints:

1 <= arr.length <= 1000
-1000 <= arr[i] <= 1000

```go
// O(n) time. O(n) space.
// 用两个map，一个统计次数，一个统计次数的次数。
func uniqueOccurrences(arr []int) bool {
	m := make(map[int]int)
	mm := make(map[int]int)
	for _, a := range arr {
		m[a] += 1
	}
	for _, v := range m {
		mm[v] += 1
	}
	for _, v := range mm {
		if v > 1 {
			return false
		}
	}
	return true
}
```