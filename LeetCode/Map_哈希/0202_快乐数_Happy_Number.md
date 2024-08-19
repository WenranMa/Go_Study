# 202. 快乐数

### 简单

编写一个算法来判断一个数 n 是不是快乐数。

「快乐数」 定义为：

- 对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和。
- 然后重复这个过程直到这个数变为 1，也可能是 无限循环 但始终变不到 1。
- 如果这个过程 结果为 1，那么这个数就是快乐数。
- 如果 n 是 快乐数 就返回 true ；不是，则返回 false 。

### 示例 1：

输入：n = 19

输出：true

解释：
- 1^2 + 9^2 = 82
- 8^2 + 2^2 = 68
- 6^2 + 8^2 = 100
- 1^2 + 0^2 + 0^2 = 1

### 示例 2：

输入：n = 2

输出：false
 
### 提示：

1 <= n <= 2^31 - 1

### 解：
1. 对每一个循环的当前数字求各位平方和
2. 如果平方和等于1，那么返回true
3. 如果平方和不为1，放入map中
4. 如果map中出现过这个数，说明进入循环，return false;

```go
func isHappy(n int) bool {
	m := make(map[int]int)
	for {
		n = getSum(n)
		if n == 1 {
			return true
		}
		if _, ok := m[n]; ok {
			return false
		}
		m[n] = 1
	}
}

func getSum(n int) int {
	res := 0
	for n > 0 {
		a := n % 10
		res += a * a
		n = n / 10
	}
	return res
}
```