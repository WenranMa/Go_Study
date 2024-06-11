# 76. 最小覆盖子串

### 困难

给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。

注意：

- 对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
- 如果 s 中存在这样的子串，我们保证它是唯一的答案。
 
### 示例 1：

    输入：s = "ADOBECODEBANC", t = "ABC"
    输出："BANC"
    解释：最小覆盖子串 "BANC" 包含来自字符串 t 的 'A'、'B' 和 'C'。

### 示例 2：

    输入：s = "a", t = "a"
    输出："a"
    解释：整个字符串 s 是最小覆盖子串。

### 示例 3:

    输入: s = "a", t = "aa"
    输出: ""
    解释: t 中两个字符 'a' 均应包含在 s 的子串中，
    因此没有符合条件的子字符串，返回空字符串。

### 解：
两个指针，移动 r 用于判断是否找完了，并且用一个数组index记录每次找到的字符的下标。移动l用于减小长度（有的字符可能有多个）；

两个map. 一个存要找的字符和次数，另一存已经找到次数。found用于记录找到的长度，如果大于等于tl，则说明找到一个窗口。然后移动 l 减小长度。程序中的 l 不是一个指针，而已一个index下标，这样减少时可以一次跳过多个。  index[l] 和 r 构成window.

```go
func minWindow(s string, t string) string {
	sl, tl := len(s), len(t)
	hasFound := make(map[byte]int)
	needFind := make(map[byte]int)
	for i := 0; i < tl; i++ {
		c := t[i]
		hasFound[c] = 0
		needFind[c] += 1
	}
	index := []int{}
	l, r := 0, 0
	found, winSize := 0, sl
	var res string
	for ; r < sl; r += 1 {
		c := s[r]
		if _, ok := needFind[c]; !ok {
			continue
		}
		index = append(index, r)

		hasFound[c] += 1
		if hasFound[c] <= needFind[c] {
			found += 1
		}
		if found >= tl {
			b := s[index[l]]
			for hasFound[b] > needFind[b] {
				hasFound[b] -= 1
				l += 1
				b = s[index[l]]
			}
			if r-index[l]+1 <= winSize {
				winSize = r - index[l] + 1
				res = s[index[l] : r+1]
			}
		}
		//fmt.Println(string(c), r, l, index, "found: ", found, "windowsize: ", winSize)
	}
	return res
}
```