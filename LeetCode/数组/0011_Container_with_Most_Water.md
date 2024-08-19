# 11. 盛最多水的容器

### 中等

给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。

找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

返回容器可以储存的最大水量。

说明：你不能倾斜容器。

### 示例 1：
![water](/file/img/water_11.jpg)

输入：[1,8,6,2,5,4,8,3,7]

输出：49 

解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。

### 示例 2：

输入：height = [1,1]

输出：1
 
### 提示：
- n == height.length
- 2 <= n <= 105
- 0 <= height[i] <= 10^4

### 解：

题目的意思是从数组里面抽出两个数（两条线），和他们之间的距离组成一个容器，然后计算器容量。

长度短的线决定装多少水，当确定短线后，长线的下标增加（在右边则减小）只会使得容器的面积减小，所以if语句中改变短线的下标。two pointers, O(n) time, O(1) space.

```go
func maxArea(height []int) int {
    l,r:= 0, len(height) - 1
    sum:= 0
    for l< r {
        sum = max(sum, min(height[l], height[r])* (r-l))
        if height[l] < height[r] {
            l++
        } else {
            r--
        }
    }
    return sum
}
```