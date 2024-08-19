# 1317. 将整数转换为两个无零整数的和 Convert Integer to the Sum of Two No-Zero Integers

### Easy

No-Zero integer is a positive integer that does not contain any 0 in its decimal representation.

Given an integer n, return a list of two integers [A, B] where:

A and B are No-Zero integers.
A + B = n
The test cases are generated so that there is at least one valid solution. If there are many valid solutions you can return any of them.

### Example 1:

Input: n = 2
Output: [1,1]
Explanation: A = 1, B = 1. A + B = n and both A and B do not contain any 0 in their decimal representation.

### Example 2:

Input: n = 11
Output: [2,9]

Constraints:

2 <= n <= 10^4

### 解：

模拟，循环检查

```go
// for loop to check every i and n-i pair.
func getNoZeroIntegers(n int) []int {
	for i := 1; i <= n/2; i++ {
		if isNZ(i) && isNZ(n-i) {
			return []int{i, n - i}
		}
	}
	return []int{1, 1}
}

func isNZ(n int) bool {
	for n > 0 {
		if n%10 == 0 {
			return false
		}
		n = n / 10
	}
	return true
}
```