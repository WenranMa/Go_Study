# 383. 赎金信

### 简单

给你两个字符串：ransomNote 和 magazine ，判断 ransomNote 能不能由 magazine 里面的字符构成。

如果可以，返回 true ；否则返回 false 。

magazine 中的每个字符只能在 ransomNote 中使用一次。

### 示例 1：

输入：ransomNote = "a", magazine = "b"

输出：false

### 示例 2：

输入：ransomNote = "aa", magazine = "ab"

输出：false

### 示例 3：

输入：ransomNote = "aa", magazine = "aab"

输出：true
 
### 提示：

1 <= ransomNote.length, magazine.length <= 10^5

ransomNote 和 magazine 由小写英文字母组成

### 解：

map. +1 -1操作。如果map中的数小于0，则false.

```go
func canConstruct(ransomNote string, magazine string) bool {
	m := make(map[rune]int)
	for _, c := range magazine {
		m[c] += 1
	}
	for _, c := range ransomNote {
		m[c] -= 1
		if m[c] < 0 {
			return false
		}
	}
	return true
}
```