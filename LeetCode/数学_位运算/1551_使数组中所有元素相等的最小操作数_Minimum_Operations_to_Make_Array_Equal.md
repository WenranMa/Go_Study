# 1551. 使数组中所有元素相等的最小操作数 Minimum Operations to Make Array Equal

### Medium

You have an array arr of length n where arr[i] = (2 * i) + 1 for all valid values of i (i.e., 0 <= i < n).

In one operation, you can select two indices x and y where 0 <= x, y < n and subtract 1 from arr[x] and add 1 to arr[y] (i.e., perform arr[x] -=1 and arr[y] += 1). The goal is to make all the elements of the array equal. It is guaranteed that all the elements of the array can be made equal using some operations.

Given an integer n, the length of the array, return the minimum number of operations needed to make all the elements of arr equal.

### Example 1:

Input: n = 3
Output: 2
Explanation: arr = [1, 3, 5]
First operation choose x = 2 and y = 0, this leads arr to be [2, 3, 4]
In the second operation choose x = 2 and y = 0 again, thus arr = [3, 3, 3].

### Example 2:

Input: n = 6
Output: 9

Constraints:

1 <= n <= 10^4

### 解：

```go
/*
if n = 6, then array is [1,3,5,7,9,11]
steps = (1+11)/2 - 1 + (3+9)/2 - 3 + (5+7)/2 - 5
	  = 6 - 1 + 6 - 3 + 6 - 5
	  = 6 * 3 - (1 + 3 + 5)

and 1 + 3 + 5 + 7 ... = i^2, i is the numbers, which is n/2.
so the final result is:
steps = n * (n / 2) - (n / 2)^2.
	  = n^2/2 - n^2/4
	  = n^2/4 
*/
func minOperations(n int) int {
	return n*n/4
}
```