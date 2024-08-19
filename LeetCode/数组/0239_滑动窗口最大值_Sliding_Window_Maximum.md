# 239. 滑动窗口最大值

### 困难

给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。

返回 滑动窗口中的最大值 。

### 示例 1：

    输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
    输出：[3,3,5,5,6,7]
    解释：
    滑动窗口的位置                最大值
    ---------------               -----
    [1  3  -1] -3  5  3  6  7       3
    1 [3  -1  -3] 5  3  6  7       3
    1  3 [-1  -3  5] 3  6  7       5
    1  3  -1 [-3  5  3] 6  7       5
    1  3  -1  -3 [5  3  6] 7       6
    1  3  -1  -3  5 [3  6  7]      7

### 示例 2：

    输入：nums = [1], k = 1
    输出：[1]

### 解：


超时了。。。
```go
func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	ans := make([]int, n-k+1)
	ans[0] = findMax(nums[0:k])
	for i := k; i < n; i++ {
		ans[i-k+1] = findMax(nums[i-k+1 : i+1])
	}
	return ans
}

func findMax(nums []int) int {
	res := math.MinInt32
	for _, n := range nums {
		res = max(res, n)
	}
	return res
}

// 最大堆，也超时 （可能更慢）
func maxSlidingWindow(nums []int, k int) []int {
	q1, q2 := make([]int, k), make([]int, k)
	for i := 0; i < k; i++ {
		q2[i] = nums[i]
	}
	copy(q1, q2)
	buildMaxHeap(q1, k)
	n := len(nums)
	ans := make([]int, n-k+1)
	ans[0] = q1[0]
	for i := k; i < n; i++ {
		q2 = append(q2, nums[i])
		q2 = q2[1 : k+1]
		copy(q1, q2)
		buildMaxHeap(q1, k)
		ans[i-k+1] = q1[0]
	}
	return ans
}

func buildMaxHeap(a []int, heapSize int) {
	for i := heapSize / 2; i >= 0; i-- {
		maxHeapify(a, i, heapSize)
	}
}

func maxHeapify(a []int, i, heapSize int) {
	l, r, largest := i*2+1, i*2+2, i
	if l < heapSize && a[l] > a[largest] {
		largest = l
	}
	if r < heapSize && a[r] > a[largest] {
		largest = r
	}
	if largest != i {
		a[i], a[largest] = a[largest], a[i]
		maxHeapify(a, largest, heapSize)
	}
}
```