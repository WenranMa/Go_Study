# 1299. 将每个元素替换为右侧最大元素 Replace Elements with Greatest Element on Right Side

### Easy

Given an array arr, replace every element in that array with the greatest element among the elements to its right, and replace the last element with -1.

After doing so, return the array.

### Example 1:

Input: arr = [17,18,5,4,6,1]
Output: [18,6,6,6,1,-1]
Explanation: 
- index 0 --> the greatest element to the right of index 0 is index 1 (18).
- index 1 --> the greatest element to the right of index 1 is index 4 (6).
- index 2 --> the greatest element to the right of index 2 is index 4 (6).
- index 3 --> the greatest element to the right of index 3 is index 4 (6).
- index 4 --> the greatest element to the right of index 4 is index 5 (1).
- index 5 --> there are no elements to the right of index 5, so we put -1.

### Example 2:

Input: arr = [400]
Output: [-1]
Explanation: There are no elements to the right of index 0.

Constraints:

1 <= arr.length <= 10^4
1 <= arr[i] <= 10^5

### 解：

从右往左

```go
// Straight forward, from right to left. 
// 如果max < arr[i], 则交换，否则替换arr[i]为max.
// O(n) time.
func replaceElements(arr []int) []int {
	max := -1
	for i := len(arr) - 1; i >= 0; i-- {
		if max < arr[i] {
			arr[i], max = max, arr[i]
		} else {
			arr[i] = max
		}
	}
	return arr
}
```