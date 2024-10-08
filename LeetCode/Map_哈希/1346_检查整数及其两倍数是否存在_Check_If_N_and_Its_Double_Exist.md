# 1346. 检查整数及其两倍数是否存在 Check If N and Its Double Exist

### Easy

Given an array arr of integers, check if there exists two integers N and M such that N is the double of M ( i.e. N = 2 * M).

More formally check if there exists two indices i and j such that :

i != j
0 <= i, j < arr.length
arr[i] == 2 * arr[j]
 
### Example 1:

Input: arr = [10,2,5,3]
Output: true
Explanation: N = 10 is the double of M = 5,that is, 10 = 2 * 5.

### Example 2:

Input: arr = [7,1,14,11]
Output: true
Explanation: N = 14 is the double of M = 7,that is, 14 = 2 * 7.

### Example 3:

Input: arr = [3,1,7,11]
Output: false
Explanation: In this case does not exist N and M, such that N = 2 * M.

Constraints:

2 <= arr.length <= 500
-10^3 <= arr[i] <= 10^3

### 解：

map 可以一次遍历：

```go
func checkIfExist(arr []int) bool {
	m := make(map[int]int)
	for _, a := range arr {
		if _, ok := m[2*a]; ok {
			return true
		}
		if a%2 == 0 {
			if _, ok := m[a/2]; ok {
				return true
			} else {
				m[a] = 1
			}
		} else {
			m[a] = 1
		}
	}
	return false
}
```

二次遍历

```go
// Map. O(n) time, O(n) space.
// 注意0特殊处理。
func checkIfExist(arr []int) bool {
	m := make(map[int]int)
	for _, a := range arr {
		m[a] += 1
	}
	if m[0] > 1 {
		return true
	}
	for _, a := range arr {
		if _, ok := m[2*a]; ok && a != 0 {
			return true
		}
	}
	return false
}
```


