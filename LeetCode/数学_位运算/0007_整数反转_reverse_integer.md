# 7. 整数反转

### 中等

给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。

如果反转后整数超过 32 位的有符号整数的范围 [−231,  231 − 1] ，就返回 0。

假设环境不允许存储 64 位整数（有符号或无符号）。

### 示例 1：

输入：x = 123
输出：321

### 示例 2：

输入：x = -123
输出：-321

### 示例 3：

输入：x = 120
输出：21

### 示例 4：

输入：x = 0
输出：0

提示：

-2^31 <= x <= 2^31 - 1

### 解

题目要求不允许使用 64 位整数，即运算过程中的数字必须在 32 位有符号整数的范围内。所以不能用res和 INT_MAX, INT_MIN 比较

考虑 x>0 的情况，记 `INT_MAX = 2^31−1 = 2147483647`，由于 `INT_MAX = (INT_MAX / 10) * 10 + INT_MAX % 10 = (INT_MAX / 10) * 10 + 7`

​则不等式 rev * 10+digit ≤ INT_MAX

等价于 `rev * 10+digit ≤ (INT_MAX / 10) * 10 + 7`

移项得 `(rev− (INT_MAX / 10)) * 10 ≤ 7−digit`

讨论该不等式成立的条件：

若 `rev > (INT_MAX / 10)` 由于 digit≥0，不等式不成立。
若 `rev = (INT_MAX / 10)`，当且仅当 digit≤7 时，不等式成立。
若 `rev < (INT_MAX / 10)`，由于 digit≤9，不等式成立。

注意到当 rev= (INT_MAX / 10), 时若还能推入数字，则说明 x 的位数与 INT_MAX 的位数相同，且要推入的数字 digit 为 x 的最高位。由于 x 不超过 INT_MAX，因此 digit 不会超过 INT_MAX 的最高位，即 `digit ≤ 2`。所以实际上当 rev= (INT_MAX / 10) 时不等式必定成立。

因此判定条件可简化为：当且仅当 rev ≤ (INT_MAX / 10) 时，不等式成立。


```go
func reverse(x int) int {
	res := 0
	for d := 0; x != 0; x = x / 10 {
		if res < math.MinInt32/10 || res > math.MaxInt32/10 {
			return 0
		}
		d = x % 10
		res = res*10 + d
	}
	return res
}
```