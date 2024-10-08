# 1446. 连续字符 Consecutive Characters

### Easy

The power of the string is the maximum length of a non-empty substring that contains only one unique character.

Given a string s, return the power of s.

### Example 1:

Input: s = "leetcode"
Output: 2
Explanation: The substring "ee" is of length 2 with the character 'e' only.

### Example 2:

Input: s = "abbcccddddeeeeedcba"
Output: 5
Explanation: The substring "eeeee" is of length 5 with the character 'e' only.

Constraints:

1 <= s.length <= 500
s consists of only lowercase English letters.

### 解：

双指针，一个表示开头，一个遍历。遇到不同的，更新起始位置i, 计算长度

```go
// Two pointers. O(n) time, O(1) space.
// 注意边界处理
func maxPower(s string) int {
	l := len(s)
	i, j := 0, 0
	power := 0
	for j < l {
		if s[i] != s[j] {
			d := j - i
			if d > power {
				power = d
			}
			i = j
		}
		j++
	}
	if power < j-i { // “aaaa” 这种情况
		power = j - i
	}
	return power
}
```