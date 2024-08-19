# 2138. 将字符串拆分为若干长度为k的组 Divide a String Into Groups of Size k

### Easy

A string s can be partitioned into groups of size k using the following procedure:

The first group consists of the first k characters of the string, the second group consists of the next k characters of the string, and so on. Each character can be a part of exactly one group.
For the last group, if the string does not have k characters remaining, a character fill is used to complete the group.
Note that the partition is done so that after removing the fill character from the last group (if it exists) and concatenating all the groups in order, the resultant string should be s.

Given the string s, the size of each group k and the character fill, return a string array denoting the composition of every group s has been divided into, using the above procedure.

### Example 1:

	Input: s = "abcdefghi", k = 3, fill = "x"
	Output: ["abc","def","ghi"]
	Explanation:
	The first 3 characters "abc" form the first group.
	The next 3 characters "def" form the second group.
	The last 3 characters "ghi" form the third group.
	Since all groups can be completely filled by characters from the string, we do not need to use fill.
	Thus, the groups formed are "abc", "def", and "ghi".

### Example 2:

	Input: s = "abcdefghij", k = 3, fill = "x"
	Output: ["abc","def","ghi","jxx"]
	Explanation:
	Similar to the previous example, we are forming the first three groups "abc", "def", and "ghi".
	For the last group, we can only use the character 'j' from the string. To complete this group, we add 'x' twice.
	Thus, the 4 groups formed are "abc", "def", "ghi", and "jxx".

Constraints:

	1 <= s.length <= 100
	s consists of lowercase English letters only.
	1 <= k <= 100
	fill is a lowercase English letter.

### 解：

先补齐缺少的字符，然后按照k个字符一组，返回结果。
l = 10, k = 3
10 % 3 = 1
要补 k - 1 = 2个

```go
// O(n) straight forward.
func divideString(s string, k int, fill byte) []string {
	res := []string{}
	l := len(s)
	n := l % k
	post := make([]byte, k-n)
	for i, _ := range post {
		post[i] = fill
	}
	s += string(post)
	for i := 0; i < l; i += k {
		res = append(res, s[i:i+k])
	}
	return res
}
```