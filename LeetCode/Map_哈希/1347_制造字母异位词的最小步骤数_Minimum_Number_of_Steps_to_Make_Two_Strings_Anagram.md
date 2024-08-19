# 1347. 制造字母异位词的最小步骤数 Minimum Number of Steps to Make Two Strings Anagram

### Medium

You are given two strings of the same length s and t. In one step you can choose any character of t and replace it with another character.

Return the minimum number of steps to make t an anagram of s.

An Anagram of a string is a string that contains the same characters with a different (or the same) ordering.

### Example 1:

Input: s = "bab", t = "aba"
Output: 1
Explanation: Replace the first 'a' in t with b, t = "bba" which is anagram of s.

### Example 2:

Input: s = "leetcode", t = "practice"
Output: 5
Explanation: Replace 'p', 'r', 'a', 'i' and 'c' from t with proper characters to make t anagram of s.

### Example 3:

Input: s = "anagram", t = "mangaar"
Output: 0
Explanation: "anagram" and "mangaar" are anagrams. 

Constraints:

1 <= s.length <= 5 * 104
s.length == t.length
s and t consist of lowercase English letters only.

### 解：

```go
// 用map. 记录target里面的字符和count，
// 遍历source，减去map中对应字符的count.
// 加 target > source的差值到结果，如果target < source则忽略。
// O(n) time, O(n) space.


//改进，用一个map
func minSteps(s string, t string) int {
	m := make(map[rune]int)
	for _, c := range t {
		m[c] += 1
	}
	for _, c := range s {
		if val, ok := m[c]; ok && val > 0 {
			m[c] -= 1
		}
	}
	res := 0
	for _, v := range m {
		res += v
	}
	return res
}

//再改进，一个map, one pass.
func minSteps(s string, t string) int {
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		m[t[i]] += 1
		m[s[i]] -= 1
	}
	res := 0
	for _, v := range m {
		if v > 0 {
			res += v
		}
	}
	return res
}
```