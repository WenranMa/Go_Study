# 541. 反转字符串II Reverse String II

### Easy

Given a string s and an integer k, reverse the first k characters for every 2k characters counting from the start of the string.

If there are fewer than k characters left, reverse all of them. If there are less than 2k but greater than or equal to k characters, then reverse the first k characters and leave the other as original.

### Example 1:

Input: s = "abcdefg", k = 2
Output: "bacdfeg"

### Example 2:

Input: s = "abcd", k = 2
Output: "bacd"

Constraints:

1 <= s.length <= 10^4
s consists of only lowercase English letters.
1 <= k <= 10^4

### 解：

```go
// straight forward. 
// O(n) time. O(n) space.
// 注意细节，如果最后的长度小于K，也要reverse.
func reverseStr(s string, k int) string {
    l := len(s)
    sb := []byte(s)
    for i := 0; i < l; i += 2 * k {
        if i+k < l {
            sb = reverse(sb, i, i+k-1)
        } else {
            sb = reverse(sb, i, l-1)
        }
    }
    return string(sb)
}
func reverse(s []byte, start, end int) []byte {
    for i, j := start, end; i <= j; i, j = i+1, j-1 {
        s[i], s[j] = s[j], s[i]
    }
    return s
}

```