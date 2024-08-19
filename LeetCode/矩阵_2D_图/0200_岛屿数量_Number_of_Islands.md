# 200. 岛屿数量

### 中等

给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。

岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。

此外，你可以假设该网格的四条边均被水包围。

### 示例 1：

输入：

    grid = [
    ["1","1","1","1","0"],
    ["1","1","0","1","0"],
    ["1","1","0","0","0"],
    ["0","0","0","0","0"]
    ]

输出：1

### 示例 2：

输入：

    grid = [
    ["1","1","0","0","0"],
    ["1","1","0","0","0"],
    ["0","0","1","0","0"],
    ["0","0","0","1","1"]
    ]

输出：3
 
### 提示：
- m == grid.length
- n == grid[i].length
- 1 <= m, n <= 300
- grid[i][j] 的值为 '0' 或 '1'

### 解：

深度优先搜索

我们可以将二维网格看成一个无向图，竖直或水平相邻的 111 之间有边相连。

为了求出岛屿的数量，我们可以扫描整个二维网格。如果一个位置为 111，则以其为起始节点开始进行深度优先搜索。在深度优先搜索的过程中，每个搜索到的 111 都会被重新标记为 000。

最终岛屿的数量就是我们进行深度优先搜索的次数。

```go
func numIslands(grid [][]byte) int {
	row := len(grid)
	col := len(grid[0])
	visit := make([][]bool, row)
	for i := range visit {
		visit[i] = make([]bool, col)
	}
	count := 0
	for y := 0; y < row; y++ {
		for x := 0; x < col; x++ {
			if grid[y][x] == '0' || visit[y][x] {
				continue
			}
			travel(grid, y, x, row, col, visit)
			count += 1
		}
	}
	return count
}

func travel(grid [][]byte, y, x, row, col int, visit [][]bool) {
	if grid[y][x] == '0' || visit[y][x] {
		return
	}
	visit[y][x] = true
	if y > 0 && y < row {
		travel(grid, y-1, x, row, col, visit)
	}
	if y >= 0 && y < row-1 {
		travel(grid, y+1, x, row, col, visit)
	}
	if x > 0 && x < col {
		travel(grid, y, x-1, row, col, visit)
	}
	if x >= 0 && x < col-1 {
		travel(grid, y, x+1, row, col, visit)
	}
	return
}
```