# 1351. Count Negative Numbers in a Sorted Matrix

### Easy

Given a m x n matrix grid which is sorted in non-increasing order both row-wise and column-wise, return the number of negative numbers in grid.

### Example 1:

Input: grid = [[4,3,2,-1],[3,2,1,-1],[1,1,-1,-2],[-1,-1,-2,-3]]
Output: 8
Explanation: There are 8 negatives number in the matrix.

### Example 2:

Input: grid = [[3,2],[1,0]]
Output: 0

Constraints:

m == grid.length
n == grid[i].length
1 <= m, n <= 100
-100 <= grid[i][j] <= 100
 
Follow up: Could you find an O(n + m) solution?

```go
// O(m*n) solution.
// straight forward. 
func countNegatives(grid [][]int) int {
	res := 0
	for i := len(grid) - 1; i >= 0; i-- {
		for j := len(grid[0]) - 1; j >= 0; j-- {
			if grid[i][j] < 0 {
				res += 1
			} else {
				break
			}
		}
	}
	return res
}

// O(m + n). 用一个循环，因为是排序的。
/*
+ + + + -
+ + + - -
+ - - - -
+ - - - -

比如这个例子m=4, n=5。走一个负数的path即可.
1. i = 0, j = 4, res += 4 = 4.
2. i = 0, j = 3, res = 4.
3. i = 1, j = 3, res += 3 = 7.
4. i = 1, j = 2, res = 7.
5. i = 2, j = 2, res += 2 = 9.
6. i = 2, j = 1, res += 2 = 11.
7. i = 2, j = 0, res = 11.
8. i = 3, j = 0, res = 11.
*/
func countNegatives(grid [][]int) int {
    res:= 0
    m:= len(grid)
    n:= len(grid[0])
    i:= 0
    j:= n - 1
    for i < m && j >= 0 {
        if grid[i][j] < 0 {
            j-= 1
            res += m - i
        } else {
            i+= 1
        }
    } 
    return res
}

```
