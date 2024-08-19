# 566. 重塑矩阵 Reshape the Matrix

### Easy

In MATLAB, there is a handy function called reshape which can reshape an m x n matrix into a new one with a different size r x c keeping its original data.

You are given an m x n matrix mat and two integers r and c representing the number of rows and the number of columns of the wanted reshaped matrix.

The reshaped matrix should be filled with all the elements of the original matrix in the same row-traversing order as they were.

If the reshape operation with given parameters is possible and legal, output the new reshaped matrix; Otherwise, output the original matrix.

### Example 1:

Input: mat = [[1,2],[3,4]], r = 1, c = 4
Output: [[1,2,3,4]]

### Example 2:

Input: mat = [[1,2],[3,4]], r = 2, c = 4
Output: [[1,2],[3,4]] 

Constraints:

m == mat.length
n == mat[i].length
1 <= m, n <= 100
-1000 <= mat[i][j] <= 1000
1 <= r, c <= 300

### 解：

```go
func matrixReshape(mat [][]int, r int, c int) [][]int {
	row := len(mat)
	col := len(mat[0])
	if r*c != row*col {
		return mat
	}
	res := [][]int{}
	oneRow := []int{}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			oneRow = append(oneRow, mat[i][j])
			if len(oneRow) == c {
				res = append(res, oneRow)
				oneRow = nil
				oneRow = []int{}
			}
		}
	}
	return res
}

// 改进，减少使用append.
func matrixReshape(mat [][]int, r int, c int) [][]int {
	row := len(mat)
	col := len(mat[0])
	if r*c != row*col {
		return mat
	}
	res := make([][]int, r)
	for i := 0; i < r; i++ {
		res[i] = make([]int, c)
	}
	for i := 0; i < row*col; i++ {
		res[i/c][i%c] = mat[i/col][i%col]
	}
	return res
}
```
