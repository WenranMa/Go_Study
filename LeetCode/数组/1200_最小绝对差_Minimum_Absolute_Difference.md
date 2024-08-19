# 1200.  最小绝对差 Minimum Absolute Difference

### Easy

Given an array of distinct integers arr, find all pairs of elements with the minimum absolute difference of any two elements.

Return a list of pairs in ascending order(with respect to pairs), each pair [a, b] follows

a, b are from arr
a < b
b - a equals to the minimum absolute difference of any two elements in arr
 
### Example 1:

Input: arr = [4,2,1,3]
Output: [[1,2],[2,3],[3,4]]
Explanation: The minimum absolute difference is 1. List all pairs with difference equal to 1 in ascending order.

### Example 2:

Input: arr = [1,3,6,10,15]
Output: [[1,3]]

### Example 3:

Input: arr = [3,8,-10,23,19,-4,-14,27]
Output: [[-14,-10],[19,23],[23,27]]

Constraints:

2 <= arr.length <= 10^5
-10^6 <= arr[i] <= 10^6

### 解：

排序，然后两次遍历，第一次找到最小距离，第二次把结果加进去.

```go
// O(nlogn) time. Sorting.
func minimumAbsDifference(arr []int) [][]int {
	res := [][]int{}
	sort.Ints(arr)
	min := 10000000
	for i, a := range arr {
		if i > 0 && a-arr[i-1] < min {
			min = a - arr[i-1]
		}
	}
	for i, a := range arr {
		if i > 0 && a-arr[i-1] == min {
			t := []int{arr[i-1], a}
			res = append(res, t)
		}
	}
	return res
}
```

也可以一次遍历。先判断结果的前一次是不是等于这次的最小值，如果不是，删掉。

但如果时这种情况 [1,3,5,7,9,10]，答案应该时[9,10],前期结果里面会有[1,3][3,5]..等多个结果。休要循环删除，速度不一定快。

这是个错误答案，没有循环删除。
```go
func minimumAbsDifference(arr []int) [][]int {
	res := [][]int{}
	sort.Ints(arr)
	fmt.Println(arr)
	min := 10000000
	for i := 1; i < len(arr); i++ {
		if arr[i]-arr[i-1] < min {
			min = arr[i] - arr[i-1]
		}
		l := len(res)
		if l == 0 || l > 0 && res[l-1][1]-res[l-1][0] > min {
			if l > 0 {
				res = res[:l-1]
			}
			res = append(res, []int{arr[i-1], arr[i]})
		} else if arr[i]-arr[i-1] == min {
			res = append(res, []int{arr[i-1], arr[i]})
		}
	}
	return res
}
```
