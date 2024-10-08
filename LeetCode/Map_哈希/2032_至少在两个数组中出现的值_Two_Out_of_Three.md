# 2032. 至少在两个数组中出现的值 Two Out of Three

### Easy

Given three integer arrays nums1, nums2, and nums3, return a distinct array containing all the values that are present in at least two out of the three arrays. You may return the values in any order.
 
### Example 1:

Input: nums1 = [1,1,3,2], nums2 = [2,3], nums3 = [3]
Output: [3,2]
Explanation: The values that are present in at least two arrays are:
- 3, in all three arrays.
- 2, in nums1 and nums2.

### Example 2:

Input: nums1 = [3,1], nums2 = [2,3], nums3 = [1,2]
Output: [2,3,1]
Explanation: The values that are present in at least two arrays are:
- 2, in nums2 and nums3.
- 3, in nums1 and nums2.
- 1, in nums1 and nums3.

### Example 3:

Input: nums1 = [1,2,2], nums2 = [4,3,3], nums3 = [5]
Output: []
Explanation: No value is present in at least two arrays.

Constraints:

1 <= nums1.length, nums2.length, nums3.length <= 100
1 <= nums1[i], nums2[j], nums3[k] <= 100

### 解：

分别遍历三个数组，然后分别 或001，010，100 . 最后map中的值如果只等于1，2，4，则说明该值只出现在一个数组中，则跳过。

```go
// map + bit mask.
func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
	m := make(map[int]int)
	for _, c := range nums1 {
		m[c] = 1
	}
	for _, c := range nums2 {
		m[c] |= 2
	}
	for _, c := range nums3 {
		m[c] |= 4
	}
	res := []int{}
	for k, v := range m {
		if v != 1 && v != 2 && v != 4 {
			res = append(res, k)
		}
	}
	return res
}
```