# 1365.有多少小于当前数字的数字 How Many Numbers Are Smaller Than the Current Number

### Easy

Given the array nums, for each nums[i] find out how many numbers in the array are smaller than it. That is, for each nums[i] you have to count the number of valid j's such that j != i and nums[j] < nums[i].

Return the answer in an array.

### Example 1:

Input: nums = [8,1,2,2,3]
Output: [4,0,1,1,3]
Explanation: 
For nums[0]=8 there exist four smaller numbers than it (1, 2, 2 and 3). 
For nums[1]=1 does not exist any smaller number than it.
For nums[2]=2 there exist one smaller number than it (1). 
For nums[3]=2 there exist one smaller number than it (1). 
For nums[4]=3 there exist three smaller numbers than it (1, 2 and 2).

### Example 2:

Input: nums = [6,5,4,8]
Output: [2,1,0,3]
Example 3:

Input: nums = [7,7,7,7]
Output: [0,0,0,0]

Constraints:

2 <= nums.length <= 500
0 <= nums[i] <= 100

### 解：

```go
// use the numbers as index in another array.
// 哈希的思想，ind数组的值，就是小于等于index的数量。
func smallerNumbersThanCurrent(nums []int) []int {
	res := make([]int, len(nums))
	ind := make([]int, 101)
	for _, n := range nums {
		ind[n]++
	}
	for i := 1; i <= 100; i++ {
		ind[i] += ind[i-1]
	}
	for i, n := range nums {
		if n == 0 {
			res[i] = 0
		} else {
			res[i] = ind[n-1]
		}
	}
	return res
}
```