# 1281.整数的各位积和之差 Subtract the Product and Sum of Digits of an Integer

### Easy

Given an integer number n, return the difference between the product of its digits and the sum of its digits.
 
### Example 1:

Input: n = 234
Output: 15 
Explanation: 
Product of digits = 2 * 3 * 4 = 24 
Sum of digits = 2 + 3 + 4 = 9 
Result = 24 - 9 = 15

### Example 2:

Input: n = 4421
Output: 21
Explanation: 
Product of digits = 4 * 4 * 2 * 1 = 32 
Sum of digits = 4 + 4 + 2 + 1 = 11 
Result = 32 - 11 = 21

Constraints:

1 <= n <= 10^5

### 解：

一次循环就行，注意乘法初始化为1.

```go
func subtractProductAndSum(n int) int {
    p := 1
    s := 0
    for n > 0 {
        r := n % 10
        p = p * r
        s = s + r
        n = n / 10
    }
    return p - s
}
```