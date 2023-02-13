# 1572. Matrix Diagonal Sum

### Easy

Given a square matrix mat, return the sum of the matrix diagonals.

Only include the sum of all the elements on the primary diagonal and all the elements on the secondary diagonal that are not part of the primary diagonal.

### Example 1:

Input: mat = [[1,2,3],
              [4,5,6],
              [7,8,9]]
Output: 25
Explanation: Diagonals sum: 1 + 5 + 9 + 3 + 7 = 25
Notice that element mat[1][1] = 5 is counted only once.

### Example 2:

Input: mat = [[1,1,1,1],
              [1,1,1,1],
              [1,1,1,1],
              [1,1,1,1]]
Output: 8
Example 3:

Input: mat = [[5]]
Output: 5

Constraints:

n == mat.length == mat[i].length
1 <= n <= 100
1 <= mat[i][j] <= 100

```go
//O(n) time, 每次加最外面的一圈4个数。奇数长度单独加中间的数。
func diagonalSum(mat [][]int) int {
    l := len(mat)
    start := 0
    end := l - 1
    res := 0
    for start < end {
        res += mat[start][start] + mat[start][end] + mat[end][start] + mat[end][end]
        start += 1
        end -= 1
    }
    if l%2 == 1 {
        l = l / 2
        res += mat[l][l]
    }
    return res
}
```