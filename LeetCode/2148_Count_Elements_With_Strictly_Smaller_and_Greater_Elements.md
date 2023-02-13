# 2148. Count Elements With Strictly Smaller and Greater Elements

### Easy

Given an integer array nums, return the number of elements that have both a strictly smaller and a strictly greater element appear in nums.

### Example 1:

Input: nums = [11,7,2,15]
Output: 2
Explanation: The element 7 has the element 2 strictly smaller than it and the element 11 strictly greater than it.
Element 11 has element 7 strictly smaller than it and element 15 strictly greater than it.
In total there are 2 elements having both a strictly smaller and a strictly greater element appear in nums.

### Example 2:

Input: nums = [-3,3,3,90]
Output: 2
Explanation: The element 3 has the element -3 strictly smaller than it and the element 90 strictly greater than it.
Since there are two elements with the value 3, in total there are 2 elements having both a strictly smaller and a strictly greater element appear in nums.

Constraints:

1 <= nums.length <= 100
-105 <= nums[i] <= 105

```go
// 排序，不好，可以优化。
func countElements(nums []int) int {
	l := len(nums)
	if l < 3 {
		return 0
	}
	sort.Ints(nums)
	i := 0
	j := l - 1
	for i < l {
		if i+1 < l && nums[i] == nums[i+1] {
			i += 1
		} else {
			break
		}

	}
	for j > 0 {
		if j-1 > 0 && nums[j] == nums[j-1] {
			j -= 1
		} else {
			break
		}
	}
	if i < j {
		return j - i - 1
	}
	return 0
}

// 优化，找到最大最小数。
func countElements(nums []int) int {
	min := math.MaxInt
	max := math.MinInt
	for _, n := range nums {
		if min > n {
			min = n
		}
		if max < n {
			max = n
		}
	}
	cnt := 0
	for _, n := range nums {
		if n < max && n > min {
			cnt += 1
		}
	}
	return cnt
}


```