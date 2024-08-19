# 461_汉明距离_Hamming_Distance
两个整数之间的 汉明距离 指的是这两个数字对应二进制位不同的位置的数目。

给你两个整数 x 和 y，计算并返回它们之间的汉明距离。

Note:  
0 ≤ x, y < 231.

Example:

    Input: x = 1, y = 4  
    Output: 2    
    Explanation:  
    1   (0 0 0 1)  
    4   (0 1 0 0)

    输入：x = 3, y = 1
    输出：1

### 解：

异或，然后统计1的个数

```go
func hammingDistance(x int, y int) int {
	h := x ^ y
	ans := 0
	for h != 0 {
		ans += h & 1
		h = h >> 1
	}
	return ans
}
```