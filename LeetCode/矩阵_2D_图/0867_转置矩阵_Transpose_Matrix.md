# 867_转置矩阵_Transpose_Matrix
Given a matrix A, return the transpose of A. The transpose of a matrix is the matrix flipped over it's main diagonal, switching the row and column indices of the matrix.

Example:  

    Input: [[1,2,3],
            [4,5,6],
            [7,8,9]] 

    Output: [[1,4,7],
             [2,5,8],
             [3,6,9]]   

    Input: [[1,2,3],
            [4,5,6]]

    Output: [[1,4],
             [2,5],
             [3,6]]
 
Note:   
1 <= A.length <= 1000   
1 <= A[0].length <= 1000  

### 解：

以这个为例：

    Input: [[1,2,3],
            [4,5,6]]

    Output: [[1,4],
             [2,5],
             [3,6]]


把i扫描放在内层循环。1，4 -> 2，5 -> 3，6的顺序

```go
func transpose(A [][]int) [][]int {
    r := len(A) 
    c := len(A[0]) 
    ans := [][]int{}
    for j := 0; j < c; j++ {
        row := []int{}
        for i := 0; i < r; i++ { 
            row = append(row, A[i][j]) 
        }
        ans = append(ans, row)
    }
    return ans
}
```