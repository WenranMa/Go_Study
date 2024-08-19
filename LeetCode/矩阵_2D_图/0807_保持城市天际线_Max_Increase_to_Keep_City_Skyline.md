# 807_保持城市天际线_Max_Increase_to_Keep_City_Skyline

In a 2 dimensional array grid, each value grid[i][j] represents the height of a building located there. We are allowed to increase the height of any number of buildings, by any amount (the amounts can be different for different buildings). Height 0 is considered to be a building as well. 

At the end, the "skyline" when viewed from all four directions of the grid, i.e. top, bottom, left, and right, must be the same as the skyline of the original grid. A city's skyline is the outer contour of the rectangles formed by all the buildings when viewed from a distance. See the following example.

What is the maximum total sum that the height of the buildings can be increased?

Example:

    Input: grid = [[3,0,8,4],
                   [2,4,5,7],
                   [9,2,6,3],
                   [0,3,1,0]]  
    Output: 35  
    
    Explanation:   
    The grid is:  
    [ [3, 0, 8, 4],   
      [2, 4, 5, 7],  
      [9, 2, 6, 3],  
      [0, 3, 1, 0] ]  

    The skyline viewed from top or bottom is: [9, 4, 8, 7]  
    The skyline viewed from left or right is: [8, 7, 9, 3]

    The grid after increasing the height of buildings without affecting skylines is:

    gridNew = [ [8, 4, 8, 7],  
                [7, 4, 7, 7],  
                [9, 4, 8, 7],  
                [3, 3, 3, 3] ]  

Notes:  
1 < grid.length = grid[0].length <= 50.  
All heights grid[i][j] are in the range [0, 100].  
All buildings in grid[i][j] occupy the entire grid cell: that is, they are a 1 x 1 x grid[i][j] rectangular prism.

### 解：

用一次循环拿到 每行，每列的最大值。

然后再遍历矩阵，可以增加的高度 是和 行列最大值中较小的值比较。

```go
func maxIncreaseKeepingSkyline(grid [][]int) int {
	row := len(grid)
	col := len(grid[0])
	rowMax := make([]int, row)
	colMax := make([]int, col)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			rowMax[i] = max(rowMax[i], grid[i][j])
			colMax[j] = max(colMax[j], grid[i][j])
		}

	} // rowMax 每个内循环都可以确定，colMax要等外循环结束后才能确定。
	ans := 0
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			minor := min(rowMax[i], colMax[j])
			if grid[i][j] < minor {
				ans += minor - grid[i][j]
			}
		}
	}
	return ans
}
```

老解法，两次循环。
```go
func maxIncreaseKeepingSkyline(grid [][]int) int {
    rowMax := []int{}
    colMax := []int{}
    row := len(grid)
    col := len(grid[0])
    for i := 0; i < row; i++ {
        max := 0
        for j := 0; j < col; j++ {
            if max < grid[i][j] {
                max = grid[i][j]
            }
        }
        rowMax = append(rowMax, max)
    }
    for j := 0; j < col; j++ {
        max := 0
        for i := 0; i < row; i++ {
            if max < grid[i][j] {
                max = grid[i][j]
            }
        }
        colMax = append(colMax, max)
    }
    ans := 0
    for i := 0; i < row; i++ {
        for j := 0; j < col; j++ {
            min := int(math.Min(float64(rowMax[i]), float64(colMax[j])))
            if grid[i][j] < min {
                ans += min - grid[i][j]
            }
        }
    }
    return ans
}
```