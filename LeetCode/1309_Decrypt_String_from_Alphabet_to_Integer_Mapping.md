# 1309. Decrypt String from Alphabet to Integer Mapping

### Easy

You are given a string s formed by digits and '#'. We want to map s to English lowercase characters as follows:

Characters ('a' to 'i') are represented by ('1' to '9') respectively.
Characters ('j' to 'z') are represented by ('10#' to '26#') respectively.
Return the string formed after mapping.

The test cases are generated so that a unique mapping will always exist.

### Example 1:

Input: s = "10#11#12"
Output: "jkab"
Explanation: "j" -> "10#" , "k" -> "11#" , "a" -> "1" , "b" -> "2".

### Example 2:

Input: s = "1326#"
Output: "acz"

Constraints:

1 <= s.length <= 1000
s consists of digits and the '#' letter.
s will be a valid string such that mapping is always possible.

```go
// 'a' == 97
// '0' == 48
// '1' == 49
// so in the map, rule is '1' + 48 = 'a' or
// '0' to '9' is 'a' to 'i'
// 'j' == 'a' + 9 = 106
// so in the map, rule is 'a' + 9 = 'j'.
// O(n) time.
func freqAlphabets(s string) string {
	p := len(s) - 1
	res := []byte{}
	for p >= 0 {
		if s[p] == '#' {
			n := (s[p-2]-'0')*10 + (s[p-1] - '0')
			res = append([]byte{'a' + n - 1}, res...)
			p -= 3
		} else {
			res = append([]byte{s[p] + 48}, res...)
			p--
		}

	}
	return string(res)
}
```