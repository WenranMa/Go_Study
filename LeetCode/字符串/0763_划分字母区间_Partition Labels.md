# 763_划分字母区间_Partition Labels

A string S of lowercase letters is given. We want to partition this string into as many parts as possible so that each letter appears in at most one part, and return a list of integers representing the size of these parts.

Example:

    Input: S = "ababcbacadefegdehijhklij"
    Output: [9,7,8]
    Explanation:
    The partition is "ababcbaca", "defegde", "hijhklij".
    This is a partition so that each letter appears in at most one part.
    A partition like "ababcbacadefegde", "hijhklij" is incorrect, because it splits S into less parts.

Note:

    S will have length in range [1, 500].
    S will consist of lowercase letters ('a' to 'z') only.

### 解：

map记录字符的最后出现的位置，然后遍历字符串，如果当前字符的位置等于tail，说明当前字符是最后一个。

比如 abababab 这种情况，只有再最后一个b时，i == tail, 将tail-head+1加入结果中，head更新为tail+1。

```go
//Better solution O(n) time, O(n) space.
func partitionLabels(S string) []int {
    var res []int
    var tail, head int 
    lastIndex := make(map[rune]int) // map[int32]int is also ok.
    for i, c := range S {
        lastIndex[c] = i
    }
    for i, c := range S {
        tail = max(tail, lastIndex[c])
        if i == tail {
            res = append(res, tail-head+1)
            head = tail + 1
        }
    }
    return res
}
```

先不管。。
```go
func partitionLabels(S string) []int {
    l := len(S)
    res := []int{}
    tail := 0
    for head := 0; head < l; head = tail + 1 {
        tail = head
        for j := head; j <= tail; j++ {
            for k := l - 1; k > tail; k-- {
                if S[j] == S[k] && k > tail {
                    tail = k
                }
            }
        }
        res = append(res, tail-head+1)
    }
    return res
}
//O(n^3) ?
```