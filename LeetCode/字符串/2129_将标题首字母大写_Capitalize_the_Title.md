# 2129. 将标题首字母大写 Capitalize the Title

### Easy

You are given a string title consisting of one or more words separated by a single space, where each word consists of English letters. Capitalize the string by changing the capitalization of each word such that:

If the length of the word is 1 or 2 letters, change all letters to lowercase.
Otherwise, change the first letter to uppercase and the remaining letters to lowercase.
Return the capitalized title.

### Example 1:

Input: title = "capiTalIze tHe titLe"
Output: "Capitalize The Title"
Explanation:
Since all the words have a length of at least 3, the first letter of each word is uppercase, and the remaining letters are lowercase.

### Example 2:

Input: title = "First leTTeR of EACH Word"
Output: "First Letter of Each Word"
Explanation:
The word "of" has length 2, so it is all lowercase.
The remaining words have a length of at least 3, so the first letter of each remaining word is uppercase, and the remaining letters are lowercase.

### Example 3:

Input: title = "i lOve leetcode"
Output: "i Love Leetcode"
Explanation:
The word "i" has length 1, so it is lowercase.
The remaining words have a length of at least 3, so the first letter of each remaining word is uppercase, and the remaining letters are lowercase.
 
Constraints:

1 <= title.length <= 100
title consists of words separated by a single space without any leading or trailing spaces.
Each word consists of uppercase and lowercase English letters and is non-empty.

### 解：

j遍历字符串，遇到空格时，判断当前单词的长度，如果大于2 (j-i > 2)，且第一个字母是小写的，则将第一个字母转换为大写。
i指向每个单词的首字母。

```go
// O(n) time, O(n) space. two pointers.
func capitalizeTitle(title string) string {
	l := len(title)
	res := []byte(title)
	i, j := 0, 0
	for j < l {
		if res[j] == ' ' {
			if j-i > 2 && res[i] >= 97 {
				res[i] -= 32
			}
			i = j + 1
		} else if res[j] <= 90 && res[j] >= 65 {
			res[j] += 32
		}
		j++
	}
	if j-i > 2 && res[i] >= 97 {
		res[i] -= 32
	}
	return string(res)
}
```