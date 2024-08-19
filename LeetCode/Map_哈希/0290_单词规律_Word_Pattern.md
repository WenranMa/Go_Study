# 290. 单词规律

### 简单

给定一种规律 pattern 和一个字符串 s ，判断 s 是否遵循相同的规律。

这里的 遵循 指完全匹配，例如， pattern 里的每个字母和字符串 s 中的每个非空单词之间存在着双向连接的对应规律。

### 示例1:

输入: pattern = "abba", s = "dog cat cat dog"

输出: true

### 示例 2:

输入:pattern = "abba", s = "dog cat cat fish"

输出: false

### 示例 3:

输入: pattern = "aaaa", s = "dog cat cat dog"

输出: false

### 提示:

- 1 <= pattern.length <= 300
- pattern 只包含小写英文字母
- 1 <= s.length <= 3000
- s 只包含小写英文字母和 ' '
- s 不包含 任何前导或尾随对空格
- s 中每个单词都被 单个空格 分隔

### 解：

leetcode 205 题一样。必须是双向的map. 

```go
func wordPattern(pattern string, s string) bool {
	mp := make(map[byte]string)
	ms := make(map[string]byte)
	strs := strings.Split(s, " ")
	if len(pattern) != len(strs) {
		return false
	}
	for i := range pattern {
		if str, ok := mp[pattern[i]]; ok && str != strs[i] {
			return false
		}
		if p, ok := ms[strs[i]]; ok && p != pattern[i] {
			return false
		}
		mp[pattern[i]] = strs[i]
		ms[strs[i]] = pattern[i]
	}
	return true
}
```