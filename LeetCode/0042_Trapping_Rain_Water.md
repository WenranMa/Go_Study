# 42. 接雨水

### 困难

给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

### 示例 1：
![water]()

    输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
    输出：6
    解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。 

### 示例 2：

    输入：height = [4,2,0,3,2,5]
    输出：9

### 提示：
- n == height.length
- 1 <= n <= 2 * 10^4
- 0 <= height[i] <= 10^5

### 解：

方法1： 双指针，左右两个pass.
```go
func trap(height []int) int {
	water := 0
	maxHightIndex := 0
	for i := range height {
		if height[maxHightIndex] < height[i] {
			maxHightIndex = i
		}
	}

	hi := 0
	for i := 0; i < maxHightIndex; i++ {
		if hi > height[i] {
			water += hi - height[i]
		} else {
			hi = height[i]
		}
	}
	hi = 0
	for i := len(height) - 1; i > maxHightIndex; i-- {
		if hi > height[i] {
			water += hi - height[i]
		} else {
			hi = height[i]
		}
	}
	return water
} 
```

方法2： 单调栈

维护一个单调栈，单调栈存储的是下标，满足从栈底到栈顶的下标对应的数组 height 中的元素递减。

从左到右遍历数组，遍历到下标 i 时，如果栈内至少有两个元素，记栈顶元素为 top, top 的下面一个元素是 left，则一定有 `height[left]≥height[top]`。如果 `height[i]>height[top]`，则得到一个可以接雨水的区域，该区域的宽度是 `i−left−1`，高度是 `min⁡(height[left],height[i])−height[top]`，根据宽度和高度即可计算得到该区域能接的雨水量。

为了得到 left, 需要将 top 出栈。在对 top 计算能接的雨水量之后，left 变成新的 top，重复上述操作，直到栈变为空，或者栈顶下标对应的 height 中的元素大于或等于 height[i]。

在对下标 i 处计算能接的雨水量之后，将 i 入栈，继续遍历后面的下标，计算能接的雨水量。遍历结束之后即可得到能接的雨水总量。


```go
func trap(height []int) int {
	res := 0
	stack := []int{}
	for i, h := range height {
		for len(stack) > 0 && h > height[stack[len(stack)-1]] {
			lowIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			left := stack[len(stack)-1]
			width := i - left - 1
			hi := min(h, height[left]) - height[lowIndex]
			res += width * hi
		}
		stack = append(stack, i)

	}
	return res
}
```
