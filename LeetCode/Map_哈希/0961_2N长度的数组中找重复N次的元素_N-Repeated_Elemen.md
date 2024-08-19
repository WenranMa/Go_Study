# 961_2N长度的数组中找重复N次的元素_N-Repeated_Element_in_Size_2N_Array
In a array A of size 2N, there are N+1 unique elements, and exactly one of these elements is repeated N times. Return the element repeated N times.

Example:

    Input: [1,2,3,3]  
    Output: 3  
    Input: [2,1,2,5,3,2]  
    Output: 2  
    Input: [5,1,5,2,5,3,5,4]  
    Output: 5  
 
Note:  
4 <= A.length <= 10000  
0 <= A[i] < 10000  
A.length is even  

### 解：

```go
func repeatedNTimes(A []int) int {
    m := make(map[int]int)
    ans := 0
    for _, n := range A {
        m[n] += 1
        if m[n] > 1 {
            ans = n
            break
        }
    }
    return ans
}
```