# 6. Z 字形变换

### 中等

将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行 Z 字形排列。

比如输入字符串为 "PAYPALISHIRING" 行数为 3 时，排列如下：

	P   A   H   N
	A P L S I I G
	Y   I   R

之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："PAHNAPLSIIGYIR"。

### 示例 1：

	输入：s = "PAYPALISHIRING", numRows = 3
	输出："PAHNAPLSIIGYIR"

### 示例 2：

	输入：s = "PAYPALISHIRING", numRows = 4
	输出："PINALSIGYAHRPI"
	解释：
	P     I    N
	A   L S  I G
	Y A   H R
	P     I

### 示例 3：

	输入：s = "A", numRows = 1
	输出："A"

### 提示：
- 1 <= s.length <= 1000
- s 由英文字母（小写和大写）、',' 和 '.' 组成
- 1 <= numRows <= 1000

### 解：
大循环处理每一行。内部循环处理每个字符。如果不是第一行和最后一行，则每步要退回两步来插入中间的字符。O(n) time. O(n) space.

```go
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	res := []byte{}
	step := 2*numRows - 2
	for i := 0; i < numRows; i++ {
		cur := i
		for cur < len(s) {
			res = append(res, s[cur])
			cur += step
			if i > 0 && i < numRows-1 && (cur-i-i) < len(s) {
				res = append(res, s[cur-i-i])
			}
		}
	}
	return string(res)
}
```