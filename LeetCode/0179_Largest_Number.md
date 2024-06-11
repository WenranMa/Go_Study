# 179. 最大数

### 中等

给定一组非负整数 nums，重新排列每个数的顺序（每个数不可拆分）使之组成一个最大的整数。

注意：输出结果可能非常大，所以你需要返回一个字符串而不是整数。

### 示例 1：

    输入：nums = [10,2]
    输出："210"

### 示例 2：

    输入：nums = [3,30,34,5,9]
    输出："9534330"

提示：

1 <= nums.length <= 100
0 <= nums[i] <= 10^9

### 解：

重新定义排序算法：
比如 x=23，y=231，先计算x的位数100, y的位数1000，那么x*1000+y=23231，y*100+x=23123, 所以x > y. 

```go
func largestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		x, y := nums[i], nums[j]
		xbit, ybit := 10, 10
		for xbit <= x {
			xbit *= 10
		}
		for ybit <= y {
			ybit *= 10
		}
		return x*ybit+y > y*xbit+x
	})
	if nums[0] == 0 {
		return "0"
	}
	ans := []byte{}
	for _, x := range nums {
		ans = append(ans, strconv.Itoa(x)...)
	}
	return string(ans)
}
```