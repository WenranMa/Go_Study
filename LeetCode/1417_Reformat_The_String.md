# 1417. Reformat The String

### Easy

You are given an alphanumeric string s. (Alphanumeric string is a string consisting of lowercase English letters and digits).

You have to find a permutation of the string where no letter is followed by another letter and no digit is followed by another digit. That is, no two adjacent characters have the same type.

Return the reformatted string or return an empty string if it is impossible to reformat the string.

### Example 1:

Input: s = "a0b1c2"
Output: "0a1b2c"
Explanation: No two adjacent characters have the same type in "0a1b2c". "a0b1c2", "0a1b2c", "0c2a1b" are also valid permutations.

### Example 2:

Input: s = "leetcode"
Output: ""
Explanation: "leetcode" has only characters so we cannot separate them by digits.

### Example 3:

Input: s = "1229857369"
Output: ""
Explanation: "1229857369" has only digits so we cannot separate them by characters.

Constraints:

1 <= s.length <= 500
s consists of only lowercase English letters and/or digits.

```go
//Two pointers, O(n) time, O(n) space.
func reformat(s string) string {
	l, n, c := len(s), 0, 0
	for i, _ := range s {
		if s[i] <= 57 {
			n += 1
		} else {
			c += 1
		}
	}
	if n > c && n-c > 1 {
		return ""
	}
	if n < c && c-n > 1 {
		return ""
	}
	i, j, k := 0, 0, 0
	res := make([]byte, l)
	for k < l {
		for i < l && s[i] >= 97 {
			i += 1
		}
		for j < l && s[j] <= 57 {
			j += 1
		}
		if n > c {
			if k%2 == 0 {
				res[k] = s[i]
				i += 1
			} else {
				res[k] = s[j]
				j += 1
			}
		} else {
			if k%2 == 0 {
				res[k] = s[j]
				j += 1
			} else {
				res[k] = s[i]
				i += 1
			}
		}
		k += 1
	}
	return string(res)
}

// 改进，简化代码
func reformat(s string) string {
	l, n, c := len(s), 0, 0
	for i, _ := range s {
		if s[i] <= 57 {
			n += 1
		} else {
			c += 1
		}
	}
	if n > c && n-c > 1 || n < c && c-n > 1 {
		return ""
	}
	res := make([]byte, l)
	for i, j, k := 0, 0, 0; k < l; k++ {
		for i < l && s[i] >= 97 {
			i += 1
		}
		for j < l && s[j] <= 57 {
			j += 1
		}
		if n >= c && k%2 == 0 || n < c && k%2 == 1 {
			res[k] = s[i]
			i += 1
		} else {
			res[k] = s[j]
			j += 1
		}
	}
	return string(res)
}
```