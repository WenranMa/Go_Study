# 415. Add Strings

### Easy

Given two non-negative integers, num1 and num2 represented as string, return the sum of num1 and num2 as a string.

You must solve the problem without using any built-in library for handling large integers (such as BigInteger). You must also not convert the inputs to integers directly.

### Example 1:

Input: num1 = "11", num2 = "123"
Output: "134"

### Example 2:

Input: num1 = "456", num2 = "77"
Output: "533"

### Example 3:

Input: num1 = "0", num2 = "0"
Output: "0"

Constraints:

1 <= num1.length, num2.length <= 104
num1 and num2 consist of only digits.
num1 and num2 don't have any leading zeros except for the zero itself.

```go
// O(n) time, O(n) space. Straight forward.
func addStrings(num1 string, num2 string) string {
	res := []byte{}
	l1 := len(num1)
	l2 := len(num2)
	var b byte = 0
	for i, j := l1-1, l2-1; i >= 0 || j >= 0; i, j = i-1, j-1 {
		var n1 byte = 0
		if i >= 0 {
			n1 = num1[i] - '0'
		}
		var n2 byte = 0
		if j >= 0 {
			n2 = num2[j] - '0'
		}
		n := n1 + n2 + b
		res = append([]byte{n%10 + '0'}, res...)
		b = n / 10
	}
	if b > 0 {
		res = append([]byte{b + '0'}, res...)
	}
	return string(res)
}
```