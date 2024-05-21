# 67. 二进制求和

### 简单

给你两个二进制字符串 a 和 b ，以二进制字符串的形式返回它们的和。

### 示例 1：

输入:a = "11", b = "1"
输出："100"

### 示例 2：

输入：a = "1010", b = "1011"
输出："10101"

### 提示：
- 1 <= a.length, b.length <= 10^4
- a 和 b 仅由字符 '0' 或 '1' 组成
- 字符串如果不是 "0" ，就不含前导零

### 解：

```go
func addBinary(a string, b string) string {
	carry := 0
	res := ""
	for i, j := len(a)-1, len(b)-1; i >= 0 || j >= 0; i, j = i-1, j-1 {
		n1, n2 := 0, 0
		if i >= 0 {
			n1 = int(a[i] - '0')
		} else {
			n1 = 0
		}
		if j >= 0 {
			n2 = int(b[j] - '0')
		} else {
			n2 = 0
		}
		val := (n1 + n2 + carry) % 2
		carry = (n1 + n2 + carry) / 2
		res = fmt.Sprint(val) + res
	}
	if carry > 0 {
		res = "1" + res
	}
	return res
}
```