# Prime Factor Statistics

Given a positive integer N, you need to factorize all integers between (1, N].Then you have to count the number of the total prime numbers.

### Example:

input: N = 6 output：7
explain：2=2, 3=3, 4=22, 5=5, 6=23, the number of prime number : 1+1+2+1+2=7

### 解：
一个简化版的试除法来计数质因数。

计算 n 的平方根 sqrtn，用于限制循环的范围，因为大于 sqrt(n) 的因数必定与小于它的某个因数配对。

外层循环: 遍历从 2 到 sqrtn 的整数 i。

当 n 能被 i 整除时（即 n % i == 0）: 增加计数器 count，表示找到了一个质因数。用 n 除以 i，更新 n 的值，以便继续寻找下一个可能的因数。这个内循环会一直执行，直到 n 不再能被 i 整除。

步长调整: 如果当前的 i 是 2，则下一次检查的 i 应为 3（即 i += 1）。否则，由于已经检查了偶数因数（通过 2 的特殊情况处理），直接跳过后续偶数，让 i 加上 2（即 i += 2），这样可以减少检查次数，提高效率。

处理剩余的 n: 外层循环结束后，如果 n 大于 1，说明 n 本身就是一个大于 sqrt(n) 的质因数，因此将计数器 count 加一。

返回结果: 最后，函数返回计数器 count，表示 n 的质因数个数（不包括相同质因数的重复计数）。

```go
func countPrimeFactorsAndPrimes(n int) int {
	res := 0
	for i := 2; i <= n; i++ {
		res += countFactors(i)
	}
	return res
}

func countFactors(n int) int {
	count := 0
	sqrtn := int(math.Sqrt(float64(n)))
	for i := 2; i <= sqrtn; {
		for n%i == 0 {
			count += 1
			n /= i
		}
		if i == 2 {
			i += 1
		} else {
			i += 2
		}
	}
	if n > 1 {
		count += 1
	}
	return count
}
```