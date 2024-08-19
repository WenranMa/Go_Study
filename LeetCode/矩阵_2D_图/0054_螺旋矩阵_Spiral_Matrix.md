# 54. 螺旋矩阵

### 中等

给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。

### 示例 1：
![spiral1](/file/img/spiral1.jpg)

输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
输出：[1,2,3,6,9,8,7,4,5]

### 示例 2：
![spiral2](/file/img/spiral2.jpg)

输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
输出：[1,2,3,4,8,12,11,10,9,5,6,7]

### 提示：

- m == matrix.length
- n == matrix[i].length
- 1 <= m, n <= 10
- -100 <= matrix[i][j] <= 100

### 解：
第一行的最后一个作为最后一列的第一个元素。以此类推，注意奇数行或列，单独处理。

```go
func spiralOrder(matrix [][]int) []int {
	res := []int{}
	rowStart, rowEnd := 0, len(matrix)-1
	colStart, colEnd := 0, len(matrix[0])-1
	for rowStart < rowEnd && colStart < colEnd {
		for j := colStart; j < colEnd; j++ {
			res = append(res, matrix[rowStart][j])
		}
		for i := rowStart; i < rowEnd; i++ {
			res = append(res, matrix[i][colEnd])
		}
		for j := colEnd; j > colStart; j-- {
			res = append(res, matrix[rowEnd][j])
		}
		for i := rowEnd; i > rowStart; i-- {
			res = append(res, matrix[i][colStart])
		}
		rowStart += 1
		rowEnd -= 1
		colStart += 1
		colEnd -= 1
	}
	if rowStart == rowEnd {
		for j := colStart; j <= colEnd; j++ {
			res = append(res, matrix[rowStart][j])
		}
	} else if colStart == colEnd {
		for i := rowStart; i <= rowEnd; i++ {
			res = append(res, matrix[i][colStart])
		}
	}
	return res
}
```