# 973_接近远点的k个点_K_Closest_Points_to_Origin
We have a list of points on the plane.  Find the K closest points to the origin (0, 0). (Here, the distance between two points on a plane is the Euclidean distance.) You may return the answer in any order.  The answer is guaranteed to be unique (except for the order that it is in.) 

Example:  

    Input: points = [[1,3],[-2,2]], K = 1   
    Output: [[-2,2]]   
    Explanation:   
    The distance between (1, 3) and the origin is sqrt(10).   
    The distance between (-2, 2) and the origin is sqrt(8).   
    Since sqrt(8) < sqrt(10), (-2, 2) is closer to the origin.    
    We only want the closest K = 1 points from the origin, so the answer is just [[-2,2]].   
    
    Input: points = [[3,3],[5,-1],[-2,4]], K = 2   
    Output: [[3,3],[-2,4]]   
    (The answer [[-2,4],[3,3]] would also be accepted.)   
    
Note:  
1 <= K <= points.length <= 10000   
-10000 < points[i][0] < 10000   
-10000 < points[i][1] < 10000   

### 解：

排序，按平方的距离排序，然后取前K个

时间 O(nlogn)

```go
func kClosest(points [][]int, K int) [][]int {
    sort.Slice(points, func(i, j int) bool {
        a := points[i][0]*points[i][0] + points[i][1]*points[i][1]
        b := points[j][0]*points[j][0] + points[j][1]*points[j][1]
        return a < b
    })
    return points[:K]
}
```

leetcode 373?

可以用堆来优化，时间 O(nlogK)

```go
type pair struct {
	dist  int
	point []int
}

type hp []pair

func (h *hp) Len() int {
	return len(*h)
}
func (h *hp) Less(i, j int) bool {
	return (*h)[i].dist > (*h)[j].dist
}
func (h *hp) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}
func (h *hp) Push(v interface{}) {
	*h = append(*h, v.(pair))
}
func (h *hp) Pop() interface{} {
	a := *h
	v := a[len(a)-1]
	*h = a[:len(a)-1]
	return v
}

func kClosest(points [][]int, k int) (ans [][]int) {
	h := make(hp, k)
	for i, p := range points[:k] {
		h[i] = pair{p[0]*p[0] + p[1]*p[1], p}
	}
	heap.Init(&h) // O(k) 初始化堆
	for _, p := range points[k:] {
		if dist := p[0]*p[0] + p[1]*p[1]; dist < h[0].dist {
			h[0] = pair{dist, p}
			heap.Fix(&h, 0) // 效率比 pop 后 push 要快
		}
	}
	for _, p := range h {
		ans = append(ans, p.point)
	}
	return
}
```