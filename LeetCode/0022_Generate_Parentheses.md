# 22. 括号生成

### 中等

数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。

### 示例 1：

    输入：n = 3
    输出：["((()))","(()())","(())()","()(())","()()()"]

### 示例 2：

    输入：n = 1
    输出：["()"]
  
### 提示：

1 <= n <= 8

### 解：

```go
var res []string

func generateParenthesis(n int) []string {
	res = []string{}
	str := make([]byte, 2*n)
	generate(n, n, 0, str)
	return res
}

func generate(l, r, i int, str []byte) {
	if l < 0 || r < 0 {
		return
	}
	if l == 0 && r == 0 {
		res = append(res, string(str))
	} else {
		if l > 0 {
			str[i] = '('
			generate(l-1, r, i+1, str)
		}
		if r > l {
			str[i] = ')'
			generate(l, r-1, i+1, str)
		}
	}
}
```