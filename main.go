package main

func main() {

}

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
	// grid[x][y] = 1
	grid[x][y] = 2
	*queue = append(*queue, []int{x, y})
	return 1
}
