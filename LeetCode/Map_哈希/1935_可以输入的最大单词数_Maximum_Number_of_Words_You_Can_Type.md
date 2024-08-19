# 1935. 可以输入的最大单词数 Maximum Number of Words You Can Type

### Easy

There is a malfunctioning keyboard where some letter keys do not work. All other keys on the keyboard work properly.

Given a string text of words separated by a single space (no leading or trailing spaces) and a string brokenLetters of all distinct letter keys that are broken, return the number of words in text you can fully type using this keyboard.

### Example 1:

Input: text = "hello world", brokenLetters = "ad"
Output: 1
Explanation: We cannot type "world" because the 'd' key is broken.

### Example 2:

Input: text = "leet code", brokenLetters = "lt"
Output: 1
Explanation: We cannot type "leet" because the 'l' and 't' keys are broken.

### Example 3:

Input: text = "leet code", brokenLetters = "e"
Output: 0
Explanation: We cannot type either word because the 'e' key is broken.

Constraints:

1 <= text.length <= 104
0 <= brokenLetters.length <= 26
text consists of words separated by a single space without any leading or trailing spaces.
Each word only consists of lowercase English letters.
brokenLetters consists of distinct lowercase English letters.

### 解：

```go
// O(n) time and O(1) space. 
// space + 1 is the word count. res is the broken word.
func canBeTypedWords(text string, brokenLetters string) int {
	res := 0
	space := 0
	t := 0 // flag, true or false
	bs := make([]int, 26)
	for _, c := range brokenLetters {
		bs[c-'a'] = 1
	}
	for _, c := range text {
		if c == ' ' {
			if t == 1 {
				res += 1
				t = 0
			}
			space += 1
		} else if bs[c-'a'] == 1 {
			t = 1
		}

	}
	if t == 1 {
		res += 1
	}
	return space + 1 - res
}
```

map

```go
func canBeTypedWords(text string, brokenLetters string) int {
	words := strings.Split(text, " ")
	brokenMap := make(map[rune]bool)
	for _, l := range brokenLetters {
		brokenMap[l] = true
	}
	res := len(words)
	for _, word := range words {
		for _, c := range word { 
			if brokenMap[c] {
				res -= 1
				break
			}
		}
	}
	return res
}
```
