# 1550. 存在三个连续奇数的数组 Three Consecutive Odds

### Easy

Given an integer array arr, return true if there are three consecutive odd numbers in the array. Otherwise, return false.

### Example 1:

Input: arr = [2,6,4,1]
Output: false
Explanation: There are no three consecutive odds.

### Example 2:

Input: arr = [1,2,34,3,4,5,7,23,12]
Output: true
Explanation: [5,7,23] are three consecutive odds.

Constraints:

1 <= arr.length <= 1000
1 <= arr[i] <= 1000

### 解：

```go
// O(n) time. one pass.
func threeConsecutiveOdds(arr []int) bool {
	o := 0
	for _, n := range arr {
		if n%2 == 1 {
			o += 1
			if o == 3 {
				return true
			}
		} else {
			o = 0
		}
	}
	return false
}
```