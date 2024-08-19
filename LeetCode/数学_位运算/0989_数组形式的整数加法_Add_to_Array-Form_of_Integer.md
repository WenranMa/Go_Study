# 989. 数组形式的整数加法 Add to Array-Form of Integer

### Easy

The array-form of an integer num is an array representing its digits in left to right order.

For example, for num = 1321, the array form is [1,3,2,1].
Given num, the array-form of an integer, and an integer k, return the array-form of the integer num + k.

### Example 1:

Input: num = [1,2,0,0], k = 34
Output: [1,2,3,4]
Explanation: 1200 + 34 = 1234

### Example 2:

Input: num = [2,7,4], k = 181
Output: [4,5,5]
Explanation: 274 + 181 = 455

### Example 3:

Input: num = [2,1,5], k = 806
Output: [1,0,2,1]
Explanation: 215 + 806 = 1021

Constraints:

1 <= num.length <= 10^4
0 <= num[i] <= 9
num does not contain any leading zeros except for the zero itself.
1 <= k <= 10^4

### 解：

```go
// O(n) time, straight forward.
func addToArrayForm(num []int, k int) []int {
	l := len(num)
	for j := l - 1; j >= 0; j-- {
		num[j] += k % 10
		k /= 10
		if num[j] >= 10 {
			num[j] = num[j] % 10
			k += 1
		}
		if k == 0 {
			return num
		}

	}
	res := []int{}
	for k > 0 {
		res = append([]int{k % 10}, res...)
		k /= 10
	}
	res = append(res, num...)
	return res
}
```


这个不对，会溢出。。。如果数组太长的话。
```go
func addToArrayForm(num []int, k int) []int {
	a := 0
	for _, n := range num {
		a = a*10 + n
	}
	sum := a + k
	var res []int
	for sum > 0 {
		res = append(res, sum%10)
		sum /= 10
	}
	return reverse(res)
}

func reverse(num []int) []int {
	for i, j := 0, len(num)-1; i < j; i, j = i+1, j-1 {
		num[i], num[j] = num[j], num[i]
	}
	return num
}
```