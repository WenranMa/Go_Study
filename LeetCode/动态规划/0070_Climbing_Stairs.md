# 70. 爬楼梯

### 简单

假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

### 示例 1：

输入：n = 2

输出：2

解释：有两种方法可以爬到楼顶。

1. 1 阶 + 1 阶
2. 2 阶

### 示例 2：

输入：n = 3

输出：3

解释：有三种方法可以爬到楼顶。

1. 1 阶 + 1 阶 + 1 阶
2. 1 阶 + 2 阶
3. 2 阶 + 1 阶

### 备注
1 <= n <= 45

### 解：

动态规划

方法1：递归，超时
```go
func climbStairs(n int) int {
	if n == 1 || n == 0 {
		return 1
	}
	return climbStairs(n-1) + climbStairs(n-2)
}
```

方法2：循环，其实就是fabinacci数组。
```go
func climbStairs(n int) int {
	w1, w2 := 1, 2
	if n == 1 {
		return w1
	}
	if n == 2 {
		return w2
	}
	res := 0
	for s := 3; s <= n; s++ {
		res = w1 + w2
		w1 = w2
		w2 = res
	}
	return res
}
```
