# 90. 子集 II

### 中等

给你一个整数数组 nums ，其中可能包含重复元素，请你返回该数组所有可能的 
子集
（幂集）。

解集 不能 包含重复的子集。返回的解集中，子集可以按 任意顺序 排列。

### 示例 1：
    
    输入：nums = [1,2,2]
    输出：[[],[1],[1,2],[1,2,2],[2],[2,2]]

### 示例 2：

    输入：nums = [0]
    输出：[[],[0]]

### 提示：
- 1 <= nums.length <= 10
- -10 <= nums[i] <= 10

### 解：
相对于78题，多了有个for loop用于去除重复。

```go
func subsetsWithDup(nums []int) [][]int {
	res := [][]int{[]int{}}
	sort.Ints(nums)
	var sub func(nums []int, index int, row []int)
	sub = func(nums []int, index int, row []int) {
		for i := index; i < len(nums); i++ {
			row = append(row, nums[i])
			r := make([]int, len(row))
			copy(r, row)
			res = append(res, r)
			sub(nums, i+1, row)
			row = row[:len(row)-1]
			for i < len(nums)-1 && nums[i] == nums[i+1] {
				i += 1
			}
		}
	}
	sub(nums, 0, []int{})
	return res
}
```