# 1342. 将数字变成0的操作次数 Number of Steps to Reduce a Number to Zero

### Easy

Given an integer num, return the number of steps to reduce it to zero.

In one step, if the current number is even, you have to divide it by 2, otherwise, you have to subtract 1 from it.

### Example 1:

Input: num = 14
Output: 6
Explanation: 
Step 1. 14 is even; divide by 2 and obtain 7. 
Step 2. 7 is odd; subtract 1 and obtain 6.
Step 3. 6 is even; divide by 2 and obtain 3. 
Step 4. 3 is odd; subtract 1 and obtain 2. 
Step 5. 2 is even; divide by 2 and obtain 1. 
Step 6. 1 is odd; subtract 1 and obtain 0.

### Example 2:

Input: num = 8
Output: 4
Explanation: 
Step 1. 8 is even; divide by 2 and obtain 4. 
Step 2. 4 is even; divide by 2 and obtain 2. 
Step 3. 2 is even; divide by 2 and obtain 1. 
Step 4. 1 is odd; subtract 1 and obtain 0.

### Example 3:

Input: num = 123
Output: 12

Constraints:

0 <= num <= 10^6

### 解：

```go
// 末位是1，就加2，是0就加1. 最后减去一个1.
func numberOfSteps(num int) int {
	ans := 0
	if num == 0 {
		return ans
	}
	for num != 0 {
		ans += (num & 1) + 1
		num = num >> 1
	}
	return ans - 1
}

// 1 -> 001   1
// 2 -> 010   2
// 3 -> 011    3
// 4 -> 100    3
// 5 -> 101    4
// 6 -> 110    4
// 7 -> 111    5
// 8 -> 1000   4
// 9 -> 1001   5
// 10 -> 1010  5
// 11 -> 1011  6
// 12 -> 1100  5
// 13 -> 1101  6
// 14 -> 1110  6

//2进制 最高位的1 + 后面的1 + 后面的位数

```