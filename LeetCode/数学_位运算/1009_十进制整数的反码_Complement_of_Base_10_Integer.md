# 1009.十进制整数的反码  Complement of Base 10 Integer

### Easy

The complement of an integer is the integer you get when you flip all the 0's to 1's and all the 1's to 0's in its binary representation.

For example, The integer 5 is "101" in binary and its complement is "010" which is the integer 2.
Given an integer n, return its complement.

### Example 1:

Input: n = 5
Output: 2
Explanation: 5 is "101" in binary, with complement "010" in binary, which is 2 in base-10.

### Example 2:

Input: n = 7
Output: 0
Explanation: 7 is "111" in binary, with complement "000" in binary, which is 0 in base-10.

### Example 3:

Input: n = 10
Output: 5
Explanation: 10 is "1010" in binary, with complement "0101" in binary, which is 5 in base-10.

Constraints:

0 <= n < 10^9

Note: This question is the same as 476

### 解：

可以用异或, 101 ^ 111 = 010, 注意0特殊情况。

```go
func bitwiseComplement(n int) int {
	if n == 0 {
		return 1
	}
	one := 0
	res := n
	for n > 0 {
		n >>= 1
		one <<= 1
		one += 1
	}
	return res ^ one
}
```

方法2：
```go 
/* 
当前数与补码的和等于2^x - 1
5 -- 101
2 -- 010
5 + 2 = 7 -- 111

res + n = 2^x - 1
res = 2^x - 1 - n
*/
func bitwiseComplement(n int) int {
	i := 1
	res := -1
	for res < 0 {
		res = (1 << i) - 1 - n
		i++
	}
	return res
}
```