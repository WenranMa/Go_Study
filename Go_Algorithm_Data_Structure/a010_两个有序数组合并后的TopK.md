# 两个有序数组合并后的Top K 

## 方法1：合并数组。

```go
func topKofTwoSortedSlice(a, b []int, k int) []int {
	if len(a) == 0 || len(b) == 0 || k <= 0 {
		return nil
	}
	result := make([]int, k)
	indexA, indexB := 0, 0
	for i := 0; i < k; i++ {
		if indexA >= len(a) {
			result[i] = b[indexB]
			indexB++
		} else if indexB >= len(b) {
			result[i] = a[indexA]
			indexA++
		} else if a[indexA] < b[indexB] {
			result[i] = a[indexA]
			indexA++
		} else {
			result[i] = b[indexB]
			indexB++
		}
	}
	return result
}
```