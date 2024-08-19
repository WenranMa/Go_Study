# 557_反转字符串中的单词_III_Reverse Words in a String III
Given a string, you need to reverse the order of characters in each word within a sentence while still preserving whitespace and initial word order.

Example: 

    Input: "Let's take LeetCode contest"  
    Output: "s'teL ekat edoCteeL tsetnoc"

Note: In the string, each word is separated by single space and there will not be any extra space in the string.

### 解：

```go
func reverseWords(s string) string {
	arr := strings.Split(s, " ")
	ans := []string{}
	for _, w := range arr {
		chars := []byte(w)
		for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
			chars[i], chars[j] = chars[j], chars[i]
		}
		ans = append(ans, string(chars))
	}
	return strings.Join(ans, " ")
}
```