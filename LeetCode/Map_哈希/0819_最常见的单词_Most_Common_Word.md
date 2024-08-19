# 819. 最常见的单词 Most Common Word

### Easy

Given a string paragraph and a string array of the banned words banned, return the most frequent word that is not banned. It is guaranteed there is at least one word that is not banned, and that the answer is unique.

The words in paragraph are case-insensitive and the answer should be returned in lowercase.

### Example 1:

Input: paragraph = "Bob hit a ball, the hit BALL flew far after it was hit.", banned = ["hit"]
Output: "ball"
Explanation: 
"hit" occurs 3 times, but it is a banned word.
"ball" occurs twice (and no other word does), so it is the most frequent non-banned word in the paragraph. 
Note that words in the paragraph are not case sensitive,
that punctuation is ignored (even if adjacent to words, such as "ball,"), 
and that "hit" isn't the answer even though it occurs more because it is banned.

### Example 2:

Input: paragraph = "a.", banned = []
Output: "a"

Constraints:

1 <= paragraph.length <= 1000
paragraph consists of English letters, space ' ', or one of the symbols: "!?',;.".
0 <= banned.length <= 100
1 <= banned[i].length <= 10
banned[i] consists of only lowercase English letters.

### 解：

```go
// two map.
func mostCommonWord(paragraph string, banned []string) string {
    mb:= make(map[string]int)
    for _, b:= range banned {
        mb[b] = 1
    }
    mw:= make(map[string]int)
    re:= regexp.MustCompile("[!?',;. ]+") // 按符号split
    words:= re.Split(paragraph, -1) // -1表示所有substring
    for _, w:= range words {
        w = strings.ToLower(w)
        if _,ok:= mb[w]; !ok {
            mw[w] += 1
        }
    }
    max:= 0
    res:= ""
    for k, v:= range mw {
        if max < v {
            max = v
            res = k
        }
    }
    return res
}

// one map, 总体思想没变~
func mostCommonWord(paragraph string, banned []string) string {
	m := make(map[string]int)
	for _, b := range banned {
		m[b] = -1
	}
	re := regexp.MustCompile("[!?',;. ]+")
	words := re.Split(paragraph, -1)
	for _, w := range words {
		w = strings.ToLower(w)
		if v, ok := m[w]; ok {
			if v != -1 {
				m[w] += 1
			}
		} else {
			m[w] += 1
		}
	}
	max := 0
	res := ""
	for k, v := range m {
		if max < v {
			max = v
			res = k
		}
	}
	return res
}


```