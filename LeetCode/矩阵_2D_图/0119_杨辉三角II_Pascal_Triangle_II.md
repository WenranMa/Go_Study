# 119. 杨辉三角 II

### 简单

给定一个非负索引 rowIndex，返回「杨辉三角」的第 rowIndex 行。

在「杨辉三角」中，每个数是它左上方和右上方的数的和。

### 示例 1:

	输入: rowIndex = 3
	输出: [1,3,3,1]

### 示例 2:

	输入: rowIndex = 0
	输出: [1]

### 示例 3:

	输入: rowIndex = 1
	输出: [1,1]
 
### 提示:
0 <= rowIndex <= 33

### 进阶：
你可以优化你的算法到 O(rowIndex) 空间复杂度吗？

### 解：
第一个数和最后一个数永远是1.

注意从后往前计算。

```go
func getRow(rowIndex int) []int {
	res := make([]int, rowIndex+1)
	res[0] = 1
	for i := 1; i <= rowIndex; i++ {
		for j := i; j > 0; j-- {
			if j == i {
				res[j] = res[j-1] // last number is always 1.
			} else {
				res[j] = res[j] + res[j-1] // fist number 1 too, not calcuate.
			}
		}
	}
	return res
}
```