# 1304.和为零的N个不同整数 Find N Unique Integers Sum up to Zero

### Easy

Given an integer n, return any array containing n unique integers such that they add up to 0.

### Example 1:

Input: n = 5
Output: [-7,-1,1,3,4]
Explanation: These arrays also are accepted [-5,-1,1,2,3] , [-3,-1,2,-2,4].

### Example 2:

Input: n = 3
Output: [-1,0,1]

### Example 3:

Input: n = 1
Output: [0]

Constraints:

1 <= n <= 1000

### 解：

奇数就多一个0，
偶数就一正一负，从1，-1开始

```go
func sumZero(n int) []int {
	res := []int{}
	if n%2 == 1 {
		res = append(res, 0)
	}
	for i := 1; i <= n/2; i++ {
		res = append(res, i)
		res = append(res, -i)
	}
	return res
}
```