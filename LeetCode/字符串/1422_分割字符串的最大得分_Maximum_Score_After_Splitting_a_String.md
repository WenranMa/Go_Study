# 1422. 分割字符串的最大得分 Maximum Score After Splitting a String

### Easy

Given a string s of zeros and ones, return the maximum score after splitting the string into two non-empty substrings (i.e. left substring and right substring).

The score after splitting a string is the number of zeros in the left substring plus the number of ones in the right substring.

### Example 1:

Input: s = "011101"
Output: 5 
Explanation: 
All possible ways of splitting s into two non-empty substrings are:
left = "0" and right = "11101", score = 1 + 4 = 5 
left = "01" and right = "1101", score = 1 + 3 = 4 
left = "011" and right = "101", score = 1 + 2 = 3 
left = "0111" and right = "01", score = 1 + 1 = 2 
left = "01110" and right = "1", score = 2 + 1 = 3

### Example 2:

Input: s = "00111"
Output: 5
Explanation: When left = "00" and right = "111", we get the maximum score = 2 + 3 = 5

### Example 3:

Input: s = "1111"
Output: 3

Constraints:

2 <= s.length <= 500
The string s consists of characters '0' and '1' only.

### 解：

```go
// Two passes, O(n) time, O(1) space.
func maxScore(s string) int {
	l := len(s)
	sz := 0
	for _, c := range s {
		if c == '0' {
			sz += 1
		}
	}
	max := 0
	left, right := 0, 0
	for i := 0; i < l-1; i++ {
		if s[i] == '0' {
			left += 1
		}
		right = (l - i - 1) - (sz - left)
		if max < left+right {
			max = left + right
		}
	}
	return max
}

//改进 two passes.
func maxScore(s string) int {
	l := len(s)
	so := 0
	for _, c := range s {
		if c == '1' {
			so += 1
		}
	}
	max := 0
	sz := 0
	for i := 0; i < l-1; i++ {
		if s[i] == '0' {
			sz += 1
		} else {
			so -= 1
		}
		if max < sz+so {
			max = sz + so
		}
	}
	return max
}

// 改进 one pass
// Max(zeroL + oneR) = Max(zeroL - oneL + oneL + oneR) = Max(zeroL - oneL) + oneTotal
func maxScore(s string) int {
	l := len(s)
	lo, lz := 0, 0
	max := -l - 1
	for i := 0; i < l; i++ {
		if s[i] == '0' {
			lz += 1
		} else {
			lo += 1
		}
		if i != l-1 && max < lz-lo {
			max = lz - lo
		}
	}
	return max + lo
}
```