# 1957. 删除字符使字符串变好 Delete Characters to Make Fancy String

### Easy

A fancy string is a string where no three consecutive characters are equal.

Given a string s, delete the minimum possible number of characters from s to make it fancy.

Return the final string after the deletion. It can be shown that the answer will always be unique.

### Example 1:

Input: s = "leeetcode"
Output: "leetcode"
Explanation:
Remove an 'e' from the first group of 'e's to create "leetcode".
No three consecutive characters are equal, so return "leetcode".

### Example 2:

Input: s = "aaabaaaa"
Output: "aabaa"
Explanation:
Remove an 'a' from the first group of 'a's to create "aabaaaa".
Remove two 'a's from the second group of 'a's to create "aabaa".
No three consecutive characters are equal, so return "aabaa".

### Example 3:

Input: s = "aab"
Output: "aab"
Explanation: No three consecutive characters are equal, so return "aab".

Constraints:

1 <= s.length <= 10^5
s consists only of lowercase English letters.

### 解：

Leetcode 80 一样。但是这个不需要in place. 

```go
// O(n) time. One pass.
func makeFancyString(s string) string {
	res := []byte{}
	c := 1
	for i, _ := range s {
		if i > 0 && s[i] == s[i-1] {
			c += 1
			if c >= 3 {
				continue
			}
		} else {
			c = 1
		}
		res = append(res, s[i])
	}
	return string(res)
}
```