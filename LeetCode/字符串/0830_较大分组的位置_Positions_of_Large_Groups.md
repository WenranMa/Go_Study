# 830. 较大分组的位置 Positions of Large Groups

### Easy

In a string s of lowercase letters, these letters form consecutive groups of the same character.

For example, a string like s = "abbxxxxzyy" has the groups "a", "bb", "xxxx", "z", and "yy".

A group is identified by an interval [start, end], where start and end denote the start and end indices (inclusive) of the group. In the above example, "xxxx" has the interval [3,6].

A group is considered large if it has 3 or more characters.

Return the intervals of every large group sorted in increasing order by start index.

### Example 1:

Input: s = "abbxxxxzzy"
Output: [[3,6]]
Explanation: "xxxx" is the only large group with start index 3 and end index 6.

### Example 2:

Input: s = "abc"
Output: []
Explanation: We have groups "a", "b", and "c", none of which are large groups.

### Example 3:

Input: s = "abcdddeeeeaabbbcd"
Output: [[3,5],[6,9],[12,14]]
Explanation: The large groups are "ddd", "eeee", and "bbb".

Constraints:

1 <= s.length <= 1000
s contains lowercase English letters only.

### 解：

双指针，遍历字符串，如果不等，更新start, 更新结果

```go
// 直接用end - start即可~
func largeGroupPositions(s string) [][]int {
	res := [][]int{}
	l := len(s)
	if l <= 1 {
		return res
	}
	start := 0
	end := 1
	for end < l {
		if s[start] != s[end] {
			if end-start >= 3 {
				res = append(res, []int{start, end - 1})
			}
			start = end
		}
		end += 1
	}
	if end-start >= 3 {
		res = append(res, []int{start, l - 1})
	}
	return res
}
```