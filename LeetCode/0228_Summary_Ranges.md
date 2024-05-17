# 228. Summary Ranges

### 简单

You are given a sorted unique integer array nums.

A range [a,b] is the set of all integers from a to b (inclusive).

Return the smallest sorted list of ranges that cover all the numbers in the array exactly. That is, each element of nums is covered by exactly one of the ranges, and there is no integer x such that x is in one of the ranges but not in nums.

Each range [a,b] in the list should be output as:

"a->b" if a != b
"a" if a == b
 
### Example 1:

Input: nums = [0,1,2,4,5,7]

Output: ["0->2","4->5","7"]

Explanation: The ranges are:

[0,2] --> "0->2"

[4,5] --> "4->5"

[7,7] --> "7"

### Example 2:

Input: nums = [0,2,3,4,6,8,9]

Output: ["0","2->4","6","8->9"]

Explanation: The ranges are:

[0,0] --> "0"

[2,4] --> "2->4"

[6,6] --> "6"

[8,9] --> "8->9"
 
### Constraints:

0 <= nums.length <= 20

-231 <= nums[i] <= 231 - 1

All the values of nums are unique.

nums is sorted in ascending order.

### 解：

```go
func summaryRanges(nums []int) []string {
	res := []string{}
	if len(nums) == 0 {
		return res
	}
	start := 0
	for i := 1; i < len(nums); i++ {
		diff := nums[i] - nums[i-1]
		if diff > 1 {
			res = append(res, getRange(nums[start], nums[i-1]))
			start = i
		}
	}
	res = append(res, getRange(nums[start], nums[len(nums)-1]))
	return res
}

func getRange(start, to int) string {
	if start == to {
		return fmt.Sprint(start)
	}
	return fmt.Sprintf("%d->%d", start, to)
}
```

