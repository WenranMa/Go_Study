# 75. 颜色分类

### 中等

给定一个包含红色、白色和蓝色、共 n 个元素的数组 nums ，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。

我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。

必须在不使用库内置的 sort 函数的情况下解决这个问题。

### 示例 1：

输入：nums = [2,0,2,1,1,0]
输出：[0,0,1,1,2,2]

### 示例 2：

输入：nums = [2,0,1]
输出：[0,1,2]

### 提示：
- n == nums.length
- 1 <= n <= 300
- nums[i] 为 0、1 或 2

### 进阶：

你能想出一个仅使用常数空间的一趟扫描算法吗？

### 解：
1. 单指针遍历两遍。
2. 三指针遍历一边。
3. 长度为3的数组记录。
```go
// 1 pointer, two passes.
func sortColors(nums []int) {
	l := len(nums)
	p := 0
	for i := range nums {
		if nums[i] == 0 {
			nums[i], nums[p] = nums[p], nums[i]
			p += 1
		}
	}
	for i := p; i < l; i++ {
		if nums[i] == 1 {
			nums[i], nums[p] = nums[p], nums[i]
			p += 1
		}
	}
}

// 3 pointers.
func sortColors(nums []int) {
	l := len(nums)
	red, blue:= 0, l -1
    i:= 0
	for i <= blue  {
		if nums[i] == 0 {
			nums[i], nums[red] = nums[red], nums[i]
			red += 1
            i += 1
		} else if  nums[i] == 2 {
            nums[i], nums[blue] = nums[blue], nums[i]
            blue -= 1
        } else {
            i += 1
        }
	}
}

// use array with length 3.
// TBD
```