# 704_二分查找_Binary_Search

Given a sorted (in ascending order) integer array nums of n elements and a target value, write a function to search target in nums. If target exists, then return its index, otherwise return -1.

Example:   
Input: nums = [-1,0,3,5,9,12], target = 9   
Output: 4   
Explanation: 9 exists in nums and its index is 4   
Input: nums = [-1,0,3,5,9,12], target = 2   
Output: -1   
Explanation: 2 does not exist in nums so return -1 
 
Note:   
You may assume that all elements in nums are unique.   
n will be in the range [1, 10000].   
The value of each element in nums will be in the range [-9999, 9999].

### 解：

```go
func search(nums []int, target int) int {
    l := 0
    r := len(nums) - 1
    for l <= r {
        m := (l + r) / 2
        if nums[m] == target {
            return m
        } else if nums[m] < target {
            l = m + 1
        } else if nums[m] > target {
            r = m - 1
        }
    }
    return -1
}
```