# 888_公平糖果交换_Fair_Candy_Swap
Alice and Bob have candy bars of different sizes: A[i] is the size of the i-th bar of candy that Alice has, and B[j] is the size of the j-th bar of candy that Bob has.

Since they are friends, they would like to exchange one candy bar each so that after the exchange, they both have the same total amount of candy.  (The total amount of candy a person has is the sum of the sizes of candy bars they have.)

Return an integer array ans where ans[0] is the size of the candy bar that Alice must exchange, and ans[1] is the size of the candy bar that Bob must exchange.

If there are multiple answers, you may return any one of them.  It is guaranteed an answer exists.

Example:

    Input: A = [1,1], B = [2,2]
    Output: [1,2]

Example 2:

    Input: A = [1,2], B = [2,3]
    Output: [1,2]

Example 3:

    Input: A = [2], B = [1,3]
    Output: [2,3]

Example 4:

    Input: A = [1,2,5], B = [2,4]
    Output: [5,4]

### 解：

求和，并用map记录其中一共数组的数。

```go
func fairCandySwap(A []int, B []int) []int {
	suma := 0
	sumb := 0
	mb := make(map[int]int)
	for _, a := range A {
		suma += a
	}
	for _, b := range B {
		sumb += b
		mb[b] = 1
	}
	// solve the equation: suma - x +y = sumb - y + x
	// y = x + (sumb - suma)/2
	d := (sumb - suma) / 2
	res := []int{}
	for _, a := range A {
		b := a + d
		if _, ok := mb[b]; ok {
			res = append(res, a, b)
			break
		}
	}
	return res
}
```