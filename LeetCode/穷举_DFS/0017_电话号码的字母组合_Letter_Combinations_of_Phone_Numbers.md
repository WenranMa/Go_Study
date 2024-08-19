# 17. 电话号码的字母组合

### 中等

给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。

![key](/file/img/telephone-keypad.png)

### 示例 1：

    输入：digits = "23"
    输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]

### 示例 2：

    输入：digits = ""
    输出：[]

### 示例 3：

    输入：digits = "2"
    输出：["a","b","c"]

### 提示：
- 0 <= digits.length <= 4
- digits[i] 是范围 ['2', '9'] 的一个数字。

### 解：
```go
func letterCombinations(digits string) []string {
	res := []string{}
	if len(digits) == 0 {
		return res
	}
	letters := []string{" ", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	var combine func(digits string, n int, one []byte)
	combine = func(digits string, n int, one []byte) {
		if n == len(digits) {
			res = append(res, string(one))
		} else {
			for i := 0; i < len(letters[digits[n]-'0']); i++ {
				one = append(one, letters[digits[n]-'0'][i])
				combine(digits, n+1, one)
				one = one[:len(one)-1]
			}
		}
	}
	combine(digits, 0, []byte{})
	return res
}
```