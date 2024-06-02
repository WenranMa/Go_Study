# 215. 数组中的第K个最大元素

### 中等

给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。

请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

你必须设计并实现时间复杂度为 O(n) 的算法解决此问题。

### 示例 1:

输入: [3,2,1,5,6,4], k = 2
输出: 5

### 示例 2:

输入: [3,2,3,1,2,4,5,5,6], k = 4
输出: 4

### 提示：

1 <= k <= nums.length <= 10^5
-10^4 <= nums[i] <= 10^4

### 解：

1. 用内置接口实现最小堆。
```go
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] } // 最小堆
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Peek 返回堆顶元素，但不删除它
func (h *IntHeap) Peek() int {
	if h.Len() == 0 {
		return 0 // 或者返回一个错误，取决于你的需求
	}
	return (*h)[0]
}

func findKthLargest(nums []int, k int) int {
	h := &IntHeap{}
	heap.Init(h)

	for _, num := range nums {
		if h.Len() < k {
			heap.Push(h, num)
		} else if num > h.Peek() {
			heap.Pop(h)
			heap.Push(h, num)
		}
	}
	return (*h)[0]
}
```


2. 快排思想
```go
func findKthLargest(nums []int, k int) int {
    n := len(nums)
    return quickselect(nums, 0, n - 1, n - k)
}

func quickselect(nums []int, l, r, k int) int{
    if (l == r){
        return nums[k]
    }
    partition := nums[l]
    i := l - 1
    j := r + 1
    for (i < j) {
        for i++;nums[i]<partition;i++{}
        for j--;nums[j]>partition;j--{}
        if (i < j) {
            nums[i],nums[j]=nums[j],nums[i]
        }
		// fmt.Println(nums, partition, i, j, l, r)
    }
	
    if (k <= j){
        return quickselect(nums, l, j, k)
    }else{
        return quickselect(nums, j + 1, r, k)
    }
}


//快速排序
func quickSort(arr []int, low, high int) {
    if low < high {
        pi := partition(arr, low, high)
        quickSort(arr, low, pi-1)
        quickSort(arr, pi+1, high)
    }
}

func partition(arr []int, low, high int) int {
    pivot := arr[high]
    i := low - 1
    for j := low; j < high; j++ {
        if arr[j] <= pivot {
            i++
            arr[i], arr[j] = arr[j], arr[i]
        }
    }
    arr[i+1], arr[high] = arr[high], arr[i+1]
    return i + 1
}
```

3. 最小堆手搓
```go
func findKthLargest(nums []int, k int) int {
    heapSize := len(nums)
    buildMaxHeap(nums, heapSize)
    for i := len(nums) - 1; i >= len(nums) - k + 1; i-- {
        nums[0], nums[i] = nums[i], nums[0]
        heapSize--
        maxHeapify(nums, 0, heapSize)
    }
    return nums[0]
}

func buildMaxHeap(a []int, heapSize int) {
    for i := heapSize/2; i >= 0; i-- {
        maxHeapify(a, i, heapSize)
    }
}

func maxHeapify(a []int, i, heapSize int) {
    l, r, largest := i * 2 + 1, i * 2 + 2, i
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