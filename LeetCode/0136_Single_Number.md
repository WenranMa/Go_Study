# 136. 只出现一次的数字

### 简单

给你一个 非空 整数数组 nums ，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。

你必须设计并实现线性时间复杂度的算法来解决此问题，且该算法只使用常量额外空间。

### 示例 1 ：

输入：nums = [2,2,1]

输出：1

### 示例 2 ：

输入：nums = [4,1,2,1,2]

输出：4

### 示例 3 ：

输入：nums = [1]

输出：1

### 提示：

1 <= nums.length <= 3 * 104

-3 * 104 <= nums[i] <= 3 * 104

除了某个元素只出现一次以外，其余每个元素均出现两次。

### 解：

用异或（按位异或），即 `1^0=1`, `1^1=0`, `0^0=0`，相同的为0，不同的为1，而且异或满足交换律。所以可以用循环。最后剩下来的就是出现一次的数，0^任何数都还是任何数。
```go
func singleNumber(nums []int) int {
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		res ^= nums[i]
	}
	return res
}
```