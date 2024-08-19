# 1189. 气球的最大数量 Maximum Number of Balloons

### Easy

Given a string text, you want to use the characters of text to form as many instances of the word "balloon" as possible.

You can use each character in text at most once. Return the maximum number of instances that can be formed.

### Example 1:

Input: text = "nlaebolko"
Output: 1

### Example 2:

Input: text = "loonbalxballpoon"
Output: 2

### Example 3:

Input: text = "leetcode"
Output: 0

Constraints:

1 <= text.length <= 10^4
text consists of lower case English letters only.

```go
// O(n) time. O(1) space. Map.
// 初始化makeM()是为了这种'lloo'的case, 如果没有初始值，lloo这种情况for range 结果是1.
func maxNumberOfBalloons(text string) int {
	m := makeM()
	for i, _ := range text {
		if isB(text[i]) {
			m[text[i]] += 1
		}
	}
	m['o'] /= 2
	m['l'] /= 2
	min := len(text)
	for _, v := range m {
		if min > v {
			min = v
		}
	}
	return min
}

func makeM() map[byte]int { 
	m := make(map[byte]int)
	m['b'] = 0
	m['a'] = 0
	m['l'] = 0
	m['o'] = 0
	m['n'] = 0
	return m
}

func isB(c byte) bool {
	return c == 'b' || c == 'a' || c == 'l' || c == 'o' || c == 'n'
}
```