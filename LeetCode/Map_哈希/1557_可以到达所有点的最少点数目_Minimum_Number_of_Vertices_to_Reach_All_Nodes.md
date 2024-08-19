# 1557. 可以到达所有点的最少点数目 Minimum Number of Vertices to Reach All Nodes

### Medium

Given a directed acyclic graph, with n vertices numbered from 0 to n-1, and an array edges where edges[i] = [fromi, toi] represents a directed edge from node fromi to node toi.

Find the smallest set of vertices from which all nodes in the graph are reachable. It's guaranteed that a unique solution exists.

Notice that you can return the vertices in any order.

### Example 1:

Input: n = 6, edges = [[0,1],[0,2],[2,5],[3,4],[4,2]]
Output: [0,3]
Explanation: It's not possible to reach all the nodes from a single vertex. From 0 we can reach [0,1,2,5]. From 3 we can reach [3,4,2,5]. So we output [0,3].

### Example 2:

Input: n = 5, edges = [[0,1],[2,1],[3,1],[1,4],[2,4]]
Output: [0,2,3]
Explanation: Notice that vertices 0, 3 and 2 are not reachable from any other node, so we must include them. Also any of these vertices can reach nodes 1 and 4.
 
Constraints:

2 <= n <= 10^5
1 <= edges.length <= min(10^5, n * (n - 1) / 2)
edges[i].length == 2
0 <= fromi, toi < n
All pairs (fromi, toi) are distinct.

### 解：
leetcode 1436, 997

类似map, 但用了数组，因为确定数字是0 ~ n-1

```go
// O(n) time. O(n) space.
// 找到起点即可，起点就是只出现在数组第一位的数字。
// 用位运算标记，1起点，2终点，3即是起点也是终点。
func findSmallestSetOfVertices(n int, edges [][]int) []int {
	m := make([]int, n)
	res := []int{}
	for _, e := range edges {
		m[e[0]] |= 1
		m[e[1]] |= 2
	}
	for k, v := range m {
		if v == 1 {
			res = append(res, k)
		}
	}
	return res
}

// 改进，没必要记录起点元素。不在终点的，就是起点。
func findSmallestSetOfVertices(n int, edges [][]int) []int {
	m := make([]int, n)
	res := []int{}
	for _, e := range edges {
		m[e[1]] = 1
	}
	for k, v := range m {
		if v == 0 {
			res = append(res, k)
		}
	}
	return res
}
```