# 852_山脉数组的巅峰索引_Peak_Index_in_a_Mountain_Array

Let's call an array A a mountain if the following properties hold:

A.length >= 3  
There exists some 0 < i < A.length - 1 such that A[0] < A[1] < ... A[i-1] < A[i] > A[i+1] > ... > A[A.length - 1]  
Given an array that is definitely a mountain, return any i such that A[0] < A[1] < ... A[i-1] < A[i] > A[i+1] > ... > A[A.length - 1].  

Example:

    Input: [0,1,0]  
    Output: 1  
    Input: [0,2,1,0]  
    Output: 1  

Note: 
3 <= A.length <= 10000  
0 <= A[i] <= 10^6  
A is a mountain, as defined above. 

### 解：

```go
func peakIndexInMountainArray(A []int) int {
    l := len(A)
    for i := 0; i < l-1; i++ {
        if A[i] > A[i+1] {
            return i
        }
    }
    return 0
}
```