# 1662.  检查两个字符串数组是否相等 Check If Two String Arrays are Equivalent

### Easy

Given two string arrays word1 and word2, return true if the two arrays represent the same string, and false otherwise.

A string is represented by an array if the array elements concatenated in order forms the string.

### Example 1:

Input: word1 = ["ab", "c"], word2 = ["a", "bc"]
Output: true
Explanation:
word1 represents string "ab" + "c" -> "abc"
word2 represents string "a" + "bc" -> "abc"
The strings are the same, so return true.

### Example 2:

Input: word1 = ["a", "cb"], word2 = ["ab", "c"]
Output: false
Example 3:

Input: word1  = ["abc", "d", "defg"], word2 = ["abcddefg"]
Output: true

Constraints:

1 <= word1.length, word2.length <= 103
1 <= word1[i].length, word2[i].length <= 103
1 <= sum(word1[i].length), sum(word2[i].length) <= 103
word1[i] and word2[i] consist of lowercase letters.

### 解：

模拟

```go
func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	return strings.Join(word1, "") == strings.Join(word2, "")
}

// O(1) space, use 4 pointers.
func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	l1 := len(word1)
	l2 := len(word2)
	w1 := 0
	c1 := 0
	w2 := 0
	c2 := 0
	for w1 < l1 && w2 < l2 {
		lw1 := len(word1[w1])
		lw2 := len(word2[w2])
		if word1[w1][c1] != word2[w2][c2] {
			return false
		} else {
			c1 += 1
			c2 += 1
			if c1 >= lw1 {
				c1 = 0
				w1 += 1
			}
			if c2 >= lw2 {
				c2 = 0
				w2 += 1
			}
		}
	}
	if (w1 < l1 && w2 >= l2) || (w1 >= l1 && w2 < l2) {
		return false
	}
	return true
}
```