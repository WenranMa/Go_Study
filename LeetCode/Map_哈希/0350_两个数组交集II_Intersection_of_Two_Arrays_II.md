# 350_两个数组交集II_Intersection_of_Two_Arrays_II
Given two arrays, write a function to compute their intersection.

Example:

    Input: nums1 = [1,2,2,1], nums2 = [2,2]  
    Output: [2,2]  

    Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]  
    Output: [4,9]  

Note:  
Each element in the result should appear as many times as it shows in both arrays.
The result can be in any order.

Follow up:  
What if the given array is already sorted? How would you optimize your algorithm?  
What if nums1's size is small compared to nums2's size? Which algorithm is better?  
What if elements of nums2 are stored on disk, and the memory is limited such that you cannot load all elements into the memory at once? 

```go
func intersect(nums1 []int, nums2 []int) []int {
    m := make(map[int]int)
    ans := []int{}
    for _, n := range nums1 {
        m[n] += 1
    }
    for _, n := range nums2 {
        if m[n] > 0 {
            ans = append(ans, n)
            m[n] -= 1
        }
    }
    return ans
}
```