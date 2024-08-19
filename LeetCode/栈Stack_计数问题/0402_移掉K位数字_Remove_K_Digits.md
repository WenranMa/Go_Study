# 402.移掉K位数字 Remove K Digits

### Medium

Given string num representing a non-negative integer num, and an integer k, return the smallest possible integer after removing k digits from num.

### Example 1:

Input: num = "1432219", k = 3
Output: "1219"
Explanation: Remove the three digits 4, 3, and 2 to form the new number 1219 which is the smallest.

### Example 2:

Input: num = "10200", k = 1
Output: "200"
Explanation: Remove the leading 1 and the number is 200. Note that the output must not contain leading zeroes.

### Example 3:

Input: num = "10", k = 2
Output: "0"
Explanation: Remove all the digits from the number and it is left with nothing which is 0.

Constraints:

1 <= k <= num.length <= 105
num consists of only digits.
num does not have any leading zeros except for the zero itself.

### 解：

栈，遇到比栈顶小的字符，就弹出。然后压入当前值。

corner case：
1. k 和字符串长度相同，返回0.
2. "111" 这种情况。直接弹出k个。
3. "0202" 这种，去掉0.

```go
func removeKdigits(num string, k int) string {
	l := len(num)
	//corner case
	if k == l {
		return "0"
	}
	stack := []byte{}
	for i := 0; i < l; i++ {
		for k > 0 && len(stack) > 0 && stack[len(stack)-1] > num[i] {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, num[i])
	}
	// corner case like "1111"
	for k > 0 {
		stack = stack[:len(stack)-1]
		k--
	}
	// remove heading 0
	for len(stack) > 1 && stack[0] == '0' {
		stack = stack[1:len(stack)]
	}
	return string(stack)
}
```