# 229. 多数元素 II

### 中等

给定一个大小为 n 的整数数组，找出其中所有出现超过 ⌊ n/3 ⌋ 次的元素。

### 示例 1：

    输入：nums = [3,2,3]
    输出：[3]

### 示例 2：

    输入：nums = [1]
    输出：[1]

### 示例 3：

    输入：nums = [1,2]
    输出：[1,2]

进阶：尝试设计时间复杂度为 O(n)、空间复杂度为 O(1)的算法解决此问题。

### 解：

我们可以利用反证法推断出满足这样条件的元素最多只有两个，我们可以利用摩尔投票法的核心思想，每次选择三个互不相同的元素进行删除（或称为「抵消」）。

```go
func majorityElement(nums []int) []int {
	ans := []int{}
	a, b := 0, 0
	ac, bc := 0, 0
	for _, num := range nums {
		if ac > 0 && num == a { // 如果该元素为第一个元素，则计数加1
			ac++
		} else if bc > 0 && num == b { // 如果该元素为第二个元素，则计数加1
			bc++
		} else if ac == 0 { // 选择第一个元素
			a = num
			ac++
		} else if bc == 0 { // 选择第二个元素
			b = num
			bc++
		} else { // 如果三个元素均不相同，则相互抵消1次
			ac--
			bc--
		}
		//fmt.Println(a, ac, b, bc)
	}
	cnt1, cnt2 := 0, 0
	for _, num := range nums {
		if ac > 0 && num == a {
			cnt1++
		}
		if bc > 0 && num == b {
			cnt2++
		}
	}
	if ac > 0 && cnt1 > len(nums)/3 {
		ans = append(ans, a)
	}
	if bc > 0 && cnt2 > len(nums)/3 {
		ans = append(ans, b)
	}
	return ans
}
```