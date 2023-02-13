# 434. Number of Segments in a String

### Easy

Given a string s, return the number of segments in the string.

A segment is defined to be a contiguous sequence of non-space characters.

### Example 1:

Input: s = "Hello, my name is John"
Output: 5
Explanation: The five segments are ["Hello,", "my", "name", "is", "John"]

### Example 2:

Input: s = "Hello"
Output: 1

Constraints:

0 <= s.length <= 300
s consists of lowercase and uppercase English letters, digits, or one of the following characters "!@#$%^&* () _ +-=',.:".
The only space character in s is ' '.

```go
// O(n) time. 只有前面有空格并且字符不是' '时结果加1. 
// 注意初始化。
func countSegments(s string) int {
	res := 0
	space := true
	for _, c := range s {
		if space && c != ' ' {
			res += 1
			space = false
		} else if c == ' ' {
			space = true
		}
	}
	return res
}
```