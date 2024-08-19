# 1122_数组相对排序_Relative_Sort_Array

Given two arrays arr1 and arr2, the elements of arr2 are distinct, and all elements in arr2 are also in arr1.

Sort the elements of arr1 such that the relative ordering of items in arr1 are the same as in arr2.  Elements that don't appear in arr2 should be placed at the end of arr1 in ascending order.

Example 1:

    Input: arr1 = [2,3,1,3,2,4,6,7,9,2,19], arr2 = [2,1,4,3,9,6]
    Output: [2,2,2,1,4,3,3,9,6,7,19]

### 解：

```go
func relativeSortArray(arr1 []int, arr2 []int) []int {
    res := []int{}
    for _, n := range arr2 {
        for j := len(arr1) - 1; j >= 0; j-- {  // 注意一定是从后往前遍历，如果从前往后，会又跳过某些值的情况。
            if n == arr1[j] {
                res = append(res, n)
                arr1 = append(arr1[:j], arr1[j+1:]...) // arr[j+1:] j+1 最大可以等于l,否则溢出。
            }
        }
    }
    sort.Ints(arr1)
    res = append(res, arr1...)
    return res
}
```