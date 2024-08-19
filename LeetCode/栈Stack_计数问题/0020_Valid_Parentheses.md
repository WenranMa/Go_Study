# 20. 有效的括号

### 简单

给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
每个右括号都有一个对应的相同类型的左括号。
 
### 示例 1：

输入：s = "()"

输出：true

### 示例 2：

输入：s = "()[]{}"

输出：true

### 示例 3：

输入：s = "(]"

输出：false
 
### 提示：

1 <= s.length <= 104

s 仅由括号 '()[]{}' 组成

### 解：

stack, 不能用三个数+1或-1的方式进行，比如这种"([)]"也是错的。

```go
func isValid(s string) bool {
	stack := []rune{}
	for _, c := range s {
		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, c)
		} else {
			if len(stack) == 0 {
				return false
			}
			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if c == ')' && last != '(' {
				return false
			}
			if c == ']' && last != '[' {
				return false
			}
			if c == '}' && last != '{' {
				return false
			}
		}
	}
	return len(stack) == 0
}
```