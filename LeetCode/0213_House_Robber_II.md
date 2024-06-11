# 213. 打家劫舍 II

### 中等
你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都 围成一圈 ，这意味着第一个房屋和最后一个房屋是紧挨着的。同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警 。

给定一个代表每个房屋存放金额的非负整数数组，计算你 在不触动警报装置的情况下 ，今晚能够偷窃到的最高金额。

### 示例 1：

     输入：nums = [2,3,2]
     输出：3
     解释：你不能先偷窃 1 号房屋（金额 = 2），然后偷窃 3 号房屋（金额 = 2）, 因为他们是相邻的。

### 示例 2：

     输入：nums = [1,2,3,1]
     输出：4
     解释：你可以先偷窃 1 号房屋（金额 = 1），然后偷窃 3 号房屋（金额 = 3）。
          偷窃到的最高金额 = 1 + 3 = 4 。

### 示例 3：

     输入：nums = [1,2,3]
     输出：3

### 解：
因为收尾相连，用两次House Robber I的方法，第一次是0到end- 1，第二次是1到end，然后取两次中较大的值；O(n) time. O(1) space.

```go
func rob(nums []int) int {
	l := len(nums)
	if l == 1 {
		return nums[0]
	}
	if l == 2 {
		return max(nums[0], nums[1])
	}
	return max(robbery(nums[1:l]), robbery(nums[0:l-1]))
}

func robbery(nums []int) int {
	l := len(nums)
	if l == 2 {
		return max(nums[0], nums[1])
	}
	f1, f2 := nums[0], max(nums[0], nums[1])
	res := 0
	for i := 2; i < l; i++ {
		res = max(nums[i]+f1, f2)
		f1, f2 = f2, res
	}
	return res
}
```