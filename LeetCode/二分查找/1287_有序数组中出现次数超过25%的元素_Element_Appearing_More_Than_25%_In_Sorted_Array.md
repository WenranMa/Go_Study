# 1287.  有序数组中出现次数超过25%的元素 Element Appearing More Than 25% In Sorted Array

### Easy

Given an integer array sorted in non-decreasing order, there is exactly one integer in the array that occurs more than 25% of the time, return that integer.

### Example 1:

Input: arr = [1,2,2,6,6,6,6,7,10]
Output: 6

### Example 2:

Input: arr = [1,1]
Output: 1

Constraints:

1 <= arr.length <= 10^4
0 <= arr[i] <= 10^5

### 解：

1. 遍历，其实就是找到出现最多的元素。

```go
// O(n) time. O(1) space.
func findSpecialInteger(arr []int) int {
	cnt := 0
	max := 0
	res := 0
	for i, a := range arr {
		if i > 0 && a != arr[i-1] {
			cnt = 1
		} else {
			cnt += 1
			if max < cnt {
				max = cnt
				res = a
			}
		}
	}
	return res
}
```

2. 二分查找

SearchInts返回a[i]的最左index.
Search返回[0,n) 返回f(j)为真的index.

```go
func findSpecialInteger(arr []int) int {
    n:= len(arr)
    span := n/4+1
    for i:= 0; i< n; i+= span {
        l := sort.SearchInts(arr, arr[i])
        r:= sort.Search(n, func(j int) bool {
           return  arr[j] > arr[i]
        })
        if r - l >= span {
            return arr[i]
        }
    }
    return -1
}
```
