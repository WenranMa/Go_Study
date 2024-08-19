# 118. 杨辉三角

### 简单

给定一个非负整数 numRows，生成「杨辉三角」的前 numRows 行。

在「杨辉三角」中，每个数是它左上方和右上方的数的和。

### 示例 1:
![triangle](/file/img/PascalTriangle1.gif)

输入: numRows = 5
输出: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]

### 示例 2:

输入: numRows = 1
输出: [[1]]

### 提示:
1 <= numRows <= 30

### 解：

```go 
func generate(numRows int) [][]int {
	res := [][]int{[]int{1}}
	for i := 1; i < numRows; i++ {
		row := make([]int, i+1)
		preRow := res[i-1]
		for j := len(row) - 1; j >= 0; j-- {
			if j == len(preRow) {
				row[j] = preRow[j-1]
			} else if j == 0 {
				row[j] = preRow[0]
			} else {
				row[j] = preRow[j-1] + preRow[j]
			}
		}
		res = append(res, row)
	}
	return res
}
```