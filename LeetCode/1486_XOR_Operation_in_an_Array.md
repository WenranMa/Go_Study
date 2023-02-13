# 1486. XOR Operation in an Array

### Easy

You are given an integer n and an integer start.

Define an array nums where nums[i] = start + 2 * i (0-indexed) and n == nums.length.

Return the bitwise XOR of all elements of nums.

### Example 1:

Input: n = 5, start = 0
Output: 8
Explanation: Array nums is equal to [0, 2, 4, 6, 8] where (0 ^ 2 ^ 4 ^ 6 ^ 8) = 8.
Where "^" corresponds to bitwise XOR operator.

### Example 2:

Input: n = 4, start = 3
Output: 8
Explanation: Array nums is equal to [3, 5, 7, 9] where (3 ^ 5 ^ 7 ^ 9) = 8.

Constraints:

1 <= n <= 1000
0 <= start <= 1000
n == nums.length

```go
func xorOperation(n int, start int) int {
	res := start
	for i := 1; i < n; i++ {
		res = res ^ (start + 2*i)
	}
	return res
}

```


Example:
Start = 0 ,1
n = 1...8

image

Looking at these charts it is easy to spot a pattern for some of the rows.
The pattern repeats after every 4th number

[N % 4 == 1] Green Row: Ans = Number[N]
[N % 4 == 2] Yellow Row: Ans = 2
[N % 4 == 3] Red Row: Ans = Number[N] ^ 2
[N % 4 == 0] Blue Row: Ans = 0

Start = 2 , 3
n = 1...8

image

Looking at these charts it is easy to spot a pattern for some of the rows.
Here, the pattern also repeats after every 4th number.

[N % 4 == 1] Green Row: Ans = Number[1]
[N % 4 == 2] Yellow Row: Ans = Number[N] ^ Number[1]
[N % 4 == 3] Red Row: Ans = Number[1] ^ 2
[N % 4 == 0] Blue Row: Ans = Number[N] ^ Number[1] ^ 2