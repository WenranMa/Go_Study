# 476. Number Complement

### Easy

The complement of an integer is the integer you get when you flip all the 0's to 1's and all the 1's to 0's in its binary representation.

For example, The integer 5 is "101" in binary and its complement is "010" which is the integer 2.
Given an integer num, return its complement.

### Example 1:

Input: num = 5
Output: 2
Explanation: The binary representation of 5 is 101 (no leading zero bits), and its complement is 010. So you need to output 2.

### Example 2:

Input: num = 1
Output: 0
Explanation: The binary representation of 1 is 1 (no leading zero bits), and its complement is 0. So you need to output 0.

Constraints:

1 <= num < 2^31

Note: This question is the same as 1009

```go
/* 
当前数与补码的和等于2^x - 1
5 -- 101
2 -- 010
5 + 2 = 7 -- 111

找到111，然后异或。
注意num是0时，答案不对。
*/
func findComplement(num int) int {
	n := num
	ones := 0
	for n != 0 {
		n = n >> 1
		ones = ones<<1 + 1
	}
	return num ^ ones
}
```