# 128. 最长连续序列

### 中等

给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。

请你设计并实现时间复杂度为 O(n) 的算法解决此问题。

### 示例 1：

    输入：nums = [100,4,200,1,3,2]
    输出：4
    解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。

### 示例 2：

    输入：nums = [0,3,7,2,5,8,4,6,0,1]
    输出：9

### 提示：
- 0 <= nums.length <= 10^5
- -10^9 <= nums[i] <= 10^9

### 解：

对一个数，从n+1,n-1两个方向找是否存在，如果存在，用一个map记录是不是被访问过。

```go
func longestConsecutive(nums []int) int {
	m := make(map[int]bool)
	for _, n := range nums {
		m[n] = false
	}
	res := 0
	for _, n := range nums {
		if !m[n] {
			m[n] = true
			len := 1 + findNeighbor(m, n+1, 1)
			len += findNeighbor(m, n-1, -1)
			res = max(res, len)
		}
	}
	return res
}

func findNeighbor(m map[int]bool, next, step int) int {
	len := 0
	_, ok := m[next]
	for ok {
		m[next] = true
		len += 1
		next += step
		_, ok = m[next]
	}
	return len
}
```