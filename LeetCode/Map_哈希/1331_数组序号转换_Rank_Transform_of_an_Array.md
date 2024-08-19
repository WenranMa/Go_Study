# 1331. 数组序号转换 Rank Transform of an Array

### Easy

Given an array of integers arr, replace each element with its rank.

The rank represents how large the element is. The rank has the following rules:

Rank is an integer starting from 1.
The larger the element, the larger the rank. If two elements are equal, their rank must be the same.
Rank should be as small as possible.

### Example 1:

Input: arr = [40,10,20,30]
Output: [4,1,2,3]
Explanation: 40 is the largest element. 10 is the smallest. 20 is the second smallest. 30 is the third smallest.

### Example 2:

Input: arr = [100,100,100]
Output: [1,1,1]
Explanation: Same elements share the same rank.

### Example 3:

Input: arr = [37,12,28,9,100,56,80,5,12]
Output: [5,3,4,2,8,6,7,1,3]

Constraints:

0 <= arr.length <= 10^5
-10^9 <= arr[i] <= 10^9

### 解：

先排序，排序后用map记录 数字和index. (index这里指 rank， 所以是m[carr[i-1]] + 1, 针对的是 [10,20,20,30,30,40], index是绝对值，rank是相对位置）
再遍历 原数组，再map中找index 作为结果。

```go
// Sort and Map.
// O(nlogn) time. O(n) space.
func arrayRankTransform(arr []int) []int {
	carr := []int{}
	if len(arr) == 0 {
		return carr
	}
	carr = append(carr, arr...)   // copy of arr
	sort.Ints(carr)
	m := make(map[int]int)
	m[carr[0]] = 1
	l := len(carr)
	for i := 0; i < l; i++ {
		if _, ok := m[carr[i]]; !ok {
			m[carr[i]] = m[carr[i-1]] + 1  // 注意
		}
	}
	res := make([]int, len(arr))
	for i, a := range arr {
		res[i] = m[a]
	}
	return res
}
```