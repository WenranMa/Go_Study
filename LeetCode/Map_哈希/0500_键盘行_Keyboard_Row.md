# 500_键盘行_Keyboard_Row

Given a List of words, return the words that can be typed using letters of alphabet on only one row's of American keyboard like the image below.

Example:  
Input: ["Hello", "Alaska", "Dad", "Peace"]  
Output: ["Alaska", "Dad"]  
 
Note:  
You may use one character in the keyboard more than once.  
You may assume the input string will only contain letters of alphabet.

### 解：

```go
func findWords(words []string) []string {
    m := make(map[rune]int)

    m['q'] = 1
    m['w'] = 1
    m['e'] = 1
    m['r'] = 1
    m['t'] = 1
    m['y'] = 1
    m['u'] = 1
    m['i'] = 1
    m['o'] = 1
    m['p'] = 1

    m['a'] = 2
    m['s'] = 2
    m['d'] = 2
    m['f'] = 2
    m['g'] = 2
    m['h'] = 2
    m['j'] = 2
    m['k'] = 2
    m['l'] = 2

    m['z'] = 3
    m['x'] = 3
    m['c'] = 3
    m['v'] = 3
    m['b'] = 3
    m['n'] = 3
    m['m'] = 3

    ans := []string{}
    for _, w := range words {
        wl := strings.ToLower(w)
        flag := true
        n := m[rune(wl[0])]
        for _, l := range wl {
            if m[l] != n {
                flag = false
                break
            }
        }
        if flag {
            ans = append(ans, w)
        }
    }
    return ans
}
```