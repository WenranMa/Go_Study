# 997. 找到小镇的法官 Find the Town Judge

### Easy

In a town, there are n people labeled from 1 to n. There is a rumor that one of these people is secretly the town judge.

If the town judge exists, then:

The town judge trusts nobody.
Everybody (except for the town judge) trusts the town judge.
There is exactly one person that satisfies properties 1 and 2.
You are given an array trust where trust[i] = [ai, bi] representing that the person labeled ai trusts the person labeled bi.

Return the label of the town judge if the town judge exists and can be identified, or return -1 otherwise.

### Example 1:

Input: n = 2, trust = [[1,2]]
Output: 2

### Example 2:

Input: n = 3, trust = [[1,3],[2,3]]
Output: 3

### Example 3:

Input: n = 3, trust = [[1,3],[2,3],[3,1]]
Output: -1

Constraints:

1 <= n <= 1000
0 <= trust.length <= 104
trust[i].length == 2
All the pairs of trust are unique.
ai != bi
1 <= ai, bi <= n

### 解：

leetcode 1557, 1436, 997

```go 
// O(n) time, O(n) space. 
// 信任别人，则自己减1. 
// 别人相信自己，自己加1.
func findJudge(n int, trust [][]int) int {
	if n == 1 {
		return 1
	}
	m := make(map[int]int)
	for _, t := range trust {
		m[t[0]] -= 1
		m[t[1]] += 1
	}
	for i := 1; i <= n; i++ {
		if v, ok := m[i]; ok {
			if v == n-1 {
				return i
			}
		}
	}
	return -1
}
```