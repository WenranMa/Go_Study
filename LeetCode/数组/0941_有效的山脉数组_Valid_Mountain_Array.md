# 941. 有效的山脉数组 Valid Mountain Array

### Easy

Given an array of integers arr, return true if and only if it is a valid mountain array.

Recall that arr is a mountain array if and only if:

arr.length >= 3
There exists some i with 0 < i < arr.length - 1 such that:
arr[0] < arr[1] < ... < arr[i - 1] < arr[i]
arr[i] > arr[i + 1] > ... > arr[arr.length - 1]

### Example 1:

Input: arr = [2,1]
Output: false

### Example 2:

Input: arr = [3,5,5]
Output: false

### Example 3:

Input: arr = [0,3,2,1]
Output: true

Constraints:

1 <= arr.length <= 10^4
0 <= arr[i] <= 10^4

### 解：

前后遍历，看peak是否一致。

```go
// O(n) time. go through the array forward and backward.
// check the peak is the same index.
// [4,3,2,1] is false.
func validMountainArray(arr []int) bool {
	l := len(arr)
	p := 0
	for i := 0; i < l-1; i++ {
		if arr[i] >= arr[i+1] {
			p = i
			break
		}
	}
	for i := l - 1; i >= 1; i-- {
		if arr[i] >= arr[i-1] {
			return p != 0 && p == i
		}
	}
	return false
}
```