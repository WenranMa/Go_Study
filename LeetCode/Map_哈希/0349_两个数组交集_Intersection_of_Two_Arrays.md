# 349_两个数组交集_Intersection_of_Two_Arrays
Given two arrays, write a function to compute their intersection.

Example:

    Input: nums1 = [1,2,2,1], nums2 = [2,2]   
    Output: [2]  
    Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]  
    Output: [9,4]  

Note:  
Each element in the result must be unique.
The result can be in any order.

### 解：

注意删除已经找到的元素，为了保证输出结果是唯一。

```go
func intersection(nums1 []int, nums2 []int) []int {
    m := make(map[int]int)
    for _, n := range nums1 {
        m[n] = 1
    }
    ans := []int{}
    for _, n := range nums2 {
        if _, ok := m[n]; ok {
            ans = append(ans, n)
            delete(m, n)
        }
    }
    return ans
}
```