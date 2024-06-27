# 137. 只出现一次的数字 II

### 中等

给你一个整数数组 nums ，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次 。请你找出并返回那个只出现了一次的元素。

你必须设计并实现线性时间复杂度的算法且使用常数级空间来解决此问题。

### 示例 1：

输入：nums = [2,2,3,2]
输出：3

### 示例 2：

输入：nums = [0,1,0,1,0,1,99]
输出：99

### 提示：
- 1 <= nums.length <= 3 * 10^4
- -2^31 <= nums[i] <= 2^31 - 1
- nums 中，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次

### 解：

bit位运算。创建一个长度为sizeof(int)的数组count[sizeof(int)]，count[i]表示所有元素的1在i位出现的次数。如果count[i]是3的整数倍，则忽略；否则就把该位取出来组成答案。
O(n) time. O(1) space. 

```go
func singleNumber(nums []int) int {
	cnt := [32]int{}
	res := int32(0)
	for i := 0; i < len(nums); i++ {
		for j := 0; j < 32; j++ {
			if (int32(nums[i]) >> j & 1) == 1 {
				cnt[j] = (cnt[j] + 1) % 3
			}
		}
	}
	for j := 0; j < 32; j++ {
		res += int32(cnt[j] << j)
	}
	return int(res)
}
```

# 136. 只出现一次的数字
# 260. 只出现一次的数字 III