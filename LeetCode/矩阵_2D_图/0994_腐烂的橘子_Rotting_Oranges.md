# 994. 腐烂的橘子

### 中等

在给定的 m x n 网格 grid 中，每个单元格可以有以下三个值之一：

- 值 0 代表空单元格；
- 值 1 代表新鲜橘子；
- 值 2 代表腐烂的橘子。

每分钟，腐烂的橘子 周围 4 个方向上相邻 的新鲜橘子都会腐烂。

返回 直到单元格中没有新鲜橘子为止所必须经过的最小分钟数。如果不可能，返回 -1 。

### 示例 1：

输入：grid = [[2,1,1],[1,1,0],[0,1,1]]
输出：4

### 示例 2：

输入：grid = [[2,1,1],[0,1,1],[1,0,1]]
输出：-1
解释：左下角的橘子（第 2 行， 第 0 列）永远不会腐烂，因为腐烂只会发生在 4 个方向上。

### 示例 3：

输入：grid = [[0,2]]
输出：0
解释：因为 0 分钟时已经没有新鲜橘子了，所以答案就是 0 。

### 提示：

m == grid.length
n == grid[i].length
1 <= m, n <= 10
grid[i][j] 仅为 0、1 或 2

### 解：

```go
func orangesRotting(grid [][]int) int {
	// 统计新鲜橙子
	freshNum := 0
	// dfs 统计腐烂橙子
	queue := [][]int{}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				freshNum++
			}
			if grid[i][j] == 2 {
				queue = append(queue, []int{i, j})
			}
		}
	}
	minutes := 0
	for len(queue) > 0 {
		if freshNum == 0 {
			// 没有新鲜橙子了
			return minutes
		}
		// 过去1分钟，周围开始腐烂
		minutes++
		l := len(queue)
		for i := 0; i < l; i++ {
			rot := queue[0]
			queue = queue[1:]
			x, y := rot[0], rot[1]
			freshNum -= roting(grid, x-1, y, &queue)
			freshNum -= roting(grid, x+1, y, &queue)
			freshNum -= roting(grid, x, y-1, &queue)
			freshNum -= roting(grid, x, y+1, &queue)
		}
	}
	// 腐烂过程结束
	if freshNum > 0 {
		return -1
	} else {
		return minutes
	}
}

func roting(grid [][]int, x, y int, queue *[][]int) int {
	if x < 0 || y < 0 || x > len(grid)-1 || y > len(grid[0])-1 || grid[x][y] != 1 {
		return 0
	}
	grid[x][y] = 2
	*queue = append(*queue, []int{x, y})
	return 1
}
```