# 1323. 6和9组成的最大数字  Maximum 69 Number

### Easy

You are given a positive integer num consisting only of digits 6 and 9.

Return the maximum number you can get by changing at most one digit (6 becomes 9, and 9 becomes 6).

### Example 1:

Input: num = 9669
Output: 9969
Explanation: 
Changing the first digit results in 6669.
Changing the second digit results in 9969.
Changing the third digit results in 9699.
Changing the fourth digit results in 9666.
The maximum number is 9969.

### Example 2:

Input: num = 9996
Output: 9999
Explanation: Changing the last digit 6 to 9 results in the maximum number.

### Example 3:

Input: num = 9999
Output: 9999
Explanation: It is better not to apply any change.

### 解：

就把从高位起的第一个6改为9即可

```go
//Convert the num to an arr.
func maximum69Number(num int) int {
	arr := []int{}
	for num > 0 {
		arr = append(arr, num%10)
		num = num / 10
	}
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == 6 {
			arr[i] = 9
			break
		}
	}
	res := 0
	for i := len(arr) - 1; i >= 0; i-- {
		res *= 10
		res += arr[i]
	}
	return res
}
```