# 49. 字母异位词分组

### 中等

给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。

字母异位词 是由重新排列源单词的所有字母得到的一个新单词。

### 示例 1:

输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
输出: [["bat"],["nat","tan"],["ate","eat","tea"]]

### 示例 2:

输入: strs = [""]
输出: [[""]]

### 示例 3:

输入: strs = ["a"]
输出: [["a"]]

### 提示：
- 1 <= strs.length <= 104
- 0 <= strs[i].length <= 100
- strs[i] 仅包含小写字母

### 解：

用一个26长度的数组作为key. 如果是anagram, 数组的元素是相等的。

```go
func groupAnagrams(strs []string) [][]string {
    m:= make(map[[26]int][]string)
    for _, str:= range strs {
        cnt := [26]int{}
        for _, b:= range str {
            cnt[b - 'a'] += 1
        }
        m[cnt] = append(m[cnt], str)
    }
    res:= [][]string{}
    for _, v:= range m {
        res = append(res, v)
    }
    return res
}
```