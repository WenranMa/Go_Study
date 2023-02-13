# 367. Valid Perfect Square

### Easy

Given a positive integer num, write a function which returns True if num is a perfect square else False.

Follow up: Do not use any built-in library function such as sqrt.

### Example 1:

Input: num = 16
Output: true

### Example 2:

Input: num = 14
Output: false

Constraints:

1 <= num <= 2^31 - 1

```go
// perfect squre 等于 1 + 3 + 5 + 7 ...
// 1 + 3 + 5 .. + (2n -1) = (1 + 2n - 1) * n /2 = n*n
func isPerfectSquare(num int) bool {
	i := 1
	for num > 0 {
		num -= i
		i += 2
	}
	return num == 0
}

```