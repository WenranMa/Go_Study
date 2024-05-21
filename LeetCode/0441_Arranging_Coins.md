# 441. Arranging Coins

### Easy

You have n coins and you want to build a staircase with these coins. The staircase consists of k rows where the ith row has exactly i coins. The last row of the staircase may be incomplete.

Given the integer n, return the number of complete rows of the staircase you will build.

### Example 1:

Input: n = 5
Output: 2
Explanation: Because the 3rd row is incomplete, we return 2.

### Example 2:

Input: n = 8
Output: 3
Explanation: Because the 4th row is incomplete, we return 3.

Constraints:

1 <= n <= 2^31 - 1

```go
// O(n) time.
func arrangeCoins(n int) int {
	res := 0
	for i := 1; n >= 0; i += 1 {
		n -= i
		res += 1
	}
	if n < 0 {
		res -= 1
	}
	return res
}

// O(1) time.
/*
1+2+3+...+x = n
-> (1+x)x/2 = n
-> x^2+x = 2n
-> x^2+x+1/4 = 2n +1/4
-> (x+1/2)^2 = 2n +1/4
-> (x+0.5) = sqrt(2n+0.25)
-> x = -0.5 + sqrt(2n+0.25)
*/
func arrangeCoins(n int) int {
	return int(math.Floor(math.Sqrt(float64(2*n)+0.25) - 0.5))
}
```