# 896_单调数列_Monotonic_Array
An array is monotonic if it is either monotone increasing or monotone decreasing.

An array A is monotone increasing if for all i <= j, A[i] <= A[j].  An array A is monotone decreasing if for all i <= j, A[i] >= A[j].

Return true if and only if the given array A is monotonic.

Example:

    Input: [1,2,2,3]
    Output: true
    Input: [6,5,4,4]
    Output: true
    Input: [1,3,2]
    Output: false
    Input: [1,1,1]
    Output: true
 
Note:   
1 <= A.length <= 50000
-100000 <= A[i] <= 100000

### 解：

递增或者递减都可以。

```go
func isMonotonic(A []int) bool {
    l := len(A)
    increasing := true
    decreasing := true
    for i := 1; i < l; i++ {
        if A[i-1] > A[i] {
            increasing = false
        }
        if A[i-1] < A[i] {
            decreasing = false
        }
    }
    return increasing || decreasing
}
```