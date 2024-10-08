# 1523. 在区间范围内统计奇数数目 Count Odd Numbers in an Interval Range

### Easy

Given two non-negative integers low and high. Return the count of odd numbers between low and high (inclusive).

### Example 1:

Input: low = 3, high = 7
Output: 3
Explanation: The odd numbers between 3 and 7 are [3,5,7].

### Example 2:

Input: low = 8, high = 10
Output: 1
Explanation: The odd numbers between 8 and 10 are [9].

Constraints:

0 <= low <= high <= 10^9

```go
//O(1) time, space.
// 2 3 4 5 6 7 8    (8 - 2)/2 = 3
// 3 4 5 6 7        (7 - 3 + 2)/2 = 3
// 3 4 5 6 7 8      (8 - 3 + 1)/2 = 3
func countOdds(low int, high int) int {
	if low%2 == 1 {
		low -= 1
	}
	if high%2 == 1 {
		high += 1
	}
	return (high - low) / 2
}
```