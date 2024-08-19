# 387_字符串第一个唯一字符_First_Unique_Character_in_a_String
Given a string, find the first non-repeating character in it and return it's index. If it doesn't exist, return -1.

Examples:

    s = "leetcode"
    return 0.

    s = "loveleetcode",
    return 2.

### 解：

```go
func firstUniqChar(s string) int {
    m := make(map[rune]int)
    for _, c := range s {
        m[c] += 1
    }
    for i, c := range s {
        if m[c] == 1 {
            return i
        }
    }
    return -1
}

// Use array instead of map, Faster!
func firstUniqChar(s string) int {
    m := [26]int{}
    for _, c := range s {
        m[c-'a'] += 1
    }
    for i, c := range s {
        if m[c-'a'] == 1 {
            return i
        }
    }
    return -1
}
```