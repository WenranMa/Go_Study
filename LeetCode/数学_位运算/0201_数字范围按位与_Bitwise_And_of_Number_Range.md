# 201. 数字范围按位与

### 中等

给你两个整数 left 和 right ，表示区间 [left, right] ，返回此区间内所有数字 按位与 的结果（包含 left 、right 端点）。

### 示例 1：

输入：left = 5, right = 7
输出：4

### 示例 2：

输入：left = 0, right = 0
输出：0

### 示例 3：

输入：left = 1, right = 2147483647
输出：0

### 提示：
0 <= left <= right <= 2^31 - 1

### 解：

「Brian Kernighan 算法」，它用于清除二进制串中最右边的 1。

Brian Kernighan 算法的关键在于我们每次对 `number` 和 `number−1` 之间进行按位与运算后，number 中最右边的 1 会被抹去变成 0。

![bk](/file/img/Brian_Kernighan.png)

基于上述技巧，我们可以用它来计算两个二进制字符串的公共前缀。

其思想是，对于给定的范围 `[m,n]（m<n）`，我们可以对数字 n 迭代地应用上述技巧，清除最右边的 1，直到它小于或等于 m，此时非公共前缀部分的 1 均被消去。因此最后我们返回 n 即可。

```go
func rangeBitwiseAnd(left int, right int) int {
    for left < right {
        right &= right - 1
    }
    return right
}
```