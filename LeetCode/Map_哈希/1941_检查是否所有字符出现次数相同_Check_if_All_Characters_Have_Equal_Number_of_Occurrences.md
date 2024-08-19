# 1941. 检查是否所有字符出现次数相同 Check if All Characters Have Equal Number of Occurrences

### Easy

Given a string s, return true if s is a good string, or false otherwise.

A string s is good if all the characters that appear in s have the same number of occurrences (i.e., the same frequency).

### Example 1:

Input: s = "abacbc"
Output: true
Explanation: The characters that appear in s are 'a', 'b', and 'c'. All characters occur 2 times in s.

### Example 2:

Input: s = "aaabb"
Output: false
Explanation: The characters that appear in s are 'a' and 'b'.
'a' occurs 3 times while 'b' occurs 2 times, which is not the same number of times.

Constraints:

1 <= s.length <= 1000
s consists of lowercase English letters.

### 解：

```go
// map, O(n) time. O(n) space.
func areOccurrencesEqual(s string) bool {
	m := make(map[rune]int)
	for _, c := range s {
		m[c] += 1
	}
	t := 0 // 临时变量，判断第一个值
	for _, v := range m {
		if t == 0 {
			t = v
		}
		if t != v {
			return false
		}
	}
	return true
}
```