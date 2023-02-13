# 1637. Widest Vertical Area Between Two Points Containing No Points

### Medium

Given n points on a 2D plane where points[i] = [xi, yi], Return the widest vertical area between two points such that no points are inside the area.

A vertical area is an area of fixed-width extending infinitely along the y-axis (i.e., infinite height). The widest vertical area is the one with the maximum width.

Note that points on the edge of a vertical area are not considered included in the area.

### Example 1:

Input: points = [[8,7],[9,9],[7,4],[9,7]]
Output: 1
Explanation: Both the red and the blue area are optimal.

### Example 2:

Input: points = [[3,1],[9,0],[1,0],[1,4],[5,3],[8,8]]
Output: 3

Constraints:

n == points.length
2 <= n <= 10^5
points[i].length == 2
0 <= xi, yi <= 10^9

```go
//only need to compare the x-axis.
//sort the x-axis and then choose max (x[i] - x[i-1])
// O(nlogn) time. O(n) space.
func maxWidthOfVerticalArea(points [][]int) int {
	xs := []int{}
	l := len(points)
	for _, point := range points {
		xs = append(xs, point[0])
	}
	sort.Ints(xs)
	res := 0
	for i := 1; i < l; i++ {
		if xs[i]-xs[i-1] > res {
			res = xs[i] - xs[i-1]
		}
	}
	return res
}
```