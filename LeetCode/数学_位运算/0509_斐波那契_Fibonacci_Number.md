# 509_斐波那契_Fibonacci_Number

The Fibonacci numbers, commonly denoted F(n) form a sequence, called the Fibonacci sequence, such that each number is the sum of the two preceding ones, starting from 0 and 1. That is,

F(0) = 0,   F(1) = 1
F(N) = F(N - 1) + F(N - 2), for N > 1.
Given N, calculate F(N).

Example:

    Input: 2  
    Output: 1  
    Explanation: F(2) = F(1) + F(0) = 1 + 0 = 1.  
    Input: 3  
    Output: 2  
    Explanation: F(3) = F(2) + F(1) = 1 + 1 = 2.  
    Input: 4  
    Output: 3  
    Explanation: F(4) = F(3) + F(2) = 2 + 1 = 3.  

### 解：

递归，最慢，O(2^N)
```go
func fib(N int) int {
    if N == 0 {
        return 0
    }
    if N == 1 {
        return 1
    }
    return fib(N-1) + fib(N-2)
}
```

循环递推，O(N)
```go
func fib(N int) int {
    var a, b int = 0, 1
    for i := 0; i < N; i++ {
        a, b = b, a+b
    }
    return a
}
```

矩阵快速幂，O(logN)
```go
type matrix [2][2]int

func multiply(a, b matrix) (c matrix) {
    for i := 0; i < 2; i++ {
        for j := 0; j < 2; j++ {
            c[i][j] = a[i][0]*b[0][j] + a[i][1]*b[1][j]
        }
    }
    return
}

func pow(a matrix, n int) matrix {
    ret := matrix{{1, 0}, {0, 1}}
    for ; n > 0; n >>= 1 {
        if n&1 == 1 {
            ret = multiply(ret, a)
        }
        a = multiply(a, a)
    }
    return ret
}

func fib(n int) int {
    if n < 2 {
        return n
    }
    res := pow(matrix{{1, 1}, {1, 0}}, n-1)
    return res[0][0]
}
```