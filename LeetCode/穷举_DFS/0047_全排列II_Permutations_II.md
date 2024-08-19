# 47. 全排列 II

### 中等

给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。

### 示例 1：

    输入：nums = [1,1,2]
    输出：
    [[1,1,2],
    [1,2,1],
    [2,1,1]]

### 示例 2：

    输入：nums = [1,2,3]
    输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

### 提示：
- 1 <= nums.length <= 8
- -10 <= nums[i] <= 10
- var res [][]int

### 解：

先排序，然后比permutation只多了for循环部分，有重复的数字就跳过。

```go
var res [][]int

func permuteUnique(nums []int) [][]int {
	res = [][]int{}
    sort.Ints(nums)
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
			for i+1 <= len(nums)-1 && nums[i] == nums[i+1] {
				i += 1
			}
		}
	}
} 
```