# 46. 全排列

### 中等

给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。

### 示例 1：

输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

### 示例 2：

输入：nums = [0,1]
输出：[[0,1],[1,0]]

### 示例 3：

输入：nums = [1]
输出：[[1]]

### 提示：
- 1 <= nums.length <= 6
- -10 <= nums[i] <= 10
- nums 中的所有整数 互不相同

### 解：

```go
var res [][]int

func permute(nums []int) [][]int {
	res = [][]int{}
	visit := make([]bool, len(nums))
	perm(nums, visit, []int{})
	return res
}

func perm(nums []int, visit []bool, row []int) {
	if len(row) == len(nums) {
		r := make([]int, len(row))
		copy(r, row)
		res = append(res, r)
	}
	for i := 0; i < len(nums); i++ {
		if !visit[i] {
			row = append(row, nums[i])
			visit[i] = true
			perm(nums, visit, row)
			visit[i] = false
			row = row[:len(row)-1]
		}
	}
} 
```