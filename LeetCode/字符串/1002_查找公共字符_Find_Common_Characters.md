# 1002_查找公共字符_Find_Common_Characters
Given an array A of strings made only from lowercase letters, return a list of all characters that show up in all strings within the list (including duplicates).  For example, if a character occurs 3 times in all strings but not 4 times, you need to include that character three times in the final answer.

You may return the answer in any order.

Example:

    Input: ["bella","label","roller"]
    Output: ["e","l","l"]

    Input: ["cool","lock","cook"]
    Output: ["c","o"]

Note:

    1 <= A.length <= 100
    1 <= A[i].length <= 100
    A[i][j] is a lowercase letter

### 解：

1. 单词中每个字符的个数，用数组`chars`表示，chars每个单词都新建。
2. 如果chars小于count, 更新count.
3. count最后剩下的就是公共的，n某个字符的个数。 

```go
func commonChars(A []string) []string {
    count := [26]int{}
    for i, _ := range count {
        count[i] = math.MaxUint16
    }
    for _, word := range A {
        chars := [26]int{}
        for _, c := range word {
            chars[c-'a'] += 1
        }
        for i := 0; i < 26; i++ {
            if chars[i] < count[i] {
                count[i] = chars[i]
            }
        }
    }
    res := []string{}
    for i, n := range count {
        for j := 1; j <= n; j++ {
            res = append(res, string('a'+i))
        }
    }
    return res
}
```