# 1021_删除最外层括号_Remove Outermost Parentheses

A valid parentheses string is either empty (""), "(" + A + ")", or A + B, where A and B are valid parentheses strings, and + represents string concatenation.  For example, "", "()", "(())()", and "(()(()))" are all valid parentheses strings.

A valid parentheses string S is primitive if it is nonempty, and there does not exist a way to split it into S = A+B, with A and B nonempty valid parentheses strings.

Given a valid parentheses string S, consider its primitive decomposition: S = P_1 + P_2 + ... + P_k, where P_i are primitive valid parentheses strings.

Return S after removing the outermost parentheses of every primitive string in the primitive decomposition of S.

### Example:

    Input: "(()())(())"
    Output: "()()()"
    Explanation:
    The input string is "(()())(())", with primitive decomposition "(()())" + "(())".
    After removing outer parentheses of each part, this is "()()" + "()" = "()()()".

    Input: "(()())(())(()(()))"
    Output: "()()()()(())"
    Explanation:
    The input string is "(()())(())(()(()))", with primitive decomposition "(()())" + "(())" + "(()(()))".
    After removing outer parentheses of each part, this is "()()" + "()" + "()(())" = "()()()()(())".

    Input: "()()"
    Output: ""
    Explanation:
    The input string is "()()", with primitive decomposition "()" + "()".
    After removing outer parentheses of each part, this is "" + "" = "".

### Note:

    S.length <= 10000
    S[i] is "(" or ")"
    S is a valid parentheses string

### 解：

```go
func removeOuterParentheses(S string) string {
    counter := 0
    head := 1
    res := ""    // 或者 var res strings.Builder
    for i, c := range S {
        if c == '(' {
            counter++
        } else {
            counter--
        }
        if counter == 0 {
            res += S[head:i]  // res.WriteString(S[head:i])
            head = i + 2
        }
    }
    return res  // res.String()
}
```

也可以用栈
栈不为空，就是剩下的字符串
```go
func removeOuterParentheses(s string) string {
    var ans, st []rune
    for _, c := range s {
        if c == ')' {
            st = st[:len(st)-1]
        }
        if len(st) > 0 {
            ans = append(ans, c)
        }
        if c == '(' {
            st = append(st, c)
        }
    }
    return string(ans)
}
```