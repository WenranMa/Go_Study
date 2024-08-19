# 169. 多数元素

### 简单

给定一个大小为 n 的数组 nums ，返回其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。

你可以假设数组是非空的，并且给定的数组总是存在多数元素。

### 示例 1：

输入：nums = [3,2,3]

输出：3

### 示例 2：

输入：nums = [2,2,1,1,1,2,2]

输出：2
 
### 提示：
n == nums.length

1 <= n <= 5 * 10^4

-10^9 <= nums[i] <= 10^9
 
### 进阶：
尝试设计时间复杂度为 O(n)、空间复杂度为 O(1) 的算法解决此问题。

### 解：

方法1： map. 但空间复杂度高了。

```go
func majorityElement(nums []int) int {
	m := make(map[int]int)
	for _, n := range nums {
		m[n] += 1
	}
	for k, v := range m {
		if v > len(nums)/2 {
			return k
		}
	}
	return 0
}
```

方法2：遍历数组，记录当前数字和次数，下一个数字不同次数减1，相同则加1。如果次数为0，更新数字和次数。最后一次指向的数字就是超过一半的数字。O(n) time.
```go
func majorityElement(nums []int) int {
	res := nums[0]
	cnt := 1
	for i := 1; i < len(nums); i++ {
		if res != nums[i] {
			cnt -= 1
			if cnt <= 0 {
				res = nums[i]
				cnt = 1
			}
		} else {
			cnt += 1
		}
	}
	return res
}
```