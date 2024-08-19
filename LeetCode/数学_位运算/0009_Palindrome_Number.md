# 9. 回文数

### 简单

给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。

回文数 是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

例如，121 是回文，而 123 不是。

### 示例 1：

    输入：x = 121
    输出：true

### 示例 2：

    输入：x = -121
    输出：false
    解释：从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。

### 示例 3：

    输入：x = 10
    输出：false
    解释：从右向左读, 为 01 。因此它不是一个回文数。

### 提示：

-2^31 <= x <= 2^31 - 1

### 进阶：

你能不将整数转为字符串来解决这个问题吗？

### 解：

第一个for循环将divisor倍乘到与x相同的位数，用于去最高位。

第二个循环，high是最高位，low是最低位，比较这两位是否相等。

x= x % divisor/ 10 用于去掉最高位和最低位。

divisor/= 100 用于使divisor与x位数相同。

```go
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	divisor := 1
	for x/divisor >= 10 {
		divisor *= 10
	}
	for x > 0 {
		high := x / divisor
		low := x % 10
		if low != high {
			return false
		}
		x = x % divisor / 10
		divisor /= 100
	}
	return true
}
```
