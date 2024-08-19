# 448_数组中所有消失的数_Find_All_Numbers_Disappeared_in_an_Array

Given an array of integers where 1 ≤ a[i] ≤ n (n = size of array), some elements appear twice and others appear once. Find all the elements of [1, n] inclusive that do not appear in this array. Could you do it without extra space and in O(n) runtime? You may assume the returned list does not count as extra space.

Example:  
Input: [4,3,2,7,8,2,3,1]  
Output: [5,6]

### 解：

leetcode 268相似
leetcoce 442
leetcode 645

```go
func findDisappearedNumbers(nums []int) []int {
    l := len(nums)
    ans := []int{}
    for i := 0; i < l; {
        if nums[i] != i+1 && nums[i] != nums[nums[i]-1] {
            nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
        } else {
            i++
        }
    }
    for i, n := range nums {
        if n != i+1 {
            ans = append(ans, i+1)
        }
    }
    return ans
}
```
