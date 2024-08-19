# 1047_删除字符串中相邻的重复项_Remove_All_Adjacent_Duplicates_In_String
Given a string S of lowercase letters, a duplicate removal consists of choosing two adjacent and equal letters, and removing them.

We repeatedly make duplicate removals on S until we no longer can.

Return the final string after all such duplicate removals have been made.  It is guaranteed the answer is unique.

Example 1:

    Input: "abbaca"
    Output: "ca"
    Explanation:
    For example, in "abbaca" we could remove "bb" since the letters are adjacent and equal, and this is the only possible move.  The result of this move is that the string is "aaca", of which only "aa" is possible, so the final string is "ca".

Note:

    1 <= S.length <= 20000
    S consists only of English lowercase letters.

### 解：

```go
func removeDuplicates(S string) string {
    stack := []rune{}
    for _, c := range S {
        l := len(stack)
        if l > 0 && stack[l-1] == c {
            stack = stack[:l-1]
        } else {
            stack = append(stack, c)
        }
    }
    return string(stack)
}

/*
func removeDuplicates(S string) string {
    b := Stack{}
    for _, c := range S {
        if b.peek() == c {
            b.pop()
        } else if b.peek() != c {
            b.push(c)
        }
    }
    return string(b.bs)
}

type Stack struct {
    bs []rune
}

func (s *Stack) push(char rune) {
    s.bs = append(s.bs, char)
}

func (s *Stack) pop() rune {
    l := len(s.bs)
    if l > 0 {
        char := s.bs[l-1]
        s.bs = s.bs[:l-1]
        return char
    }
    return 0
}

func (s *Stack) peek() rune {
    l := len(s.bs)
    if l > 0 {
        return s.bs[l-1]
    }
    return 0
}
*/
```