# 977_有序数组的平方_Squares_of_a_Sorted_Array
Given an array of integers A sorted in non-decreasing order, return an array of the squares of each number, also in sorted non-decreasing order.

Example:

    Input: [-4,-1,0,3,10]  
    Output: [0,1,9,16,100]  
    Input: [-7,-3,2,3,11]  
    Output: [4,9,9,49,121]  
 
Note:  
1 <= A.length <= 10000  
-10000 <= A[i] <= 10000
A is sorted in non-decreasing order.  

### 解：

双指针，先找到正负数的分界，然后负数的index递减，正数的index递增。

```go
func sortedSquares(A []int) []int {
    l := len(A)
    i := 0 // for positive
    for i < l && A[i] < 0 {
        i++
    }
    j := i - 1 //for negative
    ans := []int{}
    for i < l || j >= 0 {
        if j < 0 || i < l && A[i]*A[i] <= A[j]*A[j] {
            ans = append(ans, A[i]*A[i])
            i++
        } else {
            ans = append(ans, A[j]*A[j])
            j--
        }
    }
    return ans
}
```