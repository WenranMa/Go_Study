# 861_反转矩阵后的得分_Score_After_Flipping_Matrix

We have a two dimensional matrix A where each value is 0 or 1. A move consists of choosing any row or column, and toggling each value in that row or column: changing all 0s to 1s, and all 1s to 0s. After making any number of moves, every row of this matrix is interpreted as a binary number, and the score of the matrix is the sum of these numbers. Return the highest possible score.

Example: 

    Input: [[0,0,1,1],
            [1,0,1,0],
            [1,1,0,0]]  
    Output: 39  
    
    Explanation:
    1. [[1,1,0,0],
        [1,0,1,0],
        [1,1,0,0]]  

    2. [[1,1,1,0],
        [1,0,0,0],
        [1,1,1,0]]   

    3. [[1,1,1,1],
        [1,0,0,1],
        [1,1,1,1]]  

    0b1111 + 0b1001 + 0b1111 = 15 + 9 + 15 = 39  
 
Note:  
1 <= A.length <= 20  
1 <= A[0].length <= 20  
A[i][j] is 0 or 1.  

### 解：

为了得到最高的分数，矩阵的每一行的最左边的数都必须为 1。为了做到这一点，我们可以翻转那些最左边的数不为 1 的那些行，而其他的行则保持不动。

当将每一行的最左边的数都变为 1 之后，就只能进行列翻转了。为了使得总得分最大，我们要让每个列中 1 的数目尽可能多。因此，我们扫描除了最左边的列以外的每一列，如果该列 0 的数目多于 1 的数目，就翻转该列，其他的列则保持不变。

```go
func matrixScore(A [][]int) int {
    A = flip(A)
    ans := 0
    for _, row := range A {
        r := 0
        for _, n := range row {
            r = r << 1
            r += n
        }
        ans += r
    }
    return ans
}

func flip(A [][]int) [][]int {
    r := len(A)
    c := len(A[0])
    //change 1st col to 1s. 也就是以0开头的行，都反转。
    for i := 0; i < r; i++ {
        if A[i][0] == 0 {
            for j := 0; j < c; j++ {
                A[i][j] ^= 1
            }
        }
    }
    //change the rest 0s to 1s, when 0s are more than 1s
    for j := 1; j < c; j++ {
        n := 0
        for i := 0; i < r; i++ {
            n += A[i][j]
        }
        if n <= r/2 { // 0s more than 1s
            for i := 0; i < r; i++ {
                A[i][j] ^= 1
            }
        }
    }
    return A
}
```