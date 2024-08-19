# 268. 丢失的数字

### 简单

给定一个包含 [0, n] 中 n 个数的数组 nums ，找出 [0, n] 这个范围内没有出现在数组中的那个数。

### 示例 1：

	输入：nums = [3,0,1]
	输出：2
	解释：n = 3，因为有 3 个数字，所以所有的数字都在范围 [0,3] 内。2 是丢失的数字，因为它没有出现在 nums 中。

### 示例 2：

	输入：nums = [0,1]
	输出：2
	解释：n = 2，因为有 2 个数字，所以所有的数字都在范围 [0,2] 内。2 是丢失的数字，因为它没有出现在 nums 中。

### 示例 3：

	输入：nums = [9,6,4,2,3,5,7,0,1]
	输出：8
	解释：n = 9，因为有 9 个数字，所以所有的数字都在范围 [0,9] 内。8 是丢失的数字，因为它没有出现在 nums 中。

### 示例 4：

	输入：nums = [0]
	输出：1
	解释：n = 1，因为有 1 个数字，所以所有的数字都在范围 [0,1] 内。1 是丢失的数字，因为它没有出现在 nums 中。

### 提示：
- n == nums.length
- 1 <= n <= 104
- 0 <= nums[i] <= n
- nums 中的所有数字都 独一无二

进阶：你能否实现线性时间复杂度、仅使用额外常数空间的算法解决此问题?

### 解：

leetcode 41
leetcode 442
leetcode 448
leetcode 645

基数排序?

```go
//[0,n]一共n+1个数，少一个，所以数组长度是n, 最大index是n-1

// 两种写法：
// 1 最后排序成 [1, 2, 3, 4, ... ]的形式，1在对应index 0.
// `nums[i] < len(nums)` 这个条件用于不包括0，也就是1-n少一个数，数组长度是n-1
func missingNumber(nums []int) int {
	for i := 0; i < len(nums); {
		if nums[i] > 0 && nums[i] < len(nums) && nums[i] != i+1 && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		} else {
			i++
		}
	}
	for i, n := range nums {
		if n != i+1 {
			return i + 1
		}
	}
	return 0
}

// 也可以排序成 [0, 1, 2, 3, ...] 这种 0在最前面。
func missingNumber(nums []int) int {
    l := len(nums)
    for i := 0; i < l; {
        if nums[i] < l && nums[i] != i {
            nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
        } else {
            i++
        }
    }
    for i, n := range nums {
        if n != i {
            return i
        }
    }
    return l
}
```
