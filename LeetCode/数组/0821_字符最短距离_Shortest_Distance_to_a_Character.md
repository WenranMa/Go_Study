# 821_字符最短距离_Shortest_Distance_to_a_Character
Given a string S and a character C, return an array of integers representing the shortest distance from the character C in the string.

### Example:

    Input: S = "loveleetcode", C = 'e'   
    Output: [3, 2, 1, 0, 1, 0, 0, 1, 2, 2, 1, 0]   

Note:  
S string length is in [1, 10000].   
C is a single character, and guaranteed to be in string S.   
All letters in S and C are lowercase.  

### 解：

两次正反向遍历。

第一次遍历从左到右，可以计算出每个字符距离他左边第一个c的距离，不一定是最小的。

第二次遍历从右到左，可以计算出每个字符距离他右边第一个c的距离，和上一次的距离比较，取最小值。

比如这个例子：
第一次遍历：[12 13 14 0 1 0 0 1 2 3 4 0]
第二次遍历：[3 2 1 0 1 0 0 1 2 2 1 0]

```go
func shortestToChar(s string, c byte) []int {
	l := len(s)
	dis := make([]int, l)
	idx := -l
	for i := range s {
		if s[i] == c {
			idx = i
		}
		dis[i] = i - idx
	}
	idx = 2 * l
	for i := l - 1; i >= 0; i-- {
		if s[i] == c {
			idx = i
		}
		dis[i] = min(dis[i], idx-i)
	}
	return dis
}
```

老解法
```go
func shortestToChar(S string, C byte) []int {
    l := len(S)
    ans := []int{}
    pos := []int{}
    for i, c := range S {
        if c == rune(C) {
            pos = append(pos, i)
        }
    }
    for i, _ := range S {
        min := l
        for _, p := range pos {
            d := 0
            if i < p {
                d = p - i
            } else {
                d = i - p
            }
            if d < min {
                min = d
            }
        }
        ans = append(ans, min)
    }
    return ans
}
```