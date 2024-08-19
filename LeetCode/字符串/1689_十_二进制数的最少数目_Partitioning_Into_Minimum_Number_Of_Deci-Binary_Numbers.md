# 1689.  十_二进制数的最少数目 Partitioning Into Minimum Number Of Deci-Binary Numbers

### Medium

A decimal number is called deci-binary if each of its digits is either 0 or 1 without any leading zeros. For example, 101 and 1100 are deci-binary, while 112 and 3001 are not.

Given a string n that represents a positive decimal integer, return the minimum number of positive deci-binary numbers needed so that they sum up to n.

### Example 1:

Input: n = "32"
Output: 3
Explanation: 10 + 11 + 11 = 32

### Example 2:

Input: n = "82734"
Output: 8

### Example 3:

Input: n = "27346209830709182346"
Output: 9

Constraints:

1 <= n.length <= 105
n consists of only digits.
n does not contain any leading zeros and represents a positive integer.

### 解：

贪心，遍历，每一位减一。但实际不用，只要返回最大的那一位就行。

```go
//返回最大的那一位就行。
func minPartitions(n string) int {
	var m rune = '0'
	for _, c := range n {
		if c > m {
			m = c
		}
	}
	return int(m - '0')
}
```