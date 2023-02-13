# 409. Longest Palindrome

### Easy

Given a string s which consists of lowercase or uppercase letters, return the length of the longest palindrome that can be built with those letters.

Letters are case sensitive, for example, "Aa" is not considered a palindrome here.

### Example 1:

Input: s = "abccccdd"
Output: 7
Explanation:
One longest palindrome that can be built is "dccaccd", whose length is 7.

### Example 2:

Input: s = "a"
Output: 1

### Example 3:

Input: s = "bb"
Output: 2

Constraints:

1 <= s.length <= 2000
s consists of lowercase and/or uppercase English letters only.

```go
// O(n) time, O(n) space.
// 用map. 偶数个数直接加，奇数的减一，最后补一个1即可。
func longestPalindrome(s string) int {
	res := 0
	m := make(map[rune]int)
	for _, c := range s {
		m[c] += 1
	}
	odd := false
	for _, v := range m {
		if v%2 == 0 {
			res += v
		} else {
			odd = true
			res += v - 1
		}
	}
	if odd {
		res += 1
	}
	return res
}
```