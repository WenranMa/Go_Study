# 921_使括号有效的最小添加_Minimum_Add_to_Make_Parentheses_Valid

Given a string S of '(' and ')' parentheses, we add the minimum number of parentheses ( '(' or ')', and in any positions ) so that the resulting parentheses string is valid.

Formally, a parentheses string is valid if and only if:

- It is the empty string, or
- It can be written as AB (A concatenated with B), where A and B are valid strings, or
- It can be written as (A), where A is a valid string.

Given a parentheses string, return the minimum number of parentheses we must add to make the resulting string valid.

Example:

    Input: "())"
    Output: 1

    Input: "((("
    Output: 3

    Input: "()"
    Output: 0

    Input: "()))(("
    Output: 4

Note:

    S.length <= 1000
    S only consists of '(' and ')' characters.

### 解：

```go
// Double counter, simulate stack, O(1) space, O(N) time.
func minAddToMakeValid(S string) int {
    c1 := 0
    c2 := 0
    for _, c := range S {
        if c == '(' {
            c1 += 1
        } else {
            if c1 == 0 {
                c2 += 1
            } else if c1 > 0 {
                c1 -= 1
            }
        }
    }
    return c1 + c2
}
```