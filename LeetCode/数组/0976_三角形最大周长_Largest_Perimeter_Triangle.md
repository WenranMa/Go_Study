# 976_三角形最大周长_Largest_Perimeter_Triangle
Given an array A of positive lengths, return the largest perimeter of a triangle with non-zero area, formed from 3 of these lengths.

If it is impossible to form any triangle of non-zero area, return 0.

Example 1:

    Input: [2,1,2]   
    Output: 5    
    Input: [1,2,1]    
    Output: 0      
    Input: [3,2,3,4]    
    Output: 10   
    Input: [3,6,2,3]    
    Output: 8    

Note:   
3 <= A.length <= 10000   
1 <= A[i] <= 10^6

### 解：

排序 + 从尾部遍历

```go
func largestPerimeter(A []int) int {
    l := len(A)
    if l <= 2 {
        return 0
    }
    sort.Ints(A)
    ans := 0
    for i := l - 1; i >= 2; i-- {
        if A[i-1]+A[i-2] > A[i] {
            ans = A[i-2] + A[i-1] + A[i]
            break
        }
    }
    return ans
}