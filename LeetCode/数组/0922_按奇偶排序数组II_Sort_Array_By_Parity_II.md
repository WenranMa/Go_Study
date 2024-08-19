# 922_按奇偶排序数组II_Sort_Array_By_Parity_II

Given an array A of non-negative integers, half of the integers in A are odd, and half of the integers are even. Sort the array so that whenever A[i] is odd, i is odd; and whenever A[i] is even, i is even. You may return any answer array that satisfies this condition.

### Example:

    Input: [4,2,5,7].  
    Output: [4,5,2,7].  
    Explanation: [4,7,2,5], [2,5,4,7], [2,7,4,5] would also have been accepted.  
 
Note:

    2 <= A.length <= 20000.  
    A.length % 2 == 0.  
    0 <= A[i] <= 1000.  

### 解：

双指针，一个指向偶数index，一个指向奇数index，然后交换。

```go
func sortArrayByParityII(A []int) []int {
    i, j := 0, 1
    l := len(A)
    for i < l || j < l {
        for ; i < l; i = i + 2 {
            if A[i]&1 == 1 {
                break
            }
        }
        for ; j < l; j = j + 2 {
            if A[j]&1 == 0 {
                break
            }
        }
        if i < l && j < l {
            A[i], A[j] = A[j], A[i]
        }
    }
    return A
}
```