# 59. 螺旋矩阵 II

### 中等

给你一个正整数 n ，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。

### 示例 1：
![spiral](/file/img/spiraln.jpg)

    输入：n = 3
    输出：[[1,2,3],[8,9,4],[7,6,5]]

### 示例 2：

    输入：n = 1
    输出：[[1]]

### 提示：
1 <= n <= 20

### 解：
注意边界控制，行的最后一个都不算，属于列的第一个，同样列的最后一个也在下一行中进行赋值。

```go
func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	start, end := 0, n-1
	num := 1
	for start < end {
		for j := start; j < end; j++ {
			matrix[start][j] = num
			num += 1
		}
		for i := start; i < end; i++ {
			matrix[i][end] = num
			num += 1
		}
		for j := end; j > start; j-- {
			matrix[end][j] = num
			num += 1
		}
		for i := end; i > start; i-- {
			matrix[i][start] = num
			num += 1
		}
		start += 1
		end -= 1
	}
	if start == end {
		matrix[start][end] = num
	}
	return matrix
}
```