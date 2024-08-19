# 373.查找和最小的 K 对数字 Find K Pairs with Smallest Sums

### Medium

You are given two integer arrays nums1 and nums2 sorted in non-decreasing order and an integer k.

Define a pair (u, v) which consists of one element from the first array and one element from the second array.

Return the k pairs (u1, v1), (u2, v2), ..., (uk, vk) with the smallest sums.

### Example 1:

- Input: nums1 = [1,7,11], nums2 = [2,4,6], k = 3
- Output: [[1,2],[1,4],[1,6]]
- Explanation: The first 3 pairs are returned from the sequence: [1,2],[1,4],[1,6],[7,2],[7,4],[11,2],[7,6],[11,4],[11,6]

### Example 2:

- Input: nums1 = [1,1,2], nums2 = [1,2,3], k = 2
- Output: [[1,1],[1,1]]
- Explanation: The first 2 pairs are returned from the sequence: [1,1],[1,1],[1,2],[2,1],[1,2],[2,2],[1,3],[1,3],[2,3]

### Constraints:

- 1 <= nums1.length, nums2.length <= 105
- -109 <= nums1[i], nums2[i] <= 109
- nums1 and nums2 both are sorted in non-decreasing order.
- 1 <= k <= 104
- k <= nums1.length * nums2.length

```go
// O(n^2) time, 超时
func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
    res:= [][]int{}
    for i:= 0; i<len(nums1); i++ {
        for j:= 0; j<len(nums2); j++ {
            res = append(res, []int{nums1[i], nums2[j]})
        }
    }
    sort.Slice(res, func(i,j int) bool {
        return  res[i][0] + res[i][1] <  res[j][0] + res[j][1]
    })
    return res[0:k]
}
```


堆排序，但是超内存了。。
```go
func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int { 
    h:= [][]int{}
    for i:=0; i<len(nums1) && i<k; i++ {
        for j:=0; j<len(nums2) && j<k; j++ {
            h = append(h, []int{nums1[i], nums2[j]})
        }
    }
    heapSize := len(h)
    buildMaxHeap(h, heapSize)
    for i := len(h) - 1; i >= 0; i-- {    
        h[0], h[i] = h[i], h[0]
        heapSize--
        maxHeapify(h, 0, heapSize)
    }
    return h[0:k]
}

func buildMaxHeap(a [][]int, heapSize int) {
    for i := heapSize/2; i >= 0; i-- {
        maxHeapify(a, i, heapSize)
    }
}

func maxHeapify(a [][]int, i, heapSize int) {
    l, r, largest := i * 2 + 1, i * 2 + 2, i
    if l < heapSize && a[l][0] + a[l][1] >  a[largest][0] + a[largest][1] {
        largest = l
    }
    if r < heapSize && a[r][0] + a[r][1] > a[largest][0] + a[largest][1] {
        largest = r
    }
    if largest != i {
        a[i], a[largest] = a[largest], a[i]
        maxHeapify(a, largest, heapSize)
    }
}
```



--------

下面的不对


或者用heap  或优先级队列。

```go
// An Item is something we manage in a priority queue.
type Item struct {
	x int
    y int
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].x + pq[i].y > pq[j].x + pq[j].y
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
// func (pq *PriorityQueue) update(item *Item, value string, priority int) {
// 	item.value = value
// 	item.priority = priority
// 	heap.Fix(pq, item.index)
// }

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
    res := [][]int{}
    
    pq := make(PriorityQueue, k)
	//heap.Init(&pq)
    
    for i, a:= range nums1 {
        for j, b:= range nums2 {
            index := i*len(nums1) + j
            item := &Item{
                x:    a,
                y:     b,
                index:    index,
            }
            heap.Push(&pq, item) 
            if index == 0 {
                heap.Init(&pq)
            }
            heap.Fix(&pq, item.index)
            if index >= k {
                heap.Pop(&pq)
            }

        }
    }


	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		res = append(res, []int{item.x, item.y})
	}
   
    return res
}
```