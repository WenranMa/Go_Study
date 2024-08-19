# 918. 环形子数组的最大和

### 中等

给定一个长度为 n 的环形整数数组 nums ，返回 nums 的非空 子数组 的最大可能和 。

环形数组 意味着数组的末端将会与开头相连呈环状。形式上， nums[i] 的下一个元素是 nums[(i + 1) % n] ， nums[i] 的前一个元素是 nums[(i - 1 + n) % n] 。

子数组 最多只能包含固定缓冲区 nums 中的每个元素一次。形式上，对于子数组 nums[i], nums[i + 1], ..., nums[j] ，不存在 i <= k1, k2 <= j 其中 k1 % n == k2 % n 。

### 示例 1：

    输入：nums = [1,-2,3,-2]
    输出：3
    解释：从子数组 [3] 得到最大和 3

### 示例 2：

    输入：nums = [5,-3,5]
    输出：10
    解释：从子数组 [5,5] 得到最大和 5 + 5 = 10

### 示例 3：

    输入：nums = [3,-2,2,-3]
    输出：3
    解释：从子数组 [3] 和 [3,-2,2] 都可以得到最大和 3

### 提示：
- n == nums.length
- 1 <= n <= 3 * 10^4
- -3 * 10^4 <= nums[i] <= 3 * 10^4​

### 解：
参考leetcode 53题。

最大值可以出现再两个地方
1. 正常数组：num[i:j]
2. 环形数组：num[0:i] + num[j:n]

所以可以
1. 可以和53题一样计算mx. 
2. 然后取反，计算正常数组下最小值。同时计算数组总和。
3. 比较mx和sum-mi的最大值。
4. 注意如果mx<0，则数组全部都是负数，则返回mx。

```go
func maxSubarraySumCircular(nums []int) int {
	mx := -30001
	mi := 30001
	sumMax := 0
	sumMin := 0
	sum := 0
	for _, n := range nums {
		sumMax += n
		mx = max(mx, sumMax)
		if sumMax < 0 {
			sumMax = 0
		}
		sumMin += n
		mi = min(mi, sumMin)
		if sumMin > 0 {
			sumMin = 0
		}
		sum += n
	}
	if mx < 0 {
		return mx
	}
	return max(mx, sum-mi)
}
```