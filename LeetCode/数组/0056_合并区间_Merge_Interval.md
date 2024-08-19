# 56. 合并区间

### 中等

以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。

### 示例 1：

输入：intervals = [[1,3],[2,6],[8,10],[15,18]]

输出：[[1,6],[8,10],[15,18]]

解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].

### 示例 2：

输入：intervals = [[1,4],[4,5]]

输出：[[1,5]]

解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。
 
### 提示：
1 <= intervals.length <= 104

intervals[i].length == 2

0 <= starti <= endi <= 104

### 解：

比insert interval 多了一步排序。按start index排序。

```go
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	return insert(intervals[1:len(intervals)], intervals[0])
}

func insert(intervals [][]int, newInterval []int) [][]int {
	res := [][]int{}
	for _, in := range intervals {
		if in[1] < newInterval[0] {
			res = append(res, in)
		} else if newInterval[1] < in[0] {
			res = append(res, newInterval)
			newInterval = in
		} else {
			newInterval = []int{min(newInterval[0], in[0]), max(newInterval[1], in[1])}
		}
	}
	res = append(res, newInterval)
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```