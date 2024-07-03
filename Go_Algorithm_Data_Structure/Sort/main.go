package main

import (
	"fmt"
)

// Selection Sort
// 分有序区（第一个元素开始）和无序区（从第二元素开始遍历）
// 从无序区找出最小的元素移动到有序区。O(N^2) time

func SelectSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

// Insert Sort
// 前面为排序好的，但不一定是最小的序列，后面为待排序的，每读入一个新数据（arr[i]），从后往前一次比较(inner loop)。
// O(N^2)time
func InsertSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
}

// Bubble Sort
// 相邻两两比较，大的往后排，小的往前排，
// 一轮下来，最大的数会冒到最前面，然后下一轮，剩下的数会冒到第二位，依次类推，直到冒到最末尾。O(N^2) time
func BubbleSort(arr []int) {
	// for i := 0; i < len(arr); i++ {
	// 	for j := 0; j < len(arr)-i-1; j++ {
	// 		if arr[j] > arr[j+1] {
	// 			arr[j], arr[j+1] = arr[j+1], arr[j]
	// 		}
	// 	}
	// }
	for i := len(arr) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// Shell Sort --- TBD

// Merge Sort, O(nlogn) time, O(nlogn) space?
func merge(nums []int, l, mid, r int) {
	left := make([]int, mid-l+1) // 注意下标大小
	right := make([]int, r-mid)
	copy(left, nums[l:mid+1]) //right exclusive.
	copy(right, nums[mid+1:r+1])

	i, j, k := 0, 0, l // 注意K的下标 不一定是零
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			nums[k] = left[i]
			k, i = k+1, i+1
		} else {
			nums[k] = right[j]
			k, j = k+1, j+1
		}
	}
	for j < len(right) {
		nums[k] = right[j]
		k, j = k+1, j+1
	} // 剩余数据的处理
	for i < len(left) {
		nums[k] = left[i]
		k, i = k+1, i+1
	}
}

func mergeSort(nums []int, l, r int) {
	if l >= r {
		return
	}
	mid := (l + r) / 2
	mergeSort(nums, l, mid)
	mergeSort(nums, mid+1, r)
	merge(nums, l, mid, r)
}

func MergeSort(nums []int) {
	mergeSort(nums, 0, len(nums)-1)
}

// Quick Sort, O(nlogn)time, O(1) space.
/*
思想：先从数列中取出一个数作为基准数，CLRS中选取的数组最后一个数，将比这个数大的数全放到它的右边，小于或等于它的数全放到它的左边，然后递归。
运行时间与划分是否对称有关，如果对称，则与合并排序一样快。
最坏情况是每一次都划分成一个n-1个元素和一个0个元素（因为参考值不算），就是说每次data[to]都是最小的。O(n^2).
最佳情况是每次都是可以分解成两个元素个数非0的数组，两个数组长度可以不等，这样可以形成高度为logn的递归数。O(nlogn) time.
*/
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

func QuickSort(arr []int) {
	quickSort(arr, 0, len(arr)-1)
}

// Heap sort
/*
堆，二叉堆，一棵完全二叉树。树中每个节点与数组中存放该节点值的元素相对应。树的每一层都是填满的，除了最后一层。
节点在堆中的高度定义为从本节点到叶子节点的最长简单下降路径上边的数目，而堆的高度就是树根的高度。
因此，在高度为h的堆中，共有h+1层，最多有2^(h+1)-1个元素，最少有2^h个元素。而n个元素的堆的高度为log2n的下取整。
整个堆如果有n个元素，则叶子是n/2的下取整+1, +2, +3,...n. 比如n=9，9个元素，则叶子的下标是5 6 7 8 9。注意此时堆的下标是以1开始的。
实际编程中下标为0开始，所以如果是n个元素，叶子应该是n/2的下取整, +1, +2, +3,...n-1.剩下的就是非叶子节点。
最后整个heapsort的时间为O(nlogn). 因为n-1次调用max-heapify(这句存疑). O(1)Space.
*/
func HeapSort(nums []int) {
	heapSize := len(nums)
	buildMaxHeap(nums, heapSize)
	for i := len(nums) - 1; i >= 0; i-- {
		nums[0], nums[i] = nums[i], nums[0] // 把最大的元素放到最后
		heapSize--
		maxHeapify(nums, 0, heapSize)
	}
}

func buildMaxHeap(a []int, heapSize int) {
	for i := heapSize/2 - 1; i >= 0; i-- { // len/2-1 就是对所有非叶子的节点下沉
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
func main() {
	arr := []int{23, 0, 12, 56, 34, -1, 55}
	InsertSort(arr)
	fmt.Println(arr)

	arr = []int{23, 0, 12, 56, 34, -1, 55}
	SelectSort(arr)
	fmt.Println(arr)

	arr = []int{23, 0, 12, 56, 34, -1, 55}
	BubbleSort(arr)
	fmt.Println(arr)

	arr = []int{23, 0, 12, 56, 34, -1, 55}
	MergeSort(arr)
	fmt.Println(arr)

	arr = []int{23, 0, 12, 56, 34, -1, 55}
	QuickSort(arr)
	fmt.Println(arr)

	arr = []int{23, 0, 12, 56, 34, -1, 55}
	HeapSort(arr)
	fmt.Println(arr)
}
