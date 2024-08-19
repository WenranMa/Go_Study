# 890_查找和替换模式_Find_and_Replace_Pattern
You have a list of words and a pattern, and you want to know which words in words matches the pattern.

A word matches the pattern if there exists a permutation of letters p so that after replacing every letter x in the pattern with p(x), we get the desired word.

(Recall that a permutation of letters is a bijection from letters to letters: every letter maps to another letter, and no two letters map to the same letter.)

Return a list of the words in words that match the given pattern.

You may return the answer in any order.

Example:

    Input: words = ["abc","deq","mee","aqq","dkd","ccc"], pattern = "abb"
    Output: ["mee","aqq"]
    Explanation: "mee" matches the pattern because there is a permutation {a -> m, b -> e, ...}.
    "ccc" does not match the pattern because {a -> c, b -> c, ...} is not a permutation,
    since a and b map to the same letter.

Note:

    1 <= words.length <= 50
    1 <= pattern.length = words[i].length <= 20

### 解：

```go
//Two maps
// O(N*K) time, space
func findAndReplacePattern(words []string, pattern string) []string {
    res := []string{}
    for _, s := range words {
        if check(s, pattern) {
            res = append(res, s)
        }
    }
    return res
}

func check(s, pattern string) bool {
    mptos := make(map[byte]byte)
    mstop := make(map[byte]byte)
    l := len(s)
    for i := 0; i < l; i++ {
        if _, ok := mptos[pattern[i]]; !ok {
            mptos[pattern[i]] = s[i]
        }
        if _, ok := mstop[s[i]]; !ok {
            mstop[s[i]] = pattern[i]
        }
        if mptos[pattern[i]] != s[i] || mstop[s[i]] != pattern[i] {
            return false
        }
    }
    return true
}
```