# 189. 轮转数组

### 中等

给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。

### 示例 1:
输入: nums = [1,2,3,4,5,6,7], k = 3

输出: [5,6,7,1,2,3,4]

解释:

向右轮转 1 步: [7,1,2,3,4,5,6]

向右轮转 2 步: [6,7,1,2,3,4,5]

向右轮转 3 步: [5,6,7,1,2,3,4]

### 示例 2:

输入：nums = [-1,-100,3,99], k = 2

输出：[3,99,-1,-100]

解释: 

向右轮转 1 步: [99,-1,-100,3]

向右轮转 2 步: [3,99,-1,-100]
 

### 提示：

1 <= nums.length <= 105
-231 <= nums[i] <= 231 - 1
0 <= k <= 105
 

### 进阶：

- 尽可能想出更多的解决方案，至少有 三种 不同的方法可以解决这个问题。
- 你可以使用空间复杂度为 O(1) 的 原地 算法解决这个问题吗？

### 解：

方法1：不是O(1) space，需要额外内存
```go
func rotate(nums []int, k int) {
	l := len(nums)
	k = k % l
	copy(nums, append(nums[l-k:l], nums[0:l-k]...))
}
```

方法2：先整体反转，再反转两个子数组。注意k可以大于数组的长度。
```go
func rotate(nums []int, k int) {
	l := len(nums)
	k = k % l
	reverse(nums)
	reverse(nums[0:k])
	reverse(nums[k:l])
}

func reverse(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
```