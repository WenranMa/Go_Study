# 442_数组中的重复数据_Find_All_Duplicates_in_an_Array
Given an array of integers, 1 ≤ a[i] ≤ n (n = size of array), some elements appear twice and others appear once. Find all the elements that appear twice in this array. Could you do it without extra space and in O(n) runtime?

Example:
Input: [4,3,2,7,8,2,3,1]
Output: [2,3]

### 解：

leetcode 268 missing number类似。
leetcode 448
leetcode 645

数组长度是n, 最大的数也是n，数组下标范围是0到n-1.

注意不要越界，所以for循环的条件是i < l.

判断条件里面，注意nums[nums[i]-1], 因为nums[i]最大是n, 所以n-1不会越界。

```go 
func findDuplicates(nums []int) []int {
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
            ans = append(ans, n)
        }
    }
    return ans
}
```