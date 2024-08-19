# 883_三维体积投影面积_Projection_Area_of_3D_Shapes

On a N * N grid, we place some 1 * 1 * 1 cubes that are axis-aligned with the x, y, and z axes. Each value v = grid[i][j] represents a tower of v cubes placed on top of grid cell (i, j). Now we view the projection of these cubes onto the xy, yz, and zx planes. A projection is like a shadow, that maps our 3 dimensional figure to a 2 dimensional plane. Here, we are viewing the "shadow" when looking at the cubes from the top, the front, and the side. Return the total area of all three projections.

### Example:

    Input: [[2]]  
    Output: 5  
    
    Input: [[1,2],[3,4]]  
    Output: 17  
    
    Explanation:   
    Here are the three projections ("shadows") of the shape made with each axis-aligned plane.  
    
    Input: [[1,0],[0,2]]  
    Output: 8   
    
    Input: [[1,1,1],[1,0,1],[1,1,1]]  
    Output: 14  
    
    Input: [[2,2,2],[2,1,2],[2,2,2]]  
    Output: 21   
 
Note:   
1 <= grid.length = grid[0].length <= 50  
0 <= grid[i][j] <= 50

### 解：

每行最大元素的和 + 每列最大元素的和 + 非零元素个数

```go
func projectionArea(grid [][]int) int {
    ans := 0
    l := len(grid)
    for i := 0; i < l; i++ {
        m := 0
        for j := 0; j < l; j++ {
            if grid[i][j] > m {
                m = grid[i][j]
            }
            if grid[i][j] != 0 {
                ans += 1
            }
        }
        ans += m
    }
    for j := 0; j < l; j++ {
        m := 0
        for i := 0; i < l; i++ {
            if grid[i][j] > m {
                m = grid[i][j]
            }
        }
        ans += m
    }
    return ans
}
```
