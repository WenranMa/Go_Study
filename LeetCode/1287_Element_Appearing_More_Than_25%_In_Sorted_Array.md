# 1287. Element Appearing More Than 25% In Sorted Array

### Easy

Given an integer array sorted in non-decreasing order, there is exactly one integer in the array that occurs more than 25% of the time, return that integer.

### Example 1:

Input: arr = [1,2,2,6,6,6,6,7,10]
Output: 6

### Example 2:

Input: arr = [1,1]
Output: 1

Constraints:

1 <= arr.length <= 10^4
0 <= arr[i] <= 10^5

```go
// O(n) time. O(1) space.
func findSpecialInteger(arr []int) int {
	cnt := 0
	max := 0
	res := 0
	for i, a := range arr {
		if i > 0 && a != arr[i-1] {
			cnt = 1
		} else {
			cnt += 1
			if max < cnt {
				max = cnt
				res = a
			}
		}
	}
	return res
}


```