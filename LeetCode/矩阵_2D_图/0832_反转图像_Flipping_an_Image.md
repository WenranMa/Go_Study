# 832_反转图像_Flipping_an_Image
Given a binary matrix A, we want to flip the image horizontally, then invert it, and return the resulting image. To flip an image horizontally means that each row of the image is reversed.  For example, flipping [1, 1, 0] horizontally results in [0, 1, 1]. To invert an image means that each 0 is replaced by 1, and each 1 is replaced by 0. For example, inverting [0, 1, 1] results in [1, 0, 0].

### Example:
 
    Input: [[1,1,0],
            [1,0,1],
            [0,0,0]]

    Output: [[1,0,0],
             [0,1,0],
             [1,1,1]]
    
    Explanation: First reverse each row: [[0,1,1],
                                          [1,0,1],
                                          [0,0,0]].

    Then, invert the image: [[1,0,0],
                             [0,1,0],
                             [1,1,1]]  
    
    Input: [[1,1,0,0],
            [1,0,0,1],
            [0,1,1,1],
            [1,0,1,0]]

    Output: [[1,1,0,0],
             [0,1,1,0],
             [0,0,0,1],
             [1,0,1,0]]  
    
    Explanation: First reverse each row: [[0,0,1,1],
                                          [1,0,0,1],
                                          [1,1,1,0],
                                          [0,1,0,1]].

    Then invert the image: [[1,1,0,0],
                            [0,1,1,0],
                            [0,0,0,1],
                            [1,0,1,0]]  

Notes:  
1 <= A.length = A[0].length <= 20  
0 <= A[i][j] <= 1

### 解

水平反转和一零反转可以在同一个循环中完成。

```go
func flipAndInvertImage(A [][]int) [][]int {
    for _, row := range A {
        l := len(row)
        for j, k := 0, l-1; j <= k; {
            row[j], row[k] = row[k], row[j]
            row[j], row[k] = row[j]^1, row[k]^1
            j++
            k--
        }
    }
    return A
}
```