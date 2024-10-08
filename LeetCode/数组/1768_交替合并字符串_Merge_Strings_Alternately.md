# 1768.交替合并字符串 Merge Strings Alternately

### Easy

You are given two strings word1 and word2. Merge the strings by adding letters in alternating order, starting with word1. If a string is longer than the other, append the additional letters onto the end of the merged string.

Return the merged string.

### Example 1:

Input: word1 = "abc", word2 = "pqr"
Output: "apbqcr"
Explanation: The merged string will be merged as so:
word1:  a   b   c
word2:    p   q   r
merged: a p b q c r

### Example 2:

Input: word1 = "ab", word2 = "pqrs"
Output: "apbqrs"
Explanation: Notice that as word2 is longer, "rs" is appended to the end.
word1:  a   b 
word2:    p   q   r   s
merged: a p b q   r   s

### Example 3:

Input: word1 = "abcd", word2 = "pq"
Output: "apbqcd"
Explanation: Notice that as word1 is longer, "cd" is appended to the end.
word1:  a   b   c   d
word2:    p   q 
merged: a p b q c   d

Constraints:

1 <= word1.length, word2.length <= 100
word1 and word2 consist of lowercase English letters.

### 解

可以归为双指针问题。

```go
// O(n) time. straight forward.
func mergeAlternately(word1 string, word2 string) string {
	l1 := len(word1)
	l2 := len(word2)
	l := l1
	if l < l2 {
		l = l2
	}
	res := []byte{}
	for i := 0; i < l; i++ {
		if i < l1 {
			res = append(res, word1[i])
		}
		if i < l2 {
			res = append(res, word2[i])
		}
	}
	return string(res)
}
```


```go
// O(n) time. straight forward.
func mergeAlternately(word1 string, word2 string) string {
	l1 := len(word1)
	l2 := len(word2)
	res := []byte{}
	for i, j := 0, 0; i < l1 || j < l2; i, j = i+1, j+1 {
		if i < l1 {
			res = append(res, word1[i])
		}
		if j < l2 {
			res = append(res, word2[j])
		}
	}
	return string(res)
}
```