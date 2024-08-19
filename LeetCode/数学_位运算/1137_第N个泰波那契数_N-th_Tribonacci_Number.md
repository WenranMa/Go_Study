# 1137 第N个泰波那契数 N-th Tribonacci Number

### Easy

The Tribonacci sequence Tn is defined as follows: 

T0 = 0, T1 = 1, T2 = 1, and Tn+3 = Tn + Tn+1 + Tn+2 for n >= 0.

Given n, return the value of Tn.

### Example 1:

Input: n = 4
Output: 4
Explanation:
T_3 = 0 + 1 + 1 = 2
T_4 = 1 + 1 + 2 = 4

### Example 2:

Input: n = 25
Output: 1389537

Constraints:

0 <= n <= 37
The answer is guaranteed to fit within a 32-bit integer, ie. answer <= 2^31 - 1.


### 解：

```go
func tribonacci(n int) int {
	a := 0
	b := 1
	c := 1
	if n == 0 {
		return a
	} else if n == 1 {
		return b
	} else if n == 2 {
		return c
	}
	for n >= 3 {
		a, b, c = b, c, a+b+c
		n -= 1
	}
	return c
}
```

```go
// 递归方法超时，因为过多的重复计算。
func tribonacci(n int) int {
	// if n == 0 {
	//     return 0
	// } else if n == 1 {
	//     return 1
	// } else if n == 2 {
	//     return 1
	// } else {
	//     return tribonacci(n - 1) + tribonacci(n - 2 ) + tribonacci(n - 3)
	// }
	a := 0
	b := 1
	c := 1
	if n == 0 {
		return a
	} else if n == 1 {
		return b
	} else if n == 2 {
		return c
	}
	res := 0
	for n >= 3 {
		res = a + b + c
		a = b
		b = c
		c = res
		n -= 1
	}
	return res
}

/*
n = 3, res = 0 + 1 + 1 = 2
n = 4, res = 1 + 1 + 2 = 4
n = 5, res = 1 + 2 + 4 = 7
n = 6, res = 2 + 4 + 7 = 13
n = 7, res = 4 + 7 + 13 = 24
*/
```