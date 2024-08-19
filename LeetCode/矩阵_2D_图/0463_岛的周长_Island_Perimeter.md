# 463_岛的周长_Island_Perimeter

You are given a map in form of a two-dimensional integer grid where 1 represents land and 0 represents water.

Grid cells are connected horizontally/vertically (not diagonally). The grid is completely surrounded by water, and there is exactly one island (i.e., one or more connected land cells).

The island doesn't have "lakes" (water inside that isn't connected to the water around the island). One cell is a square with side length 1. The grid is rectangular, width and height don't exceed 100. Determine the perimeter of the island.

Example:  
Input:  
[[0,1,0,0],  
 [1,1,1,0],  
 [0,1,0,0],  
 [1,1,0,0]]

Output: 16

Explanation: The perimeter is the 16 yellow stripes in the image below:

![](/file/img/leetcode_463.png)

```go
func islandPerimeter(grid [][]int) int {
    h := len(grid)
    w := len(grid[0])
    ans := 0
    for i, _ := range grid {
        for j, _ := range grid[i] {
            if grid[i][j] == 1 {
                edge := 4 - checkNeighbor(grid, i, j, h, w)
                ans += edge
            }

        }
    }
    return ans
}

func checkNeighbor(grid [][]int, i, j, h, w int) int {
    n := 0
    if i-1 >= 0 && grid[i-1][j] == 1 {
        n += 1
    }
    if i+1 < h && grid[i+1][j] == 1 {
        n += 1
    }
    if j-1 >= 0 && grid[i][j-1] == 1 {
        n += 1
    }
    if j+1 < w && grid[i][j+1] == 1 {
        n += 1
    }
    return n
}
```