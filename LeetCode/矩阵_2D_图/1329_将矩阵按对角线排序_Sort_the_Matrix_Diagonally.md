# 1329. 将矩阵按对角线排序 Sort the Matrix Diagonally

### Medium

A matrix diagonal is a diagonal line of cells starting from some cell in either the topmost row or leftmost column and going in the bottom-right direction until reaching the matrix's end. For example, the matrix diagonal starting from mat[2][0], where mat is a 6 x 3 matrix, includes cells mat[2][0], mat[3][1], and mat[4][2].

Given an m x n matrix mat of integers, sort each matrix diagonal in ascending order and return the resulting matrix.

### Example 1:

Input: mat = [[3,3,1,1],[2,2,1,2],[1,1,1,2]]
Output: [[1,1,1,1],[1,2,2,2],[1,2,3,3]]

### Example 2:

Input: mat = [[11,25,66,1,69,7],[23,55,17,45,15,52],[75,31,36,44,58,8],[22,27,33,25,68,4],[84,28,14,11,5,50]]
Output: [[5,17,4,1,52,7],[11,11,25,45,8,69],[14,23,25,44,58,15],[22,27,31,36,50,66],[84,28,75,33,55,68]]

Constraints:

m == mat.length
n == mat[i].length
1 <= m, n <= 100
1 <= mat[i][j] <= 100

### 解：

```go
// Straight forward. 对角存下来，排序，再放回去。
// O(M*NlogN) time. M=max(rows, cols),N=min(rows, cols)
func diagonalSort(mat [][]int) [][]int {
	rows := len(mat)
	cols := len(mat[0])
	for i := rows - 1; i >= 0; i-- {
		t := []int{}
		for j, k := i, 0; j < rows && k < cols; j, k = j+1, k+1 {
			t = append(t, mat[j][k])
		}
		sort.Ints(t)
		for j, k := i, 0; j < rows && k < cols; j, k = j+1, k+1 {
			mat[j][k] = t[k]
		}
	}
	for j := 1; j < cols; j++ {
		t := []int{}
		for i, k := 0, j; i < rows && k < cols; i, k = i+1, k+1 {
			t = append(t, mat[i][k])
		}
		sort.Ints(t)
		for i, k := 0, j; i < rows && k < cols; i, k = i+1, k+1 {
			mat[i][k] = t[i]
		}
	}
	return mat
}
```