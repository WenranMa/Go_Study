# 796. 旋转字符串 Rotate String

### Easy

Given two strings s and goal, return true if and only if s can become goal after some number of shifts on s.

A shift on s consists of moving the leftmost character of s to the rightmost position.

For example, if s = "abcde", then it will be "bcdea" after one shift.

### Example 1:

Input: s = "abcde", goal = "cdeab"
Output: true

### Example 2:

Input: s = "abcde", goal = "abced"
Output: false

Constraints:

1 <= s.length, goal.length <= 100
s and goal consist of lowercase English letters.

### 解：

```go
// 相加找substring.
func rotateString(s string, goal string) bool {
	if len(goal) != len(s) {
		return false
	}
	temp := s + s
	return strings.Contains(temp, goal)
}
```