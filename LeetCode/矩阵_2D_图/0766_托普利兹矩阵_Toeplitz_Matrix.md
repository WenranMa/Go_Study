# 766_托普利兹矩阵_Toeplitz_Matrix

A matrix is Toeplitz if every diagonal from top-left to bottom-right has the same element. Now given an M x N matrix, return True if and only if the matrix is Toeplitz.
 
Example:

    Input:  
    matrix = [  
    [1,2,3,4],  
    [5,1,2,3],  
    [9,5,1,2]]  
    Output: True  
    
    Explanation:  
    In the above grid, the diagonals are:   
    "[9]", "[5, 5]", "[1, 1, 1]", "[2, 2, 2]", "[3, 3]", "[4]".   
    In each diagonal all elements are the same, so the answer is True.   

    Input:  
    matrix = [  
    [1,2],  
    [2,2]]  
    Output: False  
   
    Explanation: The diagonal "[1, 2]" has different elements.  

Note:  
matrix will be a 2D array of integers.  
matrix will have a number of rows and columns in range [1, 20].  
matrix[i][j] will be integers in range [0, 99]. 

### 解：

分两个大循环，第一个从左下角开始，也就是最后一行，第一列。

    [1,2,3,4]  
    [5,1,2,3]  
    [9,5,1,2]

比如这个，第一个循环把，9，5，1搞定。

第二个循环负责剩下的。2，3，4

```go
func isToeplitzMatrix(matrix [][]int) bool {
    row := len(matrix)
    col := len(matrix[0])
    for r := row - 1; r >= 0; r-- { // 从左下角开始，最后一行，第一列。
        n := matrix[r][0]
        for i, j := r, 0; i < row && j < col; i, j = i+1, j+1 {
            if matrix[i][j] != n {
                return false
            }
        }
    }
    for c := 1; c < col; c++ {  // 
        n := matrix[0][c]
        for i, j := 0, c; i < row && j < col; i, j = i+1, j+1 {
            if matrix[i][j] != n {
                return false
            }
        }
    }
    return true
}
```