# 1394. 找出数组中的幸运数 Find Lucky Integer in an Array

### Easy

Given an array of integers arr, a lucky integer is an integer that has a frequency in the array equal to its value.

Return the largest lucky integer in the array. If there is no lucky integer return -1.

### Example 1:

Input: arr = [2,2,3,4]
Output: 2
Explanation: The only lucky number in the array is 2 because frequency[2] == 2.

### Example 2:

Input: arr = [1,2,2,3,3,3]
Output: 3
Explanation: 1, 2 and 3 are all lucky numbers, return the largest of them.

### Example 3:

Input: arr = [2,2,2,3,3]
Output: -1
Explanation: There are no lucky numbers in the array.

Constraints:

1 <= arr.length <= 500
1 <= arr[i] <= 500

### 解：

```go
//O(n) time. O(n) space. map
func findLucky(arr []int) int {
	m := make(map[int]int)
	for _, a := range arr {
		m[a] += 1
	}
	max := -1
	for k, v := range m {
		if k == v && k > max {
			max = k
		}
	}
	return max
}
```