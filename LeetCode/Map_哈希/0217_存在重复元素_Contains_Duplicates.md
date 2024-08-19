# 217. 存在重复元素

### 简单

给你一个整数数组 nums 。如果任一值在数组中出现 至少两次 ，返回 true ；如果数组中每个元素互不相同，返回 false 。
 
### 示例 1：

    输入：nums = [1,2,3,1]
    输出：true

### 示例 2：

    输入：nums = [1,2,3,4]
    输出：false

### 示例 3：

    输入：nums = [1,1,1,3,3,4,3,2,4,2]
    输出：true

### 解：
```go
func containsDuplicate(nums []int) bool {
	m := make(map[int]bool)
	for _, n := range nums {
		if _, ok := m[n]; ok {
			return true
		}
		m[n] = true
	}
	return false
}
```